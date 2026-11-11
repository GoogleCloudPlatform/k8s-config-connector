### [2026-05-17] Move mockgcp parametermanager away from grpc-gateway
- **Context**: Migrating mockparametermanager to use official protos and httptogrpc.
- **Problem**: No special complexities encountered for this service.
- **Solution**: Successfully moved mockparametermanager to use official protos and httptogrpc.
- **Impact**: Verified with targeted E2E tests for Parameter and ParameterVersion.

### [2026-11-11] Prepopulate direct KRM types, identity, and judgement queue for ParameterManagerTemplate
- **Context**: Prepopulating ParameterManagerTemplate under `apis/parametermanager/v1alpha1` (Step 1a).
- **Protos & Dependencies**: ParameterManagerTemplate proto (`google.cloud.parametermanager.v1.Template`) was added in googleapis commit `4fabd6e4b356307d319b55a86a3f200e9c52b594`. Updated `apis/git.versions` to this commit, and upgraded `cloud.google.com/go/parametermanager` to `v1.1.0` in root `go.mod` so `mapper.generated.go` compiles cleanly with `pb.Template`.
- **Identity & References**: Generated `parametermanagertemplate_identity.generated.go` and `parametermanagertemplate_reference.generated.go` with pattern `projects/{project}/locations/{location}/templates/{template}`. Conformance tests pass.
- **Release Bundle**: Added `ParameterManagerTemplate` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.

