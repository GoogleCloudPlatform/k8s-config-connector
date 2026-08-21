# Journal: ComputeInstanceGroup transition to direct KRM types

## Observations & Key Learnings

1. **Non-pointer string fields mapping mismatch**:
   - In KRM types, the `zone` field in `ComputeInstanceGroupSpec` is a required non-pointer string field, while in GCP `InstanceGroup` proto message, the `zone` field is represented as a pointer string (`*string`).
   - The default `generate-mapper` tool outputs standard assignment statements (like `out.Zone = in.Zone`) which fail compilation on type mismatches when assigning pointer strings to non-pointer strings.
   - Hand-coding the Spec mappers (`ComputeInstanceGroupSpec_v1beta1_FromProto` and `ComputeInstanceGroupSpec_v1beta1_ToProto`) in `instancegroup_mapper.go` perfectly bypassed this limitation.

2. **Pointer to value fields in nested arrays**:
   - The KRM field `namedPort` has required non-pointer sub-fields (`name string` and `port int32`) to conform strictly to the CRD schema.
   - However, in GCP proto message `NamedPort`, these fields are pointer types.
   - We handcoded custom `InstanceGroupNamedPort_v1beta1_FromProto` and `InstanceGroupNamedPort_v1beta1_ToProto` mapper functions to safely unpack the pointers and handle default/fallback values without generating compile-time assignment errors.

3. **Separate status mapping**:
   - Because `ComputeInstanceGroupStatus` only maps standard fields (`selfLink` and `size`) and doesn't require direct-controller common status patterns (like `observedState` or `externalRef`), we hand-coded the status mapping functions to achieve 100% schema compatibility without altering the status structure.
