# HTTP Lifecycle Report: `clouddeployrollout-maximal`

- **Primary Resource:** `CloudDeployRollout` (`clouddeploy.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `clouddeployrollout-max-${uniqueId}`
- **Source Log:** `_http.log` (41 total HTTP entries, 37 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:67f3d59d4cc362f2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `CloudDeployDeliveryPipeline` | 27 | `CREATE`, `DELETE` |
| `CloudDeployRelease` | 4 | `CREATE`, `DELETE` |
| `CloudDeployTarget` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 10 | `CREATE` | `GET` `404` | `.../locations/us-central1/deliveryPipelines/clouddeploydeliverypipeline-${uniqueId}/releases/clouddeployrelease-${uniqueId}/rollouts/`<br>`clouddeployrollout-max-${uniqueId}` | Pre-create existence check |
| 11 | `CREATE` | `POST` `200` | `.../locations/us-central1/deliveryPipelines/clouddeploydeliverypipeline-${uniqueId}/releases/clouddeployrelease-${uniqueId}/`<br>`rollouts`<br>`?rolloutId=clouddeployrollout-max-${uniqueId}` | LRO: create -> done (#12)<br>Created resource |
| 13 | `DELETE` | `GET` `200` | `.../locations/us-central1/deliveryPipelines/clouddeploydeliverypipeline-${uniqueId}/releases/clouddeployrelease-${uniqueId}/rollouts/`<br>`clouddeployrollout-max-${uniqueId}` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_clouddeployrollout-maximal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
