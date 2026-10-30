# HTTP Lifecycle Report: `dataplexentry-minimal`

- **Primary Resource:** `DataplexEntry` (`dataplex.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dpe-${uniqueId}`
- **Source Log:** `_http.log` (20 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:b48e4ce2a78fdc0b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.entrySource.description` | **from:** `"minimal entry"`<br>**to:** `"updated entry"` | Yes (#10 `PATCH`) | Yes |
| `spec.entrySource.displayName` | **from:** `"Entry ${uniqueId}"`<br>**to:** `"Entry Updated ${uniqueId}"` | Yes (#10 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryGroup` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`CreateEntryGroup` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GRPC` `404` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryType` | Pre-create existence check |
| 5 | `CREATE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`CreateEntryType` | LRO: create -> done (#6)<br>Composite create follow-up (`GRPC`) |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/entries/`<br>`dpe-${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/`<br>`entries`<br>`?entryId=dpe-${uniqueId}` | Composite create follow-up (`POST`) |
| 9 | `CREATE` | `GET` `200` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/entries/`<br>`dpe-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 10 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/entries/`<br>`dpe-${uniqueId}` | updateMask: entrySource |
| 11 | `UPDATE` | `GET` `200` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/entries/`<br>`dpe-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/entryGroups/dpeg-${uniqueId}/entries/`<br>`dpe-${uniqueId}` | Deleted resource |
| 13 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryType` | Pre-delete read |
| 14 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryType` | Pre-delete read |
| 15 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`DeleteEntryType` | LRO: delete -> done (#16)<br>Deleted resource |
| 17 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryGroup` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`GetEntryGroup` | Pre-delete read |
| 19 | `DELETE` | `GRPC` `200` | `/google.cloud.dataplex.v1.CatalogService/`<br>`DeleteEntryGroup` | LRO: delete -> done (#20)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (3 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
