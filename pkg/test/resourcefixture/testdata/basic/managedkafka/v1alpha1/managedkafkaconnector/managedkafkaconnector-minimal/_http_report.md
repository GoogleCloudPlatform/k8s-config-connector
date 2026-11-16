# HTTP Lifecycle Report: `managedkafkaconnector-minimal`

- **Primary Resource:** `ManagedKafkaConnector` (`managedkafka.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `managedkafkaconnector-${uniqueId}`
- **Source Log:** `_http.log` (32 total HTTP entries, 26 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:18a45841393e6395`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.configs.heartbeats.topic.`<br>`.replication.factor` | **from:** *(unset)*<br>**to:** `"1"` | Yes (#18 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeNetwork` | 9 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 5 | `CREATE`, `DELETE` |
| `ManagedKafkaCluster` | 8 | `CREATE`, `DELETE` |
| `ManagedKafkaConnectCluster` | 4 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 15 | `CREATE` | `GET` `404` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/connectors/`<br>`managedkafkaconnector-${uniqueId}` | Pre-create existence check |
| 16 | `CREATE` | `POST` `200` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/`<br>`connectors`<br>`?connectorId=managedkafkaconnector-${uniqueId}` | Created resource |
| 17 | `UPDATE` | `GET` `200` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/connectors/`<br>`managedkafkaconnector-${uniqueId}` | Pre-update read |
| 18 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/connectors/`<br>`managedkafkaconnector-${uniqueId}` | updateMask: configs |
| 19 | `DELETE` | `GET` `200` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/connectors/`<br>`managedkafkaconnector-${uniqueId}` | Pre-delete read |
| 20 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/connectClusters/connectcluster-${uniqueId}/connectors/`<br>`managedkafkaconnector-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_managedkafkaconnector-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
