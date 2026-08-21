# Journal: BigQueryReservationCapacityCommitment Direct Controller Implementation

## Author
codebot-robot

## Date
2026-06-08

## Summary
Successfully implemented the direct controller, mock endpoints, and E2E test fixtures for `BigQueryReservationCapacityCommitment`.

## Observations & Implementation Details
- **MockGCP Registration Bug Fix**: During testing, we found that `mockbigqueryreservation` was never registered with the global `mockgcpregistry` used by the test normalizer because it was missing `init() { mockgcpregistry.Register(New) }` in `service.go` and was not listed in `mockgcp/register.go`. This prevented timestamp normalization from running. We corrected this by registering the service with `mockgcpregistry` and removing its duplicate manual registration from `mock_http_roundtrip.go`.
- **Normalization Rules**: Added rules in both `mockgcp/mockbigqueryreservation/normalize.go` and `tests/e2e/normalize.go` to filter dynamic `commitmentStartTime` and `commitmentEndTime` fields. This ensures that HTTP logs and status files are perfectly deterministic.
- **Controller Structure**: Implemented the direct controller under `pkg/controller/direct/bigqueryreservation/capacitycommitment_controller.go` using standard patterns, integrating it with the generated mapping functions and `common.NormalizeReferences`.
- **Promotion & Backporting**: Since `BigQueryReservationCapacityCommitment` is brownfield (with existing legacy/terraform controller), we kept the default controller as legacy/terraform for gradual opt-in, while adding `direct` as a supported controller and forcing direct controllers in `tests/e2e/unified_test.go` to test both implementations.
