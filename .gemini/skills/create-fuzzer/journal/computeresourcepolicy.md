# ComputeResourcePolicy Fuzzer Journal

## Observations & Learnings

- **Fuzzer Naming Convention:** The fuzzer file is placed at `pkg/controller/direct/compute/computeresourcepolicy_fuzzer.go` following the descriptive full lowercase kind name format, as specified in the issue requirements and instructions.
- **Detailed Field Mapping Comments:** We added a structured, comprehensive comment comparing the KRM Spec and Status type fields with their corresponding fuzzer/proto field mapping paths directly above the `SpecField` declarations to ensure high maintainability and reviewer clarity. This detailed documentation maps:
  - All 31 spec field variants under daily, hourly, and weekly cycles.
  - Disk consistency, placement, and schedule sub-policy blocks.
  - Snapshot properties including labels, storage locations, guest flush, and chain name.
  - Creation timestamp, status, and ID fields in `ObservedState`.
- **Verification:** The fuzzer compiles successfully and passes centrally via the central fuzz testing suite (`go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`).
