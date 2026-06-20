# ComputeRouter Fuzzer Journal

## Observations & Learnings

- **Existing Fuzzer Refinement:** The direct controller for `ComputeRouter` already had a partial implementation in `pkg/controller/direct/compute/computerouter_fuzzer.go` that compiled and passed basic round-trip checks.
- **Spec Fields Mapping Documentation:** Added clear, structured comments above the `f.SpecField` calls detailing the 1:1 and nested mapping between `ComputeRouterSpec` in KRM and `pb.Router` in the proto.
- **Nested Field Registrations:** Added registrations for nested fields under the `bgp.advertised_ip_ranges` list (such as `.bgp.advertised_ip_ranges[].description` and `.bgp.advertised_ip_ranges[].range`) to match best practices from other mature fuzzers like `ComputeSubnetwork`.
- **Validation:** Successfully verified all changes with the central fuzz test runner using `FOCUS=ComputeRouter go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`.
