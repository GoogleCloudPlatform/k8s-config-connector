# HTTP Lifecycle Report: `appengineflexibleappversionbasic`

- **Primary Resource:** `AppEngineFlexibleAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `appengineflexibleappversion-${uniqueId} (resourceID: v${uniqueId})`
- **Source Log:** `_http.log` (11 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:c05a88ec34dc527b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.manualScaling.instances` | **from:** `1`<br>**to:** `2` | Yes (#5 `POST`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `/v1/apps/${projectId}/services/default/`<br>`versions` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `POST` `200` | `/v1/apps/${projectId}/services/default/`<br>`versions` | LRO: create -> done (#6..#8)<br>Updated resource (`POST`) |
| 7 | `UPDATE` | `GET` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `GET` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}` | LRO: delete -> done (#11)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`POST`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_appengineflexibleappversionbasic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
