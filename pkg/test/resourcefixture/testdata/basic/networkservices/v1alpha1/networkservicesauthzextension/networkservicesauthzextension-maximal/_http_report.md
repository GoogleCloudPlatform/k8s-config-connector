# HTTP Lifecycle Report: `networkservicesauthzextension-maximal`

- **Primary Resource:** `NetworkServicesAuthzExtension` (`networkservices.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `authzext-max-${uniqueId}`
- **Source Log:** `_http.log` (23 total HTTP entries, 14 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:c30eb4ce0a5b00d7`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.authority` | **from:** `"ext-service.example.com:443"`<br>**to:** `"ext-service-updated.example.com:443"` | Yes (#13 `PATCH`) | Yes |
| `spec.description` | **from:** `"Initial description for AuthzExtension"`<br>**to:** `"Updated description for AuthzExtension"` | Yes (#13 `PATCH`) | Yes |
| `spec.failOpen` | **from:** `false`<br>**to:** `true` | Yes (#13 `PATCH`) | Yes |
| `spec.forwardHeaders` | **from:** `["x-my-header"]`<br>**to:** `["x-my-header", "x-another-header"]` | Yes (#13 `PATCH`) | Yes |
| `spec.metadata.initialKey` | **from:** `"initialVal"`<br>**to:** *(unset)* | Yes (#13 `PATCH`) | Yes |
| `spec.metadata.updatedKey` | **from:** *(unset)*<br>**to:** `"updatedVal"` | Yes (#13 `PATCH`) | Yes |
| `spec.timeout` | **from:** `"0.100s"`<br>**to:** `"0.200s"` | Yes (#13 `PATCH`) | Yes |
| `metadata.labels.key-update` | **from:** *(unset)*<br>**to:** `"val-update"` | Yes (#13 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeBackendService` | 7 | `CREATE`, `DELETE` |
| `ComputeHealthCheck` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 9 | `CREATE` | `GET` `404` | `.../locations/us-central1/authzExtensions/`<br>`authzext-max-${uniqueId}` | Pre-create existence check |
| 10 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`authzExtensions`<br>`?authzExtensionId=authzext-max-${uniqueId}` | LRO: create -> done (#11)<br>Created resource |
| 12 | `UPDATE` | `GET` `200` | `.../locations/us-central1/authzExtensions/`<br>`authzext-max-${uniqueId}` | Pre-update read |
| 13 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/authzExtensions/`<br>`authzext-max-${uniqueId}` | LRO: update -> done (#14)<br>updateMask: authority, description, failOpen, forwardHeaders, labels, metadata, timeout |
| 15 | `DELETE` | `GET` `200` | `.../locations/us-central1/authzExtensions/`<br>`authzext-max-${uniqueId}` | Pre-delete read |
| 16 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/authzExtensions/`<br>`authzext-max-${uniqueId}` | LRO: delete -> done (#17)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 8/8 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkservicesauthzextension-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
