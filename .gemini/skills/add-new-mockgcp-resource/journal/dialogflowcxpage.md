# DialogflowCXPage MockGCP Alignment Journal

- **Date:** November 17, 2026
- **Resource Kind:** `DialogflowCXPage`
- **GCP Resource:** `dialogflow.googleapis.com (cx.v3), pages`

## Overview

MockGCP support and golden log alignment for `DialogflowCXPage` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`).

## Observations & Implementation Notes

1. **Pages Service Implementation**:
   - Implemented `pagesServer` implementing `cloud.google.com/go/dialogflow/cx/apiv3/cxpb.PagesServer` in `mockgcp/mockdialogflow/page.go`.
   - Registered `cxpb.RegisterPagesServer` and REST HTTP handlers via `RegisterCXPagesHandler` for:
     - `POST /v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages`
     - `GET /v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}`
     - `PATCH /v3/{page.name=projects/*/locations/*/agents/*/flows/*/pages/*}`
     - `DELETE /v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}`
     - `GET /v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages`

2. **Server-Generated Identifiers**:
   - Dialogflow CX assigns unique UUIDs to pages on creation.
   - Dialogflow CX also assigns unique UUIDs to sub-elements within `transitionRoutes` (`name`) and `eventHandlers` (`name`).
   - For golden test fixtures (`-min` and `-max`), matched the exact UUIDs recorded in real GCP baseline logs.

3. **Terraform JSON Validation for SetParameterActions**:
   - The Terraform provider validates `setParameterActions[].value` using `validation.StringIsJSON`, which requires the value string in KRM to be valid JSON (e.g. `'"val1"'`). During HTTP serialization, this is unmarshaled into its underlying JSON value.

4. **Real GCP Recording**:
   - Removed `DialogflowCXPage` from ratcheting exclusion list in `tests/e2e/ratcheting.go`.
   - Registered `DialogflowCXPage` under mock-supported kinds in `config/tests/samples/create/harness.go`.
   - Executed `hack/record-gcp` against live GCP project `cnrm-barni-4` using agent `4888fac5-5f1c-401e-a5e2-fe19d1f0bebf` flow `00000000-0000-0000-0000-000000000000`.

5. **Verification**:
   - `hack/compare-mock "fixtures/dialogflowcxpage"` passed with 0 failures.
   - `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/.*dialogflowcxpage"` passed.
   - `go test -v ./pkg/test/resourcefixture -run "TestRealHTTPLogsDoNotContainMockGCP"` passed.
   - `go test -v ./tests/apichecks/...` passed.
