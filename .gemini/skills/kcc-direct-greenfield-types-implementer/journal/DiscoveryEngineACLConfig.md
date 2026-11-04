# DiscoveryEngineACLConfig Greenfield Types Implementation Notes

- **Resource Kind**: `DiscoveryEngineACLConfig`
- **GCP Proto**: `google.cloud.discoveryengine.v1alpha.AclConfig`
- **Identity Implementation**: Generated via `generate-identity` command in `apis/discoveryengine/generate.sh`.
- **Observations**:
  - The proto package is `v1alpha`, requiring `--allow-non-ga-proto` in `generate-types`.
  - Added `DiscoveryEngineACLConfig` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Recorded untriaged alpha missing fields exception in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
