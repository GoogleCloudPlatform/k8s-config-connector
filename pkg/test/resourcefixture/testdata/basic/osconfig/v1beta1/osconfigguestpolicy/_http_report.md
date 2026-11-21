# HTTP Lifecycle Report: `osconfigguestpolicy`

- **Primary Resource:** `OSConfigGuestPolicy` (`osconfig.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `osconfigguestpolicy-${uniqueId}`
- **Source Log:** `_http.log` (7 total HTTP entries, 0 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:c6a2fd4b8f988694`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.assignment.groupLabels` | **from:** `[{"labels": {"app": "old", "env": "prod"}}, {"labels": {"app"...`<br>**to:** `[{"labels": {"app": "new", "env": "prod"}}, {"labels": {"app"...` | Yes (#4 `PATCH`) | Yes |
| `spec.assignment.instances` | **from:** `[{"name": "computeinstance1-${uniqueId}"}, {"name": "computei...`<br>**to:** `[{"name": "computeinstance2-${uniqueId}"}]` | No | No (mismatch) |
| `spec.recipes` | **from:** `[{"artifacts": [{"allowInsecure": true, "id": "ansible", "rem...`<br>**to:** `[{"artifacts": [{"allowInsecure": true, "id": "ansible", "rem...` | No | No (mismatch) |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../guestPolicies`<br>`?guestPolicyId=osconfigguestpolicy-${uniqueId}` | Created resource |
| 3 | `CREATE` | `GET` `200` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 4 | `UPDATE` | `PATCH` `200` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Updated resource (`PATCH`) |
| 5 | `UPDATE` | `GET` `200` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Deleted resource |
| 7 | `DELETE` | `GET` `404` | `.../guestPolicies/`<br>`osconfigguestpolicy-${uniqueId}` | Post-delete 404 confirmation |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 1/3 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `spec.assignment.instances`, `spec.recipes`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `spec.assignment.instances`, `spec.recipes`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [ ] **V3 (End-State):** 1/3 updated fields verified in final GCP state; mismatch on: `spec.assignment.instances`, `spec.recipes`.
  - `<hint_for_agent>`: END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `spec.assignment.instances`, `spec.recipes`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
