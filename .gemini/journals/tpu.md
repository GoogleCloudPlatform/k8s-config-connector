# TPU Journal

### [2026-06-15] TPUQueuedResource Direct Types Scaffold and IdentityV2
- **Context**: Implementing Greenfield types, CRD, and IdentityV2 for `TPUQueuedResource` (Kind: `TPUQueuedResource`, GCP Resource: `QueuedResource`)
- **Problem**: 
  1. The issue instructions suggested using service `google.cloud.tpu.v1`. However, TPU `v1` protobuf definitions do not contain the `QueuedResource` message. It is only defined in TPU `v2` (and `v2alpha1`).
  2. Running the generator for `QueuedResource` using `google.cloud.tpu.v2` generated types that references the `Node` message, but the `Node` message is registered as a KCC Resource Kind (`TPUVirtualMachine`) and thus its Go struct representation isn't automatically generated inline under `Node`.
  3. `gcpurls.Template` failed with panic `field "queued_resource" not found in struct v1alpha1.TPUQueuedResourceIdentity` when initializing with `{queued_resource}` segment in the template while Go field was `QueuedResource`.
- **Solution**:
  1. Updated `apis/tpu/v1alpha1/generate.sh` to use `--service google.cloud.tpu.v2` for `TPUQueuedResource`.
  2. Manually defined `Node` and `NodeObservedState` structs inside `apis/tpu/v1alpha1/queuedresource_types.go` mapping to `google.cloud.tpu.v2.Node` so the generator can resolve the reachable references during types pruning.
  3. Named the Go struct field matching `{queued_resource}` exactly as `Queued_resource` inside `TPUQueuedResourceIdentity`.
- **Impact**: Enables smooth generation of all related nested structs of TPU Node and correct parsing of TPU QueuedResource URL format.
