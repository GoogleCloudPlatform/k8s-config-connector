# HTTP Lifecycle Report: `memorystoreaclpolicy-maximal`

- **Primary Resource:** `MemorystoreACLPolicy` (`memorystore.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `memorystoreaclpolicy-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2c28212230caaceb`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.rules` | **from:** `[{"rule": "on +@all ~*", "username": "default"}, {"rule": "on...`<br>**to:** `[{"rule": "on +@read ~*", "username": "default"}, {"rule": "o...` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/aclPolicies/`<br>`memorystoreaclpolicy-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/aclPolicies`<br>`?aclPolicyId=memorystoreaclpolicy-${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/us-central1/aclPolicies/`<br>`memorystoreaclpolicy-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/aclPolicies/`<br>`memorystoreaclpolicy-${uniqueId}` | LRO: update -> done (#5)<br>updateMask: rules |
| 6 | `DELETE` | `GET` `200` | `.../locations/us-central1/aclPolicies/`<br>`memorystoreaclpolicy-${uniqueId}` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/aclPolicies/`<br>`memorystoreaclpolicy-${uniqueId}` | LRO: delete -> done (#8)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_memorystoreaclpolicy-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
