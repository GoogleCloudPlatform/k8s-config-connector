# HTTP Lifecycle Report: `gatewaysecuritypolicyrule-minimal`

- **Primary Resource:** `NetworkSecurityGatewaySecurityPolicyRule` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `gsprule-min-${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:7e0b2aef87b2e209`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.priority` | **from:** `100`<br>**to:** `101` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `NetworkSecurityGatewaySecurityPolicy` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-min-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/`<br>`rules`<br>`?gatewaySecurityPolicyRuleId=gsprule-min-${uniqueId}` | LRO: create -> done (#6)<br>Created resource |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-min-${uniqueId}` | Pre-update read |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-min-${uniqueId}` | LRO: update -> done (#9)<br>updateMask: priority |
| 10 | `DELETE` | `GET` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-min-${uniqueId}` | Pre-delete read |
| 11 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-min-${uniqueId}` | LRO: delete -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_gatewaysecuritypolicyrule-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
