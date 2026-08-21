# ComputeVPNGateway Direct Types Journal

- Kind: ComputeVPNGateway
- Date: 2026-06-13

## Observations
- The baseline CRD for ComputeVPNGateway does not include `spec.projectRef` or `spec.location` and uses `region` instead.
- We must define `Region string json:"region"` without `ProjectRef` or `Location` in the KRM spec structure to preserve strict schema-compatibility.
- Renamed `VpnInterfaces` in Go struct to `VPNInterfaces` matching proto's `VPNInterfaces` (`vpn_interfaces` snake_case) to allow automatic mapper generator to pick up and generate slice map function cleanly.
- `VpnGatewayVpnGatewayInterface` fields required handcoded mappings due to `Id` field being `*uint32` in proto but `*int64` (type: integer) in CRD, and `interconnect_attachment` being `*string` (represented as `interconnectAttachmentRef` reference in CRD spec).
- Fuzzer required using `NewKRMTypedSpecFuzzer` and registering with `RegisterKRMSpecFuzzer` since status maps only empty/standard metadata.
- Fuzzer field path matching syntax uses empty square brackets `[]` instead of `[*]` to represent wildcard indices of arrays/slices (e.g. `.vpn_interfaces[].ipv6_address`).
