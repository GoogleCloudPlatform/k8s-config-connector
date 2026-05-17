## 2026-05-16

- The Go files in mocknetapp had already been migrated to use httptogrpc in a previous commit, but the generated protos and Makefile entries were left behind. Make sure to delete the generated protobuf files and update the Makefile.

## 2026-05-16 - datastream migration

- When migrating `mockdatastream` away from `httpmux` to `httptogrpc`, I found that `httptogrpc.grpcMux` does not support `RewriteError`. Previously, `mux.RewriteError` was used to manipulate the error response (e.g., removing `error.Errors` on a 404). This can be safely removed when using `httptogrpc`.

## 2026-05-17 - bigquery migration

- BigQuery v2 protos are located in `cloud.google.com/go/bigquery/v2/apiv2/bigquerypb`. This required a `go get` as it was not in the main `bigquery` module.
- `httptogrpc` was enhanced to support `x-http-code` and `x-expires` headers from gRPC metadata, allowing mocks to return custom HTTP status codes (like 204 No Content).
- `httptogrpc` was also enhanced to match the standard GCP error structure, including the `errors` array for 404 (NotFound) errors, which is expected by BigQuery clients.
- Some BigQuery methods were combined in the official proto (e.g., `UpdateOrPatchDatasetRequest`), requiring updates to the server implementation.
- Many fields changed from pointers to direct values or wrappers (`wrapperspb`).
