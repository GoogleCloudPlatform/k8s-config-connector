# HTTP Lifecycle Report: `dataformworkflowconfig-minimal`

- **Primary Resource:** `DataformWorkflowConfig` (`dataform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dfworkflowconfig-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 13 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:30d22d9b67fe2224`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.cronSchedule` | **from:** *(unset)*<br>**to:** `"0 6 * * *"` | Yes (#11 `PATCH`) | Yes |
| `spec.timeZone` | **from:** *(unset)*<br>**to:** `"UTC"` | Yes (#11 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataformReleaseConfig` | 4 | `CREATE`, `DELETE` |
| `DataformRepository` | 4 | `CREATE`, `DELETE` |
| `IAMServiceAccount` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/workflowConfigs/`<br>`dfworkflowconfig-${uniqueId}` | Pre-create existence check |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/`<br>`workflowConfigs`<br>`?workflowConfigId=dfworkflowconfig-${uniqueId}` | Created resource |
| 10 | `UPDATE` | `GET` `200` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/workflowConfigs/`<br>`dfworkflowconfig-${uniqueId}` | Pre-update read |
| 11 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/workflowConfigs/`<br>`dfworkflowconfig-${uniqueId}` | updateMask: cronSchedule, timeZone |
| 12 | `DELETE` | `GET` `200` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/workflowConfigs/`<br>`dfworkflowconfig-${uniqueId}` | Pre-delete read |
| 13 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/dfrepo-${uniqueId}/workflowConfigs/`<br>`dfworkflowconfig-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataformworkflowconfig-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
