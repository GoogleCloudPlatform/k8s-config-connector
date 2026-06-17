### [2026-05-26] DataprocSession Scaffolding Unreachable Types
- **Context**: Scaffolding direct types for DataprocSession (KCC issue 8682).
- **Problem**: Inner types from the proto like `JupyterConfig`, `SparkConnectConfig`, and `SessionStateHistory` were pruned by `generate-crds` because the initial generated `DataprocSessionSpec` did not embed them. Also, the Greenfield reviewer requires scalar primitives to be pointers, but `generate-types` generated `Location string`.
- **Solution**: Manually copied the definitions for the unreachable types from `types.generated.go` into `session_types.go` and updated `DataprocSessionSpec` and `DataprocSessionObservedState` to reference them. Changed `Location string` to `Location *string` to satisfy the review invariant.
- **Impact**: Agents must check for pruned/unreachable types in `types.generated.go` and manually include them in the `_types.go` file if they correspond to proto fields that are intended to be exposed. Additionally, verify primitive types like `Location` are pointers.

### [2026-06-05] DataprocSessionTemplate Scaffolding and SparkConnectConfig Resolution
- **Context**: Implementing direct types for `DataprocSessionTemplate` under `apis/dataproc/v1alpha1`.
- **Problem**: The GCP API proto defines `SparkConnectConfig` as an empty message `message SparkConnectConfig {}`. The generator identifies this and, because of an annotation in `session_types.go`, believes the type is non-generated but it was actually never defined. Using `SparkConnectConfig` in the new `DataprocSessionTemplateSpec` would lead to compilation errors.
- **Solution**: Explicitly defined `SparkConnectConfig` as `type SparkConnectConfig struct {}` with the `// +kcc:proto=google.cloud.dataproc.v1.SparkConnectConfig` annotation inside `sessiontemplate_types.go`. This satisfies the generator while providing the correct type definition for compilation.
- **Impact**: When implementing Dataproc resources that utilize `SparkConnectConfig`, the struct is now fully defined and compiled.

