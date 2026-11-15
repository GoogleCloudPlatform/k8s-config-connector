# HTTP Lifecycle Report: `appenginefirewallrulenamepriority`

- **Primary Resource:** `AppEngineFirewallRule` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `200`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ac43128da80b2a14`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.action` | **from:** `"ALLOW"`<br>**to:** `"DENY"` | Yes (#4 `PATCH`) | Yes |
| `spec.description` | **from:** `"Firewall rule using metadata.name as priority"`<br>**to:** `"Updated firewall rule using metadata.name"` | Yes (#4 `PATCH`) | Yes |
| `spec.sourceRange` | **from:** `"192.0.2.0/24"`<br>**to:** `"198.51.100.0/24"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`200` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `/v1/apps/${projectId}/firewall/`<br>`ingressRules` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`200` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`200` | updateMask: action, description, sourceRange |
| 5 | `DELETE` | `GET` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`200` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `/v1/apps/${projectId}/firewall/ingressRules/`<br>`200` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_appenginefirewallrulenamepriority.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
