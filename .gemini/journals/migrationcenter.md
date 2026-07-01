### [2026-06-29] MigrationCenterGroup Types, CRD, and IdentityV2 Scaffolding
- **Context**: Implementing initial types, CRD, and IdentityV2 for `MigrationCenterGroup` under `apis/migrationcenter/v1alpha1/`.
- **Problem**: The Migration Center API/resource is missing from `cloudassetinventory_names.jsonl` (not supported by CAIS yet), which causes CAIS URL template matches validation `TestRegisteredTemplatesMatchCAI` in `pkg/gcpurls/registry_test.go` to fail on missing registered templates.
- **Solution**: Added the URL template `//migrationcenter.googleapis.com/projects/{}/locations/{}/groups/{}` as an exception under `ignoredTemplates` in `pkg/gcpurls/registry_test.go`.
- **Impact**: Ensures that `gcpurls` verification tests compile and pass smoothly without requiring support in CAIS metadata yet.
