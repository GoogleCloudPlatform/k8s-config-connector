# HTTP Lifecycle Report: `computemanagedsslcertificate`

- **Primary Resource:** `ComputeManagedSSLCertificate` (`compute.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `computemanagedsslcertificate-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:57bac3c39130f0a9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../global/sslCertificates/`<br>`${sslCertificateID}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../global/sslCertificates` | LRO: insert -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../global/sslCertificates/`<br>`${sslCertificateID}` | Post-create read (matches `create.yaml`) |
| 5 | `DELETE` | `DELETE` `200` | `.../global/sslCertificates/`<br>`${sslCertificateID}` | LRO: delete -> done (#6)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computemanagedsslcertificate.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
