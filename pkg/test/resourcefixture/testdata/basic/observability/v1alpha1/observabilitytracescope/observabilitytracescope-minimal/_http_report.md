# HTTP Lifecycle Report: `observabilitytracescope-minimal`

- **Primary Resource:** `ObservabilityTraceScope` (`observability.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `observabilitytracescope-min-${uniqueId}`
- **Source Log:** `_http.log` (4 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e4dcdf9fe5fbab78`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/traceScopes/`<br>`observabilitytracescope-min-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/traceScopes`<br>`?traceScopeId=observabilitytracescope-min-${uniqueId}` | Created resource |
| 3 | `DELETE` | `GET` `200` | `.../locations/global/traceScopes/`<br>`observabilitytracescope-min-${uniqueId}` | Pre-delete read |
| 4 | `DELETE` | `DELETE` `200` | `.../locations/global/traceScopes/`<br>`observabilitytracescope-min-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_observabilitytracescope-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
