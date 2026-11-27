# NetAppVolume Journal Entry

### [2026-11-27] Greenfield: Prepopulate direct KRM types, identity, and judgement queue for NetAppVolume

- **Context**: Prepopulating direct KRM types, IdentityV2, and `judgement_queue.yaml` for `NetAppVolume` under `netapp.cnrm.cloud.google.com/v1alpha1` following Greenfield Step 1a.
- **Client Library**: Updated `cloud.google.com/go/netapp` from `v1.16.0` to `v1.21.1` in both `go.mod` and `mockgcp/go.mod` to support newer fields.
- **Protobuf Field Name Mapping**: In `cloud.google.com/go/netapp/apiv1/netapppb`, `SimpleExportPolicyRule` contains fields like `Kerberos_5ReadOnly` with an underscore (`Kerberos_5*`), whereas `controllerbuilder` generates Go field names as `Kerberos5*`. To resolve compilation errors in `mapper.generated.go`, handcrafted mapping functions `SimpleExportPolicyRule_FromProto` and `SimpleExportPolicyRule_ToProto` were created in `pkg/controller/direct/netapp/netappvolume_mapper.go`. `controllerbuilder` automatically detects existing mapping functions and skips generating conflicting duplicates in `mapper.generated.go`.
- **Release Bundle**: Added `NetAppVolume` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Validation**:
  - `apis/netapp/...` unit tests pass.
  - `pkg/cli/preview/...` passes with envtest.
  - `tests/apichecks/...` passes.
  - `ONLY_TEST_APIGROUPS=netapp.cnrm.cloud.google.com dev/ci/presubmits/tests-e2e-fixtures` passes.
  - `make fmt` and `make generate` clean.
