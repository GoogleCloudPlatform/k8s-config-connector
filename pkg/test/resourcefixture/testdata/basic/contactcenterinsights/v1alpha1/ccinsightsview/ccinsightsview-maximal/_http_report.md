# HTTP Lifecycle Report: `ccinsightsview-maximal`

- **Primary Resource:** `CCInsightsView` (`contactcenterinsights.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ccinsightsview-maximal-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:705eb7a823aad4e6`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"My CCInsights View"`<br>**to:** `"My Updated CCInsights View"` | Yes (#4 `PATCH`) | Yes |
| `spec.value` | **from:** `"agent_id=\"my-agent\""`<br>**to:** `"agent_id=\"updated-agent\""` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/views/`<br>`ccinsightsview-maximal-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/views` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/us-central1/views/`<br>`${viewId}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/views/`<br>`${viewId}` | updateMask: displayName, value |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/views/`<br>`${viewId}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/views/`<br>`${viewId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ccinsightsview-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
