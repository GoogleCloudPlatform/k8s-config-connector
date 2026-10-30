# HTTP Lifecycle Report: `migrationcenterimportjob-minimal`

- **Primary Resource:** `MigrationCenterImportJob` (`migrationcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mcij-${uniqueId}`
- **Source Log:** `_http.log` (16 total HTTP entries, 7 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:627dcfab0a72ff9f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"ImportJob ${uniqueId}"`<br>**to:** `"ImportJob Updated ${uniqueId}"` | Yes (#9 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `MigrationCenterSource` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/us-central1/importJobs/`<br>`mcij-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../locations/us-central1/importJobs`<br>`?importJobId=mcij-${uniqueId}` | LRO: create -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../locations/us-central1/importJobs/`<br>`mcij-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 9 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/importJobs/`<br>`mcij-${uniqueId}` | LRO: update -> done (#10)<br>updateMask: displayName |
| 11 | `UPDATE` | `GET` `200` | `.../locations/us-central1/importJobs/`<br>`mcij-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/importJobs/`<br>`mcij-${uniqueId}` | LRO: delete -> done (#13)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Cross-stanza continuity mismatch in `CREATE` phase on `assetSource`: stanza #6 (`POST`) request has `"projects/${projectId}/locations/us-central1/sources/mcs-${un...`, LRO response (#7) has `"projects/${projectNumber}/locations/us-central1/sources/mcs-...`, and stanza #8 (`GET`) has `"projects/${projectId}/locations/us-central1/sources/mcs-${un...`.
  - `<hint_for_agent>`: HAND-EDITED LOG DETECTED (CROSS-STANZA MISMATCH): The mutation request, completed LRO response, and subsequent GET response disagree on a field value. Do not hand-edit _http.log; re-record against RealGCP.
