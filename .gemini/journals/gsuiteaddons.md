# GSuiteAddons Service Journal

### [2026-06-15] GSuiteAddonsDeployment Greenfield Types Scaffolding
- **Context**: Implementing Greenfield direct types, CRD, and IdentityV2 for `GSuiteAddonsDeployment` (mapped to `google.cloud.gsuiteaddons.v1.Deployment` / `v1alpha1`).
- **Problem**: 
  1. The GCP URL template is `projects/{project}/deployments/{deployment}`, which is a global (project-level) resource with no `{location}` segment. However, the generator defaults to injecting a `Location` field into `GSuiteAddonsDeploymentSpec`.
  2. GSuiteAddons uses `google.protobuf.Value` and `google.protobuf.ListValue` representing dynamic/schemaless JSON configurations. Because these types are recursive, controller-gen prunes them entirely in the CRD and leaves empty types (`type: ` omitted). This caused the KCC Go clients generator to panic with `unhandled type: ` inside `pkg/crd/fielddesc/fielddesc.go`.
- **Solution**:
  1. Removed the `Location` field from `GSuiteAddonsDeploymentSpec` and implemented IdentityV2 correctly without locations.
  2. Added `// +kubebuilder:validation:XPreserveUnknownFields` annotations on all the recursive fields (`Values`, `StructValue`, `ListValue`, and `OpenLinkURLPrefixes`) in `types.generated.go`.
  3. Patched the client generator in `pkg/crd/fielddesc/fielddesc.go` to gracefully fall back to `schemalessToDescription` if `props.Type == ""` (representing untyped/schemaless OpenAPI properties), rather than panicking.
- **Impact**: Unblocks GSuiteAddonsDeployment client generation and establishes a robust pattern for any KCC resources using recursive/dynamic protobuf types like `google.protobuf.Value`.
