# Journal: ComputePacketMirroring transition to direct KRM types

## Observations & Key Learnings

1. **Handling Value-type Nested Structs**:
   - The baseline CRD for `ComputePacketMirroring` defined several nested structs like `collectorIlb`, `mirroredResources`, and `network` as non-pointer (value) structs.
   - The automatic mapping generator `generate-mapper` treats annotated types as pointers, which leads to pointer-to-value assignment compilation errors. Handcoding the mapper functions cleanly solves these value-vs-pointer assignment constraints.

2. **References without `kind` Field**:
   - The baseline CRD utilized standard reference fields without `kind` fields in their OpenAPI schema.
   - We utilized the package `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` and its `ProjectRef` for the `projectRef` field, which strictly lacks a `kind` field. This ensured absolute compatibility with the baseline schema.
   - We reused local `ForwardingRuleRef`, `InstanceRef`, `ComputeNetworkRef`, and `ComputeSubnetworkRef` definitions within the same `v1beta1` package to match other references perfectly.

3. **Integer Type Conversion**:
   - The `priority` field is an `*int64` in KRM, but represents a `*uint32` in the underlying Compute Proto message. Handcoded mapper helpers efficiently managed this pointer-to-pointer integer conversion.
