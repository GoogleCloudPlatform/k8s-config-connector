# HTTP Lifecycle Report: `apihubapi-maximal`

- **Primary Resource:** `APIHubAPI` (`apihub.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `api-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:135def3d73d97077`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A comprehensive APIHubAPI for E2E testing"`<br>**to:** `"A comprehensive APIHubAPI for E2E testing - Updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal APIHubAPI"`<br>**to:** `"Maximal APIHubAPI Updated"` | Yes (#4 `PATCH`) | Yes |
| `spec.documentation.externalURI` | **from:** `"https://example.com/docs"`<br>**to:** `"https://example.com/docs/v2"` | Yes (#4 `PATCH`) | Yes |
| `spec.owner.displayName` | **from:** `"API Owner"`<br>**to:** `"API New Owner"` | Yes (#4 `PATCH`) | Yes |
| `spec.owner.email` | **from:** `"owner@example.com"`<br>**to:** `"new-owner@example.com"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/apis/`<br>`api-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/apis`<br>`?apiId=api-${uniqueId}` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/us-central1/apis/`<br>`api-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/apis/`<br>`api-${uniqueId}` | updateMask: description, displayName, documentation, owner |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/apis/`<br>`api-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/apis/`<br>`api-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_apihubapi-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
