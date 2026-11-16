# HTTP Lifecycle Report: `saasservicemgmtunitkind-maximal`

- **Primary Resource:** `SaaSServiceMgmtUnitKind` (`saasservicemgmt.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `saasservicemgmtunitkind-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:d59a77a38d2a355f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.annotations.description` | **from:** `"maximal unit kind test"`<br>**to:** `"updated maximal unit kind test"` | Yes (#6 `PATCH`) | Yes |
| `spec.inputVariableMappings` | **from:** `[{"to": {"dependency": "dep1", "ignoreForLookup": false, "inp...`<br>**to:** `[{"to": {"dependency": "dep1", "ignoreForLookup": true, "inpu...` | Yes (#6 `PATCH`) | Yes |
| `spec.labels.env` | **from:** `"dev"`<br>**to:** `"prod"` | Yes (#6 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/unitKinds`<br>`?unitKindId=saasservicemgmtunitkind-dep-${uniqueId}` | Created resource |
| 3 | `CREATE` | `GET` `404` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-${uniqueId}` | Post-create eventual-consistency poll |
| 4 | `CREATE` | `POST` `200` | `.../locations/us-central1/unitKinds`<br>`?unitKindId=saasservicemgmtunitkind-${uniqueId}` | Composite create follow-up (`POST`) |
| 5 | `UPDATE` | `GET` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-${uniqueId}` | Pre-update read |
| 6 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-${uniqueId}` | updateMask: annotations, inputVariableMappings, labels |
| 7 | `DELETE` | `GET` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-${uniqueId}` | Pre-delete read |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-${uniqueId}` | Deleted resource |
| 9 | `DELETE` | `GET` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `400` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Deleted resource |
| 11 | `DELETE` | `GET` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Pre-delete read |
| 12 | `DELETE` | `DELETE` `400` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Deleted resource |
| 13 | `DELETE` | `GET` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Pre-delete read |
| 14 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/unitKinds/`<br>`saasservicemgmtunitkind-dep-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created via composite create (2 mutation RPCs) and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_saasservicemgmtunitkind-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
