### 2026-06-13 Direct Types for ComputeSharedVPCServiceProject
- **Context**: Implementing direct types and generate.sh configuration for ComputeSharedVPCServiceProject (Issue #9885).
- **Problem**: 
  1. The resource does not correspond to a standard CRUD resource with a standard message; instead, it matches to `google.cloud.compute.v1.ProjectsEnableXpnResourceRequest`.
  2. The generator created `projectsenablexpnresourcerequest_types.go` because of the mapped message name. Renaming it to `computesharedvpcserviceproject_types.go` caused the generator to recreate the file on subsequent runs.
  3. Under the hood, using `*refsv1beta1.ProjectRef` led to a controller-gen panic during deepcopy generation: `panic: interface conversion: types.Type is nil, not *types.Named`.
- **Solution**:
  1. Configured `--resource ComputeSharedVPCServiceProject:ProjectsEnableXpnResourceRequest` in `apis/compute/v1beta1/generate.sh`.
  2. Maintained the filename as `projectsenablexpnresourcerequest_types.go` to satisfy the generator's expectations.
  3. Replaced `*refsv1beta1.ProjectRef` with `refs.ProjectRef` (non-pointer) imported from `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"`. This resolved the controller-gen deepcopy generation panic and generated the perfect backward-compatible CRD schema.
- **Impact**: Ensures that when other xpn-related or association-like resources in Compute are migrated, they can be configured using their request protos with appropriate filename mapping and use the standard `refs.ProjectRef` reference type to avoid controller-gen issues.
