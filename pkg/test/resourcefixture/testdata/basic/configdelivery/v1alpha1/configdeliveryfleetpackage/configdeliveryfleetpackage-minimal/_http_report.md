# HTTP Lifecycle Report: `configdeliveryfleetpackage-minimal`

- **Primary Resource:** `ConfigDeliveryFleetPackage` (`configdelivery.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cdfpackage-${uniqueId}`
- **Source Log:** `_http.log` (18 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:330f41eb6c81688d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ConfigDeliveryRelease` | 6 | `CREATE`, `DELETE` |
| `ConfigDeliveryResourceBundle` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/fleetPackages/`<br>`cdfpackage-${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`fleetPackages`<br>`?fleetPackageId=cdfpackage-${uniqueId}` | LRO: create -> done (#9)<br>Created resource |
| 10 | `DELETE` | `GET` `200` | `.../locations/us-central1/fleetPackages/`<br>`cdfpackage-${uniqueId}` | Pre-delete read |
| 11 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/fleetPackages/`<br>`cdfpackage-${uniqueId}`<br>`?allowMissing=true&force=true` | LRO: delete -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_configdeliveryfleetpackage-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
