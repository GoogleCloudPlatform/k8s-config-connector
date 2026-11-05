# Journal: Match real GCP behavior in MockGCP for DocumentAIProcessorDefaultVersion

## Overview
This journal documents the findings, design decisions, and learnings from matching MockGCP behavior with real GCP for `DocumentAIProcessorDefaultVersion` (`documentai.cnrm.cloud.google.com/v1alpha1`).

## Observations & Findings

1. **Service and Proto Identification**:
   - The GCP service is Document AI (`documentai.googleapis.com` and `{region}-documentai.googleapis.com`).
   - The proto service is `google.cloud.documentai.v1.DocumentProcessorService`.
   - In GCP, `default_processor_version` is a field on the `Processor` resource (`projects/{project}/locations/{location}/processors/{processor}`).
   - The API method to update it is `SetDefaultProcessorVersion`, which is an LRO returning `SetDefaultProcessorVersionMetadata` and `SetDefaultProcessorVersionResponse`.
   - In KCC, `DocumentAIProcessorDefaultVersion` is managed by the Terraform/legacy controller (`google_document_ai_processor_default_version`).
   - Terraform reads the version using `GetProcessor` and sets the version via `POST ...:setDefaultProcessorVersion`.
   - Processor IDs in GCP Document AI are server-generated hex strings and cannot be specified by the user.

2. **Ratcheting & Missing Fields**:
   - `DocumentAIProcessorDefaultVersion` was removed from `tests/e2e/ratcheting.go` to enforce re-reconciliation testing.
   - Removed `.spec.version` from `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` as it is now exercised by the basic test fixture.

3. **Real GCP Recording**:
   - Recorded against live GCP project `cnrm-barni-4` using `./hack/record-gcp "fixtures/^documentaiprocessordefaultversion-basic$"`.
   - Start timestamp: Fri Oct 23 17:40:39 PDT 2026.
   - Finish timestamp: Fri Oct 23 17:41:36 PDT 2026.
   - Successfully verified resource acquisition/reconciliation, re-reconciliation, deletion, and wrote golden output to `_http.log`.
   - Committed the recorded golden files in their own baseline commit.

4. **MockGCP Implementation & Alignment**:
   - Added `{region}-documentai.googleapis.com` pattern to `ExpectedHosts()` in `mockdocumentai/service.go`.
   - Implemented `SetDefaultProcessorVersion` in `mockdocumentai/processors.go`.
   - Updated `GetProcessor` to auto-populate defaults on pre-created test processors while returning `NotFound` for non-existent processors (allowing normal creation flow for `DocumentAIProcessor`).
   - Added `mockdocumentai/normalize.go` implementing `mockgcpregistry.SupportsNormalization`.
   - Added `mockdocumentai/service_test.go` verifying service registration and expected hosts.
   - Enabled `DocumentAIProcessorDefaultVersion` in `config/tests/samples/create/harness.go`.
   - Executed `./hack/compare-mock "fixtures/^documentaiprocessordefaultversion-basic$"`:
     - `_http_mock.log` was generated.
     - `TestGoldenLogAlignment` passed with 0 errors across all documentai fixtures.
