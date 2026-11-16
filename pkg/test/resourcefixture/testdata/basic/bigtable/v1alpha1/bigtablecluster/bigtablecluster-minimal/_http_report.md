# HTTP Lifecycle Report: `bigtablecluster-minimal`

- **Primary Resource:** `BigtableCluster` (`bigtable.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `bigtablecluster-${uniqueId}`
- **Source Log:** `_http.log` (19 total HTTP entries, 6 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:05f4de23e2c6c59c`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.serveNodes` | **from:** `1`<br>**to:** `2` | Yes (#12 `GRPC`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigtableInstance` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListInstances` | Pre-create read |
| 5 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-create read |
| 7 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-create read |
| 8 | `CREATE` | `GRPC` `404` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`GetCluster` | Pre-create existence check |
| 9 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`CreateCluster` | LRO: create -> done (#10)<br>Created resource |
| 11 | `UPDATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`GetCluster` | Pre-update read |
| 12 | `UPDATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`PartialUpdateCluster` | LRO: PartialUpdateCluster -> done (#13)<br>updateMask: serveNodes |
| 14 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`GetCluster` | Pre-delete read |
| 15 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`GetCluster` | Pre-delete read |
| 16 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`DeleteCluster` | Deleted resource |
| 18 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigtablecluster-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
