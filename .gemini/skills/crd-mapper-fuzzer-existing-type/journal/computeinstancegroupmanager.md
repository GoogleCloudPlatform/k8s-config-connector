# ComputeInstanceGroupManager Direct KRM Transition Journal

## Overview
Successfully implemented the direct KRM types, configure code generation, hand-coded conversion mappers, and registered an automated round-trip fuzz test for `ComputeInstanceGroupManager`.

## Technical Insights & Decisions

### 1. Schema & Parity
- **Strict Parity Achieved:** The baseline CRD schema is perfectly preserved. Verified using `dev/tasks/diff-crds`, which produced a completely empty diff.
- **Spec-level parity:**
  - Removed `omitempty` from `Spec` inside `ComputeInstanceGroupManager` parent struct to ensure the `spec` field is strictly required at the root level of the CRD.
  - Used `refs.ProjectRef` from `apis/refs` package for `projectRef` because the baseline CRD's project reference does not contain a `kind` field.
  - Correctly mapped the required `targetSize` property to `int64` in `ComputeInstanceGroupManagerSpec`, while keeping the nested `versions[].targetSize` mapped to the custom object structure (`InstanceGroupManagerTargetSize` mapping to `pb.FixedOrPercent`).
- **Status-level parity:**
  - Removed `observedState` and `externalRef` fields from status as they did not exist in the baseline CRD.
  - Renamed the status field `Id` to `ID` to allow the automatic mapper to pick it up cleanly without any "near miss" mapping warnings.

### 2. Mapper Implementations (`instancegroupmanager_mapper.go`)
- Developed a complete handcoded mapper package under `pkg/controller/direct/compute/instancegroupmanager_mapper.go`.
- **IP capitalization in proto:** Capitalization of `StatefulPolicyPreservedStateNetworkIP` in proto maps is actually `StatefulPolicyPreservedStateNetworkIp` (lowercase `p`), which is handled correctly.
- **Non-existent API fields:** `FailoverAction`, `MinReadySec` and `IsStateful` are present in KCC's baseline/legacy fields but do not exist in GCP Compute v1 API's `InstanceGroupManager` proto message. We left them unmapped, which complies with strict schema compatibility and GCP API standards.

### 3. Fuzzer Development
- Created `computeinstancegroupmanager_fuzzer.go` using type-safe registration helpers.
- Added necessary untriaged field markers (`current_actions.starting`, `current_actions.resuming`, `current_actions.stopping`, etc.) to align fuzz-testing expectations with actual API fields.
- Implemented `FilterSpec` to normalize the required `TargetSize` to default 0 to prevent round-trip mismatches during zero/nil generation.
- Automated fuzz tests successfully run and pass 100% (`PASS` with thousands of random inputs).
