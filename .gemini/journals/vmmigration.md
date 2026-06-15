### [2026-06-15] VMMigrationGroup Types Generation
- **Context**: Implementing the direct KRM types and Identity V2 for VMMigrationGroup in KCC.
- **Problem**: When running `generate.sh` for VMMigrationGroup, the types `Group` and `GroupObservedState` generated from proto were originally placed in `types.generated.go` but commented out as unreachable because they weren't explicitly referenced by the top-level spec or status.
- **Solution**: Explicitly defined the spec and observed state fields in `group_types.go` utilizing the corresponding proto fields and `+kcc:proto:field` mappings, which made them reachable and cleared the commented-out unreachable sections.
- **Impact**: Demonstrates a clean greenfield resource types setup where all proto-defined properties of `google.cloud.vmmigration.v1.Group` are properly mapped directly into the KRM spec and status fields.
