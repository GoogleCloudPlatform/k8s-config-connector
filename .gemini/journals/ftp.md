### [2026-11-13] FTPServer Direct Greenfield Types Prepopulation
- **Context**: Prepopulating direct KRM types, identity, mapper, and judgement queue for `FTPServer` under `ftp.cnrm.cloud.google.com/v1alpha1`.
- **Proto & Service Details**:
  - Service: `google.cloud.ftp.v1` (`ftp.googleapis.com`)
  - Resource: `FTPServer:Server`
  - Proto path: `google/cloud/ftp/v1/service.proto` (and `google/cloud/ftp/v1/ftp_v1.yaml`)
  - Go client library: `cloud.google.com/go/ftp/apiv1/ftppb`
- **Actions & Configurations**:
  1. Proto availability: The pinned googleapis SHA in `apis/git.versions` predated `google/cloud/ftp/v1`. Added `service.proto` and `ftp_v1.yaml` to `mockgcp/apis/google/cloud/ftp/v1/`, copied them in `dev/tools/controllerbuilder/generate-proto.sh`, and registered the proto in `kccProtoOverrides` in `dev/tools/controllerbuilder/pkg/sourcelinks/sourcelinks.go`.
  2. Registered `ftp` API (`Cloud FTP API`) with prefix `FTP` in `dev/tools/controllerbuilder/serviceprefixes/service-prefixes.yaml`.
  3. Created `apis/ftp/generate.sh` configuring `generate-types`, `generate-identity`, and `generate-mapper` with `--prepopulate-resource`.
  4. Added `FTPServer` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  5. Generated Go protobuf files for `ftp` v1 into `pkg/controller/direct/ftp/pb` and mapped the package in `dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go` to keep `go.mod` untouched and avoid transitive dependency conflicts.
  6. Regenerated channel manifests via `dev/tasks/generate-channels` and operator golden tests via `WRITE_GOLDEN_OUTPUT=1 go test ./operator/pkg/controllers/configconnector/...`.
  7. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for the new `ftpservers.ftp.cnrm.cloud.google.com` CRD.
- **Impact**: Enables Greenfield Step 1a for `FTPServer` with verified conformance tests and passing `apichecks`.
