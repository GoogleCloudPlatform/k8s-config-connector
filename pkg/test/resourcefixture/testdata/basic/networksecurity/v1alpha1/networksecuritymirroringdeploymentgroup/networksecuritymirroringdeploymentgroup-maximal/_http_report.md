# HTTP Lifecycle Report: `networksecuritymirroringdeploymentgroup-maximal`

- **Primary Resource:** `NetworkSecurityMirroringDeploymentGroup` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `networksecuritymirroringdeploymentgroup-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:26f7537dadb2c6d2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A sample mirroring deployment group"`<br>**to:** `"An updated sample mirroring deployment group"` | Yes (#9 `PATCH`) | Yes |
| `spec.labels.env` | **from:** `"test"`<br>**to:** `"updated-test"` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/global/mirroringDeploymentGroups/`<br>`networksecuritymirroringdeploymentgroup-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../locations/global/`<br>`mirroringDeploymentGroups`<br>`?mirroringDeploymentGroupId=networksecuritymirroringdeploymentgroup-${uniqueId}` | LRO: create -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../locations/global/mirroringDeploymentGroups/`<br>`networksecuritymirroringdeploymentgroup-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 9 | `UPDATE` | `PATCH` `200` | `.../locations/global/mirroringDeploymentGroups/`<br>`networksecuritymirroringdeploymentgroup-${uniqueId}` | LRO: update -> done (#10)<br>updateMask: description, labels |
| 11 | `UPDATE` | `GET` `200` | `.../locations/global/mirroringDeploymentGroups/`<br>`networksecuritymirroringdeploymentgroup-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/global/mirroringDeploymentGroups/`<br>`networksecuritymirroringdeploymentgroup-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networksecuritymirroringdeploymentgroup-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
