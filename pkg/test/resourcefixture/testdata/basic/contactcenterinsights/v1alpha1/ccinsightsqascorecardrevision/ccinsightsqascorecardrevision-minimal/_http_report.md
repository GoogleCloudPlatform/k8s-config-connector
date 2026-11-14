# HTTP Lifecycle Report: `ccinsightsqascorecardrevision-minimal`

- **Primary Resource:** `CCInsightsQAScorecardRevision` (`contactcenterinsights.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `rev-minimal-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 5 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:41e67e1961eae94f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `CCInsightsQAScorecard` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 4 | `CREATE` | `GET` `404` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`rev-minimal-${uniqueId}` | Pre-create existence check |
| 5 | `CREATE` | `GET` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/`<br>`revisions` | Pre-create read |
| 6 | `CREATE` | `POST` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`9642586111448304558:tuneQaScorecardRevision` | LRO: create -> pending (sync)<br>Created resource |
| 7 | `CREATE` | `GET` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`9642586111448304558` | Post-create read (matches `create.yaml`) |
| 8 | `CREATE` | `GET` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`9642586111448304558` | Post-create read (matches `create.yaml`) |
| 9 | `CREATE` | `POST` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`9642586111448304558:deploy` | Composite create follow-up (`POST`) |
| 10 | `CREATE` | `POST` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/`<br>`revisions`<br>`?qaScorecardRevisionId=rev-minimal-${uniqueId}` | Composite create follow-up (`POST`) |
| 11 | `DELETE` | `GET` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`rev-minimal-${uniqueId}` | Pre-delete read |
| 12 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/qaScorecards/ccinsightsqascorecard-${uniqueId}/revisions/`<br>`rev-minimal-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (3 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ccinsightsqascorecardrevision-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
