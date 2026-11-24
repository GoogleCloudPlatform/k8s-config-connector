# HTTP Lifecycle Report: `discoveryenginewidgetconfig-maximal`

- **Primary Resource:** `DiscoveryEngineWidgetConfig` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `widgetconfig-${uniqueId} (resourceID: default_search_widget_config)`
- **Source Log:** `_http.log` (13 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:a704cb3021166080`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.accessSettings.`<br>`.allowPublicAccess` | **from:** `false`<br>**to:** `true` | Yes (#8 `PATCH`) | Yes |
| `spec.accessSettings.`<br>`.allowlistedDomains` | **from:** `["example.com"]`<br>**to:** `["updated.example.com"]` | Yes (#8 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Widget Config Maximal"`<br>**to:** `"Updated Widget Config Maximal"` | Yes (#8 `PATCH`) | Yes |
| `spec.enableResultScore` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.uiBranding.logo.url` | **from:** `"https://example.com/logo.png"`<br>**to:** `"https://updated.example.com/logo.png"` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.`<br>`.disableUserEventsCollection` | **from:** `false`<br>**to:** `true` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.`<br>`.enableAutocomplete` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.`<br>`.enableQualityFeedback` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.enableSafeSearch` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.`<br>`.enableSearchAsYouType` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.`<br>`.generativeAnswerConfig.`<br>`.resultCount` | **from:** `5`<br>**to:** `3` | Yes (#8 `PATCH`) | Yes |
| `spec.uiSettings.interactionType` | **from:** `"SEARCH_ONLY"`<br>**to:** `"SEARCH_WITH_ANSWER"` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DiscoveryEngineDataStore` | 4 | `CREATE`, `DELETE` |
| `DiscoveryEngineEngine` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/widgetConfigs/`<br>`default_search_widget_config` | Pre-create read |
| 6 | `CREATE` | `PATCH` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/widgetConfigs/`<br>`default_search_widget_config` | Created resource |
| 7 | `UPDATE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/widgetConfigs/`<br>`default_search_widget_config` | Pre-update read |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/widgetConfigs/`<br>`default_search_widget_config` | updateMask: accessSettings, displayName, enableResultScore, uiBranding, uiSettings |
| 9 | `DELETE` | `GET` `200` | `.../locations/global/collections/default_collection/dataStores/datastore-${uniqueId}/widgetConfigs/`<br>`default_search_widget_config` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 12/12 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryenginewidgetconfig-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
