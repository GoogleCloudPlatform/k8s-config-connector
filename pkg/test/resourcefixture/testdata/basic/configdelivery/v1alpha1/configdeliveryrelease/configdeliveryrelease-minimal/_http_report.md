# HTTP Lifecycle Report: `configdeliveryrelease-minimal`

- **Primary Resource:** `ConfigDeliveryRelease` (`configdelivery.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cdrelease-${uniqueId} (resourceID: v1-0-0)`
- **Source Log:** `_http.log` (12 total HTTP entries, 6 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:41f837d81c55cb62`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.env` | **from:** *(unset)*<br>**to:** `"test"` | No (0 Update RPCs) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ConfigDeliveryResourceBundle` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/`<br>`releases`<br>`?releaseId=v1-0-0` | LRO: create -> done (#6)<br>Created resource |
| 7 | `CREATE` | `GET` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | Post-create read (matches `create.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/resourceBundles/cdrbundle-${uniqueId}/releases/`<br>`v1-0-0` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 0 mutation RPCs occurred in `UPDATE` phase despite 1 modified field(s) in `update.yaml` (metadata.labels.env).
  - `<hint_for_agent>`: SILENT UPDATE NO-OP DETECTED: `update.yaml` modified metadata.labels.env, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Cross-stanza continuity mismatch in `CREATE` phase on `lifecycle`: stanza #5 (`POST`) request has `2`, LRO response (#6) has `"PUBLISHED"`, and stanza #7 (`GET`) has `2`.
  - `<hint_for_agent>`: HAND-EDITED LOG DETECTED (CROSS-STANZA MISMATCH): The mutation request, completed LRO response, and subsequent GET response disagree on a field value. Do not hand-edit _http.log; re-record against RealGCP.
