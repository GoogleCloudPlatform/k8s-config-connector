# AppEngineServiceSplitTraffic MockGCP Alignment Journal

- **Date:** October 28, 2026
- **Resource Kind:** `AppEngineServiceSplitTraffic`
- **GCP Service:** `appengine.googleapis.com` (`google.appengine.v1.Services`)

## Overview

Existing `mockappengine` mock service only supported `google.appengine.v1.Firewall`. We implemented `google.appengine.v1.Services` (`ServicesServer`) in `mockgcp/mockappengine/services.go` and registered it on the gRPC server and HTTP mux.

## Real GCP Recording & Findings

1. **App Engine Application & Service Dependency**:
   - In real GCP, App Engine service traffic splitting operates on `apps/{project}/services/{service}` via `UpdateService` (`PATCH`) and `GetService` (`GET`).
   - In an App Engine application, the `default` service exists as the primary service. To test traffic splitting in real GCP, versions `v1` and `v2` were deployed to the `default` service in test project `cnrm-barni-4`.
2. **`migrateTraffic` Parameter**:
   - In Terraform provider (`resource_app_engine_service_split_traffic.go`), the update URL is templated as `{{AppEngineBasePath}}apps/{{project}}/services/{{service}}?migrateTraffic={{migrate_traffic}}&updateMask=split`.
   - When `migrateTraffic` is not explicitly set in the resource spec, `{{migrate_traffic}}` expands to `""` (`migrateTraffic=`), causing GCP to reject the request with a 400 Bad Request (`Invalid value at 'migrate_traffic' (TYPE_BOOL), ""`). Manifests must specify `migrateTraffic: false` (or `true`).
3. **Terraform Read Behavior and Re-reconciliation**:
   - In the legacy Terraform provider, `resourceAppEngineServiceSplitTrafficRead` only populates `project` and `service` into the Terraform schema state, and does not read back `split` or `migrate_traffic`.
   - Consequently, during re-reconciliation, the Terraform controller always detects a diff between desired spec and state, triggering an unexpected `PATCH`. Re-reconciliation testing was therefore excluded in `tests/e2e/ratcheting.go` with a clarifying comment until migration to direct controller occurs.
4. **Operation Long-Running Operations (LRO)**:
   - `UpdateService` returns an LRO `apps/{project}/operations/{operationId}` of type `google.appengine.v1.OperationMetadataV1`.
   - The operation polling path is `GET https://appengine.googleapis.com/v1/apps/{project}/operations/{operationId}`.

## Implementation Details

- **Services Server**: Implemented `GetService`, `UpdateService`, `DeleteService`, and `ListServices` in `mockgcp/mockappengine/services.go`.
- **Default Service Auto-population**: App Engine applications contain a `default` service by definition; `GetService` initializes the default service in mock storage if not already present.
- **Operations Route**: Added `grpcMux.AddOperationsPath("/v1/{prefix=apps/*}/operations/{name}", conn)` to `service.go`.
- **Normalization**: Added `ConfigureVisitor` to `mockappengine/normalize.go` replacing dynamic timestamps (`.metadata.insertTime`).

## Validation

- Recorded authentic real GCP golden logs (`_http.log`) and golden objects with `hack/record-gcp "fixtures/^appengineservicesplittrafficbasic$"` using project `cnrm-barni-4`.
- Ran `hack/compare-mock "fixtures/^appengineservicesplittrafficbasic$"` to generate `_http_mock.log` and verify full mock reconciliation.
- Executed `TestGoldenLogAlignment` and `TestRealHTTPLogsDoNotContainMockGCP`, passing with 0 errors and 0 diffs.
