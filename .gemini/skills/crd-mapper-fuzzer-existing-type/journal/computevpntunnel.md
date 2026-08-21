# ComputeVPNTunnel direct KRM types migration

This journal entry captures the learnings and implementation details during the transition of `ComputeVPNTunnel` to direct KRM types.

## Learnings & Observations

### 1. Rebuilding Validation Rules
- **Problem:** After generating types and CRDs, the `oneOf` legacy constraint on `spec.sharedSecret` (requiring exactly one of `value` or `valueFrom`) was lost.
- **Root Cause:** The `oneOf` schema validator is handled by `scripts/add-validation-to-crds/parse-crds.go` when the field signature is `"value,valueFrom"`. However, this rule was hard-coded for a specific list of kinds (`AlloyDBUser`, `ComputeInstance`, `ContainerCluster`, and `MonitoringUptimeCheckConfig`).
- **Solution:** Added `"ComputeVPNTunnel"` to the list of kinds inside `parse-crds.go` to restore the legacy validation block perfectly.

### 2. Lossless Proto Mapping of Pointer Fields (Region and Refs)
- **Problem:** In protobuf, regional string and reference URL fields are pointers (e.g. `Region *string`, `VpnGateway *string`). In KRM, we map them as required or optional non-pointer `string` fields or custom reference structs (e.g., `Region string`, `VpnGatewayRef *VpntunnelVpnGatewayRef`). If these fields are unset (`nil`) in the original proto, KRM maps them to `""` or `nil`, but during `ToProto`, they get mapped to pointers to empty strings `&""`, leading to round-trip diff failures in the fuzzing framework.
- **Solution:** Modified `ToProto` mappings to only assign string pointers to proto if the string is non-empty:
  ```go
  if in.Region != "" {
      out.Region = &in.Region
  }
  ```
  And similarly for references:
  ```go
  if in.VpnGatewayRef != nil && in.VpnGatewayRef.External != "" {
      out.VpnGateway = &in.VpnGatewayRef.External
  }
  ```
  This resolves fuzzer failures losslessly.
