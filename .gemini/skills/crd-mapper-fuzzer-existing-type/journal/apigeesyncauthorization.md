# ApigeeSyncAuthorization KRM Types Journal

## Observations & Learnings

1. **Brownfield Resource with Existing `v1alpha1` CRD**:
   - `ApigeeSyncAuthorization` is a Terraform-managed brownfield resource (`tf2crd: "true"`).
   - Its baseline CRD is at `v1alpha1` with no `v1beta1` version.

2. **Apigee Proto Service & Client Library**:
   - Apigee APIs do not have standard `cloud.google.com/go/apigee/apiv1/apigeepb` protobuf Go packages in Google Cloud Go SDK; instead, MockGCP defines `mockgcp.cloud.apigee.v1` in `mockgcp/apis/mockgcp/cloud/apigee/v1/service.proto`, and direct controllers use `google.golang.org/api/apigee/v1`.
   - As a result, `generate-mapper` cannot be used for `apigee`, and `apis/apigee/generate.sh` only runs `generate-types` targeting `mockgcp.cloud.apigee.v1`.

3. **Multi-Resource `generate-types` in `generate.sh`**:
   - In `apis/apigee/generate.sh`, both `ApigeeAPIProduct` and `ApigeeSyncAuthorization` must be passed as multiple `--resource` arguments to `generate-types` in `v1alpha1` so that `types.generated.go` correctly retains all resources without overwriting.

4. **Skipping Proto Struct Emission in `types.generated.go`**:
   - Adding `// +kcc:spec:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1SyncAuthorization` on `ApigeeSyncAuthorizationSpec` signals to `controllerbuilder` that `GoogleCloudApigeeV1SyncAuthorization` is mapped by the KRM type, avoiding redundant struct generation in `types.generated.go`.

5. **Strict Schema Compatibility**:
   - The baseline CRD only defines `identities` (required array of strings) and `resourceID` (optional string) in `spec`.
   - The status only defines `conditions`, `etag` (optional string), and `observedGeneration` (optional int64).
   - Direct controller conventions like `projectRef`, `organizationRef`, `externalRef`, and `observedState` were omitted to maintain 100% schema compatibility.
   - Top-level doc comments were omitted from `ApigeeSyncAuthorizationSpec` and `ApigeeSyncAuthorizationStatus` because the original CRD did not have descriptions on `spec` and `status`.
   - `dev/tasks/diff-crds --base master` verified zero diff against the baseline CRD.
