# HTTP Lifecycle Report: `configdeliveryvariant-maximal`

- **Primary Resource:** `ConfigdeliveryVariant` (`configdelivery.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `cd-variant-${uniqueId}`
- **Source Log:** `_http.log` (21 total HTTP entries, 12 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:a3bcf1bd3fc96a5f`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.resources` | **from:** `["apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test-cm...`<br>**to:** `["apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test-cm...` | Yes (#11 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ConfigDeliveryRelease` | 6 | `CREATE`, `DELETE` |
| `ConfigDeliveryResourceBundle` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/variants/`<br>`cd-variant-${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/`<br>`variants`<br>`?variantId=cd-variant-${uniqueId}` | LRO: create -> done (#9)<br>Created resource |
| 10 | `UPDATE` | `GET` `200` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/variants/`<br>`cd-variant-${uniqueId}` | Pre-update read |
| 11 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/variants/`<br>`cd-variant-${uniqueId}` | LRO: update -> done (#12)<br>updateMask: resources |
| 13 | `DELETE` | `GET` `200` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/variants/`<br>`cd-variant-${uniqueId}` | Pre-delete read |
| 14 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/resourceBundles/cd-bundle-${uniqueId}/releases/v1-0-0/variants/`<br>`cd-variant-${uniqueId}` | LRO: delete -> done (#15)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_configdeliveryvariant-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
