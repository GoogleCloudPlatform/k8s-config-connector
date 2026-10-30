# HTTP Lifecycle Report: `networkconnectivityserviceconnectiontoken-minimal`

- **Primary Resource:** `NetworkConnectivityServiceConnectionToken` (`networkconnectivity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ncsct-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:fbe0a160256baa02`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `.spec` or `.metadata.labels` differences between `create.yaml` and `update.yaml`._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GRPC` `404` | `/google.cloud.networkconnectivity.v1.CrossNetworkAutomationService/`<br>`GetServiceConnectionToken` | Pre-create existence check |
| 6 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.CrossNetworkAutomationService/`<br>`CreateServiceConnectionToken` | LRO: create -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.CrossNetworkAutomationService/`<br>`GetServiceConnectionToken` | Post-create read (matches `create.yaml`) |
| 9 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.CrossNetworkAutomationService/`<br>`GetServiceConnectionToken` | Post-create read (matches `create.yaml`) |
| 10 | `DELETE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.CrossNetworkAutomationService/`<br>`DeleteServiceConnectionToken` | LRO: delete -> done (#11)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `.spec` or `.metadata.labels` changes between `create.yaml` and `update.yaml`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkconnectivityserviceconnectiontoken-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
