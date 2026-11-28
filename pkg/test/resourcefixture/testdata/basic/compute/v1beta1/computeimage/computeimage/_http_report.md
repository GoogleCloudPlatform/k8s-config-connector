# HTTP Lifecycle Report: `computeimage`

- **Primary Resource:** `ComputeImage` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computeimage-test-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ce0a47dadba79a58`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.label-one` | **from:** *(unset)*<br>**to:** `"value-one"` | Yes (#9 `POST`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeDisk` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../global/images/${imageID}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../global/images` | LRO: insert -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../global/images/${imageID}` | Post-create read (matches `create.yaml`) |
| 9 | `UPDATE` | `POST` `200` | `.../global/images/${imageID}/setLabels` | LRO: setLabels -> done (#10)<br>Updated resource (`POST`) |
| 11 | `UPDATE` | `GET` `200` | `.../global/images/${imageID}` | Post-update read (matches `update.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../global/images/${imageID}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`POST`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeimage.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
