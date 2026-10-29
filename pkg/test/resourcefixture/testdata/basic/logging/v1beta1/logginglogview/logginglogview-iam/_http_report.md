# HTTP Lifecycle Report: `logginglogview-iam`

- **Primary Resource:** `IAMPolicyMember` (`iam.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `iampolicymember-${uniqueId}`
- **Source Log:** `_http.log` (21 total HTTP entries, 16 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:4b661822900ac088`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `IAMServiceAccount` | 5 | `CREATE`, `DELETE` |
| `LoggingLogBucket` | 5 | `CREATE`, `DELETE` |
| `LoggingLogView` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 10 | `CREATE` | `POST` `200` | `.../locations/global/buckets/logginglogbucket-${uniqueId}/views/`<br>`logginglogview-${uniqueId}:getIamPolicy` | Pre-create read |
| 11 | `CREATE` | `POST` `200` | `.../locations/global/buckets/logginglogbucket-${uniqueId}/views/`<br>`logginglogview-${uniqueId}:getIamPolicy` | Pre-create read |
| 12 | `CREATE` | `POST` `200` | `.../locations/global/buckets/logginglogbucket-${uniqueId}/views/`<br>`logginglogview-${uniqueId}:setIamPolicy` | Created resource |
| 13 | `DELETE` | `POST` `200` | `.../locations/global/buckets/logginglogbucket-${uniqueId}/views/`<br>`logginglogview-${uniqueId}:getIamPolicy` | Pre-delete read |
| 14 | `DELETE` | `POST` `200` | `.../locations/global/buckets/logginglogbucket-${uniqueId}/views/`<br>`logginglogview-${uniqueId}:setIamPolicy` | Teardown (`POST`) |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_logginglogview-iam.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
