# HTTP Lifecycle Report: `bigtablebackup-minimal`

- **Primary Resource:** `BigtableBackup` (`bigtable.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `bigtablebackup-${uniqueId}`
- **Source Log:** `_http.log` (24 total HTTP entries, 13 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:6bb07a8a9dec2403`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigtableInstance` | 6 | `CREATE`, `DELETE` |
| `BigtableTable` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListInstances` | Pre-create read |
| 5 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-create read |
| 7 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-create read |
| 13 | `CREATE` | `GRPC` `404` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`GetBackup` | Pre-create existence check |
| 14 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`CreateBackup` | LRO: create -> done (#15)<br>Created resource |
| 16 | `CREATE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`GetBackup` | Post-create read (matches `create.yaml`) |
| 17 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`GetBackup` | Pre-delete read |
| 18 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`GetBackup` | Pre-delete read |
| 19 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableTableAdmin/`<br>`DeleteBackup` | Deleted resource |
| 23 | `DELETE` | `GRPC` `200` | `/google.bigtable.admin.v2.BigtableInstanceAdmin/`<br>`ListClusters` | Pre-delete read |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigtablebackup-minimal.golden.yaml` match `create.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
