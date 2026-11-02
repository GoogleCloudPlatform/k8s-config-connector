# ComputeMachineImage MockGCP Alignment Journal

## Overview
This document records observations and learnings from implementing MockGCP support and aligning real GCP behavior for `ComputeMachineImage`.

## Key Observations
1. **MachineImages Proto Availability**:
   - The proto definitions and gRPC client/server interfaces for `MachineImages` were already present in `mockgcp/generated/mockgcp/cloud/compute/v1` (`pb.RegisterMachineImagesServer` and `pb.RegisterMachineImagesHandler`).
2. **ComputeInstance Dependency Fixes**:
   - The test fixture `machineimageautogen` had legacy schema violations for `ComputeInstance` (`bootDisk` as a list rather than an object with `initializeParams`, and missing `zone`).
   - The fixture directory was renamed to `computemachineimageautogen` to comply with the basic test case naming standard that requires directory names to contain the target resource kind name.
   - In real GCP, Debian 11 images are no longer available in the `debian-cloud` project in 2026; updating to `debian-cloud/debian-12` resolved the instance creation failure.
3. **Response Structure for MachineImages**:
   - When a machine image is created from a source instance, GCE populates `instanceProperties`, `sourceInstanceProperties`, and `savedDisks`.
   - `instanceProperties.disks` includes `savedState: "DISK_SAVED_STATE_UNSPECIFIED"`.
   - `sourceInstanceProperties.disks` includes `kind: "compute#savedAttachedDisk"`, `storageBytes`, `storageBytesStatus: "UP_TO_DATE"`, and `diskType: "pd-standard"`.
   - The licenses field under `sourceInstanceProperties.disks` references `projects/${projectId}/global/licenses/debian-12-bookworm`.
   - `metadata` in `instanceProperties` and `sourceInstanceProperties` contains `{fingerprint: "x4iOk7cA2GY=", kind: "compute#metadata"}` even when no custom metadata was specified.
   - `tags` contains `{fingerprint: "42WmSpB8rSM="}`.
4. **Golden Alignment Placeholder**:
   - Added `ComputeMachineImage` with placeholder `${machineImageID}` to `getPlaceholdersForKind` in `pkg/test/resourcefixture/golden_alignment_test.go`.
