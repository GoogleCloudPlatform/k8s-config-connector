### [2026-11-17] Prepopulating CloudDeployRelease
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CloudDeployRelease` using `google.cloud.deploy.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `Release` in `google.cloud.deploy.v1` is considered non-declarative by default because the API defines `CreateRelease` without an `UpdateRelease` (create-only execution resource). Passing `--allow-non-declarative` to `generate-types` is required to scaffold it.
  2. Identity pattern is `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/releases/{release}` with ancestor segment `deliveryPipelines/{delivery_pipeline}`.
  3. `Release.target_artifacts` is an output-only map (`map<string, TargetArtifact>`) where `TargetArtifact` contains only output-only fields in proto. Under `--prepopulate-resource`, `TargetArtifact` is generated as an empty struct, requiring `clouddeployreleases.clouddeploy.cnrm.cloud.google.com` in `knownInvalidCRDs` in `tests/apichecks/crds_test.go` until triaged in Step 1b.
  4. Generated `clouddeployrelease_identity.generated.go`, `clouddeployrelease_reference.generated.go`, and conformance tests cleanly.
  5. Added `CloudDeployRelease` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` and synchronized `tests/apichecks/testdata/exceptions/` golden files (`alpha-missingfields.txt`, `no_refs_in_status.txt`, `refs_not_representable.txt`, `spec_dislike_etag.txt`).
