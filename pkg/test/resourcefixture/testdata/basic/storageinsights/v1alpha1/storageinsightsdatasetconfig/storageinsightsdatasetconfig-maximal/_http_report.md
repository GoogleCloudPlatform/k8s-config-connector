# HTTP Lifecycle Report: `storageinsightsdatasetconfig-maximal`

- **Primary Resource:** `StorageInsightsDatasetConfig` (`storageinsights.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `datasetconfigmax${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ebd680c1a53f0303`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal Storage Insights Dataset Config"`<br>**to:** `"Updated Maximal Storage Insights Dataset Config"` | Yes (#8 `PATCH`) | Yes |
| `spec.includeNewlyCreatedBuckets` | **from:** `true`<br>**to:** `false` | Yes (#8 `PATCH`) | Yes |
| `spec.retentionPeriodDays` | **from:** `30`<br>**to:** `14` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/datasetConfigs/`<br>`datasetconfigmax${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`datasetConfigs`<br>`?datasetConfigId=datasetconfigmax${uniqueId}` | LRO: create -> done (#6)<br>Created resource |
| 7 | `CREATE` | `GET` `200` | `.../locations/us-central1/datasetConfigs/`<br>`datasetconfigmax${uniqueId}` | Post-create read (matches `create.yaml`) |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/datasetConfigs/`<br>`datasetconfigmax${uniqueId}` | LRO: update -> done (#9)<br>updateMask: description, includeNewlyCreatedBuckets, retentionPeriodDays |
| 10 | `UPDATE` | `GET` `200` | `.../locations/us-central1/datasetConfigs/`<br>`datasetconfigmax${uniqueId}` | Post-update read (matches `update.yaml`) |
| 11 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/datasetConfigs/`<br>`datasetconfigmax${uniqueId}` | LRO: delete -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_storageinsightsdatasetconfig-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
