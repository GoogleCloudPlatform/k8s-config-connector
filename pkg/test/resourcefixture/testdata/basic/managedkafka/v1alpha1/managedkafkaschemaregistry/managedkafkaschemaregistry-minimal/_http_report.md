# HTTP Lifecycle Report: `managedkafkaschemaregistry-minimal`

- **Primary Resource:** `ManagedKafkaSchemaRegistry` (`managedkafka.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `mkschemaregistry-${uniqueId} (resourceID: mkschemaregistry_${uniqueId})`
- **Source Log:** `_http.log` (1 total HTTP entries, 0 dependency entries)
- **Validation Status:** FAIL (4/7 checks passed)
- **Integrity Hash:** `sha256:00045b26d8836474`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `metadata.labels.updated` | **from:** *(unset)*<br>**to:** `"true"` | No (0 Update RPCs) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `400` | `.../locations/us-central1/schemaRegistries/`<br>`mkschemaregistry_${uniqueId}` | Pre-create read |

## 3. Automated Validation Checks

- [ ] **V1 (Create):** No creation mutation RPC (`POST`/`PUT`/`PATCH`) found for primary resource in `CREATE` phase.
  - `<hint_for_agent>`: Verify that the controller's Create() method invokes the GCP API and that the primary resource is not being skipped or filtered out.
- [ ] **V2 (Update):** 0 mutation RPCs occurred in `UPDATE` phase despite 1 modified field(s) in `update.yaml` (metadata.labels.updated).
  - `<hint_for_agent>`: SILENT UPDATE NO-OP DETECTED: `update.yaml` modified metadata.labels.updated, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.
- [x] **V3 (End-State):** Final GCP response matches `update.yaml`.
- [ ] **V4 (Delete):** No `DELETE` RPC recorded for primary resource in `DELETE` phase.
  - `<hint_for_agent>`: MISSING DELETE RPC: The primary resource was not deleted in _http.log. Ensure DeleteResources() ran and the DELETE stanza was not truncated or hand-deleted.
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
