package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
)

// ThreadPreview is the interface for the GraphQL type ThreadPreview.
type ThreadPreview interface {
	Repository(context.Context) (*RepositoryResolver, error)
	Title() string
	Author(context.Context) (*Actor, error)
	Body() string
	BodyText() string
	BodyHTML() string
	Kind(context.Context) (ThreadKind, error)
	RepositoryComparison(context.Context) (*RepositoryComparisonResolver, error)
	Diagnostics(context.Context, *graphqlutil.ConnectionArgs) (DiagnosticConnection, error)
}

// ThreadOrThreadPreviewConnection is the interface for the GraphQL type ThreadOrThreadPreviewConnection.
type ThreadOrThreadPreviewConnection interface {
	Nodes(context.Context) ([]ToThreadOrThreadPreview, error)
	TotalCount(context.Context) (int32, error)
	PageInfo(context.Context) (*graphqlutil.PageInfo, error)
}

type ToThreadOrThreadPreview struct {
	Thread        Thread
	ThreadPreview ThreadPreview
}

func (v ToThreadOrThreadPreview) thread() interface {
	Repository(ctx context.Context) (*RepositoryResolver, error)
	Kind(context.Context) (ThreadKind, error)
	RepositoryComparison(context.Context) (*RepositoryComparisonResolver, error)
} {
	switch {
	case v.Thread != nil:
		return v.Thread
	case v.ThreadPreview != nil:
		return v.ThreadPreview
	default:
		panic("invalid ToThreadOrThreadPreview")
	}
}

func (v ToThreadOrThreadPreview) Repository(ctx context.Context) (*RepositoryResolver, error) {
	return v.thread().Repository(ctx)
}

func (v ToThreadOrThreadPreview) RepositoryComparison(ctx context.Context) (*RepositoryComparisonResolver, error) {
	return v.thread().RepositoryComparison(ctx)
}

func (v ToThreadOrThreadPreview) Diagnostics(ctx context.Context, args *graphqlutil.ConnectionArgs) (DiagnosticConnection, error) {
	switch {
	case v.Thread != nil:
		return v.Thread.Diagnostics(ctx, &ThreadDiagnosticConnectionArgs{ConnectionArgs: *args})
	case v.ThreadPreview != nil:
		return v.ThreadPreview.Diagnostics(ctx, args)
	default:
		panic("invalid ToThreadOrThreadPreview")
	}
}

func (v ToThreadOrThreadPreview) ToThread() (Thread, bool) { return v.Thread, v.Thread != nil }
func (v ToThreadOrThreadPreview) ToThreadPreview() (ThreadPreview, bool) {
	return v.ThreadPreview, v.ThreadPreview != nil
}
