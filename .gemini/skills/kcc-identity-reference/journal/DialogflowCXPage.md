# DialogflowCXPage Identity and Reference Migration Journal

## Observations
- `DialogflowCXPage` is a child resource under `DialogflowCXFlow`, with canonical GCP URL: `projects/{project}/locations/{location}/agents/{agent}/flows/{flow}/pages/{page}` and host `dialogflow.googleapis.com`.
- The resource uses server-generated IDs in GCP (the proto `google.cloud.dialogflow.cx.v3.Page` has no client-specified ID parameter in `CreatePageRequest`, and `spec.resourceID` is optional for acquisition only). Hence, it implements `identity.ServerGeneratedIdentity`.
- `DialogflowCXPageSpec` specifies the parent via `spec.parent` (formatted as `projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>`). Using `${CONTROLLERBUILDER} generate-identity` in `apis/dialogflowcx/generate.sh` produced `identity-root-unknown` and `identity-server-generated-id` judgement queue entries, which were triaged and resolved in `apis/dialogflowcx/judgement_queue.yaml`.
- Followed the preferred deterministic generation pattern with handwritten overrides:
  - `apis/dialogflowcx/v1alpha1/dialogflowcxpage_identity.go` implements custom `getIdentityFromDialogflowCXPageSpec` (parsing `spec.parent` with `gcpurls.Template`) and `GetIdentity` (supporting both `status.externalRef` and `status.name` fallback and drift detection).
  - `apis/dialogflowcx/v1alpha1/dialogflowcxpage_reference.go` implements `Normalize` with legacy fallback via `refs.NormalizeWithFallback`, checking readiness and falling back to `status.name` or `getIdentityFromDialogflowCXPageSpec`.
  - The remaining identity struct, templates, methods, registration, and reference methods were emitted deterministically by `generate-identity` into `_identity.generated.go`, `_reference.generated.go`, and `_identity_generated_test.go`.
- Unit tests in `dialogflowcxpage_identity_test.go` cover conformance, string/URL formatting, host, parent string, server-generated ID defaulting, drift detection, validation, and reference normalization.
- Ran `WRITE_GOLDEN_OUTPUT=1 go test -v ./pkg/cli/powertools/cais/...` to update fixture golden identities.
- Verified zero schema difference using `dev/tasks/diff-crds`.
