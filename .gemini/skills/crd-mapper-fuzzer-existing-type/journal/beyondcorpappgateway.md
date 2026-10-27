# BeyondCorpAppGateway Direct KRM Types Journal

## Learnings & Observations

### 1. Schema Compatibility with Baseline CRD
- The baseline CRD `beyondcorpappgateways.beyondcorp.cnrm.cloud.google.com` is `v1alpha1` managed by Terraform (`cnrm.cloud.google.com/tf2crd: "true"`).
- Spec fields:
  - `displayName`: optional string (`*string`)
  - `hostType`: optional string (`*string`)
  - `projectRef`: required Project reference without a kind field, using `refs.ProjectRef` from `apis/refs`.
  - `region`: required string (`string` scalar value)
  - `resourceID`: optional string (`*string`)
  - `type`: optional string (`*string`)
- Status fields:
  - `allocatedConnections`: `[]AppGatewayAllocatedConnectionStatus`
    - `ingressPort`: optional int32 (`*int32`)
    - `pscUri`: optional string (`*string`, with JSON tag `json:"pscUri,omitempty"`)
  - `conditions`: standard KRM conditions
  - `observedGeneration`: optional int64 (`*int64`)
  - `state`: optional string (`*string`)
  - `uri`: optional string (`*string`)
- Top-level doc comments were omitted on `BeyondCorpAppGatewaySpec` and `BeyondCorpAppGatewayStatus` to prevent generating `spec` and `status` descriptions that would trigger `_no_description` diffs in `diff-crds`.
- Baseline labels `cnrm.cloud.google.com/tf2crd: "true"` and `cnrm.cloud.google.com/stability-level: alpha` were preserved.
- Running `dev/tasks/diff-crds` produced zero diffs.

### 2. Multi-Resource Service generate.sh Pattern
- The BeyondCorp service (`apis/beyondcorp`) hosts multiple resources (`BeyondCorpClientConnectorService` and `BeyondCorpAppGateway`) with different proto packages (`google.cloud.beyondcorp.clientconnectorservices.v1` and `google.cloud.beyondcorp.appgateways.v1`).
- Each resource uses a resource-specific generated types file: `clientconnectorservice_types.generated.go` and `appgateway_types.generated.go`.
- `apis/beyondcorp/generate.sh` was updated to generate types for `BeyondCorpAppGateway` with `--include-skipped-output`, preserving skipped and unreachable types as comments in `appgateway_types.generated.go` for future reference.
