# HTTP Lifecycle Report: `dataplexdatascan-maximal`

- **Primary Resource:** `DataplexDataScan` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dataplexdatascan-maximal-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 10 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:b5554bdd6c95065c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.dataProfileSpec.rowFilter` | **from:** `"id IS NOT NULL"`<br>**to:** `"id IS NOT NULL AND TRUE"` | Yes (#11 `PATCH`) | Yes |
| `spec.dataProfileSpec.`<br>`.samplingPercent` | **from:** `50.5`<br>**to:** `80` | Yes (#11 `PATCH`) | Yes |
| `spec.description` | **from:** `"Maximal description"`<br>**to:** `"Maximal description updated"` | Yes (#11 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Data Scan"`<br>**to:** `"Maximal Data Scan Updated"` | Yes (#11 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryDataset` | 5 | `CREATE`, `DELETE` |
| `BigQueryTable` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/dataScans/`<br>`dataplexdatascan-maximal-${uniqueId}`<br>`?view=FULL` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/dataScans`<br>`?dataScanId=dataplexdatascan-maximal-${uniqueId}` | LRO: create -> done (#9)<br>Created resource |
| 10 | `CREATE` | `GET` `200` | `.../locations/us-central1/dataScans/`<br>`dataplexdatascan-maximal-${uniqueId}`<br>`?view=FULL` | Post-create read (matches `create.yaml`) |
| 11 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/dataScans/`<br>`dataplexdatascan-maximal-${uniqueId}` | LRO: update -> done (#12)<br>updateMask: dataProfileSpec, description, displayName |
| 13 | `UPDATE` | `GET` `200` | `.../locations/us-central1/dataScans/`<br>`dataplexdatascan-maximal-${uniqueId}`<br>`?view=FULL` | Post-update read (matches `update.yaml`) |
| 14 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/dataScans/`<br>`dataplexdatascan-maximal-${uniqueId}` | LRO: delete -> done (#15)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataplexdatascan-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
