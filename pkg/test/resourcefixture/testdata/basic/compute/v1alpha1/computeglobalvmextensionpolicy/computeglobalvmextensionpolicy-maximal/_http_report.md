# HTTP Lifecycle Report: `computeglobalvmextensionpolicy-maximal`

- **Primary Resource:** `ComputeGlobalVMExtensionPolicy` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeglobalvmextensionpolicy-maximal-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:f367e78cb2cd50c9`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"Initial description"`<br>**to:** `"Updated description"` | Yes (#5 `PATCH`) | Yes |
| `spec.extensionPolicies.ops-agent.`<br>`.pinnedVersion` | **from:** `"2.57.0"`<br>**to:** `"2.58.0"` | Yes (#5 `PATCH`) | Yes |
| `spec.extensionPolicies.ops-agent.`<br>`.stringConfig` | **from:** `"sample config"`<br>**to:** `"updated config"` | Yes (#5 `PATCH`) | Yes |
| `spec.instanceSelectors` | **from:** `[{"labelSelector": {"inclusionLabels": {"env": "test"}}}]`<br>**to:** `[{"labelSelector": {"inclusionLabels": {"env": "prod"}}}]` | Yes (#5 `PATCH`) | Yes |
| `spec.priority` | **from:** `500`<br>**to:** `600` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../global/vmExtensionPolicies/`<br>`${globalVmExtensionPolicyID}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../global/vmExtensionPolicies` | LRO: insert -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../global/vmExtensionPolicies/`<br>`${globalVmExtensionPolicyID}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../global/vmExtensionPolicies/`<br>`${globalVmExtensionPolicyID}` | LRO: update -> done (#6)<br>Updated resource (`PATCH`) |
| 7 | `UPDATE` | `GET` `200` | `.../global/vmExtensionPolicies/`<br>`${globalVmExtensionPolicyID}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `POST` `200` | `.../global/vmExtensionPolicies/${globalVmExtensionPolicyID}/`<br>`delete` | LRO: update -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 5/5 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_computeglobalvmextensionpolicy-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
