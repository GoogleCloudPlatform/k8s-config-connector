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
