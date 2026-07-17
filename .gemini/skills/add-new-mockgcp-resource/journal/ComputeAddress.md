# ComputeAddress MockGCP Alignment Journal

- **Date:** June 19, 2026
- **Resource Kind:** `ComputeAddress`
- **GCP Resource:** `compute, addresses`

## Overview

Existing mock implementation for `ComputeAddress` already exists under `mockgcp/mockcompute` (split into `regionaladdress.go` and `globaladdress.go`).

## Discrepancies and Findings

Upon review of the `SetLabels` operations in MockGCP for both regional and global `ComputeAddress`, it was identified that `LabelFingerprint` was not being updated on label changes:
- In real GCP, updating labels via `SetLabels` computes and updates the `LabelFingerprint`.
- Without this update, the mock's stored `Address` resource would retain an outdated or missing `LabelFingerprint`.

## Solution

We updated the `SetLabels` method in both files to compute and update `LabelFingerprint`:
- **File:** `mockgcp/mockcompute/regionaladdress.go`
- **File:** `mockgcp/mockcompute/globaladdress.go`
- **Code Change:**
  ```go
  obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels() // (or GetGlobalSetLabelsRequestResource)
  obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))
  ```

## Validation

All E2E fixture tests for `ComputeAddress` pass cleanly against the updated MockGCP backend:
- `TestAllInSeries/fixtures/regionalcomputeaddress`
- `TestAllInSeries/fixtures/globalcomputeaddress`
- `TestAllInSeries/fixtures/computeaddressipcollection`
