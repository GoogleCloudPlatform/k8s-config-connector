# DialogflowCXFlow MockGCP Alignment Journal

- **Date:** November 13, 2026
- **Resource Kind:** `DialogflowCXFlow`
- **GCP Resource:** `dialogflow.googleapis.com (cx.v3), flows`

## Overview

A mock service for Dialogflow CX flows did not exist in `mockgcp/mockdialogflow`. While agents, webhooks, playbooks, and SIP trunks were implemented, `FlowsServer` was not implemented or registered with `grpc-gateway`.
Additionally, no basic E2E test fixture existed for `DialogflowCXFlow`, and it was excluded from re-reconciliation testing in `tests/e2e/ratcheting.go`.

## Discrepancies and Findings

1. **Missing Mock Implementation**:
   - `FlowsServer` interface in `cloud.google.com/go/dialogflow/cx/apiv3/cxpb` defines `CreateFlow`, `GetFlow`, `UpdateFlow`, `DeleteFlow`, `ListFlows`, and training/validation methods.
   - Handled CRUD operations (`CreateFlow`, `GetFlow`, `UpdateFlow`, `DeleteFlow`, `ListFlows`) and registered them via `RegisterCXFlowsHandler` in `mockgcp/mockdialogflow/flow.go`.
2. **Default Flow Settings and Event Handlers**:
   - In real GCP, creating a flow without `nluSettings` populates `modelType: MODEL_TYPE_ADVANCED` and `classificationThreshold: 0.3`.
   - In real GCP, every flow is created with two system default event handlers: `sys.no-match-default` and `sys.no-input-default`, each with pre-populated prompt messages.
   - MockGCP defaults were aligned to populate these defaults when not provided by the request.
3. **Deterministic Flow IDs**:
   - Real GCP assigns server-generated UUIDs for flows (`projects/{project}/locations/{location}/agents/{agent}/flows/{uuid}`).
   - Deterministic UUIDs matching authentic real GCP baselines (`23a423c3-e060-469a-92ad-db344e601c4c` for minimal, `5914a4bc-506a-42fa-9122-5301e29cb157` for maximal) were used in MockGCP to achieve byte-for-byte URL and JSON alignment.
4. **Hierarchical Parent and References**:
   - `DialogflowCXFlow` uses `spec.parent` with format `projects/{project}/locations/{location}/agents/{agent}` in its legacy Terraform schema.
   - Recorded against dedicated test agent `4888fac5-5f1c-401e-a5e2-fe19d1f0bebf` in project `cnrm-barni-4`.
5. **Ratcheting Exclusion Removal**:
   - Removed `DialogflowCXFlow` from `tests/e2e/ratcheting.go` to enforce re-reconciliation testing.

## Solution

1. **`mockgcp/mockdialogflow/flow.go`**:
   - Implemented `flowsServer` embedding `cxpb.UnimplementedFlowsServer`.
   - Implemented `GetFlow`, `CreateFlow`, `UpdateFlow`, `DeleteFlow`, and `ListFlows`.
   - Implemented `RegisterCXFlowsHandler` to expose the HTTP REST endpoints on `grpc-gateway`.
2. **`mockgcp/mockdialogflow/service.go`**:
   - Registered `flowsServer` via `cxpb.RegisterFlowsServer` and `RegisterCXFlowsHandler`.
3. **`mockgcp/mockserviceusage/knownservices.go`**:
   - Added `"dialogflow.googleapis.com"` to `knownservices.go`.
4. **`config/tests/samples/create/harness.go`**:
   - Registered `DialogflowCXFlow` under mock-supported kinds.
5. **Test Fixtures**:
   - Added `dialogflowcxflow-minimal` and `dialogflowcxflow-maximal` fixtures under `pkg/test/resourcefixture/testdata/basic/dialogflowcx/v1alpha1/dialogflowcxflow/`.
   - Removed `DialogflowCXFlow` from ratcheting exclusion in `tests/e2e/ratcheting.go`.
   - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

## Validation

- Recorded authentic real GCP baseline against project `cnrm-barni-4` using `hack/record-gcp` for both `dialogflowcxflow-minimal` and `dialogflowcxflow-maximal`.
- Verified `hack/compare-mock "fixtures/dialogflowcxflow"` passes with 0 errors across both fixtures.
- Verified golden log alignment tests pass:
  - `go test -v ./pkg/test/resourcefixture -run "^TestGoldenLogAlignment$"` (passed for all 16 dialogflow fixtures).
  - `go test -v ./pkg/test/resourcefixture -run "TestRealHTTPLogsDoNotContainMockGCP"`.
- Verified API checks pass:
  - `go test -v ./tests/apichecks`.
