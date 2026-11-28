# HTTP Lifecycle Report: `computetargetpool`

- **Primary Resource:** `ComputeTargetPool` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computetargetpool0-${uniqueId}`
- **Source Log:** `_http.log` (63 total HTTP entries, 41 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:45d11f02d7ee24d3`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.healthChecks` | **from:** `[{"httpHealthCheckRef": {"name": "computehttphealthcheck-${un...`<br>**to:** *(unset)* | Yes (#40 `POST`) | Yes |
| `spec.instances` | **from:** `[{"name": "computeinstance1-${uniqueId}"}, {"name": "computei...`<br>**to:** `[{"name": "computeinstance2-${uniqueId}"}]` | No | No (mismatch) |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeHTTPHealthCheck` | 7 | `CREATE`, `DELETE` |
| `ComputeInstance` | 24 | `CREATE`, `UPDATE`, `DELETE` |
| `ComputeInstanceTemplate` | 9 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 8 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 12 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 13 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 15 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 24 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 32 | `CREATE` | `GET` `404` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Pre-create existence check |
| 33 | `CREATE` | `POST` `200` | `.../regions/us-central1/targetPools` | LRO: insert -> done (#34)<br>Created resource |
| 35 | `CREATE` | `GET` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Post-create read (matches `create.yaml`) |
| 36 | `CREATE` | `GET` `404` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Post-create eventual-consistency poll |
| 37 | `CREATE` | `POST` `200` | `.../regions/us-central1/targetPools` | LRO: insert -> done (#38)<br>Composite create follow-up (`POST`) |
| 39 | `CREATE` | `GET` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Post-create read (matches `create.yaml`) |
| 40 | `UPDATE` | `POST` `200` | `.../regions/us-central1/targetPools/${targetPoolID}/`<br>`addInstance` | LRO: AddInstance -> done (#41)<br>Updated resource (`POST`) |
| 44 | `UPDATE` | `GET` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Post-update read (matches `update.yaml`) |
| 45 | `DELETE` | `DELETE` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | LRO: delete -> done (#46)<br>Deleted resource |
| 47 | `DELETE` | `GET` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | Pre-delete read |
| 48 | `DELETE` | `DELETE` `200` | `.../regions/us-central1/targetPools/`<br>`${targetPoolID}` | LRO: delete -> done (#49)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 1/2 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `spec.instances`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `spec.instances`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [ ] **V3 (End-State):** 1/2 updated fields verified in final GCP state; mismatch on: `spec.instances`.
  - `<hint_for_agent>`: END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `spec.instances`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
