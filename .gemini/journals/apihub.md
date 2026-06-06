# ApiHub Service Journal

### 2026-05-24 Initial Scaffolding and Identity for APIHubAPI
- **Context**: Greenfield implementation of APIHubAPI types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: Scarce local build cache for `googleapis.pb` caused `generate-types` to fail.
- **Solution**: Executed `./dev/tools/controllerbuilder/generate-proto.sh` to compile googleapis proto definitions into `.build/googleapis.pb` using the pinned Google API SHA in `apis/git.versions`.
- **Impact**: Provides clean scaffolding, deepcopy, CRD, and client generation workflow. Also maps system defined attribute references to KRM `APIHubAttributeValueRef`.

### 2026-06-05 Scaffolding and Identity for ApiHubRuntimeProjectAttachment
- **Context**: Greenfield implementation of ApiHubRuntimeProjectAttachment types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: `gcpurls.Template` failed with a runtime panic during initialization because the snake_case template variable `{runtime_project_attachment}` did not match standard Go field casing mapping when defined as camelCase `RuntimeProjectAttachment string`.
- **Solution**: Defined the Identity Go struct field name with underscores: `Runtime_project_attachment string`. This correctly matches the `gcpurls.Template` parser's expectations.
- **Impact**: Other direct resource types with snake_case variables in their GCP URL template formats must follow this pattern to avoid initialization panics.
