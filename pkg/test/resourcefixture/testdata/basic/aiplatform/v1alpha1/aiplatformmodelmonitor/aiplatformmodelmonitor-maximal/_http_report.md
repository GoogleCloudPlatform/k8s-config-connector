# HTTP Lifecycle Report: `aiplatformmodelmonitor-maximal`

- **Primary Resource:** `AIPlatformModelMonitor` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `modelmonitor-max-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:794dca9fd5d299cb`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"modelmonitor-max-${uniqueId}"`<br>**to:** `"modelmonitor-max-${uniqueId}-updated"` | Yes (#13 `GRPC`) | Yes |
| `spec.modelMonitoringSchema.`<br>`.featureFields` | **from:** `[{"dataType": "string", "name": "feature1"}, {"dataType": "fl...`<br>**to:** `[{"dataType": "string", "name": "feature1"}, {"dataType": "fl...` | Yes (#13 `GRPC`) | Yes |
| `spec.notificationSpec.`<br>`.emailConfig.userEmails` | **from:** `["alerts@example.com"]`<br>**to:** `["alerts-updated@example.com"]` | Yes (#13 `GRPC`) | Yes |
| `spec.notificationSpec.`<br>`.enableCloudLogging` | **from:** `true`<br>**to:** `false` | Yes (#13 `GRPC`) | Yes |
| `spec.tabularObjective.`<br>`.featureDriftSpec.`<br>`.defaultCategoricalAlertCondition.`<br>`.threshold` | **from:** `0.1`<br>**to:** `0.2` | Yes (#13 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AIPlatformModel` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Pre-create existence check |
| 6 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`CreateModelMonitor` | LRO: create -> pending (sync)<br>Created resource |
| 7 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Post-create eventual-consistency poll |
| 8 | `CREATE` | `GRPC` `400` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`CreateModelMonitor` | Composite create follow-up (`GRPC`) |
| 9 | `CREATE` | `GRPC` `404` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Post-create eventual-consistency poll |
| 10 | `CREATE` | `GRPC` `400` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`CreateModelMonitor` | Composite create follow-up (`GRPC`) |
| 11 | `CREATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Post-create read (matches `create.yaml`) |
| 12 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Pre-update read |
| 13 | `UPDATE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`UpdateModelMonitor` | LRO: update -> done (sync)<br>updateMask: displayName, modelMonitoringSchema, notificationSpec, tabularObjective |
| 14 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Pre-delete read |
| 15 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`GetModelMonitor` | Pre-delete read |
| 16 | `DELETE` | `GRPC` `200` | `/google.cloud.aiplatform.v1beta1.ModelMonitoringService/`<br>`DeleteModelMonitor` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformmodelmonitor-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
