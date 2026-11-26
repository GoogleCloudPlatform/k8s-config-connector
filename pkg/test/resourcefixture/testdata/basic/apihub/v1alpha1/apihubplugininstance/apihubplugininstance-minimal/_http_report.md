# HTTP Lifecycle Report: `apihubplugininstance-minimal`

- **Primary Resource:** `APIHubPluginInstance` (`apihub.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `apihubplugininstance-minimal-${uniqueId}`
- **Source Log:** `_http.log` (13 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:eb956167065238b2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"Minimal Plugin Instance ${uniqueId}"`<br>**to:** `"Updated Minimal Plugin Instance ${uniqueId}"` | Yes (#7 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `APIHubPlugin` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/instances/`<br>`apihubplugininstance-minimal-${uniqueId}` | Pre-create existence check |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/`<br>`instances`<br>`?pluginInstanceId=apihubplugininstance-minimal-${uniqueId}` | LRO: create -> done (#5)<br>Created resource |
| 6 | `CREATE` | `GET` `200` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/instances/`<br>`apihubplugininstance-minimal-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 7 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/instances/`<br>`apihubplugininstance-minimal-${uniqueId}` | updateMask: displayName |
| 8 | `DELETE` | `GET` `200` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/instances/`<br>`apihubplugininstance-minimal-${uniqueId}` | Pre-delete read |
| 9 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/plugins/apihubplugin-${uniqueId}/instances/`<br>`apihubplugininstance-minimal-${uniqueId}` | LRO: delete -> done (#10)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_apihubplugininstance-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
