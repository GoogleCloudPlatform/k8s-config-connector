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
  5. Added `generate-mapper` with `--prepopulate-resource` in `apis/dlp/generate.sh`, generated `pkg/controller/direct/dlp/mapper.generated.go`, added `cloud.google.com/go/dlp v1.38.0` to `go.mod`, and defined `Empty` and `timestamppb.Timestamp` mapping helpers in `pkg/controller/direct/dlp/mappings.go`.
- **Impact**: Enables clean mechanical scaffolding of `DLPContentPolicy` and downstream controller/fuzzer generation without regressions on existing DLP resources.

### [2026-11-11] DLPContentPolicy Judgement Queue Triage and Type Finalization
- **Context**: Triaging `apis/dlp/judgement_queue.yaml` and finalizing direct KRM types and identity for `DLPContentPolicy`.
- **Findings & Actions**:
  1. Parent & Location (`location-or-parent-ref`): Inlined `*parent.ProjectAndLocationRef` in `DLPContentPolicySpec`, replacing separate `ProjectRef` and `Location` fields while preserving CRD field names and enforcing immutability on `location`.
  2. Identity (`identity-server-generated-id`): Verified `CreateContentPolicyRequest` defines `content_policy_id` parameter which supports client-specified IDs, conforming to `identity.IdentityV2`.
  3. BigQuery destination references (`possible-reference-by-name`): Triaged `.spec.loggingConfigs[].logToBigQuery.{projectID,datasetID,tableID}` as flat string fields matching the proto definition and direct mapper generation.
  4. Well-known proto type (`well-known-proto-type`): Deferred `google.protobuf.Empty` (`anyInfoType`) per standard team-policy item guidelines.
  5. Nested output-only comment (`possible-output-only-pattern-in-comment`): Kept `.spec.inspectConfig.customInfoTypes[].storedType.createTime` in Spec as an optional field to match proto message and direct mapper.
- **Impact**: Fully resolves the judgement queue for `DLPContentPolicy`, preparing the resource for controller and test fixture implementation.

### [2026-11-11] DLPContentPolicy Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing direct controller, round-trip fuzzer, and recording authentic E2E fixtures for `DLPContentPolicy`.
- **Findings & Actions**:
  1. Generated direct controller (`dlpcontentpolicy_controller.go`) and fuzzer (`dlpcontentpolicy_fuzzer.generated.go`) using `controllerbuilder` one-shot generation flags with `--prepopulate-resource`.
  2. Registered `common.NormalizeGreenfieldReferences` in `AdapterForObject`.
  3. Triaged unimplemented custom info type subfields (`file_label_info_type` and `metadata_key_value_expression`) in the fuzzer via `f.Unimplemented_NotYetTriaged`.
  4. Discovered empirically that real GCP DLP API requires `inspect_config` during creation (`inspect_config is required.` 400 error if omitted), and reflected this requirement in minimal and maximal fixtures.
  5. Recorded authentic `_http.log` and `_generated_object_*.golden.yaml` on real GCP (`cnrm-barni-4`) for `dlpcontentpolicy-minimal` and `dlpcontentpolicy-maximal`, passing all 7 HTTP lifecycle report validation checks.
- **Impact**: Completes Step 2 for `DLPContentPolicy` with verified round-trip fuzzing and authentic E2E fixtures.


