# HTTP Lifecycle Report: `appenginefirewallruledefault`

- **Primary Resource:** `AppEngineFirewallRule` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `appenginefirewallruledefault-${uniqueId} (resourceID: 2147483647)`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:1d9ab932c08ec626`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.action` | **from:** `"DENY"`<br>**to:** `"ALLOW"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`2147483647` | Pre-create read |
| 2 | `CREATE` | `PATCH` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`2147483647` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`2147483647` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`2147483647` | updateMask: action |
| 5 | `DELETE` | `GET` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`2147483647` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_appenginefirewallruledefault.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
