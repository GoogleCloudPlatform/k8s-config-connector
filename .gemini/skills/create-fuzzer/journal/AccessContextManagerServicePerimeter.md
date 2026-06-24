# Fuzzer Journal: AccessContextManagerServicePerimeter

## Observations and Learnings

### 1. Reachability & Type Pruning
- By default, the type generator/pruner (`prunetypes.go`) comments out types in `types.generated.go` (like `ServicePerimeter` and sub-structs) if they are not referenced inside the handwritten API types (e.g. `serviceperimeter_types.go`).
- To make them reachable:
  - Annotate the root spec structure with `// +kcc:spec:proto=google.identity.accesscontextmanager.v1.ServicePerimeter`.
  - For any nested structures (like `ServicePerimeterConfig`), annotate them with `// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig` (and so on) so they match with the protobuf message definitions and their mappers are automatically generated.

### 2. Handling Protobuf Oneof Fields
- Protobuf `oneof` fields (like `Source` in `ServicePerimeterConfig_IngressSource`) are represented differently in Go (as wrapper structs) and cannot be directly assigned or generated via simple field mapping.
- To resolve this, define custom mapping functions in `mappers.go` (e.g., `AccessContextManagerServicePerimeterIngressSource_FromProto` and `AccessContextManagerServicePerimeterIngressSource_ToProto`). The mapper generator automatically skips generating functions that are already handwritten in the package, allowing seamless overrides.

### 3. Fuzzer Implementation, Status vs. ObservedState
- Avoid adding new fields (like `ObservedState`) to the resource's K8s Status schema if the schema does not already support it. Adding schema fields in a fuzzer PR is a breaking schema change that we want to avoid.
- Instead, annotate the existing Status struct with `// +kcc:status:proto=...`, map status fields (like `createTime`, `updateTime`) directly using `Status_FromProto` and `Status_ToProto` mapper functions, and pass these to `NewKRMTypedFuzzer`.
- Register spec fields with `f.SpecField(".field")`, status fields with `f.StatusField(".field")`, and identity fields with `f.IdentityField(".field")`. Avoid using direct set insertion on internal fuzzer fields.
