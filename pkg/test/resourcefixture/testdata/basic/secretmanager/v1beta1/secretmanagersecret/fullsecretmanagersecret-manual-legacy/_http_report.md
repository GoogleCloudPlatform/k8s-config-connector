# HTTP Lifecycle Report: `fullsecretmanagersecret-manual-legacy`

- **Primary Resource:** `SecretManagerSecret` (`secretmanager.cnrm.cloud.google.com/v1beta1`)
- **Resource Name:** `secretmanagersecret-${uniqueId}`
- **Source Log:** `_http.log` (46 total HTTP entries, 39 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:9d24916f5f989852`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.annotations.bar` | **from:** *(unset)*<br>**to:** `"secretmanagersecret-bar"` | Yes (#26 `PATCH`) | Yes |
| `spec.rotation.nextRotationTime` | **from:** `"2035-10-02T15:01:23Z"`<br>**to:** `"2035-10-03T15:01:23Z"` | Yes (#26 `PATCH`) | Yes |
| `spec.topics` | **from:** `[{"topicRef": {"name": "topic-${uniqueId}"}}]`<br>**to:** `[{"topicRef": {"name": "topic-2-${uniqueId}"}}]` | No | No (mismatch) |
| `spec.ttl` | **from:** `"100s"`<br>**to:** `"110s"` | No | No (mismatch) |
| `metadata.labels.label-two` | **from:** *(unset)*<br>**to:** `"value-two"` | Yes (#26 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `KMSCryptoKey` | 12 | `CREATE`, `DELETE` |
| `KMSKeyRing` | 4 | `CREATE`, `DELETE` |
| `PubSubTopic` | 22 | `CREATE`, `DELETE` |
| `Service` | 1 | `CREATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 23 | `CREATE` | `GET` `404` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | Pre-create existence check |
| 24 | `CREATE` | `POST` `200` | `.../secrets`<br>`?secretId=secretmanagersecret-${uniqueId}` | Created resource |
| 25 | `UPDATE` | `GET` `200` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | Pre-update read |
| 26 | `UPDATE` | `PATCH` `200` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | updateMask: annotations, expireTime, labels, rotation.nextRotationTime, topics |
| 27 | `UPDATE` | `GET` `200` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 28 | `DELETE` | `GET` `200` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | Pre-delete read |
| 29 | `DELETE` | `DELETE` `200` | `.../secrets/`<br>`secretmanagersecret-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [ ] **V2 (Update):** 3/5 mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `spec.topics`, `spec.ttl`.
  - `<hint_for_agent>`: DROPPED UPDATE FIELD(S): `update.yaml` modified `spec.topics`, `spec.ttl`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.
- [ ] **V3 (End-State):** 3/5 updated fields verified in final GCP state; mismatch on: `spec.topics`, `spec.ttl`.
  - `<hint_for_agent>`: END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `spec.topics`, `spec.ttl`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
