# HTTP Lifecycle Report: `configdeployment-minimal`

- **Primary Resource:** `ConfigDeployment` (`configdeployment.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `configdeployment-${uniqueId}`
- **Source Log:** `_http.log` (24 total HTTP entries, 15 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:11fb34204c926a18`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.terraformBlueprint.`<br>`.gcsSource.object` | **from:** `"blueprint.zip"`<br>**to:** `"blueprint-v2.zip"` | Yes (#14 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMServiceAccount` | 7 | `CREATE`, `DELETE` |
| `StorageBucket` | 8 | `CREATE`, `UPDATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `404` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-create existence check |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/deployments`<br>`?deploymentId=configdeployment-${uniqueId}` | LRO: create -> done (#12)<br>Created resource |
| 13 | `UPDATE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-update read |
| 14 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | LRO: update -> done (#16)<br>updateMask: terraformBlueprint |
| 17 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 18 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> done (#19)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_configdeployment-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
