# HTTP Lifecycle Report: `networkconnectivitypolicybasedroute-minimal`

- **Primary Resource:** `NetworkConnectivityPolicyBasedRoute` (`networkconnectivity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ncpbr-${uniqueId}`
- **Source Log:** `_http.log` (15 total HTTP entries, 7 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:2f23d3734d1a6bf0`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.env` | **from:** *(unset)*<br>**to:** `"updated"` | No (0 Update RPCs) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 5 | `CREATE` | `GRPC` `404` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`GetPolicyBasedRoute` | Pre-create existence check |
| 6 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`CreatePolicyBasedRoute` | LRO: create -> done (#7)<br>Created resource |
| 8 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`GetPolicyBasedRoute` | Post-create read (matches `create.yaml`) |
| 9 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`GetPolicyBasedRoute` | Post-create read (matches `create.yaml`) |
| 10 | `CREATE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`GetPolicyBasedRoute` | Post-create read (matches `create.yaml`) |
| 11 | `DELETE` | `GRPC` `200` | `/google.cloud.networkconnectivity.v1.PolicyBasedRoutingService/`<br>`DeletePolicyBasedRoute` | LRO: delete -> done (#12)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 0 mutation RPCs occurred in `UPDATE` phase despite 1 modified field(s) in `update.yaml` (metadata.labels.env).
  - `<hint_for_agent>`: SILENT UPDATE NO-OP DETECTED: `update.yaml` modified metadata.labels.env, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.
- [x] **V3 (End-State):** Final GCP response and `_generated_object_networkconnectivitypolicybasedroute-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
