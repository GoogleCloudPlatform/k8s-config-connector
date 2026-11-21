# HTTP Lifecycle Report: `fullmemorystoreinstance`

- **Primary Resource:** `MemorystoreInstance` (`memorystore.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `memorystore-${uniqueId}`
- **Source Log:** `_http.log` (62 total HTTP entries, 39 dependency entries)
- **Validation Status:** FAIL (6/7 checks passed)
- **Integrity Hash:** `sha256:b2137a91db6ad95b`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.automatedBackupConfig.`<br>`.fixedFrequencySchedule.`<br>`.startTime.hours` | **from:** `2`<br>**to:** `3` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.automatedBackupConfig.`<br>`.retention` | **from:** `"345600s"`<br>**to:** `"432000s"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.deletionProtectionEnabled` | **from:** `true`<br>**to:** `false` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.engineConfigs.maxclients` | **from:** *(unset)*<br>**to:** `"10000"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.labels.key1` | **from:** `"val1"`<br>**to:** *(unset)* | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.labels.key2` | **from:** `"val2"`<br>**to:** *(unset)* | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.labels.keyx` | **from:** *(unset)*<br>**to:** `"valx"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.maintenancePolicy.`<br>`.weeklyMaintenanceWindow` | **from:** `[{"day": "MONDAY", "startTime": {"hours": 5, "minutes": 0, "n...`<br>**to:** `[{"day": "WEDNESDAY", "startTime": {"hours": 6, "minutes": 0,...` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.maintenanceVersion` | **from:** `"MEMORYSTORE_20260313_01_02"`<br>**to:** *(unset)* | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.persistenceConfig.aofConfig.`<br>`.appendFsync` | **from:** `"EVERY_SEC"`<br>**to:** *(unset)* | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.persistenceConfig.mode` | **from:** `"AOF"`<br>**to:** `"RDB"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.persistenceConfig.rdbConfig.`<br>`.rdbSnapshotPeriod` | **from:** *(unset)*<br>**to:** `"ONE_HOUR"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.persistenceConfig.rdbConfig.`<br>`.rdbSnapshotStartTime` | **from:** *(unset)*<br>**to:** `"2029-10-02T15:01:23Z"` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.replicaCount` | **from:** `2`<br>**to:** `1` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |
| `spec.shardCount` | **from:** `6`<br>**to:** `12` | Yes (#29 `PATCH`, #31 `PATCH`, #33 `PATCH`, #35 `PATCH`, #37 `PATCH`, #39 `PATCH`, #41 `PATCH`, #43 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 9 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 5 | `CREATE`, `DELETE` |
| `KMSCryptoKey` | 11 | `CREATE`, `DELETE` |
| `KMSKeyRing` | 4 | `CREATE`, `DELETE` |
| `NetworkConnectivityServiceConnectionPolicy` | 10 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 25 | `CREATE` | `GET` `404` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | Pre-create existence check |
| 26 | `CREATE` | `POST` `200` | `.../locations/us-central1/instances`<br>`?instanceId=memorystore-${uniqueId}` | LRO: create -> done (#27)<br>Created resource |
| 28 | `UPDATE` | `GET` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | Pre-update read |
| 29 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#30)<br>updateMask: automatedBackupConfig |
| 31 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#32)<br>updateMask: deletionProtectionEnabled |
| 33 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#34)<br>updateMask: engineConfigs |
| 35 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#36)<br>updateMask: labels |
| 37 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#38)<br>updateMask: maintenancePolicy |
| 39 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#40)<br>updateMask: persistenceConfig |
| 41 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#42)<br>updateMask: replicaCount |
| 43 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: update -> done (#44)<br>updateMask: shardCount |
| 45 | `DELETE` | `GET` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | Pre-delete read |
| 46 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/instances/`<br>`memorystore-${uniqueId}` | LRO: delete -> done (#47)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 15/15 mutated fields from `update.yaml` sent across 8 `UPDATE` mutation RPCs.
- [x] **V3 (End-State):** Final GCP response and `_generated_object_fullmemorystoreinstance.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [ ] **V7 (Authenticity & Hand-Edit Integrity):** Protobuf schema violation: stanza #26 (`POST`) request body failed protobuf schema `google.cloud.memorystore.v1.Instance`: proto: (line 1:231): unknown field "ipVersion".
  - `<hint_for_agent>`: HAND-EDITED OR HALLUCINATED PROTO PAYLOAD DETECTED: One or more JSON payloads or @type URLs in _http.log violate the protobuf schema in .build/googleapis.pb. Do not hand-edit _http.log; re-record against RealGCP.
