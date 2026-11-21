# HTTP Lifecycle Report: `appengineservicesplittrafficsingleversion`

- **Primary Resource:** `AppEngineServiceSplitTraffic` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `appengineservicesplittraffic-${uniqueId} (resourceID: default)`
- **Source Log:** `_http.log` (7 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:042f3bb2c034a93b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.split.allocations.v1` | **from:** `"0.5"`<br>**to:** *(unset)* | Yes (#5 `PATCH`) | Yes |
| `spec.split.allocations.v2` | **from:** `"0.5"`<br>**to:** `"1"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `200` | `/v1/apps/${projectId}/services/default` | Pre-create read |
| 2 | `CREATE` | `PATCH` `200` | `/v1/apps/${projectId}/services/default` | LRO: update -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `/v1/apps/${projectId}/services/default` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `/v1/apps/${projectId}/services/default` | LRO: update -> done (#6)<br>updateMask: split |
| 7 | `UPDATE` | `GET` `200` | `/v1/apps/${projectId}/services/default` | Post-update read (matches `update.yaml`) |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_appengineservicesplittrafficsingleversion.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
