# BeyondCorpAppGateway Fuzzer Journal

## Observations & Implementation Details

1. **Service Package Structure & Isolation:**
   - BeyondCorp hosts multiple resources with distinct GCP APIs and protobuf packages:
     - `BeyondCorpAppGateway` -> `google.cloud.beyondcorp.appgateways.v1` (`cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb`)
     - `BeyondCorpClientConnectorService` -> `google.cloud.beyondcorp.clientconnectorservices.v1`
   - Following the package isolation pattern established for `BeyondCorpClientConnectorService` (and formerly `BeyondCorpClientGateway`), the controller, mappers, and fuzzer are organized in a kind-specific subpackage: `pkg/controller/direct/beyondcorp/beyondcorpappgateway/`.
   - The file names strictly follow `TestDirectResourceFileNaming` (`tests/apichecks/naming_test.go`):
     - `pkg/controller/direct/beyondcorp/beyondcorpappgateway/beyondcorpappgateway_mapper.go`
     - `pkg/controller/direct/beyondcorp/beyondcorpappgateway/beyondcorpappgateway_fuzzer.go`

2. **Proto & KRM Field Mappings:**
   - Proto message: `google.cloud.beyondcorp.appgateways.v1.AppGateway`
   - KRM Spec fields:
     - `displayName` maps to `.display_name` (string)
     - `type` maps to `.type` (`AppGateway_Type` enum via `direct.Enum_FromProto` / `direct.Enum_ToProto`)
     - `hostType` maps to `.host_type` (`AppGateway_HostType` enum via `direct.Enum_FromProto` / `direct.Enum_ToProto`)
     - `projectRef`, `region`, `resourceID` are KRM-only identity fields used for URL construction and resource acquisition.
   - KRM Status fields:
     - `state` maps to `.state` (`AppGateway_State` enum via `direct.Enum_FromProto` / `direct.Enum_ToProto`)
     - `uri` maps to `.uri` (string)
     - `allocatedConnections` maps to `.allocated_connections` (`[]*AppGateway_AllocatedConnection`)
       - `allocatedConnections[].ingressPort` maps to `.allocated_connections[].ingress_port` (int32)
       - `allocatedConnections[].pscUri` maps to `.allocated_connections[].psc_uri` (string)
     - `conditions`, `observedGeneration`, `externalRef` are KCC-internal status fields.
   - Proto fields not in KRM:
     - `.name`: Identity field (`f.Unimplemented_Identity(".name")`)
     - `.labels`: Resource labels (`f.Unimplemented_LabelsAnnotations(".labels")`)
     - `.create_time`, `.update_time`, `.uid`: GCP server-generated output-only fields (`f.Unimplemented_NotYetTriaged`)

3. **Status vs. ObservedState Handling:**
   - In `BeyondCorpAppGateway`'s brownfield type definition (`apis/beyondcorp/v1alpha1/beyondcorpappgateway_types.go`), status fields are annotated directly on `BeyondCorpAppGatewayStatus` (`// +kcc:status:proto=google.cloud.beyondcorp.appgateways.v1.AppGateway`), while `BeyondCorpAppGatewayObservedState` is an empty struct.
   - `fuzztesting.NewKRMTypedFuzzer` was configured with `krm.BeyondCorpAppGatewayStatus`, mapping `State`, `Uri`, and `AllocatedConnections` to ensure full round-trip verification of observed status fields.

4. **Acronym Casing & Hand-coded Mappers:**
   - The baseline CRD schema uses `pscUri` and `uri` in JSON tags, corresponding to `PscUri` and `Uri` Go struct fields in `apis/beyondcorp/v1alpha1/beyondcorpappgateway_types.go`.
   - Automatic `generate-mapper` has near-misses on `URI` vs `Uri` and `PSCURI` vs `PscUri`. Hand-coding the mappers cleanly in `beyondcorpappgateway_mapper.go` provides exact, type-safe mapping without requiring schema-altering acronym modifications.

5. **Registration & Verification:**
   - Registered package anonymously in `pkg/controller/direct/register/register.go`.
   - Verified with focused fuzzer test: `FOCUS=BeyondCorpAppGateway go test -count=100 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` (100/100 passes).
   - Verified that `BeyondCorpAppGateway` is removed from `TestListTypesWithoutFuzzers` output.
   - Verified `TestDirectResourceFileNaming` and `tests/apichecks/...` all pass.
