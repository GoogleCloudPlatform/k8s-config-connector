# HTTP Lifecycle Report: `networksecuritymirroringdeployment-minimal`

- **Primary Resource:** `NetworkSecurityMirroringDeployment` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mirroringdep-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ea9e5906df35501d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** *(unset)*<br>**to:** `"Updated description for minimal mirroring deployment"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1-a/`<br>`mirroringDeployments`<br>`?mirroringDeploymentId=mirroringdep-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-${uniqueId}` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: description |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1-a/mirroringDeployments/`<br>`mirroringdep-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networksecuritymirroringdeployment-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
