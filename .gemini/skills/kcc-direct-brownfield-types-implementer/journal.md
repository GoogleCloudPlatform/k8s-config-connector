### [2026-05-12] Naming Conflict for `<Kind>` vs `<ProtoMessage>` generated files
- **Context**: Implementing `NotebookRuntime` backed by the `Runtime` proto message.
- **Problem**: `generate-types` creates `runtime_types.go` based on the proto message name `Runtime`. However, KCC conventions expect `notebookruntime_types.go`. If we rename it manually but keep the generate script, re-running `generate-types` generates `runtime_types.go` again, causing a Go AST compilation error due to duplicate types during `generate-crds`.
- **Solution**: In the `generate.sh` script, add a line `rm apis/<service>/<version>/runtime_types.go` after running `generate-types` to delete the duplicate file created by the proto message name.
- **Impact**: Prevents AST compilation errors (`interface conversion: types.Type is nil, not *types.Named`) when `dev/tasks/generate-crds` runs due to duplicate files with the same type definitions.
