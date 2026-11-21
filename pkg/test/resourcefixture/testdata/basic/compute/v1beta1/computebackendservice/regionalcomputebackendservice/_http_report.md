# HTTP Lifecycle Report: `regionalcomputebackendservice`

- **Primary Resource:** `ComputeBackendService` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computeregionbackendservice-${uniqueId}`
- **Source Log:** `_http.log` (43 total HTTP entries, 31 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2f3457fdd33857ad`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.timeoutSec` | **from:** `10`<br>**to:** `5` | Yes (#28 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeHealthCheck` | 7 | `CREATE`, `DELETE` |
| `ComputeInstance` | 15 | `CREATE`, `DELETE` |
| `ComputeInstanceGroup` | 8 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 4 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 5 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 24 | `CREATE` | `GET` `404` | `.../regions/us-central1/backendServices/`<br>`${backendServiceID}` | Pre-create existence check |
| 25 | `CREATE` | `POST` `200` | `.../regions/us-central1/`<br>`backendServices` | LRO: insert -> done (#26)<br>Created resource |
| 27 | `CREATE` | `GET` `200` | `.../regions/us-central1/backendServices/`<br>`${backendServiceID}` | Post-create read (matches `create.yaml`) |
| 28 | `UPDATE` | `PUT` `200` | `.../regions/us-central1/backendServices/`<br>`${backendServiceID}` | LRO: update -> done (#29)<br>Updated resource (`PUT`) |
| 30 | `UPDATE` | `GET` `200` | `.../regions/us-central1/backendServices/`<br>`${backendServiceID}` | Post-update read (matches `update.yaml`) |
| 31 | `DELETE` | `DELETE` `200` | `.../regions/us-central1/backendServices/`<br>`${backendServiceID}` | LRO: delete -> done (#32)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_regionalcomputebackendservice.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
