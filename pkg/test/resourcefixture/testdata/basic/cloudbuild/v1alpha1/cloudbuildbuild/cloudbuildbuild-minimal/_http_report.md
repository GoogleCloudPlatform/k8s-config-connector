# HTTP Lifecycle Report: `cloudbuildbuild-minimal`

- **Primary Resource:** `CloudBuildBuild` (`cloudbuild.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cloudbuildbuild-minimal-${uniqueId}`
- **Source Log:** `_http.log` (5 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:19eccfb0ced67444`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `200` | `/google.devtools.cloudbuild.v1.CloudBuild/`<br>`CreateBuild` | LRO: create -> done (#2)<br>Created resource |
| 3 | `CREATE` | `GRPC` `200` | `/google.devtools.cloudbuild.v1.CloudBuild/`<br>`GetBuild` | Post-create read (matches `create.yaml`) |
| 4 | `DELETE` | `GRPC` `200` | `/google.devtools.cloudbuild.v1.CloudBuild/`<br>`GetBuild` | Pre-delete read |
| 5 | `DELETE` | `GRPC` `200` | `/google.devtools.cloudbuild.v1.CloudBuild/`<br>`GetBuild` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_cloudbuildbuild-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
