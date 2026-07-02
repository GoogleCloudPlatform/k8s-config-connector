# Cloud IDS Migration to httptogrpc

## Migration to httptogrpc

- Moved Cloud IDS away from `grpc-gateway` to `httptogrpc`.
- Kept the custom proto generated from Discovery JSON (`mockgcp.cloud.ids.v1`) instead of switching to the official Go library (`cloud.google.com/go/ids/apiv1/idspb`).
- Rationale: The official library is currently missing the `PatchEndpoint` method and the `threatExceptions` field in the `Endpoint` message, both of which are required by Config Connector and present in the API discovery document.
- Updated `service.go` to use `httptogrpc.NewGRPCMux` and `mux.AddService(pb.NewProjectsLocationsEndpointsServerClient(conn))`.
- Updated `Register` to use `pb.RegisterProjectsLocationsEndpointsServerServer`.
- Verified that PATCH requests are correctly routed and handled by the mock using `httptogrpc`.

## Challenges

- Encountered "no space left on device" errors during `go test`. Resolved by cleaning go cache and setting `GOTMPDIR=/tmp`.
- Missing `protoc` binary in the environment prevented re-generating protos, so deleted `.pb.go` files were restored using `git checkout`.
