## GKE Hub Migration (2026-05-17)

- GKE Hub uses two different versions: `v1beta` and `v1beta1`.
- `v1beta` was not found in `cloud.google.com/go/gkehub`, so it was imported from `google.golang.org/genproto/googleapis/cloud/gkehub/v1beta`.
- `v1beta1` was imported from the official SDK: `cloud.google.com/go/gkehub/apiv1beta1/gkehubpb`.
- Both were registered in `NewHTTPMux` using `httptogrpc`.
- Removed `mux.RewriteError` as per skill instructions.