# 3. Add the KRM and API mapper

## 3.1 Generate the API and proto mapper (repeat-run safe) 

Run the following command

```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/controllerbuilder

go run . generate-mapper \
   --service google.storage.v1  \
   --api-version "storage.cnrm.cloud.google.com/v1beta1"
```

**Note**: We suggest using the same proto for your mock GCP and for you type-generation tool to generate the Config Connector API and mapper to avoid mismatch in schema definitions. There are some exceptions where you need to [replace the proto go package](../../../dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go).

### Simple path

If no comments marked `MISSING` and all the mapper functions look good, you are done. You can move to “Add the controller” (step 4).

### Custom path

> **Warning:** `mapper.generated.go` is overwritten by code generators; custom mapping logic must be placed in `<resource>_mappings.go`. Never manually edit generated files.

If you need custom mapping logic (for example, handling reference types, complex field transforms, or fields marked `MISSING` by the generator), implement the custom methods in `<resource>_mappings.go` under `pkg/controller/direct/<service>/`.

The code generator detects existing mapping methods and automatically skips generating duplicates in `mapper.generated.go` when a custom implementation exists.

## 3.2 Add manual mapper code

Create a `<resource>_mappings.go` under `pkg/controller/direct/<service>/` when the auto-generated code cannot fulfill the mappings.

### Implementing custom mapping methods

* Follow the naming convention as the auto-generated code: implement `<Type>_ToProto(mapCtx *direct.MapContext, in *krm.<Type>) *pb.<Type>` and `<Type>_FromProto(mapCtx *direct.MapContext, in *pb.<Type>) *krm.<Type>`.
* Place all custom methods in `<resource>_mappings.go`. Once implemented, rerun `generate-mapper`, and the generator will omit the auto-generated versions from `mapper.generated.go`.

### Zero Defaulting Invariant in Conversion Mappers

Conversion mappers must be pure, lossless 1:1 translations between KRM and protobuf messages:
* **Do NOT synthesize or assign default values in mappers:** If a field is omitted or `nil` in KRM, `ToProto` must leave the proto field unset/nil. Never fabricate or assign default values in mappers.
* **Do NOT assume defaults in `FromProto`:** When converting from proto back to KRM, only map values that are actually present in the proto message.
* Defaulting is the responsibility of Kubernetes admission defaulting / CRD schemas or GCP server-side behavior. Injecting default values into conversion mappers violates round-trip fuzzer invariants and creates false diffs during reconciliation loops.

## 3.3 (Optional) Add fuzzer test

Note: For Beta resources, a fuzzer test is required for both the `Spec` and `Status` / `ObservedState`. To generate a fuzzer test, modify the following example and run it:

```bash
export REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/controllerbuilder

go run main.go prompt \
  <<EOF > $REPO_ROOT/pkg/controller/direct/bigquerydatatransfer/transferconfig_fuzzer.go
// +tool:fuzz-gen
// proto.message: google.cloud.bigquery.datatransfer.v1.TransferConfig
// api.group: bigquerydatatransfer.cnrm.cloud.google.com
EOF
```


## Exit Criteria

If you are developing a Beta or more stable resource version, you should meet the following requirements: 

* No `MISSING` comments left in the code
* No `/*NOTYET*/` comments left in the code.
* For Beta resources, a fuzzer test exists for both Spec and Status of the resource. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/f313b00c52f09c4a52a2eb5fe2c15fa4b30a05fd/pkg/controller/direct/discoveryengine/fuzzers.go#L26-L47)
* Each mapper method shall reflect in the `_http.log` as the value from `create.yaml` and `update.yaml` recorded in the `_http.log` POST and PUT/PATCH method.
