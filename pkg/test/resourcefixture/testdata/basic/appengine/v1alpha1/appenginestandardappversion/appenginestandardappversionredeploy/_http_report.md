# HTTP Lifecycle Report: `appenginestandardappversionredeploy`

- **Primary Resource:** `AppEngineStandardAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `appenginestandardappversion-${uniqueId} (resourceID: v${uniqueId})`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:0662affceecea500`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.automaticScaling.`<br>`.maxConcurrentRequests` | **from:** *(unset)*<br>**to:** `20` | Yes (#5 `POST`) | Yes |
| `spec.automaticScaling.`<br>`.standardSchedulerSettings.`<br>`.maxInstances` | **from:** *(unset)*<br>**to:** `2` | Yes (#5 `POST`) | Yes |
| `spec.automaticScaling.`<br>`.standardSchedulerSettings.`<br>`.targetCpuUtilization` | **from:** *(unset)*<br>**to:** `0.6` | Yes (#5 `POST`) | Yes |
| `spec.deployment.files` | **from:** `[{"name": "main.py", "sourceUrl": "https://storage.googleapis...`<br>**to:** `[{"name": "main.py", "sourceUrl": "https://storage.googleapis...` | Yes (#5 `POST`) | Yes |
| `spec.entrypoint.shell` | **from:** `"gunicorn -b :$PORT main:app"`<br>**to:** `"gunicorn -b :$PORT --workers 2 main:app"` | Yes (#5 `POST`) | Yes |
| `spec.envVariables.STAGE` | **from:** `"create"`<br>**to:** `"update"` | Yes (#5 `POST`) | Yes |
| `spec.handlers` | **from:** `[{"script": {"scriptPath": "auto"}, "securityLevel": "SECURE_...`<br>**to:** `[{"redirectHttpResponseCode": "REDIRECT_HTTP_RESPONSE_CODE_30...` | Yes (#5 `POST`) | Yes |
| `spec.instanceClass` | **from:** `"B2"`<br>**to:** `"F2"` | Yes (#5 `POST`) | Yes |
| `spec.manualScaling.instances` | **from:** `1`<br>**to:** *(unset)* | Yes (#5 `POST`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `/v1/apps/${projectId}/services/default/`<br>`versions` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `POST` `200` | `/v1/apps/${projectId}/services/default/`<br>`versions` | LRO: create -> done (#6)<br>Updated resource (`POST`) |
| 7 | `UPDATE` | `GET` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}`<br>`?view=FULL` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `/v1/apps/${projectId}/services/default/versions/`<br>`v${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 9/9 mutated fields from `update.yaml` sent in `UPDATE` phase (`POST`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_appenginestandardappversionredeploy.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
