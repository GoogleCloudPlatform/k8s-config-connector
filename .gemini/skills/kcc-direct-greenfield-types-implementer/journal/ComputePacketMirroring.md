# ComputePacketMirroring KRM Types Implementation Journal

## Resource Details
- **GroupKind**: `ComputePacketMirroring` (`compute.cnrm.cloud.google.com`)
- **API Version**: `v1alpha1`
- **Proto Mapping**: `google.cloud.compute.v1.PacketMirroring`

## Key Findings & Decisions

### 1. References Mapping
The original DCL-based `v1beta1` resource used a custom-constructed schema for external and local resource references. In our direct `v1alpha1` KRM types implementation, we utilized canonical reference structures:
- `collectorIlb.urlRef` -> `computev1beta1.ForwardingRuleRef` (mapping to the proto field `collector_ilb.url`)
- `network.urlRef` -> `computev1beta1.ComputeNetworkRef` (mapping to the proto field `network.url`)
- `mirroredResources.instances.urlRef` -> `computev1beta1.InstanceRef` (mapping to the proto field `mirrored_resource_info.instances.url`)
- `mirroredResources.subnetworks.urlRef` -> `computev1beta1.ComputeSubnetworkRef` (mapping to the proto field `mirrored_resource_info.subnetworks.url`)

### 2. Output-Only / Observed Fields Mapping
Output-only (server-generated/readonly) fields of `PacketMirroring` such as `creation_timestamp`, `id`, `kind`, `region`, and `self_link` were successfully placed in `ComputePacketMirroringObservedState` mapped to their corresponding proto definitions.

Additionally, nested output-only fields like `canonicalURL` in collector, network, instances, and subnetworks were cleanly represented in their nested `ObservedState` sub-structs (e.g., `PacketMirroringCollectorIlbObservedState` for the `canonicalUrl` field of GCE forwarding rules).

### 3. Stability Level
The resource was successfully registered as `stability-level=alpha` for `v1alpha1`, ensuring appropriate graduation boundaries.
