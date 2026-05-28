When migrating a service like `metastore` that was previously registered manually in `mockgcp/mock_http_roundtrip.go`, it is a good opportunity to move it to the newer registration pattern:
1. Add an `init` function in the service's `service.go` that calls `mockgcpregistry.Register(New)`.
2. Add a blank import for the service package in `mockgcp/register.go`.
3. Remove the manual registration and import from `mockgcp/mock_http_roundtrip.go`.

This keeps `mock_http_roundtrip.go` smaller and makes the registration process more consistent across services.
