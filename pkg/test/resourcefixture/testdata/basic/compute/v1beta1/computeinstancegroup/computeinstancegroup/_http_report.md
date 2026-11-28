# HTTP Lifecycle Report: `computeinstancegroup`

- **Primary Resource:** `ComputeInstanceGroup` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computeinstancegroup-${uniqueId}`
- **Source Log:** `_http.log` (59 total HTTP entries, 41 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:dd7dd70500d5c588`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.instances` | **from:** `[{"name": "computeinstance1-${uniqueId}"}, {"name": "computei...`<br>**to:** `[{"name": "computeinstance2-${uniqueId}"}]` | No | No (mismatch) |
| `spec.namedPort` | **from:** `[{"name": "http", "port": 8080}, {"name": "https", "port": 84...`<br>**to:** `[{"name": "https", "port": 8443}, {"name": "fred", "port": 85...` | Yes (#37 `POST`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeInstance` | 31 | `CREATE`, `UPDATE`, `DELETE` |
| `ComputeInstanceTemplate` | 9 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 4 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 8 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 9 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 11 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 20 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 28 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Pre-create existence check |
| 29 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/instanceGroups` | LRO: compute.instanceGroups.insert -> done (#30)<br>Created resource |
| 33 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Post-create read (matches `create.yaml`) |
| 35 | `UPDATE` | `GET` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Pre-update read |
| 37 | `UPDATE` | `POST` `200` | `.../zones/us-central1-a/instanceGroups/${instanceGroupID}/`<br>`setNamedPorts` | LRO: compute.instanceGroups.setNamedPorts -> done (#38)<br>Updated resource (`POST`) |
| 41 | `UPDATE` | `GET` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Post-update read (matches `update.yaml`) |
| 43 | `DELETE` | `GET` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Pre-delete read |
| 45 | `DELETE` | `GET` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | Pre-delete read |
| 47 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/instanceGroups/`<br>`${instanceGroupID}` | LRO: compute.instanceGroups.delete -> done (#48)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 1/2 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `spec.instances`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `spec.instances`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [ ] **V3 (End-State):** 1/2 updated fields verified in final GCP state; mismatch on: `spec.instances`.
  - `<hint_for_agent>`: END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `spec.instances`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
