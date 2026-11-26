# HTTP Lifecycle Report: `designcenterapplication-maximal`

- **Primary Resource:** `DesignCenterApplication` (`designcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `designcenterapplication-maximal-${uniqueId}`
- **Source Log:** `_http.log` (11 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:72b67b5bb6420795`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal application description for testing"`<br>**to:** `"Updated maximal application description for testing"` | Yes (#6 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Application Display Name"`<br>**to:** `"Updated Maximal Application Display Name"` | Yes (#6 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DesignCenterSpace` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/applications/`<br>`designcenterapplication-maximal-${uniqueId}` | Pre-create existence check |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/`<br>`applications`<br>`?applicationId=designcenterapplication-maximal-${uniqueId}` | Created resource |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/applications/`<br>`designcenterapplication-maximal-${uniqueId}` | Pre-update read |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/applications/`<br>`designcenterapplication-maximal-${uniqueId}` | updateMask: description, displayName |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/applications/`<br>`designcenterapplication-maximal-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/spaces/dc-space-${uniqueId}/applications/`<br>`designcenterapplication-maximal-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_designcenterapplication-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
