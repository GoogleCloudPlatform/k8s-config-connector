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
