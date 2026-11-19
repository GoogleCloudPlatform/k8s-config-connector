# HTTP Lifecycle Report: `alloydbinstancepsc`

- **Primary Resource:** `AlloyDBInstance` (`alloydb.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `alloydbinstance-${uniqueId}`
- **Source Log:** `_http.log` (85 total HTTP entries, 58 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:2fe01ea6b5600426`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.pscInstanceConfig.`<br>`.allowedConsumerProjectRefs` | **from:** *(unset)*<br>**to:** `[{"external": "${projectId}"}]` | Yes (#51 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscAutoConnectionPolicyState` | **from:** `"ENABLED"`<br>**to:** `"DISABLED"` | Yes (#51 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscAutoConnections` | **from:** `[{"consumerNetworkRef": {"name": "computenetwork-${uniqueId}"...`<br>**to:** `[{"consumerNetworkRef": {"name": "computenetwork-${uniqueId}"...` | Yes (#51 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscAutoDNSState` | **from:** `"PSC_AUTO_DNS_STATE_ENABLED"`<br>**to:** `"PSC_AUTO_DNS_STATE_DISABLED"` | Yes (#51 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscInterfaceConfigs` | **from:** `[{"networkAttachmentResourceRef": {"name": "na-${uniqueId}"}}]`<br>**to:** `[{"networkAttachmentResourceRef": {"name": "na-2-${uniqueId}"}}]` | Yes (#51 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AlloyDBCluster` | 6 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 27 | `CREATE`, `DELETE` |
| `ComputeNetworkAttachment` | 14 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 10 | `CREATE`, `DELETE` |
| `Project` | 1 | `UPDATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 28 | `CREATE` | `GET` `404` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-create existence check |
| 29 | `CREATE` | `POST` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/`<br>`instances`<br>`?instanceId=alloydbinstance-${uniqueId}` | LRO: create -> done (#30)<br>Created resource |
| 31 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-update read |
| 33 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 34 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 35 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 36 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 37 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 38 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 39 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 40 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 41 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 42 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 43 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 44 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 45 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 46 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 47 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 48 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 49 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 50 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 51 | `UPDATE` | `PATCH` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: update -> done (#52)<br>updateMask: pscInstanceConfig.allowedConsumerProjects, pscInstanceConfig.pscInterfaceConfigs, pscInstanceConfig.pscAutoConnections, pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 53 | `DELETE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-delete read |
| 54 | `DELETE` | `DELETE` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: delete -> done (#55)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_alloydbinstancepsc.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
