# HTTP Lifecycle Report: `globalcomputebackendservicestrongsessionaffinitycookie`

- **Primary Resource:** `ComputeBackendService` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computebackendservice-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:21612a362a3dc334`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.strongSessionAffinityCookie.`<br>`.name` | **from:** `"my-session-cookie"`<br>**to:** `"updated-session-cookie"` | Yes (#9 `PUT`) | Yes |
| `spec.strongSessionAffinityCookie.`<br>`.path` | **from:** `"/"`<br>**to:** `"/updated"` | Yes (#9 `PUT`) | Yes |
| `spec.strongSessionAffinityCookie.`<br>`.ttl.seconds` | **from:** `3600`<br>**to:** `7200` | Yes (#9 `PUT`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeHealthCheck` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../global/backendServices/`<br>`${backendServiceID}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../global/backendServices` | LRO: insert -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../global/backendServices/`<br>`${backendServiceID}` | Post-create read (matches `create.yaml`) |
| 9 | `UPDATE` | `PUT` `200` | `.../global/backendServices/`<br>`${backendServiceID}` | LRO: update -> done (#10)<br>Updated resource (`PUT`) |
| 11 | `UPDATE` | `GET` `200` | `.../global/backendServices/`<br>`${backendServiceID}` | Post-update read (matches `update.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../global/backendServices/`<br>`${backendServiceID}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PUT`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_globalcomputebackendservicestrongsessionaffinitycookie.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
