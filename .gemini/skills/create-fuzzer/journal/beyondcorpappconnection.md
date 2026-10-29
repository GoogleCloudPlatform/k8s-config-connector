# Fuzzer Journal: BeyondCorpAppConnection

## Observations & Learnings

### 1. Spec-Only Resource
`BeyondCorpAppConnection` does not have an `ObservedState` in its Status (Status only contains standard `conditions`, `observedGeneration`, and `externalRef`). In accordance with the skill guidelines for spec-only resources, we instantiated the fuzzer using `fuzztesting.NewKRMTypedSpecFuzzer` and registered it via `fuzztesting.RegisterKRMSpecFuzzer`.

### 2. Multi-Proto Service Group
BeyondCorp resources are split across distinct protobuf packages under `google.cloud.beyondcorp`:
- `google.cloud.beyondcorp.appconnections.v1` (`BeyondCorpAppConnection`)
- `google.cloud.beyondcorp.clientconnectorservices.v1` (`BeyondCorpClientConnectorService`)
- `google.cloud.beyondcorp.appconnectors.v1` (`BeyondCorpAppConnector`)
- `google.cloud.beyondcorp.appgateways.v1` (`BeyondCorpAppGateway`)

Running `controllerbuilder generate-mapper` with `--service google.cloud.beyondcorp.appconnections.v1` targets `pkg/controller/direct/beyondcorp/mapper.generated.go` directly under package `beyondcorp`. We updated `apis/beyondcorp/generate.sh` to include this `generate-mapper` command.

### 3. Int32 vs Go Int Type Mismatches
In `AppConnectionApplicationEndpoint` (`port`) and `AppConnectionGateway` (`ingressPort`), KRM uses Go `int` to ensure OpenAPI schema compatibility (`type: integer` without `format: int32`). The protobuf definition uses `int32`. Handcoded mappers in `pkg/controller/direct/beyondcorp/beyondcorpappconnection_mappings.go` handle conversion using:
- `int(in.GetPort())` / `int32(in.Port)`
- `direct.LazyPtr(int(in.GetIngressPort()))` / `int32(direct.ValueOf(in.IngressPort))`

The generator automatically detected these existing functions and skipped generating them in `mapper.generated.go`.

### 4. Output-Only Proto Fields Located in Spec
In protobuf, `AppConnection.Gateway` fields `ingress_port` and `uri` are marked `OUTPUT_ONLY`. In Config Connector, these fields reside under `spec.gateway`. They are mapped as `f.SpecField` in `beyondcorpappconnection_fuzzer.go` so they are fully verified by the round-trip fuzz test.

### 5. Package Registration
Imported `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/beyondcorp"` in alphabetical order in `pkg/controller/direct/register/register.go` so the fuzzer executes under the central fuzz test suite.
