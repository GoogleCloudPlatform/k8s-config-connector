# HTTP Lifecycle Report: `saasservicemgmtsaas-maximal`

- **Primary Resource:** `SaaSServiceMgmtSaaS` (`saasservicemgmt.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `saasservicemgmtsaas-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:1da38d684d2ad92e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.locations` | **from:** `[{"name": "us-central1"}]`<br>**to:** `[{"name": "us-central1"}, {"name": "us-east1"}]` | Yes (#4 `PATCH`) | Yes |
| `metadata.labels.env` | **from:** `"dev"`<br>**to:** `"prod"` | Yes (#4 `PATCH`) | Yes |
| `metadata.labels.team` | **from:** `"alpha"`<br>**to:** `"beta"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/saas/`<br>`saasservicemgmtsaas-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/saas`<br>`?saasId=saasservicemgmtsaas-${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/us-central1/saas/`<br>`saasservicemgmtsaas-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/saas/`<br>`saasservicemgmtsaas-${uniqueId}` | updateMask: labels, locations |
| 5 | `DELETE` | `GET` `200` | `.../locations/us-central1/saas/`<br>`saasservicemgmtsaas-${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/saas/`<br>`saasservicemgmtsaas-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_saasservicemgmtsaas-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
