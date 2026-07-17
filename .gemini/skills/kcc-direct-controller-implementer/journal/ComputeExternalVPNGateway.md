# ComputeExternalVPNGateway Direct Controller Journal

## Details
- **Kind**: `ComputeExternalVPNGateway`
- **GCP Service**: `compute`
- **Go API Package**: `cloud.google.com/go/compute/apiv1` / `computepb`
- **KRM Version**: `v1beta1`

## Learnings & Observations

### 1. Immutability & SetLabels Pattern
- The spec fields for `ComputeExternalVPNGateway` (`description`, `redundancyType`, `interface`) are completely immutable on GCP.
- However, metadata labels can be updated.
- To handle this, we implemented a custom comparison in the `Update` method:
  - Cloned and cleared `labels` from both `actual` and `desired` state.
  - Ran `tags.DiffForTopLevelFields` on spec/other fields first.
  - If a change is found in any field other than labels, returned a descriptive error ("ComputeExternalVPNGateway is immutable and cannot be updated").
  - If only labels changed, executed a standard `SetLabelsExternalVpnGatewayRequest` via the `ExternalVpnGatewaysClient` REST client, and waited for the LRO to complete.

### 2. MockGCP Integration & Harness Whitelist
- To run and verify the E2E fixture tests against the hermetic MockGCP layer, the GroupKind `compute.cnrm.cloud.google.com/ComputeExternalVPNGateway` must be explicitly listed in the whitelist switch within `config/tests/samples/create/harness.go`.
- Without this, the test runner defaults to skipping the test when running against the mock target, logging `"gk ComputeExternalVPNGateway.compute.cnrm.cloud.google.com not supported by mock gcp"`.
