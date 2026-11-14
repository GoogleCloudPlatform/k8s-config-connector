# `controllerbuilder` Code Generation Backlog

These backlog items capture generator improvements identified while bulk-generating 73 greenfield resources across 35 GCP services ([PR #14066](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14066)).

---

## 1. Auto-detect Recursive Protobuf Messages in `generate-types` (`typegenerator.go`)

### Problem
`controller-gen` (`crd` generator) rejects recursive Go struct definitions (`unsupported recursive type`) unless the recursion cycle is broken using `apiextensionsv1.JSON`. When generating types from protobuf definitions that contain direct or indirect self-references, `generate-types` currently emits recursive Go structs in `types.generated.go`, which causes `dev/tasks/generate-crds` and `TestAllCRDsShouldConvertToYAML` (`pkg/crd/template`) to fail until the recursive structs are manually moved to a hand-written file and broken with `apiextensionsv1.JSON`.

### Examples Encountered in PR #14066
- `google.protobuf.Value` / `google.protobuf.ListValue` (`ListValue.values` -> `[]Value` -> `ListValue`) in `apis/datalineage/v1alpha1`, `apis/dialogflow/v1alpha1`, and `apis/grafeas/v1alpha1`
- `google.cloud.dataplex.v1.Schema.SchemaField` (`Schema_SchemaField.Fields` -> `[]Schema_SchemaField`) in `apis/dataplex/v1alpha1`
- `google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent` (`AdditionalCases`), `InlineSchema` (`Items`), and `Playbook.Step` (`Steps`) in `apis/dialogflow/v1alpha1`
- `google.cloud.retail.v2.Product` (`ProductObservedState.Variants` -> `[]ProductObservedState`) in `apis/retail/v1alpha1`

### Proposed Fix
In `dev/tools/controllerbuilder/pkg/codegen/typegenerator.go`:
- Track the active ancestor chain (call stack of `protoreflect.MessageDescriptor`s) while traversing message fields.
- When a field's message type is already in the current ancestor chain (or is a known recursive well-known type like `google.protobuf.Value` / `google.protobuf.Struct` / `google.protobuf.ListValue`), automatically emit `*apiextensionsv1.JSON` (or `[]apiextensionsv1.JSON` for repeated fields) and add `"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"` to the generated file imports.
- Update `mappergenerator.go` to handle or skip fields whose KRM type was rewritten to `apiextensionsv1.JSON` due to cycle breaking.

---

## 2. Generate or Provide Shared Mappers for Well-Known `google.type.*` and `google.protobuf.*` Messages (`mappergenerator.go`)

### Problem
When a service's protobuf messages reference common types from `google.type.*` or `google.protobuf.*` (such as `google.type.PostalAddress`, `google.type.Date`, `google.type.Decimal`, `google.type.Expr`, `google.protobuf.Int32Value`, `google.protobuf.Any`, `google.protobuf.Value`), `generate-types` generates local KRM structs for them in `apis/<service>/<version>/types.generated.go`, and `generate-mapper` emits calls to `<Type>_FromProto` / `<Type>_ToProto` (e.g., `PostalAddress_FromProto`, `Date_FromProto`, `Decimal_FromProto`, `Int32Value_FromProto`, `Expr_FromProto`).

However, `generate-mapper` does not generate the corresponding `<Type>_FromProto` / `<Type>_ToProto` definitions because those messages reside in `google.type` / `google.protobuf` rather than the target `--service` proto package, leaving `pkg/controller/direct/<service>/mapper.generated.go` with undefined symbol errors until hand-written helpers are added.

### Examples Encountered in PR #14066
- `pkg/controller/direct/cloudchannel`: `PostalAddress_FromProto`/`_ToProto`, `Decimal_FromProto`/`_ToProto`, `Date_FromProto`/`_ToProto`
- `pkg/controller/direct/retail`: `Int32Value_FromProto`/`_ToProto`
- `pkg/controller/direct/securitycenter`: `Expr_FromProto`/`_ToProto`
- `pkg/controller/direct/grafeas`: `Any_FromProto`/`_ToProto`
- `pkg/controller/direct/dialogflow`: `Value_FromProto`/`_ToProto`

### Proposed Fix
Either:
1. Teach `dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go` to resolve the Go protobuf package for `google.type.*` (`google.golang.org/genproto/googleapis/type/...`) and `google.protobuf.*` (`google.golang.org/protobuf/types/known/...`) and emit the mapper functions directly in `mapper.generated.go` (importing the appropriate proto package alias), **or**
2. Provide generic/shared conversion helpers in `pkg/controller/direct` and emit calls to `direct.<Type>_FromProto` / `direct.<Type>_ToProto`.

---

## 3. Fix `[]byte` Oneof Wrapper Dereferencing in `generate-mapper` (`mappergenerator.go`)

### Problem
When generating a oneof wrapper helper function `<Struct>_<OneofField>_ToProto` for a protobuf `bytes` field inside a `oneof`, the KRM field type is `[]byte` (a slice, not a pointer). However, `mappergenerator.go` unconditionally dereferences the parameter (`*in`) when constructing the protobuf oneof wrapper struct:
```go
func Example_RawContent_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Example_RawContent {
	if in == nil {
		return nil
	}
	return &pb.Example_RawContent{RawContent: *in} // invalid operation: cannot indirect in (variable of type []byte)
}
```

### Examples Encountered in PR #14066
- `pkg/controller/direct/dialogflow`: `ConversationDataset_RawContent_ToProto`, `Generator_RawContent_ToProto`, `ResponseMessage_MixedAudio_SegmentObservedState_Audio_ToProto`
- `pkg/controller/direct/discoveryengine`: `Document_Content_RawBytes_ToProto`

### Proposed Fix
In `dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go`:
- When generating `<Struct>_<OneofField>_ToProto`, inspect whether the input parameter type is a pointer (`*T`) vs. a slice (`[]byte`). Only emit `*in` when `in` is a pointer; emit `in` directly when `in` is `[]byte`.

---

## 4. Deduplicate Oneof Helper Functions Across API Versions in `generate-mapper` (`mappergenerator.go`)

### Problem
When an API group has multiple versions (e.g., both `v1alpha1` and `v1beta1`) that define the same KRM struct containing a protobuf `oneof` field, `generate-mapper` processes both versions and emits the oneof helper function `<Struct>_<Field>_ToProto` (which only depends on primitive/proto types in its signature, e.g., `(mapCtx *direct.MapContext, in *string) *pb.<Struct>_<Field>`) once per API version into the same `pkg/controller/direct/<service>/mapper.generated.go` file. This causes `redeclared in this block` compilation errors.

### Examples Encountered in PR #14066
- `pkg/controller/direct/metastore`: `MetadataExportObservedState_DestinationGcsUri_ToProto`, `ScalingConfig_InstanceSize_ToProto`, `ScalingConfig_ScalingFactor_ToProto` (emitted twice for `v1alpha1` and `v1beta1`)
- `pkg/controller/direct/dataform`: `InvocationConfig_QueryPriority_ToProto`
- `pkg/controller/direct/netapp`: `TieringPolicy_TierAction_ToProto`
- `pkg/controller/direct/vectorsearch`: `DedicatedInfrastructure_Mode_ToProto`

### Proposed Fix
In `dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go`:
- Track emitted helper function names in a package-wide set across all processed KRM versions so that version-independent oneof wrapper helpers (`<Struct>_<Field>_ToProto`) are emitted at most once per `mapper.generated.go` file.
