### [2026-06-05] DLPDiscoveryConfig Proto Compilation and Scaffolding
- **Context**: Implementing initial direct KRM types, CRD, and IdentityV2 for `DLPDiscoveryConfig` resource under `apis/dlp/v1alpha1`.
- **Problem**: The proto messages for DLP reside under `google.privacy.dlp.v2` package (in `google/privacy/dlp/v2/dlp.proto` and `storage.proto`), which is not located under `google/cloud/` and was missing from the standard `dev/tools/controllerbuilder/generate-proto.sh` compilation script. This resulted in `generate-types` failing with `failed to find the proto message google.cloud.dlp.v1.DiscoveryConfig`.
- **Solution**:
  1. Updated `dev/tools/controllerbuilder/generate-proto.sh` to compile `google/privacy/dlp/v2/*.proto` files:
     ```bash
     ${THIRD_PARTY}/googleapis/google/privacy/dlp/v2/*.proto \
     ```
  2. Configured the `generate.sh` script with `--service google.privacy.dlp.v2` to correctly point the generator to the proper service definition.
- **Impact**: Enables future agents to successfully compile and scaffold direct resources under the Google Privacy DLP API group.

### [2026-11-10] DLPContentPolicy Proto Overlay and Prepopulation
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `DLPContentPolicy` under `apis/dlp/v1alpha1`.
- **Problem**: The pinned googleapis commit in `apis/git.versions` does not contain `ContentPolicy` which was added in upstream googleapis master. Furthermore, completely replacing `dlp.proto` with upstream master modified existing types and comments for `DiscoveryConfig`, causing missing reference violations in pre-existing CRDs. Additionally, `ContentPolicy` includes a parameterless `any_info_type` proto oneof (`google.protobuf.Empty`), producing an empty object in OpenAPI schema.
- **Solution**:
  1. Created a proto overlay `mockgcp/apis/google/privacy/dlp/v2/dlp.proto` that preserves the pinned base for existing DLP resources while adding `ContentPolicy`, its RPCs, and its request/response/verdict messages.
  2. Registered the overlay in `dev/tools/controllerbuilder/generate-proto.sh` and in `kccProtoOverrides` in `dev/tools/controllerbuilder/pkg/sourcelinks/sourcelinks.go`.
  3. Added `DLPContentPolicy` to `dev/tasks/build-release-bundle` blocked kinds list.
  4. Added `dlpcontentpolicies.dlp.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` due to the empty oneof option (`spec.rules[].conditions[].infoTypeCondition.anyInfoType`).
- **Impact**: Enables clean mechanical scaffolding of `DLPContentPolicy` without regressions on existing DLP resources.
