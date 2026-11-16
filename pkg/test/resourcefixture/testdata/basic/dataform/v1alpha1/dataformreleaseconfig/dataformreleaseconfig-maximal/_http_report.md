# HTTP Lifecycle Report: `dataformreleaseconfig-maximal`

- **Primary Resource:** `DataformReleaseConfig` (`dataform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataformreleaseconfig-max-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:17061ba3d43d515a`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.cronSchedule` | **from:** `"0 6 * * *"`<br>**to:** `"0 12 * * *"` | Yes (#6 `PATCH`) | Yes |
| `spec.disabled` | **from:** `false`<br>**to:** `true` | Yes (#6 `PATCH`) | Yes |
| `spec.gitCommitish` | **from:** `"main"`<br>**to:** `"branch1"` | Yes (#6 `PATCH`) | Yes |
| `spec.timeZone` | **from:** `"America/New_York"`<br>**to:** `"America/Los_Angeles"` | Yes (#6 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataformRepository` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/releaseConfigs/`<br>`dataformreleaseconfig-max-${uniqueId}` | Pre-create existence check |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/`<br>`releaseConfigs`<br>`?releaseConfigId=dataformreleaseconfig-max-${uniqueId}` | Created resource |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/releaseConfigs/`<br>`dataformreleaseconfig-max-${uniqueId}` | Pre-update read |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/releaseConfigs/`<br>`dataformreleaseconfig-max-${uniqueId}` | updateMask: cronSchedule, disabled, gitCommitish, timeZone |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/releaseConfigs/`<br>`dataformreleaseconfig-max-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/dataformrepository-max-${uniqueId}/releaseConfigs/`<br>`dataformreleaseconfig-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataformreleaseconfig-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
