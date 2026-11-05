# OSLoginSSHPublicKey Migration Journal

## Observations & Design Decisions

- **API Endpoint Selection (`ImportSshPublicKey` vs `CreateSshPublicKey`)**:
  The Google Cloud OS Login API provides both `ImportSshPublicKey` (`POST /v1/users/{user}:importSshPublicKey`) and `CreateSshPublicKey` (`POST /v1/users/{user}/sshPublicKeys`). The Terraform Google Beta provider used `ImportSshPublicKey` because it associates the key with the user's login profile and supports specifying `projectId` (represented by `spec.project` in KRM). To preserve behavior parity, payload structures, and ensure `spec.project` is honored, the direct controller calls `ImportSshPublicKey` using the official GAPIC REST client (`cloud.google.com/go/oslogin/apiv1`).

- **Service-Generated Resource ID Compliance**:
  `OSLoginSSHPublicKey` relies on server-computed SHA-256 fingerprints as resource identifiers. The direct controller adheres to the 4 service-generated resource ID rules:
  1. `Find()` guards against an empty `a.id.Fingerprint` and returns `false, nil` prior to resource creation without making invalid GET requests.
  2. `Create()` does not set `Fingerprint` or `Name` in the creation request payload.
  3. The server-assigned fingerprint is extracted from the returned `LoginProfile.SshPublicKeys` entry matching `spec.key`, updating `a.id.Fingerprint` and `status.fingerprint`.

- **Update Payload and Field Parity**:
  `UpdateSshPublicKey` uses PATCH semantics with `updateMask`. To achieve exact field parity with legacy Terraform behavior and prevent unintentional updates, the direct controller constructs an `SshPublicKey` update proto populated strictly with fields present in `updateMask` (e.g., `expirationTimeUsec`). Any modification to immutable fields (such as `key`) is detected and rejected with an explicit error.

- **Real GCP Permission Constraint**:
  In real GCP, the OS Login API enforces that authenticated credentials match the target user in the request path (`users/{user}`). Test runs using a service account (`overseer-kcc-tester@...`) targeting `test-user-...@example.com` are rejected by live GCP with `403 PERMISSION_DENIED: End user credentials must match the user specified in the request.` All controller behaviors, round-trip re-reconciliations, and migration takeover flows were validated hermetically against MockGCP using `hack/compare-mock` and `TestMigrationToDirect`.
