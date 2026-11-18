# HTTP Lifecycle Report: `sqlinstance-psc-networkattachment-direct`

- **Primary Resource:** `SQLInstance` (`sql.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `sqlinstance-psc-na-direct-${uniqueId}`
- **Source Log:** `_http.log` (34 total HTTP entries, 23 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:444ab6a045902a90`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.settings.ipConfiguration.`<br>`.pscConfig` | **from:** `[{"allowedConsumerProjects": ["${projectId}"], "pscEnabled": ...`<br>**to:** `[{"allowedConsumerProjects": ["${projectId}"], "networkAttach...` | Yes (#19 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 10 | `CREATE`, `DELETE` |
| `ComputeNetworkAttachment` | 7 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 6 | `CREATE`, `UPDATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 13 | `CREATE` | `GET` `404` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | Pre-create existence check |
| 14 | `CREATE` | `POST` `200` | `.../instances` | LRO: CREATE -> done (#15)<br>Created resource |
| 16 | `CREATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 17 | `CREATE` | `GET` `200` | `.../instances/sqlinstance-psc-na-direct-${uniqueId}/`<br>`users` | Post-create read (matches `create.yaml`) |
| 18 | `UPDATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | Pre-update read |
| 19 | `UPDATE` | `PUT` `200` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | LRO: UPDATE -> done (#21)<br>Updated resource (`PUT`) |
| 22 | `UPDATE` | `GET` `200` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 23 | `DELETE` | `DELETE` `200` | `.../instances/`<br>`sqlinstance-psc-na-direct-${uniqueId}` | LRO: DELETE -> done (#25)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_sqlinstance-psc-networkattachment-direct.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
