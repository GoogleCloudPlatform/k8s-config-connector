# HTTP Lifecycle Report: `alloydbinstancepsc`

- **Primary Resource:** `AlloyDBInstance` (`alloydb.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `alloydbinstance-${uniqueId}`
- **Source Log:** `_http.log` (71 total HTTP entries, 42 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:18d773d07860f062`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.pscInstanceConfig.`<br>`.pscAutoConnectionPolicyState` | **from:** `"ENABLED"`<br>**to:** `"DISABLED"` | Yes (#40 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscAutoDNSState` | **from:** `"PSC_AUTO_DNS_STATE_ENABLED"`<br>**to:** `"PSC_AUTO_DNS_STATE_DISABLED"` | Yes (#40 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AlloyDBCluster` | 6 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 24 | `CREATE`, `DELETE` |
| `ComputeNetworkAttachment` | 7 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 16 | `CREATE` | `GET` `404` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-create existence check |
| 17 | `CREATE` | `POST` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/`<br>`instances`<br>`?instanceId=alloydbinstance-${uniqueId}` | LRO: create -> done (#18)<br>Created resource |
| 19 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-update read |
| 20 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 21 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 22 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 23 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 24 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 25 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 26 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 27 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 28 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 29 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 30 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 31 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 32 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 33 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 34 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 35 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 36 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 37 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 38 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 39 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 40 | `UPDATE` | `PATCH` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: update -> done (#41)<br>updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 42 | `DELETE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-delete read |
| 43 | `DELETE` | `DELETE` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: delete -> done (#44)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_alloydbinstancepsc.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
