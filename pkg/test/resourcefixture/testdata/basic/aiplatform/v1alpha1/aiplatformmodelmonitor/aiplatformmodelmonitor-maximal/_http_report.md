# HTTP Lifecycle Report: `aiplatformmodelmonitor-maximal`

- **Primary Resource:** `AIPlatformModelMonitor` (`aiplatform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `aiplatformmodelmonitor-max-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:86808272c39e4368`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"modelmonitor-max-${uniqueId}"`<br>**to:** `"modelmonitor-upd-${uniqueId}"` | Yes (#9 `PATCH`) | Yes |
| `spec.notificationSpec.`<br>`.emailConfig.userEmails` | **from:** `["alert-${uniqueId}@example.com"]`<br>**to:** `["updated-${uniqueId}@example.com"]` | Yes (#9 `PATCH`) | Yes |
| `spec.tabularObjective.`<br>`.featureDriftSpec.`<br>`.defaultNumericAlertCondition.`<br>`.threshold` | **from:** `0.3`<br>**to:** `0.5` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AIPlatformModel` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/us-central1/modelMonitors/`<br>`aiplatformmodelmonitor-max-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`modelMonitors`<br>`?modelMonitorId=aiplatformmodelmonitor-max-${uniqueId}` | LRO: create -> done (#7)<br>Created resource |
| 8 | `UPDATE` | `GET` `200` | `.../locations/us-central1/modelMonitors/`<br>`aiplatformmodelmonitor-max-${uniqueId}` | Pre-update read |
| 9 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/modelMonitors/`<br>`aiplatformmodelmonitor-max-${uniqueId}` | LRO: update -> done (sync)<br>updateMask: displayName, notificationSpec, tabularObjective |
| 10 | `DELETE` | `GET` `200` | `.../locations/us-central1/modelMonitors/`<br>`aiplatformmodelmonitor-max-${uniqueId}` | Pre-delete read |
| 11 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/modelMonitors/`<br>`aiplatformmodelmonitor-max-${uniqueId}` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_aiplatformmodelmonitor-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
