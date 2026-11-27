# HTTP Lifecycle Report: `designcentercomponent-maximal`

- **Primary Resource:** `DesignCenterComponent` (`designcenter.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `designcentercomponent-maximal-${uniqueId}`
- **Source Log:** `_http.log` (14 total HTTP entries, 8 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:37b46a5f8c706c37`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"Maximal Component Initial"`<br>**to:** `"Maximal Component Updated"` | Yes (#8 `PATCH`) | Yes |
| `spec.inputVariableAliases` | **from:** `[{"alias": "runtime_name", "variable": "agent-runtime-1_displ...`<br>**to:** `[{"alias": "updated_runtime_name", "variable": "agent-runtime...` | Yes (#8 `PATCH`) | Yes |
| `spec.useAsRootModule` | **from:** `false`<br>**to:** `true` | Yes (#8 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `DesignCenterApplicationTemplate` | 4 | `CREATE`, `DELETE` |
| `DesignCenterSpace` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GET` `404` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/components/`<br>`designcentercomponent-maximal-${uniqueId}` | Pre-create existence check |
| 6 | `CREATE` | `POST` `200` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/`<br>`components`<br>`?componentId=designcentercomponent-maximal-${uniqueId}` | Created resource |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/components/`<br>`designcentercomponent-maximal-${uniqueId}` | Pre-update read |
| 8 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/components/`<br>`designcentercomponent-maximal-${uniqueId}` | updateMask: displayName, inputVariableAliases, useAsRootModule |
| 9 | `DELETE` | `GET` `200` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/components/`<br>`designcentercomponent-maximal-${uniqueId}` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/spaces/dc-space-max-${uniqueId}/applicationTemplates/dc-app-template-max-${uniqueId}/components/`<br>`designcentercomponent-maximal-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_designcentercomponent-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
