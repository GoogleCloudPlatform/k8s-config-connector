# Journal: MonitoringGroup Direct Controller Migration

## Observations
1. **gRPC Client Requirements**: The GCP `monitoringpb.GroupService` does not have a REST-based RESTClient equivalent in the newer monitoring API client library version. We had to instantiate standard gRPC clients via `monitoring.NewGroupClient`.
2. **GRPCClientOptions vs RESTClientOptions**: Since the `GroupClient` uses gRPC, building the client with `RESTClientOptions()` threw `WithHTTPClient is incompatible with gRPC dial options`. We correctly switched to `GRPCClientOptions()` on the `ControllerConfig` which resolved the issue.
3. **ResolveProject Interface Package Mismatch**: `ResolveProject` in `refsv1beta1` expects a `*refsv1beta1.ProjectRef` type, but `MonitoringGroupSpec.ProjectRef` uses the `apis/refs` package (`*refs.ProjectRef`). We successfully mapped `*refs.ProjectRef` to `*refsv1beta1.ProjectRef` before calling `ResolveProject` to preserve strong-typing and prevent type collisions.
4. **Parent Parameter**: `CreateGroupRequest` on the gRPC client uses `Name` rather than `Parent` as the field for the parent project resource path. We updated the struct literal accordingly.
