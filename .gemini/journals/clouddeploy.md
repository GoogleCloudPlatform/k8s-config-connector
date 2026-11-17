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
