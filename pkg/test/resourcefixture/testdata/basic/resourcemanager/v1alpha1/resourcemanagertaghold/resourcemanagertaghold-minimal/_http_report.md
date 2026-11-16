# HTTP Lifecycle Report: `resourcemanagertaghold-minimal`

- **Primary Resource:** `ResourceManagerTagHold` (`resourcemanager.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `resourcemanagertaghold-${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e283535cc9125794`

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
| 9 | `DELETE` | `DELETE` `200` | `/v3/tagValues/${tagValueID}/tagHolds/`<br>`1a169d35-7132-4512-a405-28ed1dff13ac` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_resourcemanagertaghold-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
