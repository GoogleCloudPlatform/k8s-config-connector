# DialogflowCXEntityType Identity and Reference Migration Journal

## Observations
- `DialogflowCXEntityType` is a child resource under `DialogflowCXAgent`, with canonical GCP URL: `projects/{project}/locations/{location}/agents/{agent}/entityTypes/{entityType}` and host `dialogflow.googleapis.com`.
- The resource uses server-generated IDs in GCP (the proto `google.cloud.dialogflow.cx.v3.EntityType` has no client-specified ID parameter in `CreateEntityTypeRequest`, and `spec.resourceID` is optional for acquisition only). Hence, it implements `identity.ServerGeneratedIdentity`.
- `DialogflowCXEntityTypeSpec` specifies the parent via `spec.parent` (formatted as `projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>`). Using `${CONTROLLERBUILDER} generate-identity` in `apis/dialogflowcx/generate.sh` produced `identity-root-unknown` and `identity-server-generated-id` judgement queue entries, which were triaged and resolved in `apis/dialogflowcx/judgement_queue.yaml`.
- Followed the preferred deterministic generation pattern with handwritten overrides:
  - `apis/dialogflowcx/v1alpha1/dialogflowcxentitytype_identity.go` implements custom `getIdentityFromDialogflowCXEntityTypeSpec` (parsing `spec.parent` with `gcpurls.Template`) and `GetIdentity` (supporting both `status.externalRef` and `status.name` fallback and drift detection).
  - `apis/dialogflowcx/v1alpha1/dialogflowcxentitytype_reference.go` implements `Normalize` with legacy fallback via `refs.NormalizeWithFallback`, checking readiness and falling back to `status.name` or `getIdentityFromDialogflowCXEntityTypeSpec`.
  - The remaining identity struct, templates, methods, registration, and reference methods were emitted deterministically by `generate-identity` into `_identity.generated.go`, `_reference.generated.go`, and `_identity_generated_test.go`.
- Added `apis/dialogflowcx/v1alpha1` import to `pkg/gcpurls/registry_test.go`. The URL formats matched `docs/ai/metadata/discovery_names.jsonl` without requiring any registry exceptions.
- Unit tests cover conformance, string/URL formatting, host, parent string, server-generated ID defaulting, drift detection, validation, and reference normalization (covering valid/invalid external, ready with status.externalRef, ready with status.name, ready with spec fallback, and not-ready).
- Verified zero schema difference using `dev/tasks/diff-crds`.
