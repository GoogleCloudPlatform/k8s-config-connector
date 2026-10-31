# HTTP Lifecycle Report: `composeruserworkloadsconfigmap-maximal`

- **Primary Resource:** `ComposerUserWorkloadsConfigMap` (`composer.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `composeruserworkloadsconfigmap-${uniqueId}`
- **Source Log:** `_http.log` (30 total HTTP entries, 24 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:1fc6c6dabc738d44`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.data.key1` | **from:** `"value1"`<br>**to:** `"updated-value1"` | Yes (#23 `PUT`) | Yes |
| `spec.data.key2` | **from:** `"value2"`<br>**to:** *(unset)* | Yes (#23 `PUT`) | Yes |
| `spec.data.key3` | **from:** *(unset)*<br>**to:** `"value3"` | Yes (#23 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComposerEnvironment` | 6 | `CREATE`, `DELETE` |
| `IAMPolicy` | 8 | `CREATE` |
| `IAMServiceAccount` | 10 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 20 | `CREATE` | `GET` `404` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Pre-create existence check |
| 21 | `CREATE` | `POST` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/`<br>`userWorkloadsConfigMaps` | Created resource |
| 22 | `UPDATE` | `GET` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Pre-update read |
| 23 | `UPDATE` | `PUT` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Updated resource (`PUT`) |
| 24 | `DELETE` | `GET` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Pre-delete read |
| 25 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_composeruserworkloadsconfigmap-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
