# HTTP Lifecycle Report: `networkconnectivityspoke`

- **Primary Resource:** `NetworkConnectivitySpoke` (`networkconnectivity.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `networkconnectivityspoke-${uniqueId}`
- **Source Log:** `_http.log` (34 total HTTP entries, 21 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:455bbe8e071600a3`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A sample spoke with a linked router appliance instance"`<br>**to:** `"An updated sample spoke with a linked router appliance insta...` | Yes (#21 `PATCH`) | Yes |
| `metadata.labels.label-one` | **from:** `"value-one"`<br>**to:** *(unset)* | No | Yes |
| `metadata.labels.label-two` | **from:** *(unset)*<br>**to:** `"value-two"` | Yes (#21 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeInstance` | 11 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |
| `ComputeSubnetwork` | 1 | `CREATE` |
| `NetworkConnectivityHub` | 8 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 8 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 9 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 10 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 17 | `CREATE` | `GET` `404` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | Pre-create existence check |
| 18 | `CREATE` | `POST` `200` | `.../locations/us-central1/spokes`<br>`?spokeId=networkconnectivityspoke-${uniqueId}` | LRO: create -> done (#19)<br>Created resource |
| 20 | `CREATE` | `GET` `200` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 21 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | LRO: update -> done (#22)<br>updateMask: description, labels |
| 23 | `UPDATE` | `GET` `200` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 24 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | LRO: delete -> done (#25)<br>Deleted resource |
| 26 | `DELETE` | `GET` `404` | `.../locations/us-central1/spokes/`<br>`networkconnectivityspoke-${uniqueId}` | Post-delete 404 confirmation |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 2/3 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `metadata.labels.label-one`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `metadata.labels.label-one`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkconnectivityspoke.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
