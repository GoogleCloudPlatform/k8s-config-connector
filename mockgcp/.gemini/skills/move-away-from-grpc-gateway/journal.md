## 2026-05-16

- The Go files in mocknetapp had already been migrated to use httptogrpc in a previous commit, but the generated protos and Makefile entries were left behind. Make sure to delete the generated protobuf files and update the Makefile.

## 2026-05-16 - datastream migration

- When migrating `mockdatastream` away from `httpmux` to `httptogrpc`, I found that `httptogrpc.grpcMux` does not support `RewriteError`. Previously, `mux.RewriteError` was used to manipulate the error response (e.g., removing `error.Errors` on a 404). This can be safely removed when using `httptogrpc`.

## 2026-05-17
- Moved mockgcp dataplex to httptogrpc.
- Observed that dataplex protos were under `google/cloud/dataplex` instead of `mockgcp/cloud/dataplex`. Updated the skill to reflect this possibility.
- Dataplex has multiple services (`DataplexService`, `CatalogService`, etc.). `NewHTTPMux` should register all services that were previously registered.

## 2026-05-17 - essentialcontacts migration
- Moved mockgcp essentialcontacts to httptogrpc.
- The service was straightforward and did not require long-running operations support (`AddOperationsPath`).
- Updated imports in `contact.go` and `service.go` to use `cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb`.
