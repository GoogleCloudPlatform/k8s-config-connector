### [2026-06-03] Handling Greenfield Protos Added After Pinned Commit
- **Context**: Implementing `CESApp` types (Google Cloud CES v1beta) under `ces.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The `google/cloud/ces` proto files did not exist in the pinned googleapis SHA `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb`. Therefore, compiling `googleapis.pb` with the standard script left out the CES package, causing `generate-types` to fail.
- **Solution**: Modified `dev/tools/controllerbuilder/generate-proto.sh` to checkout `google/cloud/ces/v1beta` from `origin/master` after resetting to the pinned SHA, and added the path explicitly to the protoc compile command. This allowed us to successfully compile the `googleapis.pb` descriptor containing the ces service definitions.
- **Impact**: Any subsequent developer or automated pipeline running `./apis/ces/v1alpha1/generate.sh` can successfully regenerate types without missing proto errors.

### [2026-11-11] Prepopulating CESApp with GA Proto and Recursive Schema Handling
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CESApp` using `google.cloud.ces.v1` and `--prepopulate-resource`.
- **Problem**:
  1. `google.cloud.ces.v1.App` depends on `google.protobuf.Value`, requiring external proto mapping. In generated mappers, `direct.Struct_ToProto` expects `*apiextensionsv1.JSON` while `Value.StructValue` is non-pointer `apiextensionsv1.JSON`.
  2. `App.VariableDeclaration.Schema` is a recursive JSON schema message (`SchemaWithRequired` / `SchemaObservedState`) that causes `controller-gen` to emit empty `{}` OpenAPI schemas, failing `envtest` CRD validation in `TestPreview`.
- **Solution**:
  1. Added `,google.protobuf` to `--service` in `apis/ces/generate.sh` and implemented a minimal handwritten `Value_ToProto` helper in `pkg/controller/direct/ces/mapper.go`.
  2. Defined `App_VariableDeclaration` and `App_VariableDeclarationObservedState` in `cesapp_types.go` with `+kubebuilder:validation:Schemaless`, `+kubebuilder:validation:XPreserveUnknownFields`, and `+kubebuilder:validation:Type=object` on the `Schema` field to terminate the recursion cycle safely in OpenAPI v3 schema.
- **Impact**: Clean CRD generation, passing envtest, and full idempotency with `apis/ces/generate.sh`.
