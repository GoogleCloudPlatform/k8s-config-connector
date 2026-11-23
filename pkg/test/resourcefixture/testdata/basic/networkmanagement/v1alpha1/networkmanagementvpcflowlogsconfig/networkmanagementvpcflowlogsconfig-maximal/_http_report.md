# HTTP Lifecycle Report: `networkmanagementvpcflowlogsconfig-maximal`

- **Primary Resource:** `NetworkManagementVPCFlowLogsConfig` (`networkmanagement.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `flowlogsconfig-${uniqueId}`
- **Source Log:** `_http.log` (11 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:3148065043a222b4`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A test vpc flow logs config created by kcc direct controller...`<br>**to:** `"An updated test vpc flow logs config description"` | Yes (#5 `PATCH`) | Yes |
| `spec.flowSampling` | **from:** `0.5`<br>**to:** `0.8` | Yes (#5 `PATCH`) | Yes |
| `spec.labels.label-key` | **from:** `"label-value"`<br>**to:** `"label-value-updated"` | Yes (#5 `PATCH`) | Yes |
| `spec.state` | **from:** `"ENABLED"`<br>**to:** `"DISABLED"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/`<br>`vpcFlowLogsConfigs`<br>`?vpcFlowLogsConfigId=flowlogsconfig-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | LRO: update -> done (#6..#8)<br>updateMask: description, flowSampling, labels, state |
| 7 | `UPDATE` | `GET` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `GET` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Pre-delete read |
| 10 | `DELETE` | `DELETE` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | LRO: delete -> done (#11)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkmanagementvpcflowlogsconfig-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
