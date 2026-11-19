### [2026-11-17] Prepopulating CloudDeployRelease
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CloudDeployRelease` using `google.cloud.deploy.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `Release` in `google.cloud.deploy.v1` is considered non-declarative by default because the API defines `CreateRelease` without an `UpdateRelease` (create-only execution resource). Passing `--allow-non-declarative` to `generate-types` is required to scaffold it.
  2. Identity pattern is `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/releases/{release}` with ancestor segment `deliveryPipelines/{delivery_pipeline}`.
  3. `Release.target_artifacts` is an output-only map (`map<string, TargetArtifact>`) where `TargetArtifact` contains only output-only fields in proto. Under `--prepopulate-resource`, `TargetArtifact` is generated as an empty struct, requiring `clouddeployreleases.clouddeploy.cnrm.cloud.google.com` in `knownInvalidCRDs` in `tests/apichecks/crds_test.go` until triaged in Step 1b.
  4. Generated `clouddeployrelease_identity.generated.go`, `clouddeployrelease_reference.generated.go`, and conformance tests cleanly.
  5. Added `CloudDeployRelease` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` and synchronized `tests/apichecks/testdata/exceptions/` golden files (`alpha-missingfields.txt`, `no_refs_in_status.txt`, `refs_not_representable.txt`, `spec_dislike_etag.txt`).

### [2026-11-17] Triaging Judgement Queue and Finalizing CloudDeployRelease Types
- **Context**: Triaged `apis/clouddeploy/judgement_queue.yaml` and finalized KRM types and identity for `CloudDeployRelease` following Greenfield Step 1b.
- **Learnings & Observations**:
  1. `CloudDeployRelease` is a child resource under `DeliveryPipeline` with pattern `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/releases/{release}`.
  2. Implemented `refs.Ref` interface methods on existing `apis/clouddeploy/v1beta1/clouddeploydeliverypipeline_reference.go` (`DeliveryPipelineRef`) while preserving `NormalizedExternal` for backward compatibility.
  3. Replaced `projectRef` and `location` in `CloudDeployReleaseSpec` with a single required and immutable `spec.parentRef *v1beta1.DeliveryPipelineRef`.
  4. Removed top-level `labels` from `CloudDeployReleaseSpec` (managed via KRM `metadata.labels`) and moved `etag` to `CloudDeployReleaseObservedState`.
  5. Updated resource documentation URL to `https://cloud.google.com/deploy/docs/reference/rest/v1/projects.locations.deliveryPipelines.releases`.
  6. Resolved `skaffoldConfigURI` as `reference-not-representable` (GCS archive path decomposable into bucket + path, remains string) and accepted `message-map-derived` fields.
  7. Regenerated types, identity, deepcopy, mapper, and CRD via `./apis/clouddeploy/generate.sh`, updated exceptions, and added unit tests in `clouddeployrelease_identity_test.go`.

### [2026-11-17] Direct Controller, Fuzzer, and E2E Fixtures for CloudDeployRelease
- **Context**: Implemented direct controller, fuzzer, minimal/maximal E2E fixtures, and recorded golden logs against real GCP for `CloudDeployRelease` following Greenfield Step 2.
- **Learnings & Observations**:
  1. `Release` has no `UpdateRelease` or `DeleteRelease` RPC in GCP Cloud Deploy API. The resource is immutable, and deletion removes the KRM finalizer without deleting the GCP resource. Configured fixtures with `cnrm.cloud.google.com/deletion-policy: "abandon"`.
  2. Added `CloudDeployRelease` to `IsNonDeletable()` in `pkg/test/httplogreport/fixture/loader.go`.
  3. Deleting `CloudDeployDeliveryPipeline` when a `Release` child resource exists in GCP requires `Force: true` in `DeleteDeliveryPipelineRequest`. Updated `DeliveryPipelineAdapter.Delete` to set `Force: true`.
  4. Real GCP `Release` responses return additional condition fields (`dockerVersionSupportedCondition`, `helmVersionSupportedCondition`, etc.) and `toolVersions` absent from the vendored proto schema in `.build/googleapis.pb`. Added `google.cloud.deploy.v1.Release` to `protovalidate.go`'s `unmarshalDynamicJSON` unknown field allowlist.
  5. `CreateRelease` requires a valid Skaffold archive (`skaffoldConfigUri`) during server-side validation; tests reference `gs://${projectId}-clouddeploy-test-bucket-2026/source.tar.gz`.

### [2026-11-18] Greenfield: MockGCP Implementation and Alignment for CloudDeployRelease
- **Context**: Implemented MockGCP support and golden log alignment for `CloudDeployRelease` following Greenfield Step 3 (`kcc-direct-mockgcp-implementer`).
- **Learnings & Observations**:
  1. Implemented `CreateRelease`, `GetRelease`, `ListReleases`, and `AbandonRelease` in `mockgcp/mockclouddeploy/release.go`.
  2. `CreateRelease` automatically creates `deliveryPipelineSnapshot` and `targetSnapshots` based on the parent `DeliveryPipeline` in mock storage and its serial pipeline stages.
  3. Real GCP asynchronous rendering creates `targetArtifacts` and `targetRenders` with a build ID in `projects/${projectNumber}/locations/${location}/builds/${buildId}`. Handled deterministic generation in `GetRelease` to avoid run-to-run drift.
  4. Normalized volatile nested snapshot fields (`createTime`, `updateTime`, `uid`, `etag`) in `tests/e2e/normalize.go` and `mockgcp/mockclouddeploy/normalize.go`.
  5. Added proto3 unmapped fields (`toolVersions`, `dockerVersionSupportedCondition`, etc.) and snapshot etags to `normalizeRepresentation` in `pkg/test/resourcefixture/golden_alignment_test.go`.
  6. Verified all fixtures pass both hermetic e2e tests (`hack/compare-mock`) and golden alignment tests (`TestGoldenLogAlignment`) with 0 errors.

### [2026-11-18] Prepopulating CloudDeployRollout
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CloudDeployRollout` using `google.cloud.deploy.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `Rollout` in `google.cloud.deploy.v1` is considered non-declarative by default because the API defines `CreateRollout` without `UpdateRollout` (create-only execution resource). Passing `--allow-non-declarative` to `generate-types` is required to scaffold it.
  2. `CloudDeployRollout` is a child resource under `Release` with pattern `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/releases/{release}/rollouts/{rollout}`. Identity is generated with ancestor segments `deliveryPipelines/{delivery_pipeline}` and `releases/{release}`.
  3. `status.observedState.phases[].childRolloutJobs.advanceRolloutJobs[].advanceChildRolloutJob` corresponds to `message AdvanceChildRolloutJob {}` in proto which has no fields. Added `clouddeployrollouts.clouddeploy.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with comment `// GCP known pattern: Empty oneof option (status.observedState.phases[].childRolloutJobs.advanceRolloutJobs[].advanceChildRolloutJob)`.
  4. `google.cloud.deploy.v1.Metadata` is skipped by `generatemappercommand` because its message name ends in `Metadata`. Added explicit `MetadataObservedState_v1alpha1_FromProto`, `MetadataObservedState_v1alpha1_ToProto`, `CloudRunMetadataObservedState_v1alpha1_FromProto`, `CloudRunMetadataObservedState_v1alpha1_ToProto`, `AutomationRolloutMetadataObservedState_v1alpha1_FromProto`, and `AutomationRolloutMetadataObservedState_v1alpha1_ToProto` helper functions in `pkg/controller/direct/clouddeploy/mapper.generated.go`.
  5. Added `CloudDeployRollout` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` and synchronized `tests/apichecks/testdata/exceptions/` golden files (`alpha-missingfields.txt`, `spec_dislike_etag.txt`).

### [2026-11-18] Triage Judgement Queue and Finalize Types for CloudDeployRollout
- **Context**: Triaged `apis/clouddeploy/judgement_queue.yaml` and finalized KRM types, identity, and references for `CloudDeployRollout` following Greenfield Step 1b (`judgement/SKILL.md`).
- **Learnings & Observations**:
  1. `CloudDeployRollout` is a child of `CloudDeployRelease`. Modelled the parent relationship as a single required, immutable `spec.parentRef *CloudDeployReleaseRef` per `PARENT_AND_IDENTITY.md` (Case B). Removed `ProjectRef` and `Location` from `CloudDeployRolloutSpec`.
  2. `controllerbuilder generate-identity` automatically recognizes `spec.parentRef *CloudDeployReleaseRef` and resolves `CloudDeployRolloutIdentity` via parent path suffixing (`strings.TrimSuffix(parentRef.GetExternal(), "/") + "/rollouts/" + resourceID`).
  3. `targetID` on `CloudDeployRolloutSpec` was triaged from `possible-reference-by-name` to `TargetRef *CloudDeployTargetRef` (`// +required`, `json:"targetRef"`).
  4. Moved `etag` from `CloudDeployRolloutSpec` to `CloudDeployRolloutObservedState` and removed `CloudDeployRollout` exception from `tests/apichecks/testdata/exceptions/spec_dislike_etag.txt`.
  5. Updated documentation URL in `clouddeployrollout_types.go` to `https://cloud.google.com/deploy/docs/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts`.
  6. Added unit tests for `CloudDeployRolloutIdentity` and `GetIdentity` in `clouddeployrollout_identity_test.go`.

### [2026-11-19] Greenfield: MockGCP Implementation and Alignment for CloudDeployRollout
- **Context**: Implemented MockGCP support and golden log alignment for `CloudDeployRollout` following Greenfield Step 3 (`kcc-direct-mockgcp-implementer`).
- **Learnings & Observations**:
  1. Implemented `CreateRollout`, `GetRollout`, and `ListRollouts` in `mockgcp/mockclouddeploy/rollout.go`.
  2. `CreateRollout` validates the parent release exists in mock storage and returns NotFound (404) if absent.
  3. Defaulted initial `approvalState` to `DOES_NOT_NEED_APPROVAL`, `state` to `PENDING_RELEASE`, and initialized the stable phase with deploy (PENDING) and verify (DISABLED) deployment jobs matching real GCP API responses.
  4. Added `CloudDeployRollout` to the mock-supported kinds switch in `config/tests/samples/create/harness.go`.
  5. Configured rollout array normalization (`.rollouts[]`, `.response.rollouts[]`) in `mockgcp/mockclouddeploy/normalize.go`.
  6. Generated authentic `_http_mock.log` files and confirmed 0-diff log alignment in `TestGoldenLogAlignment` and 14/14 tests passing in `tests-e2e-fixtures-suite clouddeploy`.





