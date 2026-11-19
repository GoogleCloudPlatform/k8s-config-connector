# HTTP Lifecycle Report: `computeinterconnectattachmentgroup-maximal`

- **Primary Resource:** `ComputeInterconnectAttachmentGroup` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeinterconnectattachmentgroup-maximal-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:4c1eb382c5880616`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description"`<br>**to:** `"Updated description"` | Yes (#5 `PATCH`) | Yes |
| `spec.intent.availabilitySla` | **from:** `"PRODUCTION_CRITICAL"`<br>**to:** `"PRODUCTION_NON_CRITICAL"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../global/interconnectAttachmentGroups/`<br>`${interconnectAttachmentGroupID}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../global/`<br>`interconnectAttachmentGroups` | LRO: insert -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../global/interconnectAttachmentGroups/`<br>`${interconnectAttachmentGroupID}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../global/interconnectAttachmentGroups/`<br>`${interconnectAttachmentGroupID}` | LRO: patch -> done (#6)<br>updateMask: description, intent |
| 7 | `UPDATE` | `GET` `200` | `.../global/interconnectAttachmentGroups/`<br>`${interconnectAttachmentGroupID}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../global/interconnectAttachmentGroups/`<br>`${interconnectAttachmentGroupID}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeinterconnectattachmentgroup-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
