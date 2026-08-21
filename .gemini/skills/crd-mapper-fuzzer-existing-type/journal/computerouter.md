# ComputeRouter KRM Transition Journal

## Overview
Transitioned `ComputeRouter` (KRM API version `compute.cnrm.cloud.google.com/v1beta1`) to direct KRM types and automatic mapper-generator configuration.

## Key Learnings

1. **Strict Schema Compatibility vs. Standard Greenfield Output**:
   - Initial generator run produced standard Greenfield fields (e.g. `projectRef`, `location`), whereas the baseline CRD of `ComputeRouter` used legacy patterns (`region` as a string, `networkRef` as a direct referent).
   - Hand-coded types in `router_types.go` were required to maintain 100% strict schema compatibility.
   - We preserved `NetworkRef` using `ComputeNetworkRef` directly, and restored original metadata labels (`cnrm.cloud.google.com/tf2crd=true` and `stability-level=stable`) and custom printer columns.

2. **Leveraging Renames for Automatic Mapper Generation**:
   - The Go protobuf representation of `Router` capitalized the acronym `BGP` to `BGP` (e.g. `RouterBGP` and `RouterAdvertisedIPRange`), whereas the original KRM types used lowercase `Bgp` (e.g. `RouterBgp`).
   - By renaming the Go struct fields and types in `router_types.go` (e.g., from `RouterBgp` to `RouterBGP`) while preserving the JSON tags (e.g. `json:"bgp,omitempty"` and `json:"advertisedIpRanges,omitempty"`), we aligned the field names with the Go protobuf compiler's uppercase acronym naming rules.
   - Adding `// +kcc:proto=...` annotations to these nested types allowed the controllerbuilder to automatically discover and map them without any manual intervention.

3. **Handling Type Mismatches in Mapper**:
   - Handcoded mapper functions were placed in `router_mapper.go` to handle type mismatches (e.g., `asn` and `keepaliveInterval` being `int64` in KRM but `*uint32` in proto, and `Range` being a pointer in proto but a non-pointer in KRM). This automatically overrode the generator's default implementations.

4. **KRM Round-Trip Fuzzing Integration**:
   - Configured `FilterSpec` and `FilterStatus` correctly to clear mismatched non-roundtripping fields (e.g. clearing spec fields like `Bgp`, `Network`, `Region` from `FilterStatus`, and status fields like `CreationTimestamp` and `SelfLink` from `FilterSpec`).
   - Marked untriaged/internal proto fields like `.params` and `.bgp_peers` as unimplemented.
