## 2026-05-17 - artifactregistry migration

- Migrated `mockartifactregistry` to `httptogrpc` and official Go client library `cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb`.
- `NewHTTPMux` was updated to use `httptogrpc.NewGRPCMux`.
- Removed `mux.RewriteError` as it's not supported by `httptogrpc` and generally not needed for standard behavior.
- Deleted `mockgcp/generated/mockgcp/devtools/artifactregistry/` and removed it from `mockgcp/Makefile`.
- Build verified with `go build ./mockgcp/...`.
