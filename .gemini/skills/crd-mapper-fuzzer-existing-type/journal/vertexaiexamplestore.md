# Journal Entry: VertexAIExampleStore Direct KRM Types Implementation

## Overview
Successfully configured the code generation script and aligned the generated structures for `VertexAIExampleStore` in `apis/vertexai/v1alpha1/` with zero CRD schema diffs and full workspace compilation.

## Key Observations and Solutions

1. **Re-generating Types with Comments causing Mapper Conflicts:**
   - When running `generate-types` with `--include-skipped-output`, the skipped types are output as comments inside `types.generated.go`.
   - Although commented out as `/* ... */`, Go parsing or search logic inside the `controllerbuilder`'s `generate-mapper` tool still parsed comments and recognized the annotations (e.g. `+kcc:proto=...`), which caused conflicting and duplicate mapper functions (like `ExampleStore_v1alpha1_FromProto`) to be generated in the shared `pkg/controller/direct/vertexai/mapper.generated.go` file.
   - **Solution:** Removed the `--include-skipped-output` flag from the `v1alpha1/generate.sh` type generator run to prevent comments of skipped types from being output in `types.generated.go`, while keeping the configuration clean.

2. **Ensuring Strict Schema Compatibility of Shared Common Types:**
   - The type generator special-cases `google.rpc.Status` and automatically replaces local occurrences with `apis/common.Status` or comments out the local `Status` definition.
   - This broke hand-written controllers and mappers like `vertexaispecialistpool_mapper` and `vertexaidatalabelingjob_types` which expected a local `Status` struct in package `v1alpha1`.
   - **Solution:** Created a hand-coded file `apis/vertexai/v1alpha1/status_types.go` that defines the expected local `Status`, `Any`, and legacy `ExampleStore` types with `Copyright 2026 Google LLC` and Apache 2.0 license headers. This perfectly matches the baseline CRD schema and maintains complete compilation compatibility for all existing controllers and packages.
