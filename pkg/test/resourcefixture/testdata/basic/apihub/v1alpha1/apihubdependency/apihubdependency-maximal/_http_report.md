# HTTP Lifecycle Report: `apihubdependency-maximal`

- **Primary Resource:** `APIHubDependency` (`apihub.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dep-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:708e3e8d628ef2fc`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal APIHubDependency"`<br>**to:** `"Maximal APIHubDependency Updated"` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `APIHubExternalAPI` | 8 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/us-central1/dependencies/`<br>`dep-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../locations/us-central1/dependencies`<br>`?dependencyId=dep-${uniqueId}` | Created resource |
| 7 | `CREATE` | `GET` `200` | `.../locations/us-central1/dependencies/`<br>`dep-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/dependencies/`<br>`dep-${uniqueId}` | updateMask: description |
| 9 | `UPDATE` | `GET` `200` | `.../locations/us-central1/dependencies/`<br>`dep-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/dependencies/`<br>`dep-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_apihubdependency-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
