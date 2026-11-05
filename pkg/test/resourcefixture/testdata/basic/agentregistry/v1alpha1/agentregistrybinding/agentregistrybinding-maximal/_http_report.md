# HTTP Lifecycle Report: `agentregistrybinding-maximal`

- **Primary Resource:** `AgentRegistryBinding` (`agentregistry.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `agentregistrybinding-maximal-${uniqueId}`
- **Source Log:** `_http.log` (9 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:b4bd138a86dcb120`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.authProviderBinding.`<br>`.continueURI` | **from:** `"https://example.com/callback"`<br>**to:** `"https://example.com/callback-updated"` | Yes (#5 `PATCH`) | Yes |
| `spec.authProviderBinding.scopes` | **from:** `["openid", "email"]`<br>**to:** `["openid", "email", "profile"]` | Yes (#5 `PATCH`) | Yes |
| `spec.description` | **from:** `"Original Description"`<br>**to:** `"Updated Description"` | Yes (#5 `PATCH`) | Yes |
| `spec.displayName` | **from:** `"Original Display Name"`<br>**to:** `"Updated Display Name"` | Yes (#5 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/bindings/`<br>`agentregistrybinding-maximal-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/bindings`<br>`?bindingId=agentregistrybinding-maximal-${uniqueId}` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GET` `200` | `.../locations/us-central1/bindings/`<br>`agentregistrybinding-maximal-${uniqueId}` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/bindings/`<br>`agentregistrybinding-maximal-${uniqueId}` | LRO: update -> done (#6)<br>updateMask: authProviderBinding, description, displayName |
| 7 | `UPDATE` | `GET` `200` | `.../locations/us-central1/bindings/`<br>`agentregistrybinding-maximal-${uniqueId}` | Post-update read (matches `update.yaml`) |
| 8 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/bindings/`<br>`agentregistrybinding-maximal-${uniqueId}` | LRO: delete -> done (#9)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 4/4 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_agentregistrybinding-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
