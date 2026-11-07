# HTTP Lifecycle Report: `computeorganizationsecuritypolicyrulebasic`

- **Primary Resource:** `ComputeOrganizationSecurityPolicyRule` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeorganizationsecuritypolicyrule-${uniqueId} (resourceID: 100)`
- **Source Log:** `_http.log` (18 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:d3989b95aeb7ae77`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Sample organization security policy rule"`<br>**to:** `"Updated organization security policy rule"` | Yes (#9 `POST`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/securityPolicies`<br>`?parentId=organizations/${organizationID}` | LRO: createOrganizationSecurityPolicy -> done (#2)<br>Created resource |
| 3 | `CREATE` | `GET` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | Post-create read (matches `create.yaml`) |
| 4 | `CREATE` | `GET` `400` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`getRule`<br>`?priority=100` | Post-create read (matches `create.yaml`) |
| 5 | `CREATE` | `POST` `200` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`addRule`<br>`?priority=100` | LRO: addOrganizationSecurityPolicyRule -> done (#7)<br>Composite create follow-up (`POST`) |
| 6 | `CREATE` | `GET` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | Post-create read (matches `create.yaml`) |
| 8 | `CREATE` | `GET` `200` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`getRule`<br>`?priority=100` | Post-create read (matches `create.yaml`) |
| 9 | `UPDATE` | `POST` `200` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`patchRule`<br>`?priority=100` | LRO: patchOrganizationSecurityPolicyRule -> done (#11)<br>Updated resource (`POST`) |
| 10 | `UPDATE` | `GET` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | Post-update read (matches `update.yaml`) |
| 12 | `UPDATE` | `GET` `200` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`getRule`<br>`?priority=100` | Post-update read (matches `update.yaml`) |
| 13 | `DELETE` | `POST` `200` | `.../locations/global/securityPolicies/${securityPolicyID}/`<br>`removeRule`<br>`?priority=100` | LRO: removeOrganizationSecurityPolicyRule -> done (#15)<br>Teardown (`POST`) |
| 14 | `DELETE` | `GET` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | Pre-delete read |
| 16 | `DELETE` | `GET` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | Pre-delete read |
| 17 | `DELETE` | `DELETE` `200` | `.../locations/global/securityPolicies/`<br>`${securityPolicyID}` | LRO: deleteOrganizationSecurityPolicy -> done (#18)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`POST`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeorganizationsecuritypolicyrulebasic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
