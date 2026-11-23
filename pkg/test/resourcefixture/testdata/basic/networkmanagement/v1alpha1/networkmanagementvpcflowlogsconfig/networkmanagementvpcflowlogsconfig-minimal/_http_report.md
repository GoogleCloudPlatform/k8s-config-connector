# HTTP Lifecycle Report: `networkmanagementvpcflowlogsconfig-minimal`

- **Primary Resource:** `NetworkManagementVPCFlowLogsConfig` (`networkmanagement.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `flowlogsconfig-${uniqueId}`
- **Source Log:** `_http.log` (8 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:4efc7eca8d3a1896`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/global/`<br>`vpcFlowLogsConfigs`<br>`?vpcFlowLogsConfigId=flowlogsconfig-${uniqueId}` | LRO: create -> done (#3..#5)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 6 | `DELETE` | `GET` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | Pre-delete read |
| 7 | `DELETE` | `DELETE` `200` | `.../locations/global/vpcFlowLogsConfigs/`<br>`flowlogsconfig-${uniqueId}` | LRO: delete -> done (#8)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkmanagementvpcflowlogsconfig-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
