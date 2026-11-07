# AIPlatformTensorboardRun Journal

## Resource
- Kind: `AIPlatformTensorboardRun`
- Proto Message: `google.cloud.aiplatform.v1.TensorboardRun`
- Service: `aiplatform`

## Notes & Findings
- **Proto Source Path Pinning**: `apis/aiplatform/generate.sh` pins `--proto-source-path ${REPO_ROOT}/.build/googleapis-e7513c828b8db536fa12fc5493d541f323aedf90.pb` for existing resources to preserve CRD backward compatibility against `master`.
- **Identity & Reference**: Generated `aiplatformtensorboardrun_identity.generated.go` and `aiplatformtensorboardrun_reference.generated.go` with pattern `projects/{project}/locations/{location}/tensorboards/{tensorboard}/experiments/{experiment}/runs/{run}`.
- **Blocked Kinds**: Added `AIPlatformTensorboardRun` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Apichecks**: Recorded untriaged field absences in `alpha-missingfields.txt` and `spec_dislike_etag.txt`.
