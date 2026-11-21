# HTTP Lifecycle Report: `ftpuser-maximal`

- **Primary Resource:** `FTPUser` (`ftp.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ftpumax-${uniqueId}`
- **Source Log:** `_http.log` (29 total HTTP entries, 20 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:047df9dc7207c5a1`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.storageDirectoryMappings` | **from:** `[{"bucketPrefix": "uploads", "bucketRef": {"name": "ftpbucket...`<br>**to:** `[{"bucketPrefix": "downloads", "bucketRef": {"name": "ftpbuck...` | Yes (#15 `PATCH`) | Yes |
| `spec.userCredentials` | **from:** `[{"credentialName": "cred-max", "credentialType": "PUBLIC_KEY...`<br>**to:** `[{"credentialName": "cred-max-upd", "credentialType": "PUBLIC...` | Yes (#15 `PATCH`) | Yes |
| `metadata.labels.test-label` | **from:** `"ftpuser-${uniqueId}"`<br>**to:** `"ftpuser-upd-${uniqueId}"` | Yes (#15 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `FTPServer` | 7 | `CREATE`, `DELETE` |
| `IAMServiceAccount` | 6 | `CREATE`, `DELETE` |
| `StorageBucket` | 7 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 11 | `CREATE` | `GET` `404` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/users/`<br>`ftpumax-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Pre-create existence check |
| 12 | `CREATE` | `POST` `200` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/`<br>`users`<br>`?userId=ftpumax-${uniqueId}` | LRO: create -> done (#13)<br>Created resource |
| 14 | `CREATE` | `GET` `200` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/users/`<br>`ftpumax-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Post-create read (matches `create.yaml`) |
| 15 | `UPDATE` | `PATCH` `200` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/users/`<br>`ftpumax-${uniqueId}` | LRO: update -> done (#16)<br>updateMask: labels, storageDirectoryMappings, userCredentials |
| 17 | `UPDATE` | `GET` `200` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/users/`<br>`ftpumax-${uniqueId}`<br>`?view=USER_VIEW_FULL` | Post-update read (matches `update.yaml`) |
| 18 | `DELETE` | `DELETE` `200` | `.../locations/us-west1/servers/ftpserv-max-${uniqueId}/users/`<br>`ftpumax-${uniqueId}` | LRO: delete -> done (#19)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ftpuser-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
