# BeyondCorpAppConnection Identity and Reference Journal

- **Resource**: `BeyondCorpAppConnection` (group: `beyondcorp.cnrm.cloud.google.com`, version: `v1alpha1`).
- **Generation**: Added `controllerbuilder generate-identity` to `apis/beyondcorp/generate.sh`.
- **Identity Pattern**: Matches `docs/ai/metadata/cloudassetinventory_names.jsonl`: `projects/{project}/locations/{location}/appConnections/{appConnection}` with host `beyondcorp.googleapis.com`.
- **Location vs Region**: In `BeyondCorpAppConnectionSpec`, the location parameter is named `region` (`spec.region`). `controllerbuilder generate-identity` automatically resolved `obj.Spec.Region` for the location segment.
- **Status Field & Cross-Check**: Added `ExternalRef *string` to `BeyondCorpAppConnectionStatus`. `GetIdentity` checks and cross-checks the parsed identity against `status.externalRef`.
- **Reference Normalization**: Fully generated via `controllerbuilder generate-identity` using `refs.Normalize`.
- **Verification**: Regenerated CRDs with `dev/tasks/generate-crds`. All tests in `apis/beyondcorp/v1alpha1/...`, `pkg/gcpurls/...`, `pkg/cli/powertools/cais/...`, and `tests/apichecks/...` passed cleanly.
