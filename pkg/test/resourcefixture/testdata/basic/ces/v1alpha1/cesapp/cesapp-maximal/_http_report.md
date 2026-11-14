# HTTP Lifecycle Report: `cesapp-maximal`

- **Primary Resource:** `CESApp` (`ces.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cesapp-maximal-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:521c68a3dc7f4554`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal test fixture for CESApp"`<br>**to:** `"Updated maximal test fixture for CESApp"` | Yes (#5 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"cesapp-maximal-${uniqueId}"`<br>**to:** `"cesapp-max-updated-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |
| `spec.globalInstruction` | **from:** `"You are a helpful customer service representative."`<br>**to:** `"You are an updated customer service representative."` | Yes (#5 `PATCH`) | Yes |
| `spec.metadata.env` | **from:** `"test"`<br>**to:** `"prod"` | Yes (#5 `PATCH`) | Yes |
| `spec.modelSettings.temperature` | **from:** `0.7`<br>**to:** `0.5` | Yes (#5 `PATCH`) | Yes |
| `spec.pinned` | **from:** `true`<br>**to:** `false` | Yes (#5 `PATCH`) | Yes |
| `spec.timeZoneSettings.timeZone` | **from:** `"America/New_York"`<br>**to:** `"America/Chicago"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us/apps/`<br>`cesapp-maximal-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us/apps`<br>`?appId=cesapp-maximal-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/us/apps/`<br>`cesapp-maximal-${uniqueId}` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us/apps/`<br>`cesapp-maximal-${uniqueId}` | updateMask: description, displayName, globalInstruction, metadata, modelSettings, pinned, timeZoneSettings |
| 6 | `DELETE` | `GET` `200` | `.../locations/us/apps/`<br>`cesapp-maximal-${uniqueId}` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/us/apps/`<br>`cesapp-maximal-${uniqueId}` | LRO: delete -> done (#8)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 7/7 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cesapp-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
