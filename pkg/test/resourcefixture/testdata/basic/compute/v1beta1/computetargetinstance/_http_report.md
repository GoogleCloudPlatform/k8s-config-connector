# HTTP Lifecycle Report: `computetargetinstance`

- **Primary Resource:** `ComputeTargetInstance` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computetargetinstance-${uniqueId}`
- **Source Log:** `_http.log` (21 total HTTP entries, 4 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:65c431ee985ff194`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeInstance` | 3 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 2 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/instances/`<br>`${instanceID}` | Pre-create existence check |
| 3 | `CREATE` | `GET` `200` | `.../zones/us-central1-a` | Pre-create read |
| 4 | `CREATE` | `GET` `404` | `.../global/images/debian-12` | Pre-create existence check |
| 5 | `CREATE` | `GET` `200` | `.../global/images/family/debian-12` | Pre-create read |
| 6 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/instances` | LRO: insert -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/instances/`<br>`${instanceID}` | Post-create read (matches `create.yaml`) |
| 10 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/instances/`<br>`${instanceID}` | Post-create read (matches `create.yaml`) |
| 12 | `CREATE` | `GET` `404` | `.../zones/us-central1-a/targetInstances/`<br>`${targetInstanceID}` | Post-create eventual-consistency poll |
| 13 | `CREATE` | `POST` `200` | `.../zones/us-central1-a/`<br>`targetInstances` | LRO: insert -> done (#14)<br>Composite create follow-up (`POST`) |
| 15 | `CREATE` | `GET` `200` | `.../zones/us-central1-a/targetInstances/`<br>`${targetInstanceID}` | Post-create read (matches `create.yaml`) |
| 16 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/targetInstances/`<br>`${targetInstanceID}` | LRO: delete -> done (#17)<br>Deleted resource |
| 18 | `DELETE` | `GET` `200` | `.../zones/us-central1-a/instances/`<br>`${instanceID}` | Pre-delete read |
| 20 | `DELETE` | `DELETE` `200` | `.../zones/us-central1-a/instances/`<br>`${instanceID}` | LRO: delete -> done (#21)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computetargetinstance.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Protobuf schema violation: stanza #8 (`GET`) response body failed protobuf schema `google.cloud.compute.v1.Instance`: proto: (line 1:1487): unknown field "effectiveInstanceMetadata".
  - `<hint_for_agent>`: HAND-EDITED OR HALLUCINATED PROTO PAYLOAD DETECTED: One or more JSON payloads or @type URLs in _http.log violate the protobuf schema in .build/googleapis.pb. Do not hand-edit _http.log; re-record against RealGCP.
