# HTTP Lifecycle Report: `resourcemanagertaghold-maximal`

- **Primary Resource:** `ResourceManagerTagHold` (`resourcemanager.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `resourcemanagertaghold-${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:3f4ea0000a832ab9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `TagsTagKey` | 6 | `CREATE`, `DELETE` |
| `TagsTagValue` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `POST` `200` | `/v3/tagValues/${tagValueID}/tagHolds` | LRO: create -> done (sync)<br>Created resource |
| 8 | `DELETE` | `GET` `200` | `/v3/tagValues/${tagValueID}/tagHolds` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `/v3/tagValues/${tagValueID}/tagHolds/`<br>`d33e9c4a-abe6-4847-a568-6433eb67dfb6` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_resourcemanagertaghold-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
