# HTTP Lifecycle Report: `mapmanagementmapcontextconfig-minimal`

- **Primary Resource:** `MapmanagementMapContextConfig` (`mapmanagement.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mmmcc-${uniqueId}`
- **Source Log:** `_http.log` (13 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e1549128b3cb52e7`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.alias` | **from:** `"alias-${uniqueId}"`<br>**to:** `"alias-upd-${uniqueId}"` | Yes (#7 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `MapManagementMapConfig` | 4 | `CREATE`, `DELETE` |
| `MapManagementStyleConfig` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `POST` `200` | `.../mapConfigs/63bce22e5a9e9aa653af78ef/`<br>`mapContextConfigs` | Created resource |
| 6 | `CREATE` | `GET` `200` | `.../mapConfigs/63bce22e5a9e9aa653af78ef/mapContextConfigs/`<br>`1792890820903793` | Post-create read (matches `create.yaml`) |
| 7 | `UPDATE` | `PATCH` `200` | `.../mapConfigs/63bce22e5a9e9aa653af78ef/mapContextConfigs/`<br>`1792890820903793` | updateMask: alias |
| 8 | `UPDATE` | `GET` `200` | `.../mapConfigs/63bce22e5a9e9aa653af78ef/mapContextConfigs/`<br>`1792890820903793` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `DELETE` `200` | `.../mapConfigs/63bce22e5a9e9aa653af78ef/mapContextConfigs/`<br>`1792890820903793` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_mapmanagementmapcontextconfig-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
