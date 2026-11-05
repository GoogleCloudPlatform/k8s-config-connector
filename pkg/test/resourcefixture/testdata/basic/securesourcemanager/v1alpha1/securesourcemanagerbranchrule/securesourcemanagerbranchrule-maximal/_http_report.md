# HTTP Lifecycle Report: `securesourcemanagerbranchrule-maximal`

- **Primary Resource:** `SecureSourceManagerBranchRule` (`securesourcemanager.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ssmbranchrule-maximal-${uniqueId}`
- **Source Log:** `_http.log` (85 total HTTP entries, 79 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:036ec8bf4d67d790`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.allowStaleReviews` | **from:** `false`<br>**to:** `true` | Yes (#78 `PATCH`) | Yes |
| `spec.disabled` | **from:** `false`<br>**to:** `true` | Yes (#78 `PATCH`) | Yes |
| `spec.requireCommentsResolved` | **from:** `true`<br>**to:** `false` | Yes (#78 `PATCH`) | Yes |
| `spec.requireLinearHistory` | **from:** `true`<br>**to:** `false` | Yes (#78 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `Project` | 1 | `CREATE` |
| `SecureSourceManagerInstance` | 12 | `CREATE`, `DELETE` |
| `SecureSourceManagerRepository` | 66 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 75 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-maximal-${uniqueId}` | Pre-create existence check |
| 76 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/`<br>`branchRules`<br>`?branchRuleId=ssmbranchrule-maximal-${uniqueId}` | LRO: create -> done (sync)<br>Created resource |
| 77 | `UPDATE` | `GET` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-maximal-${uniqueId}` | Pre-update read |
| 78 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-maximal-${uniqueId}` | LRO: update -> done (sync)<br>updateMask: allowStaleReviews, disabled, requireCommentsResolved, requireLinearHistory |
| 79 | `DELETE` | `GET` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-maximal-${uniqueId}` | Pre-delete read |
| 80 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-maximal-${uniqueId}` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_securesourcemanagerbranchrule-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
