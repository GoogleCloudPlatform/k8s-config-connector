# CloudDMS Journal

### 2026-05-24 CloudDMS PrivateConnection Direct Implementation and gRPC Client Options
- **Context**: [Implement direct controller for CloudDMSPrivateConnection](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8637)
- **Problem**: 
  1. The CloudDMS PrivateConnection API is completely immutable and does not expose any `UpdatePrivateConnection` method or endpoint.
  2. The Google Cloud DMS Go library (`cloud.google.com/go/clouddms/apiv1`) uses gRPC by default. Using `m.config.RESTClientOptions()` adds `WithHTTPClient`, which causes a panic/error during gRPC client creation: `WithHTTPClient is incompatible with gRPC dial options`.
- **Solution**: 
  1. Implemented immutable `Update` semantics in `privateconnection_controller.go` by comparing desired fields against actual fields, and returning an unsupported error if any spec difference is found.
  2. Used `m.config.GRPCClientOptions()` instead of `RESTClientOptions()` when initializing the `DataMigrationClient` to avoid incompatible HTTP options.
- **Impact**: Future direct controllers targeting clouddms (or other gRPC-default services) can refer to this pattern to build their clients safely and handle immutable resources correctly.
