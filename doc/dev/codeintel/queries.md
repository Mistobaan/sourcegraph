<!--
# Ranges

<object data="/dev/codeintel/diagrams/ranges.svg" type="image/svg+xml"></object>
-->

# Definitions

<a href="/dev/codeintel/diagrams/definitions.svg" target="_blank">
  <img src="/dev/codeintel/diagrams/definitions.svg">
</a>

# References

<a href="/dev/codeintel/diagrams/references.svg" target="_blank">
  <img src="/dev/codeintel/diagrams/references.svg">
</a>
<a href="/dev/codeintel/diagrams/resolve-page.svg" target="_blank">
  <img src="/dev/codeintel/diagrams/resolve-page.svg">
</a>

# Hover

<a href="/dev/codeintel/diagrams/hover.svg" target="_blank">
  <img src="/dev/codeintel/diagrams/hover.svg">
</a>

# Code Appendix

- [PositionAdjuster.AdjustPosition](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/position.go#L63:28)
- [PositionAdjuster.AdjustRange](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/position.go#L77:28)

- [QueryResolver](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/resolver.go#L73:20)
- [queryResolver.Definitions](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/query.go#L138:25)
- [queryResolver.References](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/query.go#L167:25)
- [queryResolver.Hover](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/resolvers/query.go#L236:25)

- [DecodeOrCreateCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/cursor.go#L54:6)

- [CodeIntelAPI.FindClosestDumps](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/exists.go#L18:26)
- [CodeIntelAPI.Definitions](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/definitions.go#L21:26)
- [CodeIntelAPI.Hover](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/hover.go#L13:26)
- [CodeIntelAPI.References](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L24:26)

- [Database.Definitions](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L166:25)
- [Database.References](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L189:25)
- [Database.Hover](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L213:25)
- [Database.MonikersByPosition](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L284:25)
- [Database.PackageInformation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L358:25)
- [Database.MonikerResults](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/cmd/precise-code-intel-bundle-manager/internal/database/database.go#L316:25)

- [Store.FindClosestDumps](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/store/dumps.go#L99:17)
- [Store.GetPackage](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/store/packages.go#L11:17)
- [Store.SameRepoPager](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/store/references.go#L39:17)
- [Store.PackageReferencePager](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/store/references.go#L77:17)

- [ReferencePageResolver.resolvePage](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L50:33)
- [ReferencePageResolver.handleSameDumpCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L91:33)
- [ReferencePageResolver.handleSameDumpMonikersCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L137:33)
- [ReferencePageResolver.handleDefinitionMonikersCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L218:33)
- [ReferencePageResolver.handleSameRepoCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L283:33)
- [ReferencePageResolver.handleRemoteRepoCursor](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/enterprise/internal/codeintel/api/references.go#L311:33)
