# HTTP Lifecycle Report: `vmmigrationgroup-maximal`

- **Primary Resource:** `VMMigrationGroup` (`vmmigration.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `vmmg-max-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:1a873c51733c7762`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description for VMMigrationGroup"`<br>**to:** `"Updated description for VMMigrationGroup"` | Yes (#5 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Initial Display Name"`<br>**to:** `"Updated Display Name"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/groups/`<br>`${groupID}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/groups`<br>`?groupId=${groupID}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/us-central1/groups/`<br>`${groupID}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/groups/`<br>`${groupID}` | LRO: update -> done (#6)<br>updateMask: description, displayName |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/groups/`<br>`${groupID}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/groups/`<br>`${groupID}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_vmmigrationgroup-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
