# DialogflowAgent MockGCP Alignment Journal

- **Date:** November 17, 2026
- **Resource Kind:** `DialogflowAgent`
- **GCP Resource:** `dialogflow.googleapis.com (v2/v2beta1), agents`

## Overview

A mock service for Dialogflow CX agents existed in `mockgcp/mockdialogflow/agent.go`, but Dialogflow ES (v2) agents (`DialogflowAgent`, mapped to `projects/{project}/agent`) were not yet implemented in MockGCP.
Additionally, no basic E2E test fixture existed for `DialogflowAgent` as a primary resource, and it was excluded from re-reconciliation testing in `tests/e2e/ratcheting.go`.

## Discrepancies and Findings

1. **Dialogflow ES (v2) Agents API Support**:
   - Dialogflow ES agents use singleton endpoints per project:
     - `GET https://dialogflow.googleapis.com/v2/projects/{project}/agent`
     - `POST https://dialogflow.googleapis.com/v2/projects/{project}/agent` (creates or updates)
     - `DELETE https://dialogflow.googleapis.com/v2/projects/{project}/agent`
   - In `mockgcp/generated/google/cloud/dialogflow/v2beta1/agent.pb.gw.go`, `RegisterAgentsHandler` handles `/v2beta1/{parent=projects/*}/agent` (with `service.go` rewriting `/v2/` to `/v2beta1/`).
   - Mocking requires implementing `pb.AgentsServer` (`GetAgent`, `SetAgent`, `DeleteAgent`) using `cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb`.
2. **NotFound Error Formatting**:
   - Real GCP returns 404 with error message:
     `com.google.apps.framework.request.NotFoundException: No DesignTimeAgent found for project '<project>'.`
   - Returning this exact message with `codes.NotFound` aligns with real GCP behavior.
3. **Defaulting Behavior**:
   - `apiVersion`: defaults to `API_VERSION_V2` if unspecified.
   - `tier`: defaults to `TIER_STANDARD` if unspecified.
4. **Terraform Provider Write-Only `tier` Field**:
   - The Terraform provider docs note that `tier` is not read back from the API due to consistency issues. Including `tier` in test specs causes an infinite diff loop during re-reconciliation because live state always has an empty string for `tier`. Omitting `tier` from the test spec ensures clean re-reconciliation.
5. **Missing Fields Coverage**:
   - Tested 9 previously missing fields across `dialogflowagent-minimal` and `dialogflowagent-maximal`, reducing missing fields entries in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
6. **Ratcheting Exclusion**:
   - Removed `DialogflowAgent` from `tests/e2e/ratcheting.go` to enable re-reconciliation validation.

## Solution

1. **`mockgcp/mockdialogflow/v2agent.go`**:
   - Implemented `agentsServer` implementing `pb.AgentsServer` with `GetAgent`, `SetAgent`, and `DeleteAgent`.
   - Handled project-scoped storage with key `projects/{project}/agent`.
   - Populated defaults for `apiVersion` (`API_VERSION_V2`) and `tier` (`TIER_STANDARD`).
2. **`mockgcp/mockdialogflow/service.go`**:
   - Registered `pb.RegisterAgentsServer(grpcServer, &agentsServer{MockService: s})`.
   - Added `grpcpb.RegisterAgentsHandler` to `httpmux.NewServeMux`.
3. **Test Fixtures & Harness**:
   - Added `dialogflowagent-minimal` and `dialogflowagent-maximal` fixtures under `pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowagent/`.
   - Added `DialogflowAgent` to `config/tests/samples/create/harness.go`.
   - Removed `DialogflowAgent` from ratcheting exclusion in `tests/e2e/ratcheting.go`.
   - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

## Validation

- Recorded authentic real GCP baseline against project `cnrm-barni-4` using `hack/record-gcp` for both `dialogflowagent-minimal` and `dialogflowagent-maximal`.
- Ran `hack/compare-mock "fixtures/dialogflowagent"` with 0 errors across both fixtures.
- Verified golden log alignment tests pass:
  - `go test -v ./pkg/test/resourcefixture -run "^TestGoldenLogAlignment$"`.
- Verified API checks pass:
  - `go test -v ./tests/apichecks`.
