# HTTP Lifecycle Report: `cloudbuildconnection-maximal`

- **Primary Resource:** `CloudBuildConnection` (`cloudbuild.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cloudbuildconnection-max-${uniqueId}`
- **Source Log:** `_http.log` (31 total HTTP entries, 20 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ff9c8fa12a05c1d1`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.annotations.my-key` | **from:** `"val1"`<br>**to:** `"val2"` | Yes (#16 `PATCH`) | Yes |
| `spec.disabled` | **from:** `false`<br>**to:** `true` | Yes (#16 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMPolicy` | 15 | `CREATE`, `DELETE` |
| `SecretManagerSecret` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 12 | `CREATE` | `GET` `404` | `.../locations/us-central1/connections/`<br>`cloudbuildconnection-max-${uniqueId}` | Pre-create existence check |
| 13 | `CREATE` | `POST` `200` | `.../locations/us-central1/connections`<br>`?connectionId=cloudbuildconnection-max-${uniqueId}` | LRO: create -> done (#14)<br>Created resource |
| 15 | `CREATE` | `GET` `200` | `.../locations/us-central1/connections/`<br>`cloudbuildconnection-max-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 16 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/connections/`<br>`cloudbuildconnection-max-${uniqueId}` | LRO: update -> done (#17)<br>updateMask: annotations, disabled |
| 18 | `UPDATE` | `GET` `200` | `.../locations/us-central1/connections/`<br>`cloudbuildconnection-max-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 19 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/connections/`<br>`cloudbuildconnection-max-${uniqueId}` | LRO: delete -> done (#20)<br>Deleted resource |
| 28 | `DELETE` | `GET` `200` | `.../secrets/secretmanagersecret-${uniqueId}/versions/`<br>`1` | Pre-delete read |
| 29 | `DELETE` | `POST` `200` | `.../secrets/secretmanagersecret-${uniqueId}/versions/`<br>`1:destroy` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cloudbuildconnection-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
