# HTTP Lifecycle Report: `dialogflowcxwebhook-basic`

- **Primary Resource:** `DialogflowCXWebhook` (`dialogflowcx.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dialogflowcxwebhook-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:4583f5c1913ddbe7`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.disabled` | **from:** `false`<br>**to:** `true` | Yes (#3 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"webhook-${uniqueId}"`<br>**to:** `"webhook-updated-${uniqueId}"` | Yes (#3 `PATCH`) | Yes |
| `spec.genericWebService.uri` | **from:** `"https://example.com/webhook"`<br>**to:** `"https://example.com/webhook-updated"` | Yes (#3 `PATCH`) | Yes |
| `spec.timeout` | **from:** `"10s"`<br>**to:** `"20s"` | Yes (#3 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../locations/global/agents/9fc1dd9f-b1a2-4c57-9d91-f9d7a7b263d9/`<br>`webhooks` | Created resource |
| 2 | `CREATE` | `GET` `200` | `.../locations/global/agents/9fc1dd9f-b1a2-4c57-9d91-f9d7a7b263d9/webhooks/`<br>`bd72b10f-953f-4543-8f9b-715425c14612` | Post-create read (matches `create.yaml`) |
| 3 | `UPDATE` | `PATCH` `200` | `.../locations/global/agents/9fc1dd9f-b1a2-4c57-9d91-f9d7a7b263d9/webhooks/`<br>`bd72b10f-953f-4543-8f9b-715425c14612` | updateMask: displayName, timeout, disabled, genericWebService |
| 4 | `UPDATE` | `GET` `200` | `.../locations/global/agents/9fc1dd9f-b1a2-4c57-9d91-f9d7a7b263d9/webhooks/`<br>`bd72b10f-953f-4543-8f9b-715425c14612` | Post-update read (matches `update.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../locations/global/agents/9fc1dd9f-b1a2-4c57-9d91-f9d7a7b263d9/webhooks/`<br>`bd72b10f-953f-4543-8f9b-715425c14612` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dialogflowcxwebhook-basic.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
