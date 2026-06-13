# ComputeInstance KRM Transition Journal

## Overview
We successfully implemented direct KRM types, configured the `generate.sh` pipeline, and created/registered a fuzzer for the `ComputeInstance` resource.

## Key Learnings & Design Choices

1. **Flattened Proto Structs vs KRM Fields**:
   - In GCE proto, both boot and attached disks are elements of a single list `disks []AttachedDisk`.
   - In KRM, we represent them as separate fields: `spec.bootDisk` (single struct), `spec.attachedDisk` (list of structs), and `spec.scratchDisk` (list of structs).
   - This requires manually handling them in our high-level `ComputeInstanceSpec` mapping functions.

2. **Empty Structs with Unknown Fields**:
   - `InstanceLabels` and `InstanceResourceManagerTags` are defined as empty structs with `// +kubebuilder:validation:XPreserveUnknownFields`.
   - In the proto, they correspond to standard `map[string]string`.
   - To avoid type assignment mismatches in `mapper.generated.go`, we handcoded custom mapping functions for `InstanceInitializeParams` and `InstanceParams` in `instance_mapper.go`. The generator detects these functions and correctly skips generating conflicting definitions.

3. **Duration Mappings**:
   - `pb.Duration` has fields `Nanos *int32` and `Seconds *int64`.
   - KRM uses `Nanos *int64` and `Seconds int64` (non-pointer) on `InstanceLocalSsdRecoveryTimeout` and `InstanceMaxRunDuration`.
   - Handcoding custom Duration mappers for `localSsdRecoveryTimeout` and `maxRunDuration` resolved the integer and pointer mismatches.

4. **Required Non-pointer string in Ipv6AccessConfig**:
   - `NetworkTier` in `InstanceIpv6AccessConfig` is a non-pointer `string`. In the proto `pb.AccessConfig`, it is `*string`.
   - Handcoding `InstanceIpv6AccessConfig` mappers resolved this type mismatch.
