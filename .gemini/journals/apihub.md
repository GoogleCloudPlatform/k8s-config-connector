# ApiHub Service Journal

### 2026-05-24 Initial Scaffolding and Identity for ApiHubApi
- **Context**: Greenfield implementation of ApiHubApi types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: Scarce local build cache for `googleapis.pb` caused `generate-types` to fail.
- **Solution**: Executed `./dev/tools/controllerbuilder/generate-proto.sh` to compile googleapis proto definitions into `.build/googleapis.pb` using the pinned Google API SHA in `apis/git.versions`.
- **Impact**: Provides clean scaffolding, deepcopy, CRD, and client generation workflow. Also maps system defined attribute references to KRM `APIHubAttributeValueRef`.
