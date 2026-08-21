# ComputeRouterNAT Fuzzer Journal

- **Kind**: ComputeRouterNAT
- **Date**: 2026-06-26
- **Observations**:
  - `RouternatLogConfig.Enable` is a non-pointer `bool` in KRM, which is a required field under `logConfig` in the CRD.
  - To prevent breaking CRD equivalence/compatibility with master, this field must remain `bool` in KRM instead of changing to `*bool`.
  - In the mapper, we map to Proto using `out.Enable = &in.Enable` (instead of `direct.LazyPtr(in.Enable)`) to ensure `false` values are not discarded as nil/omitted.
  - In `FilterSpec` under `computerouternat_fuzzer.go`, we normalize `in.LogConfig.Enable` to `*false` if it is `nil` and `in.LogConfig` is not nil. This ensures that the original Proto and the round-tripped Proto match perfectly.
  - No Status fields were found in the Proto that needed mapping to KRM, other than standard conditions and observed generation.
  - Identity is handled via `resourceID` in `Spec` and mapped to `.name` in Proto.
