# HTTP Lifecycle Report: `cesdeployment-maximal`

- **Primary Resource:** `CESDeployment` (`ces.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cesdeployment-maximal-${uniqueId} (resourceID: cesdep-max-${uniqueId})`
- **Source Log:** `_http.log` (12 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:1d4c1c4d4fe58c3b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.channelProfile.`<br>`.disableBargeInControl` | **from:** `false`<br>**to:** `true` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.disableDtmf` | **from:** `false`<br>**to:** `true` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.noiseSuppressionLevel` | **from:** `"moderate"`<br>**to:** `"high"` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.personaProperty.persona` | **from:** `"CONCISE"`<br>**to:** `"CHATTY"` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.webWidgetConfig.modality` | **from:** `"CHAT_AND_VOICE"`<br>**to:** `"CHAT_ONLY"` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.webWidgetConfig.`<br>`.securitySettings.allowedOrigins` | **from:** `["https://example.com"]`<br>**to:** `["https://updated.example.com"]` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.webWidgetConfig.`<br>`.securitySettings.`<br>`.enablePublicAccess` | **from:** `true`<br>**to:** `false` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.webWidgetConfig.theme` | **from:** `"LIGHT"`<br>**to:** `"DARK"` | Yes (#7 `PATCH`) | Yes |
| `spec.channelProfile.`<br>`.webWidgetConfig.webWidgetTitle` | **from:** `"Customer Support Widget"`<br>**to:** `"Updated Customer Support Widget"` | Yes (#7 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"cesdeployment-maximal-${uniqueId}"`<br>**to:** `"cesdeployment-max-updated-${uniqueId}"` | Yes (#7 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `CESApp` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us/apps/cesapp-${uniqueId}/deployments/`<br>`cesdep-max-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `POST` `200` | `.../locations/us/apps/cesapp-${uniqueId}/`<br>`deployments`<br>`?deploymentId=cesdep-max-${uniqueId}` | Created resource |
| 6 | `UPDATE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/deployments/`<br>`cesdep-max-${uniqueId}` | Pre-update read |
| 7 | `UPDATE` | `PATCH` `200` | `.../locations/us/apps/cesapp-${uniqueId}/deployments/`<br>`cesdep-max-${uniqueId}` | updateMask: channelProfile, displayName |
| 8 | `DELETE` | `GET` `200` | `.../locations/us/apps/cesapp-${uniqueId}/deployments/`<br>`cesdep-max-${uniqueId}` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/us/apps/cesapp-${uniqueId}/deployments/`<br>`cesdep-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 10/10 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cesdeployment-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
