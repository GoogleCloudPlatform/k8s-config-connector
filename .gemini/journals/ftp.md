### [2026-11-13] FTPServer Direct Greenfield Types Prepopulation
- **Context**: Prepopulating direct KRM types, identity, mapper, and judgement queue for `FTPServer` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Proto & Service Details**:
  - Service: `google.cloud.ftp.v1` (`ftp.googleapis.com`)
  - Resource: `FTPServer:Server`
  - Proto path: `google/cloud/ftp/v1/service.proto` (and `google/cloud/ftp/v1/ftp_v1.yaml`)
  - Go client library: `cloud.google.com/go/ftp/apiv1/ftppb`
- **Actions & Configurations**:
  1. Proto availability: Pinned googleapis commit `befe6dc01ccbe4fc0292c5966efca2ad1fafc2e6` directly in `apis/ftp/generate.sh` which contains `google/cloud/ftp/v1/service.proto` and `ftp_v1.yaml`.
  2. Registered `ftp` API (`Cloud FTP API`) with prefix `FTP` in `dev/tools/controllerbuilder/serviceprefixes/service-prefixes.yaml`.
  3. Created `apis/ftp/generate.sh` configuring `generate-types`, `generate-identity`, and `generate-mapper` with `--prepopulate-resource`.
  4. Added `FTPServer` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  5. Added `cloud.google.com/go/ftp v0.3.1` to `go.mod` with pinned `replace` directives for `google.golang.org/api`, `genproto`, and `protobuf` to prevent transitive dependency drift. Mappers import `cloud.google.com/go/ftp/apiv1/ftppb` directly.
  6. Regenerated channel manifests via `dev/tasks/generate-channels` and operator golden tests via `WRITE_GOLDEN_OUTPUT=1 go test ./operator/pkg/controllers/configconnector/...`.
  7. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for the new `ftpservers.ftp.cnrm.cloud.google.com` CRD.
- **Impact**: Enables Greenfield Step 1a for `FTPServer` with verified conformance tests and passing `apichecks`.

### [2026-11-13] FTPServer Direct Greenfield Types Judgement & Finalization
- **Context**: Triaged judgement queue and finalized direct KRM types, identity, and mappers for `FTPServer` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Actions & Configurations**:
  1. Inlined `*parent.ProjectAndLocationRef` in `FTPServerSpec` to model `projects/{project}/locations/{location}` parentage.
  2. Removed top-level `labels` from `FTPServerSpec` as it is reconciled via `metadata.labels`.
  3. Replaced `project` in `InternalServerConfig_AllowedConsumer` and `InternalServerConfig_DeniedConsumer` with `projectRef *refsv1beta1.ProjectRef` and implemented custom `ToProto` mappers in `pkg/controller/direct/ftp/mapper.go` for the protobuf oneof `consumer_type`.
  4. Updated documentation URLs in `ftpserver_types.go` to valid `https://cloud.google.com/cloud-ftp/docs` references and removed the guess annotation.
  5. Resolved all 5 entries in `apis/ftp/judgement_queue.yaml`.
  6. Regenerated types, CRDs, and updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Completes Greenfield Step 1b for `FTPServer` with passing unit tests and apichecks.

### [2026-11-17] FTPServer Direct Greenfield Controller & E2E Fixtures
- **Context**: Implementing direct controller, round-trip fuzzer, and minimal/maximal E2E fixtures for `FTPServer` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Learnings & API Specifics**:
  1. `ServerView`: Cloud FTP uses `ServerView` enum on `GetServer`. Default view is `SERVER_VIEW_BASIC` (excludes `external_config`, `internal_config`, and `google_managed_server_credential`). Controller explicitly requests `SERVER_VIEW_FULL` in `Find`, and re-fetches the server with `SERVER_VIEW_FULL` post-LRO in `Create` and `Update` so full state is observed.
  2. Supported locations: `ftp.googleapis.com` is available in regions such as `us-west1` (not `us-central1`).
  3. Resource name length: GCP enforces maximum 30 characters on `server_id` (`[a-z]([-a-z0-9]*[a-z0-9])?`). Fixture names shortened to `ftpmin-${uniqueId}` and `ftpmax-${uniqueId}` to ensure names fit well within the 30-char limit.
  4. Response fields: Real GCP returns `satisfiesPzi` / `satisfiesPzs`, which are omitted via `tests/e2e/normalize.go` for `ftp.googleapis.com` responses.
  5. Validation: `pkg/test/httplogreport/specdiff/differ.go` was updated to simplify reference values (e.g. `projectRef`) when comparing nested map structures in slices.
- **Impact**: Full direct controller implementation, 200-seed round-trip fuzzer, and real GCP recorded minimal and maximal fixtures with 100% CRD field coverage (all missing field exceptions cleared).

### [2026-11-17] FTPUser Direct Greenfield Types Prepopulation
- **Context**: Prepopulating direct KRM types, identity, mapper, and judgement queue for `FTPUser` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Actions & Configurations**:
  1. Configured `apis/ftp/generate.sh` to include `generate-types` and `generate-identity` for `FTPUser:User` before `generate-mapper`.
  2. Added `FTPUser` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  3. Ran `./apis/ftp/generate.sh` to emit `ftpuser_types.go`, `ftpuser_identity.generated.go`, `ftpuser_reference.generated.go`, updated `types.generated.go`, `mapper.generated.go`, CRD `ftpusers.ftp.cnrm.cloud.google.com`, and `apis/ftp/judgement_queue.yaml` entries.
  4. Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` to update `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Prepopulates Greenfield Step 1a for `FTPUser` with passing identity conformance unit tests and `apichecks`.

### [2026-11-20] FTPUser Direct Greenfield Controller & E2E Fixtures
- **Context**: Implementing direct controller, round-trip fuzzer, and minimal/maximal E2E fixtures for `FTPUser` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Learnings & API Specifics**:
  1. `UserView`: Cloud FTP uses `UserView` enum on `GetUser`. The default view is `USER_VIEW_BASIC`, which excludes `user_credentials` and `storage_directory_mappings`. Controller explicitly requests `USER_VIEW_FULL` in `Find`, and re-fetches the user with `USER_VIEW_FULL` post-LRO in `Create` and `Update` so the full state is observed.
  2. `StorageDirectoryMapping.bucket`: Cloud FTP expects just the bucket name, not `projects/{project}/buckets/{bucket}`. Implemented `StorageDirectoryMapping_ToProto` in `pkg/controller/direct/ftp/mapper.go` to trim `projects/.../buckets/` and `gs://` prefixes from `bucketRef.External`.
  3. `UserCredential.ssh_public_key_body`: GCP Cloud FTP strips trailing comments (e.g. `user@host`) from OpenSSH public keys. Implemented a `normalizeUser` callback for `common.CompareGreenfieldSpecAndLabels` that normalizes `sshPublicKeyBody` to `<algorithm> <base64-key>`, preventing spurious diffs and infinite update loops during re-reconciliation.
  4. Validation: Updated `pkg/test/httplogreport/specdiff/differ.go` to support normalizing `*Ref` field keys inside object collections, and to handle enum string differences when `foundVal` contains numeric enum encodings.
- **Impact**: Full direct controller implementation, 200-seed round-trip fuzzer, and real GCP recorded minimal and maximal fixtures with 100% CRD field coverage (all 9 missing field exceptions cleared).


