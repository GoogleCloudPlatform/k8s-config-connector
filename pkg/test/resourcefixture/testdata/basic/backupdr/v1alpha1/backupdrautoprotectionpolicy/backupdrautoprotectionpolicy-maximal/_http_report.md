# HTTP Lifecycle Report: `backupdrautoprotectionpolicy-maximal`

- **Primary Resource:** `BackupDRAutoProtectionPolicy` (`backupdr.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `backupdrautoprotectionpolicy-maximal-${uniqueId}`
- **Source Log:** `_http.log` (22 total HTTP entries, 13 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:35ff6be57417d4e9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.criteria.matchingConditions` | **from:** `[{"labelCondition": {"key": "env", "values": ["prod"]}}]`<br>**to:** `[{"labelCondition": {"key": "env", "values": ["stage"]}}]` | Yes (#12 `PATCH`) | Yes |
| `spec.description` | **from:** `"Auto protection policy maximal test"`<br>**to:** `"Updated auto protection policy maximal test"` | Yes (#12 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `BackupDRBackupPlan` | 6 | `CREATE`, `DELETE` |
| `BackupDRBackupVault` | 6 | `CREATE`, `DELETE` |
| `Project` | 1 | `UPDATE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `CREATE` | `GET` `404` | `.../locations/us-central1/autoProtectionPolicies/`<br>`backupdrautoprotectionpolicy-maximal-${uniqueId}` | Pre-create existence check |
| 8 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`autoProtectionPolicies`<br>`?autoProtectionPolicyId=backupdrautoprotectionpolicy-maximal-${uniqueId}` | LRO: create -> done (#9)<br>Created resource |
| 10 | `UPDATE` | `GET` `200` | `.../locations/us-central1/autoProtectionPolicies/`<br>`backupdrautoprotectionpolicy-maximal-${uniqueId}` | Pre-update read |
| 12 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/autoProtectionPolicies/`<br>`backupdrautoprotectionpolicy-maximal-${uniqueId}` | LRO: update -> done (#13)<br>updateMask: criteria, description |
| 14 | `DELETE` | `GET` `200` | `.../locations/us-central1/autoProtectionPolicies/`<br>`backupdrautoprotectionpolicy-maximal-${uniqueId}` | Pre-delete read |
| 15 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/autoProtectionPolicies/`<br>`backupdrautoprotectionpolicy-maximal-${uniqueId}` | LRO: delete -> done (#16)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_backupdrautoprotectionpolicy-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
