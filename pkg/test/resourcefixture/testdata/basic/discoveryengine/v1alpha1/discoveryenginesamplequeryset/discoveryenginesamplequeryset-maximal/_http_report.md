# HTTP Lifecycle Report: `discoveryenginesamplequeryset-maximal`

- **Primary Resource:** `DiscoveryEngineSampleQuerySet` (`discoveryengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `discoveryenginesamplequeryset-${uniqueId} (resourceID: sqs-${uniqueId})`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:d90d4c540e198f3f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"Maximal Sample Query Set"`<br>**to:** `"Maximal Sample Query Set Updated"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/sampleQuerySets/`<br>`sqs-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/sampleQuerySets`<br>`?sampleQuerySetId=sqs-${uniqueId}` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/global/sampleQuerySets/`<br>`sqs-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/global/sampleQuerySets/`<br>`sqs-${uniqueId}` | updateMask: displayName |
| 5 | `UPDATE` | `GET` `200` | `.../locations/global/sampleQuerySets/`<br>`sqs-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/global/sampleQuerySets/`<br>`sqs-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_discoveryenginesamplequeryset-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
