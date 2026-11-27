# HTTP Lifecycle Report: `bigquerypubsubsubscription`

- **Primary Resource:** `PubSubSubscription` (`pubsub.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `pubsubsubscription-${uniqueId}`
- **Source Log:** `_http.log` (31 total HTTP entries, 24 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:6ee878caf437ec9d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.bigqueryConfig.`<br>`.dropUnknownFields` | **from:** `true`<br>**to:** `false` | Yes (#21 `PATCH`) | Yes |
| `spec.bigqueryConfig.tableRef` | **from:** `"bigquerytable-1-${uniqueId}"`<br>**to:** `"bigquerytable-2-${uniqueId}"` | Yes (#21 `PATCH`) | Yes |
| `metadata.labels.label-one` | **from:** `"value-one"`<br>**to:** `"value-two"` | Yes (#21 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BigQueryDataset` | 5 | `CREATE`, `DELETE` |
| `BigQueryTable` | 10 | `CREATE`, `DELETE` |
| `IAMPolicy` | 4 | `CREATE` |
| `PubSubTopic` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 17 | `CREATE` | `GET` `404` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Pre-create existence check |
| 18 | `CREATE` | `PUT` `200` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Created resource |
| 19 | `CREATE` | `GET` `404` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Post-create eventual-consistency poll |
| 20 | `CREATE` | `GET` `200` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 21 | `UPDATE` | `PATCH` `200` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | updateMask: bigqueryConfig, labels |
| 22 | `DELETE` | `GET` `200` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Pre-delete read |
| 23 | `DELETE` | `DELETE` `200` | `.../subscriptions/`<br>`pubsubsubscription-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 3/3 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_bigquerypubsubsubscription.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
