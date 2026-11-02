# HTTP Lifecycle Report: `securesourcemanagerhook-maximal`

- **Primary Resource:** `SecureSourceManagerHook` (`securesourcemanager.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ssmhook-max-${uniqueId}`
- **Source Log:** `_http.log` (41 total HTTP entries, 35 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:5c26da449d499dc5`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.disabled` | **from:** `false`<br>**to:** `true` | Yes (#34 `PATCH`) | Yes |
| `spec.events` | **from:** `["PUSH"]`<br>**to:** `["PUSH", "PULL_REQUEST"]` | Yes (#34 `PATCH`) | Yes |
| `spec.pushOption.branchFilter` | **from:** `"main"`<br>**to:** `"release/*"` | Yes (#34 `PATCH`) | Yes |
| `spec.sensitiveQueryString` | **from:** `"secret=token1"`<br>**to:** `"secret=token2"` | Yes (#34 `PATCH`) | Yes |
| `spec.targetURI` | **from:** `"https://example.com/webhook"`<br>**to:** `"https://example.com/webhook-updated"` | Yes (#34 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `Project` | 1 | `CREATE` |
| `SecureSourceManagerInstance` | 6 | `CREATE`, `DELETE` |
| `SecureSourceManagerRepository` | 28 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 31 | `CREATE` | `GET` `404` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/hooks/`<br>`ssmhook-max-${uniqueId}` | Pre-create existence check |
| 32 | `CREATE` | `POST` `200` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/`<br>`hooks`<br>`?hookId=ssmhook-max-${uniqueId}` | LRO: create -> done (sync)<br>Created resource |
| 33 | `UPDATE` | `GET` `200` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/hooks/`<br>`ssmhook-max-${uniqueId}` | Pre-update read |
| 34 | `UPDATE` | `PATCH` `200` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/hooks/`<br>`ssmhook-max-${uniqueId}` | LRO: update -> done (sync)<br>updateMask: disabled, events, pushOption, sensitiveQueryString, targetUri |
| 35 | `DELETE` | `GET` `200` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/hooks/`<br>`ssmhook-max-${uniqueId}` | Pre-delete read |
| 36 | `DELETE` | `DELETE` `200` | `.../locations/asia-east1/repositories/ssmrepository-dep-${uniqueId}/hooks/`<br>`ssmhook-max-${uniqueId}` | LRO: delete -> done (sync)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_securesourcemanagerhook-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
