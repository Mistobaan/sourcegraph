package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sourcegraph/mux"
	"src.sourcegraph.com/sourcegraph/app/internal"
	"src.sourcegraph.com/sourcegraph/app/internal/tmpl"
	"src.sourcegraph.com/sourcegraph/errcode"
	"src.sourcegraph.com/sourcegraph/util/httputil"
)

const (
	routeIndex = "index"
	routePost  = "post"
)

type tumblr struct {
	// Path is the path to this blog under the liveblog root
	Path string

	// Blog is the name of the blog on Tumblr
	Blog string

	// BlogTitle is the title of the blog
	BlogTitle string

	// BlogBanner is HTML to show as the title of the blog
	BlogBanner template.HTML
}

var (
	tumblrAPIKey = os.Getenv("SG_TUMBLR_API_KEY")
	tumblrHost   = "https://api.tumblr.com"
)

var tumblrBlogs = []*tumblr{{
	Path:       "/blog/live/gopherconindia/",
	Blog:       "gopherconindia.tumblr.com",
	BlogTitle:  "GopherConIndia liveblog",
	BlogBanner: `<img class="img-responsive" src="https://s3-us-west-2.amazonaws.com/sourcegraph-assets/blog/gopherconindia_liveblog_banner.png" alt="Sourcegraph @ GopherConIndia">`,
}, {
	Path:       "/blog/live/gophercon/",
	Blog:       "gophercon.sourcegraph.com",
	BlogTitle:  "GopherCon 2014 liveblog",
	BlogBanner: `<h1 class="blog-title">GopherCon 2014 Liveblog</h1>`,
}, {
	Path:       "/blog/live/gophercon2015/",
	Blog:       "gophercon2015.tumblr.com",
	BlogTitle:  "GopherCon 2015 liveblog",
	BlogBanner: `<img class="img-responsive" src="https://s3-us-west-2.amazonaws.com/sourcegraph-assets/blog/gophercon2015_liveblog_banner.png" alt="Sourcegraph @ GopherCon 2015>`,
}}

var liveblogHandler http.Handler

func init() {
	m := http.NewServeMux()
	for _, t := range tumblrBlogs {
		m.Handle(t.Path, t.NewRouter(mux.NewRouter().PathPrefix(t.Path).Subrouter()))
	}
	liveblogHandler = m
}

type ListOptions struct {
	Page    int `url:",omitempty"`
	PerPage int `url:",omitempty"`
}

type ResponseWrapper struct {
	Meta     map[string]interface{} `json:"meta"`
	Response json.RawMessage        `json:"response"`
}

type PostsResponse struct {
	Blog       map[string]interface{} `json:"blog"`
	Posts      []*Post                `json:"posts"`
	TotalPosts int                    `json:"total_posts"`
}

const (
	typeText = "text"
)

type Post struct {
	BlogName    string   `json:"blog_name,omitempty"`
	ID          int      `json:",omitempty"`
	PostURL     string   `json:"post_url,omitempty"`
	PostAuthor  string   `json:"post_author,omitempty"`
	Type        string   `json:",omitempty"`
	Timestamp   int      `json:",omitempty"`
	Date        string   `json:",omitempty"`
	Format      string   `json:",omitempty"`
	ReblogKey   string   `json:"reblog_key,omitempty"`
	Tags        []string `json:",omitempty"`
	Bookmarklet bool     `json:",omitempty"`
	Mobile      bool     `json:",omitempty"`
	SourceURL   string   `json:"source_url,omitempty"`
	SourceTitle string   `json:"source_title,omitempty"`
	Liked       bool     `json:",omitempty"`
	State       string   `json:",omitempty"`

	Title       string
	Body        template.HTML
	Photos      []map[string]interface{}
	Caption     string
	Width       int
	Height      int
	Text        template.HTML
	Source      template.HTML
	URL         string
	Description template.HTML

	SGURL string `json:",omitempty"`
}

type PostsOpts struct {
	DisableSanitize bool
	ID              string
	Type            string
	ListOptions
}

func (t *tumblr) Posts(opt PostsOpts) (*PostsResponse, error) {
	query := make(url.Values)
	if opt.PerPage != 0 {
		query.Set("offset", strconv.Itoa(opt.PerPage*(opt.Page-1)))
	}
	if opt.PerPage != 0 {
		query.Set("limit", strconv.Itoa(opt.PerPage))
	}
	if opt.ID != "" {
		query.Set("id", opt.ID)
	}
	query.Set("api_key", tumblrAPIKey)
	var u string
	if opt.Type != "" {
		u = fmt.Sprintf("%s/v2/blog/%s/posts/%s?%s", tumblrHost, t.Blog, opt.Type, url.Values(query).Encode())
	} else {
		u = fmt.Sprintf("%s/v2/blog/%s/posts?%s", tumblrHost, t.Blog, url.Values(query).Encode())
	}
	resp, err := httputil.CachingClient.Get(u)
	if err != nil {
		return nil, err
	}

	var wrapper ResponseWrapper
	err = json.NewDecoder(resp.Body).Decode(&wrapper)
	if err != nil {
		return nil, err
	}
	if status, ok := wrapper.Meta["status"].(float64); ok && status != 200 {
		err := fmt.Errorf("Tumblr status %f, message: %s", status, wrapper.Meta["msg"])
		var s int
		if status == 404.0 {
			s = 404
		} else {
			s = 500
		}
		return nil, &errcode.HTTPErr{
			Status: s,
			Err:    err,
		}
	}
	var postsResp PostsResponse
	err = json.Unmarshal(wrapper.Response, &postsResp)
	if err != nil {
		return nil, err
	}

	if !opt.DisableSanitize {
		// sanitize
		for _, post := range postsResp.Posts {
			post.Body = sanitizeTumblrHTML(post.Body)
			post.Description = sanitizeTumblrHTML(post.Description)
			post.Text = sanitizeTumblrHTML(post.Text)
			post.Source = sanitizeTumblrHTML(post.Source)
			for _, photo := range post.Photos {
				if caption, isStr := photo["caption"].(string); isStr {
					photo["caption"] = sanitizeTumblrHTML(template.HTML(caption))
				}
			}
		}
	}

	// postprocessing
	for _, post := range postsResp.Posts {
		post.SGURL = fmt.Sprintf("%s%d", t.Path, post.ID)
	}

	return &postsResp, nil
}

func (t *tumblr) NewRouter(r *mux.Router) *mux.Router {
	r.Path("/").Methods("GET").Name(routeIndex).Handler(internal.Handler(t.serveIndex))
	r.Path("/{ID}").Methods("GET").Name(routePost).Handler(internal.Handler(t.servePost))
	return r
}

func (t *tumblr) serveIndex(w http.ResponseWriter, r *http.Request) error {
	var opt PostsOpts
	if err := schema.NewDecoder().Decode(&opt, r.URL.Query()); err != nil {
		return err
	}
	if opt.PerPage == 0 || opt.PerPage > 10 {
		opt.PerPage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	resp, err := t.Posts(opt)
	if err != nil {
		return err
	}

	return tmpl.Exec(r, w, "liveblog/index.html", http.StatusOK, nil, &struct {
		*tumblr
		Response *PostsResponse

		Limit  int
		Offset int

		tmpl.Common
	}{
		tumblr:   t,
		Response: resp,

		Limit:  opt.PerPage,
		Offset: opt.PerPage * (opt.Page - 1),
	})
}

func (t *tumblr) servePost(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	// Redirect dead reddit post to the correct id for the
	// bleve blog post:
	// https://www.reddit.com/r/golang/comments/2ygntc/
	if vars["ID"] == "113149531137" {
		vars["ID"] = "113241457917"
		varsList := make([]string, 2*len(vars))
		i := 0
		for name, val := range vars {
			varsList[i*2] = name
			varsList[i*2+1] = val
			i++
		}
		url, err := t.NewRouter(mux.NewRouter().PathPrefix(t.Path).Subrouter()).Get(routePost).URL(varsList...)
		if err != nil {
			return err
		}
		http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
		return nil
	}
	resp, err := t.Posts(PostsOpts{ID: vars["ID"]})
	if err != nil {
		return err
	}

	if len(resp.Posts) != 1 {
		return fmt.Errorf("expected 1 post, but got %d", len(resp.Posts))
	}

	return tmpl.Exec(r, w, "liveblog/post.html", http.StatusOK, nil, &struct {
		*tumblr
		Post *Post

		tmpl.Common
	}{
		tumblr: t,
		Post:   resp.Posts[0],
	})
}

// sanitizeTumblrHTML sanitizes HTML.
//
// This should NOT be relied upon to prevent reflected XSS attacks.
// Access to the Tumblr blog account must be restricted to Sourcegraph
// employees.
func sanitizeTumblrHTML(origHTML template.HTML) template.HTML {
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("class").Globally()
	return template.HTML(policy.Sanitize(string(origHTML)))
}
