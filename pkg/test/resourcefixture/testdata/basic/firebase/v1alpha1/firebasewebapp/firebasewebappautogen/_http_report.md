# HTTP Lifecycle Report: `firebasewebappautogen`

- **Primary Resource:** `FirebaseWebApp` (`firebase.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `firebasewebapp-${uniqueId}`
- **Source Log:** `_http.log` (4 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:31772a2913db8ae0`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `POST` `200` | `.../webApps` | LRO: create -> done (#2..#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../webApps/`<br>`1:${projectNumber}:web:d845db9c6118f7948b7d90` | Post-create read (matches `create.yaml`) |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_firebasewebappautogen.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
