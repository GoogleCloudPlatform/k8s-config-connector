# HTTP Lifecycle Report: `sqlbluegreendeployment-minimal`

- **Primary Resource:** `SQLBlueGreenDeployment` (`sql.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `sqlbgd-minimal-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ce61e8d2aecad948`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `SQLInstance` | 10 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/blueGreenDeployments/`<br>`sqlbgd-minimal-${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`blueGreenDeployments`<br>`?blueGreenDeploymentId=sqlbgd-minimal-${uniqueId}` | LRO: create -> pending (sync)<br>Created resource |
| 9 | `CREATE` | `GET` `200` | `.../locations/us-central1/blueGreenDeployments/`<br>`sqlbgd-minimal-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/blueGreenDeployments/`<br>`sqlbgd-minimal-${uniqueId}` | LRO: delete -> pending (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_sqlbluegreendeployment-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
