# HTTP Lifecycle Report: `computeautoscalerbasic`

- **Primary Resource:** `ComputeAutoscaler` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeautoscaler-${uniqueId}`
- **Source Log:** `_http.log` (47 total HTTP entries, 26 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:81cd06839e9f73d8`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.autoscalingPolicy.`<br>`.cooldownPeriod` | **from:** `60`<br>**to:** `90` | Yes (#33 `PATCH`) | Yes |
| `spec.autoscalingPolicy.`<br>`.cpuUtilization.target` | **from:** `0.6`<br>**to:** `0.5` | Yes (#33 `PATCH`) | Yes |
| `spec.autoscalingPolicy.`<br>`.maxReplicas` | **from:** `5`<br>**to:** `6` | Yes (#33 `PATCH`) | Yes |
| `spec.autoscalingPolicy.`<br>`.minReplicas` | **from:** `1`<br>**to:** `2` | Yes (#33 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeHealthCheck` | 7 | `CREATE`, `DELETE` |
| `ComputeInstanceGroupManager` | 8 | `CREATE`, `DELETE` |
| `ComputeInstanceTemplate` | 7 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 4 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 10 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 11 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 12 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 13 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 14 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 15 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 19 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 20 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 21 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 22 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 23 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 24 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 29 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/autoscalers/`<br>`${autoscalerID}` | Pre-create existence check |
| 30 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/autoscalers` | LRO: compute.autoscalers.insert -> done (#31)<br>Created resource |
| 32 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/autoscalers/`<br>`${autoscalerID}` | Post-create read (matches `create.yaml`) |
| 33 | `UPDATE` | `PATCH` `200` | `.../zones/us-central1-a/autoscalers`<br>`?autoscaler=${autoscalerID}` | LRO: compute.autoscalers.patch -> done (#34)<br>Updated resource (`PATCH`) |
| 35 | `UPDATE` | `GET` `200` | `.../zones/us-central1-a/autoscalers/`<br>`${autoscalerID}` | Post-update read (matches `update.yaml`) |
| 36 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/autoscalers/`<br>`${autoscalerID}` | LRO: delete -> done (#37)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeautoscalerbasic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
