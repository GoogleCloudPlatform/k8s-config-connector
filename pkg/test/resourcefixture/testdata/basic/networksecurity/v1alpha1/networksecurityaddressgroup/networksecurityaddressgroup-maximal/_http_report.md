# HTTP Lifecycle Report: `networksecurityaddressgroup-maximal`

- **Primary Resource:** `NetworkSecurityAddressGroup` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `addressgroup-max-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2d5ff09d95597051`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A maximal address group test"`<br>**to:** `"An updated maximal address group test"` | Yes (#5 `PATCH`) | Yes |
| `spec.items` | **from:** `["10.0.0.0/24"]`<br>**to:** `["10.0.0.0/24", "10.0.1.0/24"]` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.foo` | **from:** `"bar"`<br>**to:** `"baz"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/addressGroups/`<br>`addressgroup-max-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/addressGroups`<br>`?addressGroupId=addressgroup-max-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/global/addressGroups/`<br>`addressgroup-max-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/global/addressGroups/`<br>`addressgroup-max-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: description, items, labels |
| 7 | `UPDATE` | `GET` `200` | `.../locations/global/addressGroups/`<br>`addressgroup-max-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/global/addressGroups/`<br>`addressgroup-max-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networksecurityaddressgroup-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
