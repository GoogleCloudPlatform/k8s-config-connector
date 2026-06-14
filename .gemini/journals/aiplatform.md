# AIPlatform Domain Knowledge

## KCC Proto Generation Quirks

- `google.protobuf.Value` and `google.protobuf.ListValue` are somewhat problematic in `dev/tools/controllerbuilder/pkg/codegen/typegenerator.go` and `common.go`. When modifying KRM types that include `Value` fields, ensure that `ListValue` is not accidentally commented out in manually maintained files like `model_types.go`, otherwise it breaks `types.generated.go` and causes `controller-gen` to fail with `invalid type`.
- `types.generated.go` will be entirely re-written by `generate.sh`. Ensure that any `generate.sh` script invokes `generate-types` for ALL existing resources in the package (e.g. `--resource AIPlatformModel:Model`) to prevent deleting other resources from the shared `types.generated.go`.
- `apiextensionsv1` might not be imported correctly by `generate-types` if it gets generated for `StructValue`. Check `dev/tools/controllerbuilder/pkg/codegen/typegenerator.go` (and we successfully added `addImport("apiextensionsv1", ...)` inside `WriteVisitedMessages` and `WriteOutputMessages` to fix this project-wide).
