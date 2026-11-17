# HTTP Lifecycle Report: `computehealthsource-maximal`

- **Primary Resource:** `ComputeHealthSource` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computehealthsource-maximal-${uniqueId}`
- **Source Log:** `_http.log` (37 total HTTP entries, 28 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:63a39d2db1f4849b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description for maximal health source"`<br>**to:** `"Updated description for maximal health source"` | Yes (#21 `PATCH`) | Yes |
| `spec.healthAggregationPolicyRef` | **from:** `"hap-max-1-${uniqueId}"`<br>**to:** `"hap-max-2-${uniqueId}"` | Yes (#21 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeBackendService` | 7 | `CREATE`, `DELETE` |
| `ComputeHealthAggregationPolicy` | 14 | `CREATE`, `DELETE` |
| `ComputeHealthCheck` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 17 | `CREATE` | `GET` `404` | `.../regions/us-central1/healthSources/`<br>`${healthSourceID}` | Pre-create existence check |
| 18 | `CREATE` | `POST` `200` | `.../regions/us-central1/healthSources` | LRO: compute.healthSources.insert -> done (#19)<br>Created resource |
| 20 | `CREATE` | `GET` `200` | `.../regions/us-central1/healthSources/`<br>`${healthSourceID}` | Post-create read (matches `create.yaml`) |
| 21 | `UPDATE` | `PATCH` `200` | `.../regions/us-central1/healthSources/`<br>`${healthSourceID}` | LRO: compute.healthSources.patch -> done (#22)<br>Updated resource (`PATCH`) |
| 23 | `UPDATE` | `GET` `200` | `.../regions/us-central1/healthSources/`<br>`${healthSourceID}` | Post-update read (matches `update.yaml`) |
| 24 | `DELETE` | `DELETE` `200` | `.../regions/us-central1/healthSources/`<br>`${healthSourceID}` | LRO: compute.healthSources.delete -> done (#25)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computehealthsource-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
