# HTTP Lifecycle Report: `agentidentityauthprovider-maximal`

- **Primary Resource:** `AgentIdentityAuthProvider` (`agentidentity.cnrm.cloud.google.com/v1alpha1`)
- **Resource Name:** `authprovider-max-${uniqueId}`
- **Source Log:** `_http.log` (6 total HTTP entries, 0 dependency entries)
- **Validation Status:** PASS (7/7 checks passed)
- **Integrity Hash:** `sha256:8c076137651f0bc2`

## 1. Spec Diff (`create.yaml` -> `update.yaml`)

| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |
| :--- | :--- | :--- | :--- |
| `spec.allowedScopes` | **from:** `["https://www.googleapis.com/auth/userinfo.email"]`<br>**to:** `["https://www.googleapis.com/auth/userinfo.email", "https://w...` | Yes (#4 `PATCH`) | Yes |
| `spec.authProviderTypeParams.`<br>`.threeLeggedOauth.`<br>`.authorizationURL` | **from:** `"https://auth.example.com/authorize"`<br>**to:** `"https://auth.example.com/authorize2"` | Yes (#4 `PATCH`) | Yes |
| `spec.authProviderTypeParams.`<br>`.threeLeggedOauth.clientID` | **from:** `"example-client-id"`<br>**to:** `"updated-client-id"` | Yes (#4 `PATCH`) | Yes |
| `spec.authProviderTypeParams.`<br>`.threeLeggedOauth.`<br>`.defaultContinueURI` | **from:** `"https://example.com/continue"`<br>**to:** `"https://example.com/continue2"` | Yes (#4 `PATCH`) | Yes |
| `spec.authProviderTypeParams.`<br>`.threeLeggedOauth.enablePkce` | **from:** `false`<br>**to:** `true` | Yes (#4 `PATCH`) | Yes |
| `spec.authProviderTypeParams.`<br>`.threeLeggedOauth.tokenURL` | **from:** `"https://auth.example.com/oauth/token"`<br>**to:** `"https://auth.example.com/oauth/token2"` | Yes (#4 `PATCH`) | Yes |
| `spec.blockedScopes` | **from:** `["https://www.googleapis.com/auth/userinfo.profile"]`<br>**to:** `["https://www.googleapis.com/auth/userinfo.profile", "https:/...` | Yes (#4 `PATCH`) | Yes |
| `spec.description` | **from:** `"Initial description for maximal auth provider"`<br>**to:** `"Updated description for maximal auth provider"` | Yes (#4 `PATCH`) | Yes |

## 2. Primary Resource RPC Summary

| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |
| :--- | :--- | :--- | :--- | :--- |
| 1 | `CREATE` | `GET` `404` | `.../locations/us-central1/authProviders/`<br>`authprovider-max-${uniqueId}` | Pre-create existence check |
| 2 | `CREATE` | `POST` `200` | `.../locations/us-central1/`<br>`authProviders`<br>`?authProviderId=authprovider-max-${uniqueId}` | Created resource |
| 3 | `UPDATE` | `GET` `200` | `.../locations/us-central1/authProviders/`<br>`authprovider-max-${uniqueId}` | Pre-update read |
| 4 | `UPDATE` | `PATCH` `200` | `.../locations/us-central1/authProviders/`<br>`authprovider-max-${uniqueId}` | updateMask: allowedScopes, authProviderTypeParams, blockedScopes, description |
| 5 | `DELETE` | `GET` `200` | `.../locations/us-central1/authProviders/`<br>`authprovider-max-${uniqueId}` | Pre-delete read |
| 6 | `DELETE` | `DELETE` `200` | `.../locations/us-central1/authProviders/`<br>`authprovider-max-${uniqueId}` | Deleted resource |

## 3. Automated Validation Checks

- [x] **V1 (Create):** Primary resource created and post-create state matches `create.yaml`.
- [x] **V2 (Update):** 8/8 mutated fields from `update.yaml` sent in `UPDATE` phase (`PATCH` with `updateMask`).
- [x] **V3 (End-State):** Final GCP response and `_generated_object_agentidentityauthprovider-maximal.golden.yaml` match `update.yaml`.
- [x] **V4 (Delete):** Primary resource deleted cleanly (`200 OK`).
- [x] **V5 (Idempotency):** No unexpected `4xx`/`5xx` errors or duplicate mutation loops.
- [x] **V6 (Normalization):** No unnormalized project IDs, numbers, or timestamps detected.
- [x] **V7 (Authenticity & Hand-Edit Integrity):** Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified.
