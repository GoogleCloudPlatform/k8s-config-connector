# ContentWarehouse Journal

### [2026-07-07] ContentWarehouseSchema Direct Controller Setup
- **Context**: Greenfield implementation of ContentWarehouseSchema direct controller under `pkg/controller/direct/contentwarehouse/contentwarehouseschema/`.
- **Problem**: 
  1. The Google Cloud Go client library for ContentWarehouse (`cloud.google.com/go/contentwarehouse`) does not exist as a separate module in `google-cloud-go`. Instead, the Go client resides under `google.golang.org/api/contentwarehouse/v1`, which is a REST-based client.
  2. The pre-compiled Go genproto package (`google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1`) lacks several newer fields defined in the latest `.proto` schema files (specifically `RetrievalImportance` and `SchemaSource` inside `PropertyDefinition`). This mismatch causes compilation failures when generated mappers reference non-existent fields.
- **Solution**:
  1. Utilized the REST-based `google.golang.org/api/contentwarehouse/v1` client by adding a client initializer using `ControllerConfig.RESTClientOptions()`.
  2. Added a Go package import path mapping exception to `GoPackageForProto` in `dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go` to cleanly direct the proto import to `google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1`.
  3. Hand-wrote a manual mapper under `pkg/controller/direct/contentwarehouse/mapper_manual.go` to stub `Policy` mappings and provide a manual override for `PropertyDefinition` mapping that strips out the missing `RetrievalImportance` and `SchemaSources` fields. This tells the code generator to skip generating conflicting mappers, resulting in perfect, clean compilation.
  4. Utilized `common.ProtoToAPI` and `common.APIToProto` helpers to seamlessly convert between the REST client's JSON representation and the proto-based mappers.
- **Impact**: Establishes a robust, compile-safe, and standard blueprint for any future ContentWarehouse resources, routing REST-based client communication smoothly while maintaining the standard proto-mapper paradigm.
