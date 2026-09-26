# DiscoveryEngineACLConfig Greenfield Types Learnings

- `DiscoveryEngineACLConfig` maps to GCP proto `google.cloud.discoveryengine.v1beta.AclConfig`.
- The resource is located at `projects/{project}/locations/{location}/aclConfig` (singleton per project/location).
- The identity uses `gcpurls.Template[DiscoveryEngineACLConfigIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/aclConfig")`.
- URL template format was added as an exception in `pkg/gcpurls/registry_test.go` because `AclConfig` is not listed in `cloudassetinventory_names.jsonl`.
