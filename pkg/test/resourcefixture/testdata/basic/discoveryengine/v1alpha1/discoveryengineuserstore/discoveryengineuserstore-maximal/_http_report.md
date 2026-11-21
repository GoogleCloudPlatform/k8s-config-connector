# HTTP Lifecycle Report: `discoveryengineuserstore-maximal`

- **Primary Resource:** `DiscoveryEngineUserStore` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `userstore-max-${uniqueId} (resourceID: default_user_store)`
- **Source Log:** `_http.log` (8 total HTTP entries, 3 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:45b68e74ab289444`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"My User Store"`<br>**to:** `"My Updated User Store"` | Yes (#6 `PATCH`) | Yes |
| `spec.`<br>`.enableExpiredLicenseAutoUpdate` | **from:** `true`<br>**to:** `false` | Yes (#6 `PATCH`) | Yes |
| `spec.enableLicenseAutoRegister` | **from:** `true`<br>**to:** `false` | Yes (#6 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DiscoveryEngineLicenseConfig` | 3 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `200` | `.../locations/global/userStores/`<br>`default_user_store` | Pre-create read |
| 4 | `CREATE` | `PATCH` `200` | `.../locations/global/userStores/`<br>`default_user_store` | Created resource |
| 5 | `UPDATE` | `GET` `200` | `.../locations/global/userStores/`<br>`default_user_store` | Pre-update read |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/global/userStores/`<br>`default_user_store` | updateMask: displayName, enableExpiredLicenseAutoUpdate, enableLicenseAutoRegister |
| 7 | `DELETE` | `GET` `200` | `.../locations/global/userStores/`<br>`default_user_store` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryengineuserstore-maximal.golden.yaml` match `update.yaml`.
- [ ] **V4 (Delete):** No `DELETE` RPC recorded for primary resource in `DELETE` phase.
  - `<hint_for_agent>`: MISSING DELETE RPC: The primary resource was not deleted in _http.log. Ensure DeleteResources() ran and the DELETE stanza was not truncated or hand-deleted.
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
