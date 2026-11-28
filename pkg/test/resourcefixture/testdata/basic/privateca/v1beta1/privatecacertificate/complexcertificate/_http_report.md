# HTTP Lifecycle Report: `complexcertificate`

- **Primary Resource:** `PrivateCACertificate` (`privateca.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `privatecacertificate-${uniqueId}`
- **Source Log:** `_http.log` (17 total HTTP entries, 11 dependency entries)
- **Validation Status:** FAIL (4/7 checks passed)
- **Integrity Hash:** `sha256:dcbf5cec86ffcd35`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.clientAuth` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.codeSigning` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.`<br>`.emailProtection` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.ocspSigning` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.serverAuth` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `spec.config.x509Config.`<br>`.extendedKeyUsage.timeStamping` | **from:** `true`<br>**to:** *(unset)* | Yes (#12 `PATCH`) | Yes |
| `metadata.labels.key` | **from:** `"value"`<br>**to:** *(unset)* | No | No (mismatch) |
| `metadata.labels.value` | **from:** *(unset)*<br>**to:** `"key"` | Yes (#12 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `PrivateCACAPool` | 4 | `CREATE` |
| `PrivateCACertificateAuthority` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 9 | `CREATE` | `GET` `404` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/certificates/`<br>`privatecacertificate-${uniqueId}` | Pre-create existence check |
| 10 | `CREATE` | `POST` `200` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/`<br>`certificates`<br>`?certificateId=privatecacertificate-${uniqueId}&issuingCertificateAuthorityId=privatecacertificateauthority-${uniqueId}` | Created resource |
| 11 | `CREATE` | `GET` `200` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/certificates/`<br>`privatecacertificate-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 12 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/certificates/`<br>`privatecacertificate-${uniqueId}` | updateMask: labels |
| 13 | `UPDATE` | `GET` `200` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/certificates/`<br>`privatecacertificate-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 14 | `DELETE` | `POST` `200` | `.../locations/us-central1/caPools/privatecacapool-${uniqueId}/certificates/`<br>`privatecacertificate-${uniqueId}:revoke` | Teardown (`POST`) |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 7/8 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `metadata.labels.key`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `metadata.labels.key`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [ ] **V3 (End-State):** 7/8 updated fields verified in final GCP state; mismatch on: `metadata.labels.key`.
  - `<hint_for_agent>`: END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `metadata.labels.key`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.
- [ ] **V4 (Delete):** No `DELETE` RPC recorded for primary resource in `DELETE` phase.
  - `<hint_for_agent>`: MISSING DELETE RPC: The primary resource was not deleted in _http.log. Ensure DeleteResources() ran and the DELETE stanza was not truncated or hand-deleted.
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
