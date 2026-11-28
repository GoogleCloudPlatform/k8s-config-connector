# HTTP Lifecycle Report: `regiondiskresourcepolicyattachmentautogen`

- **Primary Resource:** `ComputeRegionDiskResourcePolicyAttachment` (`compute.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `computeregiondiskresourcepolicyattachment-${uniqueId} (resourceID: computeregiondiskresourcepolicyattachment${uniqueId})`
- **Source Log:** `_http.log` (33 total HTTP entries, 33 dependency entries)
- **Validation Status:** FAIL (5/7 checks passed)
- **Integrity Hash:** `sha256:f4166ce238e322bb`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

_No `update.yaml` in this fixture._

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `ComputeDisk` | 17 | `CREATE`, `DELETE` |
| `ComputeResourcePolicy` | 11 | `CREATE`, `DELETE` |
| `ComputeSnapshot` | 5 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

_No primary resource RPC entries recorded._

## 3. Automated Validation Checks

- [ ] **V1 (Create):** No creation mutation RPC (`POST`/`PUT`/`PATCH`) found for primary resource in `CREATE` phase.
  - `<hint_for_agent>`: Verify that the controller's Create() method invokes the GCP API and that the primary resource is not being skipped or filtered out.
- [x] **V2 (Update):** N/A (no `update.yaml` in fixture).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_regiondiskresourcepolicyattachmentautogen.golden.yaml` match `create.yaml`.
- [ ] **V4 (Delete):** No `DELETE` RPC recorded for primary resource in `DELETE` phase.
  - `<hint_for_agent>`: MISSING DELETE RPC: The primary resource was not deleted in _http.log. Ensure DeleteResources() ran and the DELETE stanza was not truncated or hand-deleted.
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
