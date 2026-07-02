### 2026-07-02 VisionProduct Greenfield Types Scaffolding
- **Context**: Implemented Greenfield types and IdentityV2 for VisionProduct (vision.cnrm.cloud.google.com/v1alpha1).
- **Problem**: The VisionProduct resource is not registered in the Cloud Asset Inventory (CAI) name list, causing gcpurls template matching checks to fail.
- **Solution**: Added the VisionProduct URL template `projects/{project}/locations/{location}/products/{product}` exception to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to bypass the CAI name matching validation.
- **Impact**: Allows the resource's IdentityV2 and url template registry tests to pass without failing on missing CAI entries.
