# ComputeSubnetwork Fuzzer Journal

## Observations & Learnings

- **Fuzzer Naming Convention:** The previous fuzzer implementation for `ComputeSubnetwork` was named `computesubnetwork_fuzzer.go`. In alignment with the expected path and typical pattern where the directory is `compute` and the resource name is `subnetwork`, we renamed/moved the file to `pkg/controller/direct/compute/subnetwork_fuzzer.go`.
- **Proto Mapping Details:** All Spec and Status fields have been fully mapped and validated under the central fuzz test suite. Unimplemented proto fields or internal implementation details (e.g., `allow_subnet_cidr_routes_overlap`, `utilization_details`, `log_config.enable`, and `secondary_ip_ranges[].reserved_internal_range`) have been safely excluded using `f.Unimplemented_NotYetTriaged`.
- **Field Comparisons Added:** Added a structured, comprehensive comment comparing the KRM Spec type fields with their corresponding fuzzer/proto field mapping paths to assist future maintainability and review.
