# HTTP Lifecycle Report: `modelarmortemplate-maximal`

- **Primary Resource:** `ModelArmorTemplate` (`modelarmor.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `modelarmortemplate-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:ce175e84983e05d8`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.labels.foo` | **from:** `"bar"`<br>**to:** `"updated-bar"` | Yes (#4 `GRPC`) | Yes |
| `spec.labels.key` | **from:** `"value"`<br>**to:** *(unset)* | Yes (#4 `GRPC`) | Yes |
| `spec.labels.new-key` | **from:** *(unset)*<br>**to:** `"new-value"` | Yes (#4 `GRPC`) | Yes |
| `spec.templateMetadata.`<br>`.customPromptSafetyErrorCode` | **from:** `123`<br>**to:** `789` | Yes (#4 `GRPC`) | Yes |
| `spec.templateMetadata.`<br>`.customPromptSafetyErrorMessage` | **from:** `"Blocked prompt"`<br>**to:** `"Blocked prompt updated"` | Yes (#4 `GRPC`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`GetTemplate` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`CreateTemplate` | Created resource |
| 3 | `UPDATE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`GetTemplate` | Pre-update read |
| 4 | `UPDATE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`UpdateTemplate` | updateMask: labels, templateMetadata |
| 5 | `UPDATE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`GetTemplate` | Post-update read (matches `update.yaml`) |
| 6 | `DELETE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`GetTemplate` | Pre-delete read |
| 7 | `DELETE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`GetTemplate` | Pre-delete read |
| 8 | `DELETE` | `GRPC` `200` | `/google.cloud.modelarmor.v1.ModelArmor/`<br>`DeleteTemplate` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_modelarmortemplate-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
