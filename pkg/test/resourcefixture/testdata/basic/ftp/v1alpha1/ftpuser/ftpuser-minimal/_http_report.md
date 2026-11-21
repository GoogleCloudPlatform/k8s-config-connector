# HTTP Lifecycle Report: `ftpuser-minimal`

- **Primary Resource:** `FTPUser` (`ftp.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ftpumin-${uniqueId}`
- **Source Log:** `_http.log` (35 total HTTP entries, 26 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:37d5732545d8aed8`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.storageDirectoryMappings` | **from:** `[{"bucketRef": {"name": "ftpbucket-min-${uniqueId}"}, "direct...`<br>**to:** `[{"bucketRef": {"name": "ftpbucket-min-${uniqueId}"}, "direct...` | Yes (#22 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `FTPServer` | 7 | `CREATE`, `DELETE` |
| `IAMServiceAccount` | 12 | `CREATE`, `DELETE` |
| `StorageBucket` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 18 | `CREATE` | `GET` `404` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/users/`<br>`ftpumin-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Pre-create existence check |
| 19 | `CREATE` | `POST` `200` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/`<br>`users`<br>`?userId=ftpumin-${uniqueId}` | LRO: create -> done (#20)<br>Created resource |
| 21 | `CREATE` | `GET` `200` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/users/`<br>`ftpumin-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Post-create read (matches `create.yaml`) |
| 22 | `UPDATE` | `PATCH` `200` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/users/`<br>`ftpumin-${uniqueId}` | LRO: update -> done (#23)<br>updateMask: storageDirectoryMappings |
| 24 | `UPDATE` | `GET` `200` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/users/`<br>`ftpumin-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Post-update read (matches `update.yaml`) |
| 25 | `DELETE` | `DELETE` `200` | `.../locations/us-west1/servers/ftpserv-min-${uniqueId}/users/`<br>`ftpumin-${uniqueId}` | LRO: delete -> done (#26)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ftpuser-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
