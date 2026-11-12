# HTTP Lifecycle Report: `alloydbinstancepsc`

- **Primary Resource:** `AlloyDBInstance` (`alloydb.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `alloydbinstance-${uniqueId}`
- **Source Log:** `_http.log` (55 total HTTP entries, 28 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:e18ae4a2d5f3cee2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.pscInstanceConfig.`<br>`.pscAutoConnectionPolicyState` | **from:** `"ENABLED"`<br>**to:** `"DISABLED"` | Yes (#39 `PATCH`) | Yes |
| `spec.pscInstanceConfig.`<br>`.pscAutoDNSState` | **from:** `"PSC_AUTO_DNS_STATE_ENABLED"`<br>**to:** `"PSC_AUTO_DNS_STATE_DISABLED"` | Yes (#39 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `AlloyDBCluster` | 6 | `CREATE`, `DELETE` |
| `ComputeNetwork` | 9 | `CREATE`, `DELETE` |
| `ComputeNetworkAttachment` | 7 | `CREATE`, `DELETE` |
| `ComputeSubnetwork` | 5 | `CREATE`, `DELETE` |
| `Project` | 1 | `UPDATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 16 | `CREATE` | `GET` `404` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-create existence check |
| 17 | `CREATE` | `POST` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/`<br>`instances`<br>`?instanceId=alloydbinstance-${uniqueId}` | LRO: create -> done (#18)<br>Created resource |
| 19 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-update read |
| 21 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 22 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 23 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 24 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 25 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 26 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 27 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 28 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 29 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 30 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 31 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 32 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 33 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 34 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 35 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 36 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 37 | `UPDATE` | `PATCH` `409` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 38 | `UPDATE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 39 | `UPDATE` | `PATCH` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: update -> done (#40)<br>updateMask: pscInstanceConfig.pscAutoDnsState, pscInstanceConfig.pscAutoConnectionPolicyState |
| 41 | `DELETE` | `GET` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | Pre-delete read |
| 42 | `DELETE` | `DELETE` `200` | `.../locations/europe-west1/clusters/alloydbcluster-${uniqueId}/instances/`<br>`alloydbinstance-${uniqueId}` | LRO: delete -> done (#43)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_alloydbinstancepsc.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
