# HTTP Lifecycle Report: `urllist-maximal`

- **Primary Resource:** `NetworkSecurityURLList` (`networksecurity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `urllist-max-${uniqueId}`
- **Source Log:** `_http.log` (10 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:803a44ac5ac46dc3`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.description` | **from:** `"A URL list maximal test"`<br>**to:** `"An updated URL list maximal test"` | Yes (#5 `GRPC`) | Yes |
| `spec.values` | **from:** `["google.com", "github.com"]`<br>**to:** `["google.com", "github.com", "gitlab.com"]` | Yes (#5 `GRPC`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GRPC` `404` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`GetUrlList` | Pre-create existence check |
| 2 | `CREATE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`CreateUrlList` | LRO: create -> done (#3)<br>Created resource |
| 4 | `CREATE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`GetUrlList` | Post-create read (matches `create.yaml`) |
| 5 | `UPDATE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`UpdateUrlList` | LRO: update -> done (#6)<br>updateMask: description, values |
| 7 | `UPDATE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`GetUrlList` | Post-update read (matches `update.yaml`) |
| 8 | `UPDATE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`GetUrlList` | Post-update read (matches `update.yaml`) |
| 9 | `DELETE` | `GRPC` `200` | `/google.cloud.networksecurity.v1.NetworkSecurity/`<br>`DeleteUrlList` | LRO: delete -> done (#10)<br>Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 2/2 mutated fields from `update.yaml` sent in `UPDATE` phase (`GRPC` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_urllist-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`OK`, LRO `done: true`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
