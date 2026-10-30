# HTTP Lifecycle Report: `dataformworkspace-minimal`

- **Primary Resource:** `DataformWorkspace` (`dataform.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `dfws-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 4 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:7d0b0ccf19ebc333`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `.spec` or `.metadata.labels` differences between `create.yaml` and `update.yaml`._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DataformRepository` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 3 | `CREATE` | `GET` `404` | `.../locations/us-central1/repositories/dfws-repo-${uniqueId}/workspaces/`<br>`dfws-${uniqueId}` | Pre-create existence check |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-central1/repositories/dfws-repo-${uniqueId}/`<br>`workspaces`<br>`?workspaceId=dfws-${uniqueId}` | Created resource |
| 5 | `CREATE` | `GET` `200` | `.../locations/us-central1/repositories/dfws-repo-${uniqueId}/workspaces/`<br>`dfws-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/repositories/dfws-repo-${uniqueId}/workspaces/`<br>`dfws-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `.spec` or `.metadata.labels` changes between `create.yaml` and `update.yaml`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_dataformworkspace-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
