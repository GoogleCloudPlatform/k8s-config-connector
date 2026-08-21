# ComputeRoute Direct Types Migration Journal

## Overview
Successfully implemented the direct KRM types and configured `generate.sh` for `ComputeRoute` (`apis/compute/v1beta1/route_types.go`).

## Findings & Learnings
- **Proto Mapping**: The underlying proto for `ComputeRoute` is `Route`.
- **Reference Types**:
  - `networkRef` mapped to `ComputeNetworkRef` in package `v1beta1`.
  - `nextHopILBRef` mapped to `ForwardingRuleRef` in package `v1beta1`.
  - `nextHopInstanceRef` mapped to `InstanceRef` in package `v1beta1`.
  - Since `ComputeVPNTunnel` does not have an existing canonical reference type in the repository, we hand-coded `ComputeVPNTunnelRef` inside `route_types.go` locally. It conforms perfectly to the standard signature `{external, name, namespace}` expected by KCC and matches the baseline CRD schema exactly.
- **Go Struct Field Formatting (camelCase/Proto matching)**:
  - We utilized field name capitalization to align with proto fields while keeping JSON tags compliant with the baseline CRD schema. For instance, `NextHopIP` maps directly to proto `next_hop_ip` (Go name `NextHopIP`) while the JSON tag `json:"nextHopIp,omitempty"` preserves the exact KRM schema case `nextHopIp`.
  - Similarly, using `NextHopILBRef` with JSON tag `json:"nextHopILBRef,omitempty"` matches both proto (`NextHopIlb` / `next_hop_ilb`) and KRM schema.
- **Schema Compatibility**:
  - Validated using `dev/tasks/diff-crds`. The generated CRD matched the baseline perfectly with only standard description comments, listKind, and integer formatting (e.g. `format: int32` for `priority` and `format: int64` for `observedGeneration`) added, which are fully compatible.
