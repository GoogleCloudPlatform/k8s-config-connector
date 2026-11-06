# HTTP Lifecycle Report: `gatewaysecuritypolicyrule-maximal`

- **Primary Resource:** `NetworkSecurityGatewaySecurityPolicyRule` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `gsprule-max-${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:cb8f31157007cb0e`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.applicationMatcher` | **from:** `"request.headers['user-agent'].contains('Chrome')"`<br>**to:** `"request.headers['user-agent'].contains('Firefox')"` | Yes (#8 `PATCH`) | Yes |
| `spec.basicProfile` | **from:** `"DENY"`<br>**to:** `"ALLOW"` | Yes (#8 `PATCH`) | Yes |
| `spec.description` | **from:** `"Initial description"`<br>**to:** `"Updated description"` | Yes (#8 `PATCH`) | Yes |
| `spec.priority` | **from:** `200`<br>**to:** `201` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `NetworkSecurityGatewaySecurityPolicy` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-max-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/`<br>`rules`<br>`?gatewaySecurityPolicyRuleId=gsprule-max-${uniqueId}` | LRO: create -> done (#6)<br>Created resource |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-max-${uniqueId}` | Pre-update read |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-max-${uniqueId}` | LRO: update -> done (#9)<br>updateMask: applicationMatcher, basicProfile, description, priority |
| 10 | `DELETE` | `GET` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-max-${uniqueId}` | Pre-delete read |
| 11 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/gatewaySecurityPolicies/gsp-${uniqueId}/rules/`<br>`gsprule-max-${uniqueId}` | LRO: delete -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_gatewaysecuritypolicyrule-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
