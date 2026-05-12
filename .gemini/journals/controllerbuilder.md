### [2026-05-12] controllerbuilder generate-mapper crashing on apiextensionsv1
- **Context**: Implementing direct types for `EventarcEnrollment` using `dev/tools/controllerbuilder/generate.sh`.
- **Problem**: `go run . generate-mapper` panics with `could not find import for "apiextensionsv1.JSON" (package "apiextensionsv1")` during AST inspection. It crashes universally across APIs that happen to have generated fields relying on `apiextensionsv1` but lacking explicit imports.
- **Solution**: Avoid running `generate-mapper` automatically if it crashes, or manually inject the `apiextensionsv1` import in the respective `types.generated.go` file (e.g., `discoveryengine/v1alpha1/types.generated.go` or `vertexai/v1alpha1/types.generated.go`) to allow the `go/ast` parser to succeed.
- **Impact**: The next agent should be prepared to skip `generate-mapper` and manually create mappers if `dev/tasks/generate-types-and-mappers` fails, or manually fix the missing imports in other generated files.

### [2026-05-12] controllerbuilder empty structs for types
- **Context**: Scaffolded `EventarcEnrollmentSpec` using `generate-types`.
- **Problem**: The generated struct `EventarcEnrollmentSpec` was empty except for the parent fields (Location, ProjectRef, ResourceID). The underlying proto `google.cloud.eventarc.v1.Enrollment` had many fields.
- **Solution**: Manually implement the fields in `enrollment_types.go` matching the proto fields and then manually map them in `mapper.generated.go` (if `generate-mapper` already added the empty stubs) or a new `mapper_enrollment.go`.
- **Impact**: The next agent should always review generated `_types.go` fields. If they are inexplicably missing, manually implement them by consulting the `.proto` file.
