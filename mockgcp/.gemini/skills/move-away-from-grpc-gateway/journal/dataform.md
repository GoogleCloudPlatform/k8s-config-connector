# Dataform Migration to httptogrpc

- Moved `mockgcp/mockdataform` from `grpc-gateway` to `httptogrpc`.
- Updated imports to use official client library: `cloud.google.com/go/dataform/apiv1beta1/dataformpb`.
- Fixed `proto.CloneOf` usage to `proto.Clone` and type assertion (though `proto.CloneOf` might have worked if it was available, I opted for the standard way).
- Discovered that `mockdataform` was registered manually in `mockgcp/mock_http_roundtrip.go` and was missing from `mockgcp/register.go`.
- Moved `mockdataform` to use `mockgcpregistry` for automatic registration.
- Added `mockdataform` to `mockgcp/register.go`.
- Removed manual registration from `mockgcp/mock_http_roundtrip.go`.
- Removed `httpmux.SetStatusCode(ctx, http.StatusNoContent)` as `httptogrpc` does not yet support it via the same mechanism. This might need investigation if 204 responses are required for stability of logs.
