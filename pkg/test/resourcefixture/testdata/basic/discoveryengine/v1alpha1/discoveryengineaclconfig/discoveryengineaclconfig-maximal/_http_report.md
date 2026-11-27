# HTTP Lifecycle Report: `discoveryengineaclconfig-maximal`

- **Primary Resource:** `DiscoveryEngineACLConfig` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `discoveryengineaclconfig-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:526ea92271113d8f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.idpConfig.externalIdpConfig.`<br>`.workforcePoolName` | **from:** `"locations/global/workforcePools/dummy-pool"`<br>**to:** `"locations/global/workforcePools/test-pool"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `200` | `.../locations/global/aclConfig` | Pre-create read |
| 2 | `CREATE` | `PATCH` `200` | `.../locations/global/aclConfig` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/global/aclConfig` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/aclConfig` | Updated resource (`PATCH`) |
| 5 | `DELETE` | `GET` `200` | `.../locations/global/aclConfig` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryengineaclconfig-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
