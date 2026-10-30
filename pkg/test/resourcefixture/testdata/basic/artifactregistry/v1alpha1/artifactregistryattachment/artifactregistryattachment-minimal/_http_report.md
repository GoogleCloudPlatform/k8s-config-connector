# HTTP Lifecycle Report: `artifactregistryattachment-minimal`

- **Primary Resource:** `ArtifactRegistryAttachment` (`artifactregistry.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `arattachment-${uniqueId}`
- **Source Log:** `_http.log` (18 total HTTP entries, 7 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:7ee919564bf79d10`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.updated` | **from:** *(unset)*<br>**to:** `"true"` | No (0 Update RPCs) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ArtifactRegistryRepository` | 7 | `CREATE`, `UPDATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `400` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/`<br>`attachments`<br>`?attachmentId=arattachment-${uniqueId}` | Created resource |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Pre-create existence check |
| 8 | `UPDATE` | `POST` `400` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/`<br>`attachments`<br>`?attachmentId=arattachment-${uniqueId}` | Updated resource (`POST`) |
| 9 | `UPDATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 10 | `UPDATE` | `POST` `400` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/`<br>`attachments`<br>`?attachmentId=arattachment-${uniqueId}` | Updated resource (`POST`) |
| 11 | `UPDATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 12 | `UPDATE` | `POST` `400` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/`<br>`attachments`<br>`?attachmentId=arattachment-${uniqueId}` | Updated resource (`POST`) |
| 13 | `UPDATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 14 | `UPDATE` | `POST` `400` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/`<br>`attachments`<br>`?attachmentId=arattachment-${uniqueId}` | Updated resource (`POST`) |
| 15 | `UPDATE` | `GET` `404` | `.../locations/us-west1/repositories/arrepository-${uniqueId}/attachments/`<br>`arattachment-${uniqueId}` | Post-update read (matches `update.yaml`) |

## 3. Automated Validation Checks

- [ ] **V1 (Create):** Creation RPC #6 (`POST`) returned non-2xx status `400 Bad Request`.
  - `<hint_for_agent>`: Inspect the failed CREATE request payload and GCP error response in _http.log.
- [ ] **V2 (Update):** 0 mutation RPCs occurred in `UPDATE` phase despite 1 modified field(s) in `update.yaml` (metadata.labels.updated).
  - `<hint_for_agent>`: SILENT UPDATE NO-OP DETECTED: `update.yaml` modified metadata.labels.updated, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
