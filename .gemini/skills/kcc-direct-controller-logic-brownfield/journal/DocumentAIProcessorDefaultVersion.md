# DocumentAIProcessorDefaultVersion Direct Controller Logic Journal

## Overview
Implemented the direct controller logic for `DocumentAIProcessorDefaultVersion` (brownfield resource migrating from Terraform).

## Observations & Implementation Insights

### 1. Sub-Aspect Semantics of Processor
- `DocumentAIProcessorDefaultVersion` is a 1-to-1 sub-aspect managing the `defaultProcessorVersion` attribute of a `pb.Processor` resource (`projects/{project}/locations/{location}/processors/{processor}`).
- `Find` checks for the parent processor using `GetProcessor`. If the processor exists, `Find` returns `true, nil` and stores the fetched processor in `a.actual`.
- `Update` and `Create` both call `SetDefaultProcessorVersion` when `diff` detects that the desired version does not match the actual default processor version.

### 2. Behavioral Parity with Legacy Terraform Controller
- The legacy Terraform controller does not support deleting the default processor version (emitting a warning that Google Cloud does not support deleting processor default versions).
- In the direct controller, the model is registered with `registry.CannotBeDeleted()`, and `Delete` is a no-op that returns `true, nil`.
- In `_http_mock.diff`, the direct controller uses the regional endpoint `https://us-documentai.googleapis.com:443/v1/...` with standard GAPIC query parameter formatting (`%24alt=json%3Benum-encoding%3Dint`).
- In `_final_object.diff`, the direct controller does not require the legacy `cnrm.cloud.google.com/state-into-spec: absent` annotation.
- Re-reconciliation in both `unified_test.go` and `TestMigrationToDirect` produces 0 updates/writes, confirming steady-state reconciliation.
