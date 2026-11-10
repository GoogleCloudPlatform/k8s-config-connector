# DialogflowCXFlow Identity and Reference Journal

## Observations & Learnings

- **Server-Generated ID with Optional ResourceID**:
  - `DialogflowCXFlow` uses server-generated flow IDs in GCP (`projects/{project}/locations/{location}/agents/{agent}/flows/{flow}`).
  - In `DialogflowCXFlowSpec`, `resourceID` is optional and used for acquisition/import. When creating a new flow, `resourceID` is left empty.
  - Implements `identity.ServerGeneratedIdentity` where `HasIdentitySpecified()` checks `i.Flow != ""`.

- **Hierarchical Parent Structure (`spec.parent`)**:
  - The resource uses `spec.parent` with format `projects/{project}/locations/{location}/agents/{agent}` rather than direct `spec.location` and `spec.agentRef` fields due to its legacy Terraform schema.
  - We use a scoped `gcpurls.Template` helper `dialogflowCXAgentIdentityFormat` to parse `spec.parent` into project, location, and agent segments.
  - `getIdentityFromDialogflowCXFlowSpec` and `GetIdentity` are overridden declaration-by-declaration in `dialogflowcxflow_identity.go`.

- **Status Cross-Check & Legacy Fallback**:
  - In `GetIdentity`, both `status.externalRef` and `status.name` are supported to seamlessly work with both modern direct controllers and existing legacy Terraform controllers.
  - In `DialogflowCXFlowRef.Normalize`, legacy fallback retrieves `status.name` when `status.externalRef` is not yet populated by a direct controller.

- **Deterministic Code Generation**:
  - Configured `${CONTROLLERBUILDER} generate-identity` in `apis/dialogflowcx/generate.sh` to deterministically generate identity, reference, conformance test, and deepcopy boilerplate.
  - Triaged `judgement_queue.yaml` entries for `identity-root-unknown` and `identity-server-generated-id`.
