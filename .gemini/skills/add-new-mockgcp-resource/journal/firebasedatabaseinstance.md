# Journal: FirebaseDatabaseInstance, FirebaseAndroidApp, and FirebaseWebApp MockGCP Alignment

This journal documents the findings and design decisions while matching real GCP behavior in MockGCP for `FirebaseAndroidApp`, `FirebaseWebApp`, and `FirebaseDatabaseInstance`.

## Observations & Findings

1. **Protobuf & Schema Discovery**:
   - `FirebaseAndroidApp` and `FirebaseWebApp` protobufs are defined under `mockgcp/apis/mockgcp/cloud/firebase/v1beta1/service.proto` and already had generated stubs under `mockgcp/generated/mockgcp/cloud/firebase/v1beta1/`.
   - `FirebaseDatabaseInstance` is a REST/Discovery API (`firebasedatabase.googleapis.com`). The proto was generated from `firebasedatabase-api.json` using `tools/gapic` with package `mockgcp.cloud.firebasedatabase.v1beta` and compiled into `mockgcp/generated/mockgcp/cloud/firebasedatabase/v1beta/`.
   - `mockgcp/Makefile` was updated to generate `firebasedatabase` protos alongside `firebasehosting`.

2. **Real GCP Recording (Baseline)**:
   - Real GCP tests were run against project `cnrm-barni-4` using `hack/record-gcp`.
   - `firebasedatabase.googleapis.com` was enabled via `gcloud services enable firebasedatabase.googleapis.com --project cnrm-barni-4`.
   - Real GCP rejects database IDs starting with `firebase` ("Database id cannot use the reserved prefix 'firebase'"). The fixture was updated to use `resourceID: rtdb${uniqueId}`.
   - For `FirebaseWebApp`, `deletionPolicy` was omitted so desired matches actual and no spurious update diff occurs during re-reconciliation.
   - Fixed `protovalidate.go` V7 check: when a proto package is not present in `.build/googleapis.pb` (e.g. discovery-doc APIs returning internal `@type` strings), unknown `@type` is not treated as a schema violation if the package is not in the descriptor registry.

3. **MockGCP Implementation**:
   - `mockgcp/mockfirebase`: Extended with `ProjectsAndroidAppsServer` and `ProjectsWebAppsServer`.
     - `Create`: returns long-running operation via `StartLRO` with initial pending state matching real GCP.
     - `Delete`: returns completed operation with `OperationMetadata` matching real GCP `:remove` response.
     - Deterministic `appId` generation based on display name to avoid volatile hash churn.
     - Mapped `google.firebase.service.v1beta1` type URL in `common/httpmux/json.go` and `common/operations/operations.go` so gRPC gateway cleanly resolves message descriptors.
   - `mockgcp/mockfirebasedatabase`: Implemented CRUD for `ProjectsLocationsInstancesServer`.
     - Real GCP returns 404 with message `"Namespace does not exist."`.
     - Populates default values (`USER_DATABASE`, `databaseUrl: https://<id>.firebaseio.com`, `ACTIVE` state).
     - Handles `:disable`, `:reenable`, and `DELETE`.
   - Registered `mockfirebasedatabase` in `mockgcp/register.go` and `mockserviceusage/knownservices.go`.
   - Added GroupKinds to `config/tests/samples/create/harness.go`.
