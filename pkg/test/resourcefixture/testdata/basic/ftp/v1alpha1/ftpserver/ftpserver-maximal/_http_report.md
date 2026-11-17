# HTTP Lifecycle Report: `ftpserver-maximal`

- **Primary Resource:** `FTPServer` (`ftp.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `ftpmax-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e2e0e715b0fd5daf`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.displayName` | **from:** `"ftpmax-${uniqueId}"`<br>**to:** `"ftpmax-upd-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |
| `spec.internalConfig.`<br>`.consumerAcceptList` | **from:** `[{"connectionLimit": 10, "projectRef": {"external": "projects...`<br>**to:** `[{"connectionLimit": 20, "projectRef": {"external": "projects...` | Yes (#5 `PATCH`) | Yes |
| `metadata.labels.test-label` | **from:** `"ftpserver-${uniqueId}"`<br>**to:** `"ftpserver-upd-${uniqueId}"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-west1/servers/`<br>`ftpmax-${uniqueId}`<br>`?view=SERVER_VIEW_FULL` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-west1/servers`<br>`?serverId=ftpmax-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/us-west1/servers/`<br>`ftpmax-${uniqueId}`<br>`?view=SERVER_VIEW_FULL` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-west1/servers/`<br>`ftpmax-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: displayName, internalConfig, labels |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-west1/servers/`<br>`ftpmax-${uniqueId}`<br>`?view=SERVER_VIEW_FULL` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-west1/servers/`<br>`ftpmax-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_ftpserver-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
