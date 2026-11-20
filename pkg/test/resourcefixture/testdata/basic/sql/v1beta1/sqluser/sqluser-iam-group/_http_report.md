# HTTP Lifecycle Report: `sqluser-iam-group`

- **Primary Resource:** `SQLUser` (`sql.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `sqluser-${uniqueId} (resourceID: iam-group-auth-test-group@GOOGLE.COM)`
- **Source Log:** `_http.log` (17 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:a517c6e068933677`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `SQLInstance` | 10 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users` | Pre-create read |
| 6 | `CREATE` | `DELETE` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users`<br>`?host=%&name=root` | LRO: DELETE_USER -> done (sync)<br>Created resource |
| 7 | `CREATE` | `GET` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users` | Post-create read (matches `create.yaml`) |
| 9 | `CREATE` | `POST` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users` | LRO: CREATE_USER -> done (sync)<br>Composite create follow-up (`POST`) |
| 10 | `CREATE` | `GET` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users` | Post-create read (matches `create.yaml`) |
| 12 | `DELETE` | `GET` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users` | Pre-delete read |
| 14 | `DELETE` | `DELETE` `200` | `.../instances/sqluser-dep-${uniqueId}/`<br>`users`<br>`?host=%&name=iam-group-auth-test-group@google.com` | LRO: DELETE_USER -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_sqluser-iam-group.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
