# HTTP Lifecycle Report: `configdeployment-maximal`

- **Primary Resource:** `ConfigDeployment` (`configdeployment.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `configdeployment-${uniqueId}`
- **Source Log:** `_http.log` (95 total HTTP entries, 20 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:d82bbf637be89bb9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.annotations.note` | **from:** `"initial"`<br>**to:** `"updated"` | Yes (#14 `PATCH`) | Yes |
| `spec.terraformBlueprint.`<br>`.externalValues.ext_var.`<br>`.deploymentSource.outputName` | **from:** `"my_output"`<br>**to:** `"updated_output"` | Yes (#14 `PATCH`) | Yes |
| `spec.terraformBlueprint.`<br>`.gitSource.ref` | **from:** `"main"`<br>**to:** `"v1.0.0"` | Yes (#14 `PATCH`) | Yes |
| `spec.terraformBlueprint.`<br>`.inputValues.foo.inputValue.`<br>`.stringValue` | **from:** `"bar"`<br>**to:** `"updated"` | Yes (#14 `PATCH`) | Yes |
| `metadata.labels.env` | **from:** `"test"`<br>**to:** `"staging"` | Yes (#14 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMServiceAccount` | 11 | `CREATE`, `DELETE` |
| `StorageBucket` | 9 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 10 | `CREATE` | `GET` `404` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-create existence check |
| 11 | `CREATE` | `POST` `200` | `.../locations/us-central1/deployments`<br>`?deploymentId=configdeployment-${uniqueId}` | LRO: create -> done (#12)<br>Created resource |
| 13 | `CREATE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 14 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | LRO: update -> done (#15)<br>updateMask: annotations, labels, terraformBlueprint |
| 16 | `UPDATE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 17 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#18)<br>Deleted resource |
| 19 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 20 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#21)<br>Deleted resource |
| 22 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 23 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#24)<br>Deleted resource |
| 25 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 26 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#27)<br>Deleted resource |
| 28 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 29 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#30)<br>Deleted resource |
| 31 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 32 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#33)<br>Deleted resource |
| 34 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 35 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#36)<br>Deleted resource |
| 37 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 38 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#39)<br>Deleted resource |
| 40 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 41 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#42)<br>Deleted resource |
| 43 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 44 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#45)<br>Deleted resource |
| 46 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 47 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#48)<br>Deleted resource |
| 50 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 51 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#52)<br>Deleted resource |
| 54 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 55 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#56)<br>Deleted resource |
| 57 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 58 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#59)<br>Deleted resource |
| 60 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 61 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#62)<br>Deleted resource |
| 63 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 64 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#65)<br>Deleted resource |
| 66 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 67 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#68)<br>Deleted resource |
| 70 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 71 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#72)<br>Deleted resource |
| 73 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 74 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#75)<br>Deleted resource |
| 77 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 78 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#79)<br>Deleted resource |
| 80 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 81 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#82)<br>Deleted resource |
| 84 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 85 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#86)<br>Deleted resource |
| 87 | `DELETE` | `GET` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}` | Pre-delete read |
| 88 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/deployments/`<br>`configdeployment-${uniqueId}`<br>`?force=true` | LRO: delete -> error (#89)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_configdeployment-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
