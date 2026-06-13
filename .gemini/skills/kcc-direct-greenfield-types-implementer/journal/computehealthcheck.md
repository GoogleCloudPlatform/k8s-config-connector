### 2026-06-13 ComputeHealthCheck Direct Types
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeHealthCheck`.
- **Problem**: `ComputeHealthCheck` is a brownfield resource currently managed by the Terraform-based controller and already present in `v1beta1`. The existing `healthcheck_types.go` file contains handcrafted Go types with `// +kcc:spec:proto` and skipping configurations, but was missing standard direct-controller validation metadata such as enum restrictions (`// +kubebuilder:validation:Enum`), exact status field proto tag mapping (`// +kcc:status:proto`), and `// +kcc:proto:field` annotations on the status struct fields.
- **Solution**:
  - Maintained the existing `apis/compute/v1beta1/generate.sh` configuration without adding duplicate `ComputeHealthCheck` definitions to `v1alpha1`.
  - Added `// +kcc:status:proto=google.cloud.compute.v1.HealthCheck` annotation to `ComputeHealthCheckStatus`.
  - Added standard proto field annotations (`// +kcc:proto:field`) to fields inside `ComputeHealthCheckStatus`.
  - Added `// +kubebuilder:validation:Enum` tags for enum fields like `PortSpecification` and `ProxyHeader` across all nested protocol-specific health check structs.
  - Ran `generate.sh` and `dev/tasks/generate-crds` to compile, regenerate, and validate the updated CRD and mappers.
- **Impact**: Enables `ComputeHealthCheck` to have complete, direct-controller compliant types while preserving backwards compatibility and schema-compatibility for existing stable `v1beta1` users.
