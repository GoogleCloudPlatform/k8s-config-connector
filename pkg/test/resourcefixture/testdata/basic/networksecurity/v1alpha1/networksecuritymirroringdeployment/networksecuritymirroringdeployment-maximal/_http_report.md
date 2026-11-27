# HTTP Lifecycle Report: `networksecuritymirroringdeployment-maximal`

- **Primary Resource:** `NetworkSecurityMirroringDeployment` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mirroringdep-max-${uniqueId}`
- **Source Log:** `_http.log` (46 total HTTP entries, 37 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:fb406df545b1f197`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal mirroring deployment description"`<br>**to:** `"Updated maximal mirroring deployment description"` | Yes (#27 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeBackendService` | 7 | `CREATE`, `DELETE` |
| `ComputeForwardingRule` | 9 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 9 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 5 | `CREATE`, `DELETE` |
| `NetworkSecurityMirroringDeploymentGroup` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 23 | `CREATE` | `GET` `404` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-max-${uniqueId}` | Pre-create existence check |
| 24 | `CREATE` | `POST` `200` | `.../locations/us-central1-a/`<br>`mirroringDeployments`<br>`?mirroringDeploymentId=mirroringdep-max-${uniqueId}` | LRO: create -> done (#25)<br>Created resource |
| 26 | `UPDATE` | `GET` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-max-${uniqueId}` | Pre-update read |
| 27 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-max-${uniqueId}` | LRO: update -> done (#28)<br>updateMask: description |
| 29 | `DELETE` | `GET` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-max-${uniqueId}` | Pre-delete read |
| 30 | `DELETE` | `DELETE` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-max-${uniqueId}` | LRO: delete -> done (#31)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networksecuritymirroringdeployment-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
