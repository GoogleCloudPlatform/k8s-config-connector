# Journal: Direct KRM Types Implementation for BeyondCorpAppConnection

## Learnings & Observations

### 1. Multi-Proto Services Under A Single API Service Group
BeyondCorp resources in GCP are split across distinct protobuf packages:
- `BeyondCorpClientConnectorService`: `google.cloud.beyondcorp.clientconnectorservices.v1`
- `BeyondCorpAppConnection`: `google.cloud.beyondcorp.appconnections.v1`

In `apis/beyondcorp/generate.sh`, each resource's types are generated using its respective proto package with `generate-types`. By saving generated type files as `<resource>_types.generated.go` (e.g., `clientconnectorservice_types.generated.go` and `appconnection_types.generated.go`), both resources maintain their auto-generated reference types and skipped-type comment blocks without clobbering each other.

### 2. Output-Only Proto Fields Located in Spec
In the protobuf definition for `AppConnection.Gateway`, `ingress_port` and `uri` are marked `OUTPUT_ONLY`. However, in the baseline CRD, all 4 gateway fields (`appGateway`, `ingressPort`, `type`, `uri`) reside under `spec.gateway`. Defining `AppConnectionGateway` directly in `beyondcorpappconnection_types.go` ensures strict schema parity with the baseline CRD without incorrectly surfacing those fields under `status`.

### 3. ProjectRef Without Kind Property
The baseline CRD for `BeyondCorpAppConnection` defines `projectRef` without a `kind` field. Importing and using `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` matched the OpenAPI schema signature (`external,name,namespace`), prompting `scripts/add-validation-to-crds` to generate the exact `oneOf` validation constraint without schema deviations.

### 4. Integer Type Format Compatibility
In the baseline CRD, `spec.applicationEndpoint.port` and `spec.gateway.ingressPort` are defined as `type: integer` without a format specifier. Using Go `int` rather than `int32` generates `type: integer` without `format: int32`, preserving strict backward compatibility tested by `TestCRDBackwardCompatibility`.

### 5. Zero Schema Diff
Running `./dev/tasks/diff-crds --base master` verified zero schema changes between the baseline CRD and the generated CRD.
