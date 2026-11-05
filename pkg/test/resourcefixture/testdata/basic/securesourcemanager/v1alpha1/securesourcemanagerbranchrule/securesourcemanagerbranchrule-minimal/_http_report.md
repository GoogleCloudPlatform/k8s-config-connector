# HTTP Lifecycle Report: `securesourcemanagerbranchrule-minimal`

- **Primary Resource:** `SecureSourceManagerBranchRule` (`securesourcemanager.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ssmbranchrule-minimal-${uniqueId}`
- **Source Log:** `_http.log` (41 total HTTP entries, 35 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ed99a4548df8a735`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.disabled` | **from:** *(unset)*<br>**to:** `true` | Yes (#33 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `Project` | 1 | `CREATE` |
| `SecureSourceManagerInstance` | 6 | `CREATE`, `DELETE` |
| `SecureSourceManagerRepository` | 28 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 30 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-minimal-${uniqueId}` | Pre-create existence check |
| 31 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/`<br>`branchRules`<br>`?branchRuleId=ssmbranchrule-minimal-${uniqueId}` | LRO: create -> done (sync)<br>Created resource |
| 32 | `UPDATE` | `GET` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-minimal-${uniqueId}` | Pre-update read |
| 33 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-minimal-${uniqueId}` | LRO: update -> done (sync)<br>updateMask: disabled |
| 34 | `DELETE` | `GET` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-minimal-${uniqueId}` | Pre-delete read |
| 35 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/ssmrepository-dep-${uniqueId}/branchRules/`<br>`ssmbranchrule-minimal-${uniqueId}` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_securesourcemanagerbranchrule-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
