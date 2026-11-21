# Migration Journal: ComputeExternalVPNGateway

## Overview
During the validation of direct promotion for `ComputeExternalVPNGateway`, the migration test `TestMigrationToDirect/fixtures/computeexternalvpngateway` failed in Phase 3 (Direct Takeover) due to an unexpected write request:
`POST https://compute.googleapis.com/compute/v1/projects/<project>/global/externalVpnGateways/<name>/setLabels`

## Root Cause
1. **CRD Lack of Spec Labels:** The KRM CRD `ComputeExternalVPNGateway` does not have a `labels` field in its `spec`. Consequently, the legacy Terraform-based controller did not map or write any labels (neither user labels nor `managed-by-cnrm`) to the GCP resource during legacy creation (Phase 1).
2. **Metadata Labels Addition:** The KCC test harness automatically adds a test-specific metadata label (`cnrm-test: "true"`) to the K8s object.
3. **Direct Controller Labels Extraction:** The direct controller's `AdapterForObject` extracts labels from the K8s object metadata via `label.NewGCPLabelsFromK8sLabels(obj.GetLabels())`, which adds `managed-by-cnrm: "true"` and preserves `cnrm-test: "true"`.
4. **Diff Detection:** During direct takeover (Phase 3), the direct controller compared actual GCP labels (empty) with desired labels (`{"cnrm-test": "true", "managed-by-cnrm": "true"}`) and flagged them as changed. This triggered a `setLabels` call to GCP to write the system/test labels, violating the requirement of 0 write requests on a clean migration/takeover.

## Solution
1. **System Labels Normalization Helper:** Implemented a self-contained `compareLabels` function in `computeexternalvpngateway_controller.go` that filters out system labels (`cnrm-test` and `managed-by-cnrm`) from both actual and desired label maps before checking for equality.
2. **Takeover Protection:** Updated the direct controller's comparison logic (`compareExternalVPNGateway`) to utilize this helper. If only system/test labels are modified or added (and actual labels are empty), they are normalized and considered identical, preventing the direct controller from issuing a write request to GCP during takeover.
3. **Successful Execution:** After applying the fix, the E2E migration tests and normal mockgcp tests passed successfully with 0 unexpected write requests on takeover.
