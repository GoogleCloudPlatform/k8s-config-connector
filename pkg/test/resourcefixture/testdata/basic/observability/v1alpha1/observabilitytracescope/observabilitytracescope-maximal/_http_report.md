# HTTP Lifecycle Report: `observabilitytracescope-maximal`

- **Primary Resource:** `ObservabilityTraceScope` (`observability.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `tracescope-max-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2bcbc98a98fb4470`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description for observabilitytracescope-maximal"`<br>**to:** `"Updated description for observabilitytracescope-maximal"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/traceScopes/`<br>`tracescope-max-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/traceScopes`<br>`?traceScopeId=tracescope-max-${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/global/traceScopes/`<br>`tracescope-max-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/traceScopes/`<br>`tracescope-max-${uniqueId}` | updateMask: description |
| 5 | `DELETE` | `GET` `200` | `.../locations/global/traceScopes/`<br>`tracescope-max-${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/global/traceScopes/`<br>`tracescope-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_observabilitytracescope-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
