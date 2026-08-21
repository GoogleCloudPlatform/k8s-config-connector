# ComputeInstanceGroup Fuzzer Journal

## Observations & Learnings

- **Fuzzer Naming Convention:** The fuzzer for `ComputeInstanceGroup` is named `computeinstancegroup_fuzzer.go` using its full lowercase kind name to avoid ambiguity and retain maximum clarity, aligning with the project's engineering standards.
- **Unmapped Fields:** KRM `spec.instances` is not represented in the `pb.InstanceGroup` protobuf message since instances membership is managed via separate GCP API endpoints/RPCs. This is explicitly noted in the detailed field comparison comment within the fuzzer.
- **Field Comparison Documentation:** We updated the fuzzer file to add detailed comparison comments explicitly mapping KRM Spec and Status fields to the `pb.InstanceGroup` proto fields.
- **Verification:** Ran focused fuzz tests over 100 seeds for both Spec and Status:
  `FOCUS=ComputeInstanceGroup go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`
  The tests passed flawlessly.
