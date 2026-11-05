### [2026-10-21] Direct Controller, Fuzzer, and E2E Fixtures Implementation for SaaSServiceMgmtRolloutKind
- **Context**: Implementing the Greenfield direct controller, KRM fuzzer, and real GCP test fixtures for `SaaSServiceMgmtRolloutKind`.
- **Observations**:
  - **SaasRolloutsClient**: `RolloutKind` methods (`CreateRolloutKind`, `GetRolloutKind`, `UpdateRolloutKind`, `DeleteRolloutKind`) are on `SaasRolloutsClient` (`NewSaasRolloutsRESTClient`), rather than `SaasDeploymentsClient`.
  - **RolloutOrchestrationStrategy Validation**: The live GCP API requires `rolloutOrchestrationStrategy` to be specified during creation (e.g. `Google.Cloud.Simple.AllAtOnce` or `Google.Cloud.Simple.OneLocationAtATime`), otherwise failing with `400 Bad Request: rollout kind: ... has an invalid strategy: , supported strategies: [...]`.
  - **Synchronous Operations**: Create, Update, and Delete on `SaasRolloutsClient` for RolloutKind are all synchronous calls that return the resource or error directly without using LROs.
  - **Audit Probe**: `_audit_probe.log` captures the GET REST call and audit markers are emitted to Cloud Logging.

### [2026-10-14] Direct Controller and Fuzzer Implementation for SaaSServiceMgmtRollout
- **Context**: Implementing the Greenfield direct controller, KRM fuzzer, and test fixtures for `SaaSServiceMgmtRollout`.
- **Observations**:
  - **Client Library**: Rollout operations belong to `SaasRolloutsClient` rather than `SaasDeploymentsClient`. Used `NewSaasRolloutsRESTClient` from `cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1`.
  - **Prerequisite Resources**: Real GCP requires a valid `RolloutKind` in the same location upon creation of a Rollout.
  - **Control Defaulting**: GCP defaults `control.action` to `ROLLOUT_ACTION_RUN` (enum value 1). When `spec.control` is not specified by the user, `compareRollout` preserves `maskedActual.Control` to prevent false diffs and spurious update requests during re-reconciliation.
  - **CEL Unit Filter & Mutability**: `unitFilter` takes a CEL boolean expression (e.g. `"true"`). In GCP, updating `unitFilter` is rejected unless the Rollout is in a `PAUSED` state.
  - **Synchronous Operations**: `CreateRollout`, `UpdateRollout`, and `DeleteRollout` are synchronous calls (not LROs).

### [2026-10-14] Direct Controller, E2E Fixtures, and Fuzzer Implementation for SaaSServiceMgmtTenant
- **Context**: Implementing the Greenfield direct controller, KRM fuzzer, and E2E test fixtures for `SaaSServiceMgmtTenant`.
- **Observations**:
  - **Synchronous API**: All operations (`CreateTenant`, `GetTenant`, `UpdateTenant`, `DeleteTenant`) on `SaasDeploymentsClient` are synchronous (do not return LROs).
  - **Client Selection**: Utilized GAPIC REST client (`NewSaasDeploymentsRESTClient`) with `m.config.RESTClientOptions()`.
  - **Spec Immutability**: All spec fields (`projectRef`, `location`, `resourceID`, `saasRef`, `consumerResource`) are immutable in GCP. As directed by the Greenfield controller instructions, `update.yaml` was omitted from both minimal and maximal fixtures.
  - **Reference Validation**: The GCP API accepted `saasRef` pointing to a SaaS resource format (`projects/{project}/locations/{location}/saas/kcc-test-saas`) without requiring the SaaS resource to exist beforehand.

### [2026-10-09] Direct Controller, Fuzzer, and E2E Fixtures Implementation for SaaSServiceMgmtUnit
- **Context**: Implementing the Greenfield direct controller, KRM fuzzer, and real GCP test fixtures for `SaaSServiceMgmtUnit`.
- **Observations**:
  - **UnitKind Requirement in GCP**: While proto marks `unit_kind` as optional, the live GCP API returns `400 Bad Request: resource name is empty` if `unitKind` is omitted during `CreateUnit`. Supplying `unitKindRef` resolves this.
  - **Tenant Gateway Verification**: Specifying `tenantRef` causes live GCP API calls to fail with `400 Bad Request: unable to retrieve Tenant ... from the gateway: rpc error: code = NotFound`. Unlike `unitKind`, tenant existence is strictly verified on creation. Since KCC does not currently manage `SaaSServiceMgmtTenant`, `tenantRef` was omitted from maximal test fixtures.
  - **Identifier Handling in Diff**: Restoring `maskedActual.Name = desired.Name` in `compareUnit` is essential after `mappers.OnlySpecFields` to avoid false diffs on the `name` field during re-reconciliation.
  - **Synchronous Operations**: `CreateUnit`, `UpdateUnit`, and `DeleteUnit` are all synchronous calls on `SaasDeploymentsClient`, avoiding LRO polling.

### [2026-10-01] Scaffolding Types and Identity for SaaSServiceMgmtRolloutKind
- **Context**: Greenfield scaffolding of direct KRM types, identity, reference, and generate.sh for `SaaSServiceMgmtRolloutKind`.
- **Observations**:
  - `RolloutKind` protobuf message is located in `google.cloud.saasplatform.saasservicemgmt.v1beta1` under `rollouts_resources.proto`.
  - The resource uses URL template `projects/{project}/locations/{location}/rolloutKinds/{rolloutkind}` and maps to `saasservicemgmt.googleapis.com/RolloutKind`.
  - `pb.ErrorBudget` contains proto3 optional fields `allowed_count` and `allowed_percentage` mapping to `*int32`. A manual mapper in `pkg/controller/direct/saasservicemgmt/mapper.go` was created to assign pointer fields directly.

### [2026-07-03] Direct Controller and Fuzzer Implementation for SaasServiceMgmtRelease
- **Context**: Implementing the Greenfield direct controller, KRM fuzzer, and test fixtures for `SaasServiceMgmtRelease`.
- **Observations**: 
  - **REST Go Client**: Bypassed low-level gRPC or HTTP JSON mappings by utilizing the standard GCP Go client from `cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1`.
  - **No LROs**: Verified that the `CreateRelease`, `UpdateRelease`, and `DeleteRelease` client operations are synchronous and return objects or errors directly without using long-running operations.
  - **MockGCP Skipping**: Discovered that mockgcp automatically skips resources/GroupKinds that are not in its mock registry, enabling the test suite to pass gracefully without implementing a mock service for `saasservicemgmt` in mockgcp.

### [2026-07-02] Proto Package Discrepancy for SaasServiceMgmtRelease
- **Context**: Scaffolding new types for the `SaasServiceMgmtRelease` Kind (`saasservicemgmt.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: The issue specified using `--service google.cloud.saasservicemgmt.v1`, but no such proto package or `Release` message exists under that package path in the cloned `googleapis` third_party repository.
- **Solution**: Traced the `Release` protobuf message to `google/cloud/saasplatform/saasservicemgmt/v1beta1/deployments_resources.proto` in package `google.cloud.saasplatform.saasservicemgmt.v1beta1`. Updated the `--service` flag in `generate.sh` to `google.cloud.saasplatform.saasservicemgmt.v1beta1` which successfully resolved and scaffolded the types.
- **Impact**: Future agents working on saasservicemgmt controllers/mappers should use `google.cloud.saasplatform.saasservicemgmt.v1beta1` as the protobuf package path to avoid "proto not found" compiler errors.
