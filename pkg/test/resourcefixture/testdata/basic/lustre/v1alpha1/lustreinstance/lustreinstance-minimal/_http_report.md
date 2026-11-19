# HTTP Lifecycle Report: `lustreinstance-minimal`

- **Primary Resource:** `LustreInstance` (`lustre.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `lustreinstance-minimal-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 1 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:fd224eea21563f1c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** *(unset)*<br>**to:** `"Minimal updated description"` | Yes (#6 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 2 | `CREATE` | `GET` `404` | `.../locations/us-east4-b/instances/`<br>`lustreinstance-minimal-${uniqueId}` | Pre-create existence check |
| 3 | `CREATE` | `POST` `200` | `.../locations/us-east4-b/instances`<br>`?instanceId=lustreinstance-minimal-${uniqueId}` | LRO: create -> done (#4)<br>Created resource |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-east4-b/instances/`<br>`lustreinstance-minimal-${uniqueId}` | Pre-update read |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/us-east4-b/instances/`<br>`lustreinstance-minimal-${uniqueId}` | LRO: update -> done (#7)<br>updateMask: description |
| 8 | `DELETE` | `GET` `200` | `.../locations/us-east4-b/instances/`<br>`lustreinstance-minimal-${uniqueId}` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/us-east4-b/instances/`<br>`lustreinstance-minimal-${uniqueId}` | LRO: delete -> done (#10)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_lustreinstance-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
