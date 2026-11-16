# HTTP Lifecycle Report: `livestreamdvrsession-minimal`

- **Primary Resource:** `LiveStreamDVRSession` (`livestream.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `livestreamdvrsession-minimal-${uniqueId}`
- **Source Log:** `_http.log` (29 total HTTP entries, 20 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:0d15657a264a1d4d`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.dvrWindows` | **from:** `[{"timeInterval": {"endTime": "2026-11-05T01:00:00Z", "startT...`<br>**to:** `[{"timeInterval": {"endTime": "2026-11-05T02:00:00Z", "startT...` | Yes (#16 `PATCH`) | Yes |

## Dependency Traffic Summary

| Dependency Kind | Total RPCs | Phases |
| :--- | :--- | :--- |
| `LiveStreamChannel` | 7 | `CREATE`, `DELETE` |
| `LiveStreamInput` | 7 | `CREATE`, `DELETE` |
| `StorageBucket` | 6 | `CREATE`, `DELETE` |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 12 | `CREATE` | `GET` `404` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/dvrSessions/`<br>`livestreamdvrsession-minimal-${uniqueId}` | Pre-create existence check |
| 13 | `CREATE` | `POST` `200` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/`<br>`dvrSessions`<br>`?dvrSessionId=livestreamdvrsession-minimal-${uniqueId}` | LRO: create -> done (#14)<br>Created resource |
| 15 | `UPDATE` | `GET` `200` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/dvrSessions/`<br>`livestreamdvrsession-minimal-${uniqueId}` | Pre-update read |
| 16 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/dvrSessions/`<br>`livestreamdvrsession-minimal-${uniqueId}` | LRO: update -> done (#17)<br>updateMask: dvrWindows |
| 18 | `DELETE` | `GET` `200` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/dvrSessions/`<br>`livestreamdvrsession-minimal-${uniqueId}` | Pre-delete read |
| 19 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/channels/livestreamdvrsession-min-channel-${uniqueId}/dvrSessions/`<br>`livestreamdvrsession-minimal-${uniqueId}` | LRO: delete -> done (#20)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 1/1 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_livestreamdvrsession-minimal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
