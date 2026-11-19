# DialogflowCXIntent Identity and Reference Migration Journal

## Observations
- `DialogflowCXIntent` is a child resource under `DialogflowCXAgent`, with canonical GCP URL: `projects/{project}/locations/{location}/agents/{agent}/intents/{intent}` and host `dialogflow.googleapis.com`.
- The resource uses server-generated IDs in GCP (the proto `google.cloud.dialogflow.cx.v3.Intent` has no client-specified ID parameter in `CreateIntentRequest`, and `spec.resourceID` is optional for acquisition only). Hence, it implements `identity.ServerGeneratedIdentity`.
- `DialogflowCXIntentSpec` specifies the parent via `spec.parent` (formatted as `projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>`). Adding `${CONTROLLERBUILDER} generate-identity` in `apis/dialogflowcx/generate.sh` produced `identity-root-unknown` and `identity-server-generated-id` judgement queue entries, which were triaged and resolved in `apis/dialogflowcx/judgement_queue.yaml`.
- Followed the preferred deterministic generation pattern with handwritten overrides:
  - `apis/dialogflowcx/v1alpha1/dialogflowcxintent_identity.go` implements custom `getIdentityFromDialogflowCXIntentSpec` (parsing `spec.parent` with `agentParentFormat` `gcpurls.Template`) and `GetIdentity` (supporting both `status.externalRef` and `status.name` fallback and drift detection).
  - `apis/dialogflowcx/v1alpha1/dialogflowcxintent_reference.go` implements `Normalize` with legacy fallback via `refs.NormalizeWithFallback`, checking readiness and falling back to `status.name` or `getIdentityFromDialogflowCXIntentSpec`.
  - The remaining identity struct, templates, methods, registration, and reference methods were emitted deterministically by `generate-identity` into `_identity.generated.go`, `_reference.generated.go`, and `_identity_generated_test.go`.
- The URL format `//dialogflow.googleapis.com/projects/{}/locations/{}/agents/{}/intents/{}` was verified against `docs/ai/metadata/discovery_names.jsonl` and `pkg/gcpurls/registry_test.go` passed without requiring registry exceptions.
- Unit tests cover conformance, string/URL formatting, host, parent string, server-generated ID defaulting, drift detection, validation, and reference normalization (covering valid/invalid external, ready with status.externalRef, ready with status.name, ready with spec fallback, and not-ready).
- Verified zero schema difference using `dev/tasks/diff-crds`.
