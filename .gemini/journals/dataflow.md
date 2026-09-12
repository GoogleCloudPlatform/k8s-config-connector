# Dataflow Journal

### 2026-07-28 Refactoring of DataflowJob KRM fuzzer and verification of Direct Controller execution
- **Context**: Refactoring the DataflowJob KRM round-trip fuzzer to use fluent and highly descriptive helper methods as required by our controller guidelines, and verifying direct controller integration with MockGCP.
- **Problem**: Unimplemented fields were previously directly inserted into the `UnimplementedFields` set. Standardizing them to categorization methods (e.g., `Unimplemented_Identity`, `Unimplemented_LabelsAnnotations`, `Unimplemented_NotYetTriaged`) makes the reasons for non-coverage clear and self-documenting.
- **Solution**: Modified `dataflowjob_fuzzer.go` to use:
  - `f.Unimplemented_Identity(".project_id")`
  - `f.Unimplemented_LabelsAnnotations(".labels")`
  - `f.Unimplemented_NotYetTriaged(...)` for all remaining unmapped fields.
  Vetted changes through the `fuzz-roundtrippers` presubmit suite and verified that all 6 dual-reconciler (Terraform and Direct) DataflowJob E2E test fixtures run and pass against MockGCP.
- **Impact**: Establishes a highly idiomatic and compliant implementation pattern for the DataflowJob fuzzer, making it easily maintainable and aligned with modern KCC controller-building standards.
