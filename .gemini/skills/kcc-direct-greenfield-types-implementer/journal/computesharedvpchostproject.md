# ComputeSharedVPCHostProject Types Implementation Journal

## Observations & Learnings

### 1. File-Naming Logic and Duplicate-Type compiler Panic
- **Behavior**: The generator builds the Go types file path as `<lowercase-proto-message-name>_types.go`. Since `ComputeSharedVPCHostProject` maps to the `Project` proto message in `google.cloud.compute.v1`, the generator creates `project_types.go`.
- **Constraint**: If we rename `project_types.go` to `computesharedvpchostproject_types.go` and re-run `generate.sh`, the generator will recreate `project_types.go`. Having both files in the package containing identical struct definitions causes `controller-gen`'s type loader to fail or panic with `interface conversion: types.Type is nil, not *types.Named`.
- **Guidance**: Keep the file name exactly as generated (e.g., `project_types.go` for the `Project` proto) to prevent duplicate definitions and subsequent deepcopy/CRD generation failures.

### 2. Adjusting Generated Fields for Global Resources
- **Location Field**: For project-scoped/global resources like `ComputeSharedVPCHostProject` that do not reside in any region or zone, the heuristic-generated `Location string` field must be removed to avoid requiring location specifications in KRM resources.
- **ProjectRef Field**: The generated `ProjectRef` struct should be made optional (`// +optional` and `omitempty`) rather than required, preserving backward compatibility for users using namespace-based scoping or the `cnrm.cloud.google.com/project-id` annotation without `spec.projectRef`.
- **Labels**: Ensure CRD metadata labels (`tf2crd: "true"` and `stability-level: stable`) are manually restored or specified in the root type's comments to match the legacy CRD exactly.
