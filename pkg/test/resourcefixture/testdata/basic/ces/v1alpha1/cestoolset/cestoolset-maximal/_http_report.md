# HTTP Lifecycle Report: `cestoolset-maximal`

- **Primary Resource:** `CESToolset` (`ces.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cestoolset-maximal-${uniqueId} (resourceID: cestoolset-max-${uniqueId})`
- **Source Log:** `_http.log` (12 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:01436a01e4ca1094`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Maximal CESToolset fixture"`<br>**to:** `"Updated maximal CESToolset fixture"` | Yes (#7 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"cestoolset-maximal-${uniqueId}"`<br>**to:** `"cestoolset-maximal-updated-${uniqueId}"` | Yes (#7 `PATCH`) | Yes |
| `spec.openAPIToolset.`<br>`.apiAuthentication.`<br>`.bearerTokenConfig.token` | **from:** `"$context.variables.auth_token"`<br>**to:** `"$context.variables.updated_auth_token"` | Yes (#7 `PATCH`) | Yes |
| `spec.openAPIToolset.`<br>`.ignoreUnknownFields` | **from:** `true`<br>**to:** `false` | Yes (#7 `PATCH`) | Yes |
| `spec.openAPIToolset.`<br>`.openAPISchema` | **from:** `"openapi: 3.0.0\ninfo:\n  title: Maximal Toolset API\n  versi...`<br>**to:** `"openapi: 3.0.0\ninfo:\n  title: Maximal Toolset API Updated\...` | Yes (#7 `PATCH`) | Yes |
| `spec.toolFakeConfig.`<br>`.enableFakeMode` | **from:** `true`<br>**to:** `false` | Yes (#7 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `CESApp` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us/apps/cesapp-${uniqueId}/toolsets/`<br>`cestoolset-max-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us/apps/cesapp-${uniqueId}/`<br>`toolsets`<br>`?toolsetId=cestoolset-max-${uniqueId}` | Created resource |
| 6 | `UPDATE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/toolsets/`<br>`cestoolset-max-${uniqueId}` | Pre-update read |
| 7 | `UPDATE` | `PATCH` `200` | `.../locations/us/apps/cesapp-${uniqueId}/toolsets/`<br>`cestoolset-max-${uniqueId}` | updateMask: description, displayName, openApiToolset, toolFakeConfig |
| 8 | `DELETE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/toolsets/`<br>`cestoolset-max-${uniqueId}` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/us/apps/cesapp-${uniqueId}/toolsets/`<br>`cestoolset-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 6/6 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cestoolset-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
