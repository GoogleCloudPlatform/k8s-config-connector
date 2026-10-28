# BeyondCorpAppConnection Identity and Reference Journal

- **Resource**: `BeyondCorpAppConnection` (group: `beyondcorp.cnrm.cloud.google.com`, version: `v1alpha1`).
- **Generation**: Added `controllerbuilder generate-identity` to `apis/beyondcorp/generate.sh`.
- **Identity Pattern**: Matches `docs/ai/metadata/cloudassetinventory_names.jsonl`: `projects/{project}/locations/{location}/appConnections/{appConnection}` with host `beyondcorp.googleapis.com`.
- **Location vs Region**: In `BeyondCorpAppConnectionSpec`, the location parameter is named `region` (`spec.region`). `controllerbuilder generate-identity` automatically resolved `obj.Spec.Region` for the location segment.
- **Status Cross-Check**: `BeyondCorpAppConnectionStatus` does not contain `ExternalRef` or `Name`. Per standard guidance, no status fields were added and `GetIdentity` returns `specIdentity` directly.
- **Reference Normalization**: Because `BeyondCorpAppConnection` is a legacy resource without `status.externalRef`, a hand-written `beyondcorpappconnection_reference.go` was created implementing `Normalize` with `refs.NormalizeWithFallback`, checking `status.conditions` for `Ready: True` before falling back to spec-based identity. `controllerbuilder`'s declaration-by-declaration override kept this method while generating the rest of the reference boilerplate.
- **Verification**: `dev/tasks/diff-crds` confirmed zero CRD schema drift. All tests in `apis/beyondcorp/v1alpha1/...`, `pkg/gcpurls/...`, `pkg/cli/powertools/cais/...`, and `tests/apichecks/...` passed cleanly.
