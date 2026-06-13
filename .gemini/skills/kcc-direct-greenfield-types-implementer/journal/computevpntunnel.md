# ComputeVPNTunnel Greenfield Types Implementation Journal

## Resource Overview
- **KCC Kind:** `ComputeVPNTunnel`
- **GCP Proto Resource:** `google.cloud.compute.v1.VpnTunnel`
- **API Version:** `v1beta1`

## Learnings & Implementation Details

### SharedSecret Structural Difference
In the GCP Proto API, `shared_secret` is represented as a plain string field:
```go
SharedSecret *string
```
However, in KCC, we support referencing Kubernetes secrets for sensitive fields (such as VPN shared secrets) to maintain security hygiene. Consequently, the KCC type is structured as:
```go
type VPNTunnelSharedSecret struct {
	Value     *string
	ValueFrom *VPNTunnelValueFrom
}
```
This mismatch between the Protobuf API type (`*string`) and the KRM spec type (`VPNTunnelSharedSecret`) causes type errors in the automatically generated mappers. 

### Resolution via Custom Mapper
To resolve the type discrepancy, we implemented a custom, hand-written mapper under `pkg/controller/direct/compute/vpntunnel_mapper.go`:
- Implemented `ComputeVPNTunnelSpec_v1beta1_FromProto` / `ComputeVPNTunnelSpec_v1beta1_ToProto`
- Implemented `ComputeVPNTunnelStatus_v1beta1_FromProto` / `ComputeVPNTunnelStatus_v1beta1_ToProto`

When running the `./generate.sh` generator script, KCC's build tooling automatically detects existing non-generated mapping functions and skips/comments out their automatically generated counterparts inside `mapper.generated.go`.

### Reference Structures Added
`ComputeVPNTunnel` references several other compute resources that did not yet have defined reference types under `apis/refs/v1beta1`. We added standard, backward-compatible reference definitions to `apis/refs/v1beta1/computerefs.go` for:
- `ComputeRouterRef`
- `ComputeVPNGatewayRef` (shared by peer and target gateway mappings)
- `ComputeExternalVPNGatewayRef`
