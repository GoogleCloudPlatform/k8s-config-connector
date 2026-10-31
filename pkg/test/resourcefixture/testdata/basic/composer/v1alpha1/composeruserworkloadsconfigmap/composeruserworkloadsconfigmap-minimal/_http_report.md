# HTTP Lifecycle Report: `composeruserworkloadsconfigmap-minimal`

- **Primary Resource:** `ComposerUserWorkloadsConfigMap` (`composer.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `composeruserworkloadsconfigmap-${uniqueId}`
- **Source Log:** `_http.log` (36 total HTTP entries, 32 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:32973d427a1b8b83`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComposerEnvironment` | 6 | `CREATE`, `DELETE` |
| `IAMPolicy` | 8 | `CREATE` |
| `IAMServiceAccount` | 18 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 28 | `CREATE` | `GET` `404` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Pre-create existence check |
| 29 | `CREATE` | `POST` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/`<br>`userWorkloadsConfigMaps` | Created resource |
| 30 | `DELETE` | `GET` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Pre-delete read |
| 31 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/environments/composerenvironment-${uniqueId}/userWorkloadsConfigMaps/`<br>`composeruserworkloadsconfigmap-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_composeruserworkloadsconfigmap-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
