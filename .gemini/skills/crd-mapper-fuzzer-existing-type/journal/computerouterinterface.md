# ComputeRouterInterface Journal Entry

## Key Observations & Learnings

- **Pre-existing CRD Alignment:**
  - The baseline CRD did not contain `spec.projectRef` or `spec.location`, and `spec.region` was used instead. To achieve 100% strict schema-compatibility, we avoided adding `ProjectRef` or `Location` in our `ComputeRouterInterfaceSpec` definition.
  - The generated status must not include fields like `observedState` or `externalRef` if they weren't in the original baseline CRD. 
  - `dev/tasks/diff-crds` returned a completely clean/empty output, confirming 100% structural equivalence.

- **Handling Custom Reference Types:**
  - `ComputeRouter`, `ComputeAddress`, and `ComputeSubnetwork` have canonical reference structures under `apis/compute/v1beta1/`.
  - `ComputeInterconnectAttachmentRef` and `ComputeVPNTunnelRef` did not exist as canonical reference types, so we defined them as custom local structs in `routerinterface_types.go` without a `kind` field.
  - Custom references matching `"external,name,namespace"` signature are correctly decorated with OpenAPI constraints (oneOf block) automatically.

- **Hand-coded Mappers:**
  - Fields in KRM (`interconnectAttachmentRef`, `vpnTunnelRef`) didn't match the proto field names (`linked_interconnect_attachment`, `linked_vpn_tunnel`) 1:1, so they were marked as `MISSING` by the automated generator.
  - We cleanly resolved this by handcoding `ComputeRouterInterfaceSpec_v1beta1_FromProto` and `ComputeRouterInterfaceSpec_v1beta1_ToProto` in `pkg/controller/direct/compute/computerouterinterface_mapper.go`.
  - The generator automatically detected these existing mapping functions and commented out its own conflicting generation versions.
