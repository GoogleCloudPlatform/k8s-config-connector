# ComputeMachineImage Migration Journal

## Overview
Migrated `ComputeMachineImage` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Resource Immutability:**
   - GCP Compute Engine `MachineImages` API does not support updating existing machine images (only Insert, Get, Delete, SetLabels, SetIamPolicy).
   - In KCC, all spec fields (`description`, `guestFlush`, `machineImageEncryptionKey`, `sourceInstanceRef`, `resourceID`) are marked immutable.
   - The direct controller's `Update` method checks for diffs using `common.CompareBrownfieldSpec`. If diffs are detected, it returns an error stating that `ComputeMachineImage is immutable and cannot be updated`. If no diffs are found, it updates the resource status to Ready.

2. **Source Instance URL Normalization:**
   - GCP returns the `sourceInstance` as a full URL (e.g. `https://www.googleapis.com/compute/v1/projects/.../zones/.../instances/...`).
   - Using `refs.TrimComputeURIPrefix` ensures it compares cleanly against KCC's canonical relative path format `projects/.../zones/.../instances/...` during steady-state diffing and export.

3. **Controller Configuration:**
   - Kept `DefaultController: k8s.ReconcilerTypeTerraform` and added `k8s.ReconcilerTypeDirect` to `SupportedControllers` in `static_config.go` for gradual brownfield adoption.
   - Tests automatically run both controllers and verify diffs between direct and legacy executions.
