# Network Management Migration Journal

## Migration to httptogrpc

- The migration of `networkmanagement` was straightforward.
- Updated `service.go` to use `httptogrpc.NewGRPCMux` and `mux.AddService(pb.NewReachabilityServiceClient(conn))`.
- Updated imports in `service.go` and `connectivitytest.go` to use `cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb`.
- `proto.CloneOf` was already used in the codebase and worked fine with the new proto types.
- The `longrunningpb` import was updated to `cloud.google.com/go/longrunning/autogen/longrunningpb`.
- Removed `mux.RewriteError` as it is not supported by `httptogrpc`.

## Challenges

- Encountered "no space left on device" errors during `go test`. This was resolved by clearing the go build cache and temporary files, and setting `GOTMPDIR` to a larger partition.
- `envtest` assets had to be manually set up using `KUBEBUILDER_ASSETS` to run the e2e tests locally.
