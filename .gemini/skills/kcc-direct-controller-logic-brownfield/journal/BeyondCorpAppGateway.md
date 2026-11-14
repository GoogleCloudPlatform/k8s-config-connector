# BeyondCorpAppGateway Migration Journal

## Overview
Migrated `BeyondCorpAppGateway` (`beyondcorp.cnrm.cloud.google.com/v1alpha1`) from legacy Terraform controller to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Resource Immutability:**
   - GCP BeyondCorp AppGateway API does not support update operations (there is no `UpdateAppGateway` API or method). All spec fields (`displayName`, `type`, `hostType`, etc.) are immutable in GCP.
   - In `Update()`, `common.CompareBrownfieldSpecAndLabels` detects changes to spec or metadata labels. If a diff is found, it reports the diff via `structuredreporting.ReportDiff` and returns an error explaining that `BeyondCorpAppGateway` is immutable.
   - In steady-state re-reconciliation with no diff, status is updated and no mutating requests are sent to GCP.

2. **Brownfield Labels Reconciliation:**
   - GCP labels are mapped from KRM `metadata.labels` using `label.GCPLabels(u)` in `buildDesired`.
   - Labels are compared in `Update()` via `common.CompareBrownfieldSpecAndLabels`.

3. **Client Selection:**
   - The GCP Go client library `cloud.google.com/go/beyondcorp/appgateways/apiv1` does not provide a REST client (`NewRESTClient` is absent), so `NewClient` is used with `m.config.GRPCClientOptions()`.
   - MockGCP supports both HTTP and gRPC for `beyondcorp.googleapis.com`.

4. **Dual-Testing Alignment:**
   - `pkg/controller/resourceconfig/static_config.go` was updated to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `BeyondCorpAppGateway` while keeping `DefaultController` as `k8s.ReconcilerTypeTerraform`.
   - Golden mock testing runs both Direct and Terraform controllers and passed without errors.
