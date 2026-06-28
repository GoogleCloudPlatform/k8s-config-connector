# Journal - ContainerNodePool Direct KRM Types Implementation

## Lessons Learned & Key Findings

1. **Top-Level Structs & Scaffolding**: 
   When running `generate-types` for a new resource like `ContainerNodePool`, the generator only outputs a skeleton struct `ContainerNodePoolSpec` and `ContainerNodePoolStatus` in `<kind>_types.go`. It comments out the main proto message mapping (e.g. `NodePool`) in `types.generated.go` as unreachable. We must copy and inline those commented-out fields directly into `<kind>_types.go`.

2. **Strict Schema Compatibility & Custom Types**:
   To maintain 100% schema alignment with the baseline CRD, we had to define several local custom structs (e.g. `NodePoolUpgradeSettings`, `NodePoolBlueGreenSettings`, `NodePool_PlacementPolicy`, `NodePoolNodeConfig` and `WindowsNodeConfig`). By using custom type names locally and tagging them with `// +kcc:proto=`, we avoid adding/un-commenting unsupported fields to `types.generated.go` and prevent conflict with existing shared structs in `cluster_types.go`.

3. **Int32 vs Int64 Truncation in Fuzzer**:
   When mapping integer fields from proto (such as GKE `max_pods_per_node` which is `int64`), using `*int32` or `*int` in KRM is schema-compatible (Kubernetes OpenAPI `integer` represents both). However, since the fuzzer generates arbitrary large 64-bit integers for fuzzing, using `*int32` will cause truncation during round-trip comparison. Prefer `*int` (which is 64-bit on 64-bit platforms) or `*int64` to prevent round-trip value truncation.

4. **Empty Struct Oneof Round-Tripping**:
   When a proto sub-message lives inside a oneof field (like GKE rollout policies), the Go protobuf generated code serializes empty wrappers (e.g. `standard_rollout_policy: {}`) if the wrapper object is constructed, even if the nested policy message is `nil`. Ensure that you only assign the oneof wrapper if the nested KRM struct pointer is non-nil to avoid false-positive round-trip diffs.

5. **Dummy Status Mappers for Spec-Only Fuzzers**:
   For resources that do not yet have a direct controller for reconciliation or status mapping, registering the fuzzer requires status mapping arguments. Passing dummy Status `FromProto` / `ToProto` mappers that return non-nil, empty structs (e.g., `&krm.ContainerNodePoolStatus{}` / `&pb.NodePool{}`) is an elegant way to allow spec fuzzing to pass successfully.
