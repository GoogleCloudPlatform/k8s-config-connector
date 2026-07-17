# ComputeImage Direct KRM Transition Journal

## Observations
- `ComputeImage` represents a VM image in Google Compute Engine, which is a global resource (so no `Location` or `ProjectRef` was present in the original baseline CRD schema).
- To maintain **strict schema-compatibility**:
  - We omitted `spec.projectRef` and `spec.location` fields from the spec, keeping exactly the original set of fields.
  - We omitted `status.observedState` and `status.externalRef` fields from the status because they weren't in the baseline CRD.
  - Retained all existing labels on the struct definition including `cnrm.cloud.google.com/tf2crd=true`.
- **Field Name Matching**:
  - `Image` proto fields like `guest_os_features` and `disk_size_gb` mapped to `GuestOSFeatures` and `DiskSizeGB` in Golang, while KRM uses `GuestOsFeatures` and `DiskSizeGb`.
  - We hand-coded these inside a custom mapper file `pkg/controller/direct/compute/computeimage_mapper.go` to handle custom reference mappings seamlessly.
  - Resolved `pb.RawDisk.Sha1Checksum` to `krm.ImageRawDisk.Sha1`.
- **Fuzzer Implementation**:
  - Registered fuzzer in `pkg/controller/direct/compute/computeimage_fuzzer.go`.
  - Configured `FilterStatus` to clear spec-specific fields from being randomized/asserted in status roundtrips.
  - Cleaned up empty rawDisk/encryptionKey structs in `FilterSpec`.
  - The fuzzing suite passed successfully!
