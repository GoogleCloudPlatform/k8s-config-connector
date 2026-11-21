# HTTP Lifecycle Report: `gkehubfleet-maximal`

- **Primary Resource:** `GKEHubFleet` (`gkehub.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `gkehubfleet-maximal-${uniqueId} (resourceID: default)`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:8c58fdbfe0ad1b1c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.defaultClusterConfig.`<br>`.securityPostureConfig.mode` | **from:** `"BASIC"`<br>**to:** `"DISABLED"` | Yes (#5 `PATCH`) | Yes |
| `spec.defaultClusterConfig.`<br>`.securityPostureConfig.`<br>`.vulnerabilityMode` | **from:** `"VULNERABILITY_BASIC"`<br>**to:** `"VULNERABILITY_DISABLED"` | Yes (#5 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Maximal Fleet"`<br>**to:** `"Maximal Fleet Updated"` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.test-label` | **from:** `"maximal"`<br>**to:** `"maximal-updated"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/fleets/default` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/fleets` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/global/fleets/default` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/global/fleets/default` | LRO: update -> done (#6)<br>updateMask: displayName, labels, defaultClusterConfig |
| 7 | `UPDATE` | `GET` `200` | `.../locations/global/fleets/default` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/global/fleets/default` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_gkehubfleet-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
