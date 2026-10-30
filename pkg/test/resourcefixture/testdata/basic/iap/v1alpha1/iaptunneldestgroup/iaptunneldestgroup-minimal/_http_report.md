# HTTP Lifecycle Report: `iaptunneldestgroup-minimal`

- **Primary Resource:** `IAPTunnelDestGroup` (`iap.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `iaptunneldestgroup-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:3ab9809f805fc657`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.cidrs` | **from:** `["10.10.0.0/24"]`<br>**to:** `["10.10.0.0/24", "10.10.1.0/24"]` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/destGroups/`<br>`iaptunneldestgroup-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/destGroups`<br>`?tunnelDestGroupId=iaptunneldestgroup-${uniqueId}` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/us-central1/destGroups/`<br>`iaptunneldestgroup-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/destGroups/`<br>`iaptunneldestgroup-${uniqueId}` | updateMask: cidrs |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/destGroups/`<br>`iaptunneldestgroup-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/destGroups/`<br>`iaptunneldestgroup-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_iaptunneldestgroup-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
