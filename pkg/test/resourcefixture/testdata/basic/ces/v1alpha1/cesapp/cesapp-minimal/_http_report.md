# HTTP Lifecycle Report: `cesapp-minimal`

- **Primary Resource:** `CESApp` (`ces.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cesapp-minimal-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f7a3d6da08a44ce9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"cesapp-minimal-${uniqueId}"`<br>**to:** `"cesapp-min-updated-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us/apps/`<br>`cesapp-minimal-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us/apps`<br>`?appId=cesapp-minimal-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `UPDATE` | `GET` `200` | `.../locations/us/apps/`<br>`cesapp-minimal-${uniqueId}` | Pre-update read |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us/apps/`<br>`cesapp-minimal-${uniqueId}` | updateMask: displayName |
| 6 | `DELETE` | `GET` `200` | `.../locations/us/apps/`<br>`cesapp-minimal-${uniqueId}` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/us/apps/`<br>`cesapp-minimal-${uniqueId}` | LRO: delete -> done (#8)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cesapp-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
