# BeyondCorpAppConnector Migration Journal

## Overview
Implemented the direct controller for `BeyondCorpAppConnector` under `pkg/controller/direct/beyondcorp`, migrating it from legacy Terraform to direct controller support while keeping legacy as the default controller during the transition period.

## Implementation Details
1. **Controller Architecture:**
   - Package: `pkg/controller/direct/beyondcorp`.
   - Utilizes `cloud.google.com/go/beyondcorp/appconnectors/apiv1` (`AppConnectorsClient`) initialized with `m.config.GRPCClientOptions()`.
   - Standard CRUD methods:
     - `Find`: Fetches existing AppConnector by resource name; correctly handles 404 with `direct.IsNotFound(err)`.
     - `Create`: Builds desired proto with `buildDesired`, creates resource with LRO waiting via `op.Wait(ctx)`, and updates KRM status.
     - `Update`: Builds desired proto and utilizes `common.CompareBrownfieldSpecAndLabels` with `"labels"` field name. Submits `UpdateAppConnectorRequest` with field mask and waits for LRO completion.
     - `Delete`: Submits `DeleteAppConnectorRequest`, handles 404 gracefully for idempotency, and waits for LRO completion.
     - `updateStatus`: Updates `status.state` and sets `status.externalRef`.

2. **Labels Handling:**
   - Follows `kcc-direct-brownfield-labels`:
     - Reads labels from `metadata.labels` via `label.GCPLabels(u)` in `buildDesired`.
     - Diffs spec and labels together via `common.CompareBrownfieldSpecAndLabels`.

3. **Dual Controller Registration:**
   - Added `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `BeyondCorpAppConnector` in `pkg/controller/resourceconfig/static_config.go`.
   - Preserved `DefaultController: k8s.ReconcilerTypeTerraform`.

4. **Testing and Verification:**
   - Golden object updated with `externalRef`.
   - Verified roundtrip fuzzing passes (`TestFocusedMappers`).
   - Verified CRD templates (`pkg/crd/template`) and API checks (`tests/apichecks`).
   - Verified e2e testing against MockGCP (`TestAllInSeries` runs both direct and legacy controllers successfully).
   - Golden log alignment verified (`TestGoldenLogAlignment`).
