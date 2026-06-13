### [2026-06-13] ComputeRouterPeer Direct Types Implementation
- **Context**: Implementing direct types and `generate.sh` configuration for `ComputeRouterPeer` (GCP `RouterBgpPeer`).
- **Problem**: 
  1. The existing CRD has lowercase-acronym fields like `advertisedIpRanges`, while the GCP protobuf definition uses `advertisedIPRanges` (uppercase IP). The auto-generated structs from `types.generated.go` use the proto casing, which would break schema compatibility if used directly in `routerbgppeer_types.go`.
  2. The resource references `ComputeRouter` and `ComputeRouterInterface` are still TF-based and do not have existing direct identity or reference structures.
- **Solution**: 
  1. Hand-crafted custom structs (`RouterpeerAdvertisedIpRanges`, `RouterpeerBfd`, and `RouterpeerIpAddress`) were defined in `routerbgppeer_types.go` to precisely preserve the original KCC CRD schema, ensuring 100% backward compatibility.
  2. Minimal `ComputeRouterRef` and `ComputeRouterInterfaceRef` structs with `NormalizedExternal` methods were implemented under `apis/compute/v1beta1/` to gracefully resolve referenced legacy/TF resources.
- **Impact**: Demonstrates how to balance the transition to direct types while preserving strict schema/backwards compatibility for stable resources containing lowercase-acronym JSON tags.
