# HTTP Lifecycle Report: `dataformworkflowconfig-minimal`

- **Primary Resource:** `DataformWorkflowConfig` (`dataform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dfwc-${uniqueId}`
- **Source Log:** `_http.log` (22 total HTTP entries, 16 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:7a25e6af121014ac`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.cronSchedule` | **from:** `"0 7 * * *"`<br>**to:** `"0 8 * * *"` | Yes (#14 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataformReleaseConfig` | 3 | `CREATE`, `DELETE` |
| `DataformRepository` | 5 | `CREATE`, `DELETE` |
| `IAMServiceAccount` | 8 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 11 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/workflowConfigs/`<br>`dfwc-${uniqueId}` | Pre-create existence check |
| 12 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/`<br>`workflowConfigs`<br>`?workflowConfigId=dfwc-${uniqueId}` | Created resource |
| 13 | `CREATE` | `GET` `200` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/workflowConfigs/`<br>`dfwc-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 14 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/workflowConfigs/`<br>`dfwc-${uniqueId}` | updateMask: cronSchedule |
| 15 | `UPDATE` | `GET` `200` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/workflowConfigs/`<br>`dfwc-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 16 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/dfwc-repo-${uniqueId}/workflowConfigs/`<br>`dfwc-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataformworkflowconfig-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
