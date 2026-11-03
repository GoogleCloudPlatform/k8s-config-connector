# HTTP Lifecycle Report: `sqlinstance-retainbackupsondelete-direct`

- **Primary Resource:** `SQLInstance` (`sql.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `sqlinstance-retainbackupsondelete-direct-${uniqueId}`
- **Source Log:** `_http.log` (11 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:78a34681e65ad0c2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.settings.`<br>`.retainBackupsOnDelete` | **from:** `true`<br>**to:** `false` | Yes (#7 `PUT`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../instances` | LRO: CREATE -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 5 | `CREATE` | `GET` `200` | `.../instances/sqlinstance-retainbackupsondelete-direct-${uniqueId}/`<br>`users` | Post-create read (matches `create.yaml`) |
| 6 | `UPDATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | Pre-update read |
| 7 | `UPDATE` | `PUT` `200` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | LRO: UPDATE -> done (#8)<br>Updated resource (`PUT`) |
| 9 | `UPDATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 10 | `DELETE` | `DELETE` `200` | `.../instances/`<br>`sqlinstance-retainbackupsondelete-direct-${uniqueId}` | LRO: DELETE -> done (#11)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_sqlinstance-retainbackupsondelete-direct.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
