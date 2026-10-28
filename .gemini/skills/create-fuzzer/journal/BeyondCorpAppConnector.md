# Journal: BeyondCorpAppConnector Fuzzer Implementation

## Overview
Implemented round-trip KRM fuzzer for `BeyondCorpAppConnector` (`beyondcorp.cnrm.cloud.google.com/v1alpha1`) targeting `google.cloud.beyondcorp.appconnectors.v1.AppConnector` (`cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb.AppConnector`).

## Proto & KRM Mapping Details
- **KRM Spec Mapping**:
  - `displayName`: maps to `.display_name` in proto (`direct.LazyPtr` / `direct.ValueOf`).
  - `principalInfo`: maps to `.principal_info` in proto.
    - `principalInfo.serviceAccount`: maps to `.principal_info.service_account` oneof (`*pb.AppConnector_PrincipalInfo_ServiceAccount_`).
    - `principalInfo.serviceAccount.email`: maps to `.principal_info.service_account.email` (string).
  - `projectRef`: KRM spec-only GCP project reference (not in `pb.AppConnector`).
  - `region`: KRM spec-only GCP region/location (resource identity).
  - `resourceID`: KRM resource ID (maps to resource identity name).
- **Proto Status Fields**:
  - `.create_time`: Marked as status field (`f.StatusField(".create_time")`).
  - `.update_time`: Marked as status field (`f.StatusField(".update_time")`).
  - `.uid`: Marked as status field (`f.StatusField(".uid")`).
  - `.state`: Marked as status field (`f.StatusField(".state")`).
- **Identity Fields**:
  - `.name`: Resource name/ID GCP field, marked with `f.Unimplemented_Identity(".name")`.
- **Unimplemented Fields**:
  - `.labels`: Resource labels managed via KRM `metadata.labels`, marked with `f.Unimplemented_LabelsAnnotations(".labels")`.
  - `.resource_info`: Optional connector runtime resource info, not yet triaged/supported in KRM Spec or Status, marked with `f.Unimplemented_NotYetTriaged(".resource_info")`.

## Spec-Only Fuzzing Implementation
- `BeyondCorpAppConnector` does not have an `ObservedState` sub-struct in `BeyondCorpAppConnectorStatus`.
- Implemented `BeyondCorpAppConnectorSpec_FromProto` and `BeyondCorpAppConnectorSpec_ToProto` (along with helper mappers for `AppConnectorPrincipalInfo` and `AppConnectorPrincipalInfoServiceAccount`) in `pkg/controller/direct/beyondcorp/beyondcorpappconnector_mapper.go`.
- Registered using `fuzztesting.NewKRMTypedSpecFuzzer` and `fuzztesting.RegisterKRMSpecFuzzer` in `pkg/controller/direct/beyondcorp/beyondcorpappconnector_fuzzer.go`.
- Added `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/beyondcorp"` to `pkg/controller/direct/register/register.go` in alphabetical order.

## Verification
- Verified with `FOCUS=BeyondCorpAppConnector go test -count=100 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` (all 100 runs passed).
- Verified with `go test -v ./pkg/fuzztesting/fuzztests/ -run TestListTypesWithoutFuzzers`, confirming `BeyondCorpAppConnector` is now recognized as covered by the central fuzz test suite and no longer in the missing list.
