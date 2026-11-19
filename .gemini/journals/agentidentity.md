### [2026-11-10] AgentIdentity Prepopulation
- **Context**: Prepopulating direct KRM types, identity, and `judgement_queue.yaml` for `AgentIdentityAuthProvider`.
- **Proto & Dependencies**:
  - Service: `google.cloud.agentidentity.v1`
  - Resource: `AgentIdentityAuthProvider:AuthProvider` (`projects/{project}/locations/{location}/authProviders/{auth_provider}`)
  - Pinned googleapis commit `2c4dd791cc44104f1fdc49b4cbdadb160772b11e` in `apis/agentidentity/generate.sh` via `PROTO_SHA` (which introduced the Agent Identity v1 API).
  - Added Go client dependency `cloud.google.com/go/agentidentity v0.1.0`.
- **Validation**:
  - `geAuthProvider` (`GeminiEnterpriseAuthProviderParams`) is an empty oneof option message; added `agentidentityauthproviders.agentidentity.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go`.
  - Added `AgentIdentityAuthProvider` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.

### [2026-11-11] AgentIdentityAuthProvider Types & Judgement Queue Finalization
- **Context**: Triaged `apis/agentidentity/judgement_queue.yaml` and finalized KRM types and identity for `AgentIdentityAuthProvider`.
- **Decisions**:
  - **Parent & Location**: Inlined `*parent.ProjectAndLocationRef` into `AgentIdentityAuthProviderSpec`, removing separate `projectRef` and `location` fields while preserving CRD schema keys with immutable location validation.
  - **Labels**: Removed top-level `Labels` map from `AgentIdentityAuthProviderSpec` so labels are reconciled via standard KRM `metadata.labels`.
  - **OAuth Credentials**: Resolved `clientID` and `clientSecret` under `threeLeggedOauth` and `twoLeggedOauth` as `not-applicable`; these represent external identity provider client identifiers/secrets, not GCP resource references or Secret Manager secret versions.
  - **Workload IDs**: Triaged `spec.workloadIDs` as `not-applicable` in judgement queue and deferred in `tests/apichecks/testdata/exceptions/refs_deferred.txt`; the field accepts IAM principal strings in `principal://` URI format rather than direct KCC resource references.
  - **Documentation Links**: Removed unverified REST doc URL guess annotation as no public REST reference page currently exists for `projects.locations.authProviders`.

### [2026-11-17] AgentIdentityAuthProvider Controller, Fuzzer & E2E Fixtures
- **Context**: Implementing Greenfield Step 2 direct controller, round-trip fuzzer, and E2E fixtures for `AgentIdentityAuthProvider`.
- **Implementation & Controller Decisions**:
  - Generated direct controller and fuzzer via `controllerbuilder generate-controller` and `generate-fuzzer` with `--prepopulate-resource`.
  - **Labels Handling**: Although `AuthProvider` proto includes `labels map<string, string>`, real GCP `AgentIdentity` does not store or return labels (returns empty map on GET). Used `common.CompareGreenfieldSpec` rather than `CompareGreenfieldSpecAndLabels` to prevent infinite re-reconciliation loops from spurious label differences.
  - **Input-Only Secret Fields in Update**: GCP `UpdateAuthProvider` requires that input-only secret fields (`apiKey.apiKey`, `threeLeggedOauth.clientSecret`, `twoLeggedOauth.clientSecret`) are only provided in the payload if their paths are explicitly included in `update_mask` (otherwise failing with HTTP 400). In `Update()`, unmasked secret fields are cleared to empty string before calling `UpdateAuthProvider`.
  - **Fixtures**: Implemented three fixtures covering all supported auth provider configurations:
    - `agentidentityauthprovider-minimal`: Basic API key auth provider with minimal required fields.
    - `agentidentityauthprovider-maximal`: 3-legged OAuth auth provider exercising `allowedScopes`, `blockedScopes`, `description`, `workloadIDs`, and `threeLeggedOauth` parameters (transitioning `enablePkce` from `false` to `true`).
    - `agentidentityauthprovider-twoleggedoauth`: 2-legged OAuth auth provider exercising `twoLeggedOauth` parameters.
  - **Golden Testing**: Recorded all three fixtures against real GCP (`cnrm-barni-4`) with valid `_audit_probe.log` entries and 7/7 validation checks passing.

### [2026-11-18] AgentIdentityAuthProvider MockGCP Implementation & Alignment
- **Context**: Implementing Greenfield Step 3 MockGCP and golden log alignment for `AgentIdentityAuthProvider`.
- **Implementation & Mock Decisions**:
  - Implemented `mockagentidentity` service under `mockgcp/mockagentidentity` supporting `AuthProviderService` with standard CRUD and list operations.
  - Extracted `auth_provider_service.proto` from googleapis (pinned commit `2c4dd791cc44104f1fdc49b4cbdadb160772b11e`) to `mockgcp/apis/google/cloud/agentidentity/v1/auth_provider_service.proto` and generated grpc-gateway client handler code into `mockgcp/generated/google/cloud/agentidentity/v1/auth_provider_service.pb.gw.go`.
  - Added `cloud.google.com/go/agentidentity v0.1.0` to `mockgcp/go.mod` and registered `mockagentidentity` in `mockgcp/register.go`.
  - Added `AgentIdentityAuthProvider` to supported MockGCP GroupKinds in `config/tests/samples/create/harness.go`.
  - Mock service behavior:
    - Default `State`: defaults `AuthProvider_STATE_UNSPECIFIED` (0) to `AuthProvider_ENABLED` (1).
    - `redirectUrl`: computed deterministically as `https://agentidentitycredentials.googleapis.com/v1/{name}/oauthcallback` for 3-legged OAuth providers.
    - Secret and input-only stripping: `apiKey.apiKey`, `threeLeggedOauth.clientSecret`, `twoLeggedOauth.clientSecret`, and `workload_ids` are input-only and stripped before returning responses or storage read/write to match real GCP behavior.
    - Error responses: returns 404 NOT_FOUND with "The requested resource was not found." on non-existent resource lookups, matching real GCP.
  - **Golden Alignment**: Verified all 3 test fixtures (`agentidentityauthprovider-minimal`, `agentidentityauthprovider-maximal`, and `agentidentityauthprovider-twoleggedoauth`) against MockGCP with zero discrepancies between `_http.log` and `_http_mock.log`.
