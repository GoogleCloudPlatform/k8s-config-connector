## 2026-05-18 - serviceusage migration

- Moved mockgcp serviceusage to httptogrpc.
- The official Go client for `v1` is in `cloud.google.com/go/serviceusage/apiv1/serviceusagepb`.
- However, `v1beta1` is not in the same module; it was found in `google.golang.org/genproto/googleapis/api/serviceusage/v1beta1`.
- Enhanced `httptogrpc` to set `common.MetadataKeyHttpRequestQuery` ("http.request.query") in the context metadata. This was needed for `serviceusage`'s custom field filtering logic which originally used `httpmux.MetadataKeyHttpRequestQuery`.
- Moved `MetadataKeyHttpRequestQuery` constant from `httpmux` to a new `mockgcp/common/metadata.go` to avoid circular dependencies.
- `httptogrpc` does not support `RewriteHeaders`. The `serviceusage` mock previously used it to delete the `Cache-Control` header. This was removed during migration.
