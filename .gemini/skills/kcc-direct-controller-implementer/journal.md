- ApiHubDeployment: Dropped 'Annotations' from spec since proto uses map of string to AttributeValues (wrapper message). Used IdentityV2 standard.

### [2026-05-19] Package Isolation with generate-mapper
- **Context**: Implementing isolated package for `BeyondCorpClientConnectorService` direct controller.
- **Problem**: The `dev/tools/controllerbuilder/main.go generate-mapper` command appends the API group name to the output directory and hardcodes the Go package name to the API group, violating package isolation if left as is.
- **Solution**: Output the mapper to the isolated directory, then move the nested file and run `sed` to replace the package name on the generated mapper file. This is best scripted in the resource's `generate.sh`.
- **Impact**: Ensures that generated mappers for different resources within the same API group do not collide, strictly adhering to the package isolation rules for direct controllers.

### [2026-05-19] FieldMask Generation in Direct Controllers
- **Context**: Implementing `Update` for `BeyondCorpClientConnectorService` which required an `UpdateMask`.
- **Problem**: Need an easy way to compute the updated paths for the field mask based on the desired and actual protobuf representations.
- **Solution**: Use `paths, err := common.CompareProtoMessage(updateReq, a.actual, common.BasicDiff)`. Then pass `sets.List(paths)` (from `k8s.io/apimachinery/pkg/util/sets`) to `fieldmaskpb.FieldMask.Paths`.
- **Impact**: Provides a standardized, straightforward pattern for direct KCC controllers to compute update masks accurately.

### [2026-05-19] Enum_ToProto and Enum_FromProto with Generics
- **Context**: Implementing `AutoMLDataset` direct controller mappers.
- **Problem**: Enums in `types.generated.go` are generated as `*string` (or a type alias for string), but proto expects a specific enum type (e.g., `pb.ClassificationType`). Using `direct.Enum_ToProto(mapCtx, in.ClassificationType)` results in compilation error `cannot infer U`.
- **Solution**: Use explicit type instantiation: `direct.Enum_ToProto[pb.ClassificationType](mapCtx, in.ClassificationType)`.
- **Impact**: Ensures correct type inference and compilation for enum mapping.
