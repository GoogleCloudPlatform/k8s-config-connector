# DCL (Declarative Resource Client Library) Integration

This package (`pkg/dcl`) implements Config Connector's integration with the [Declarative Resource Client Library (DCL)](https://github.com/GoogleCloudPlatform/declarative-resource-client-library).

> [!NOTE]
> Config Connector is actively prioritizing the [Direct reconciliation approach](../../docs/develop-resources/README.md). New resources and fields should be implemented as Direct controllers. For guidance on maintaining or updating DCL-backed resources, see the [DCL Field and Resource Playbook](../../docs/develop-resources/scenarios/dcl-field-and-resource.md).

## Package Overview

- **`metadata/`**: Contains the centralized registry (`metadata.go`) of DCL-backed services, API versions, and resource configurations (`Releasable`, `SupportsHierarchicalReferences`, `SupportsContainerAnnotations`).
- **`schema/`**: Manages DCL OpenAPI schema assets and runtime loading:
  - `main/assets_generate.go`: Uses `vfsgen` to embed DCL OpenAPI schemas into virtual filesystem Go code (`embed/dcl_assets_vfsdata.go`).
  - `embed/`: Contains the generated embedded virtual filesystem asset code.
  - `dclschemaloader/`: Provides the schema loader interface and implementation (`EmbedDCLSchemaLoader`) used during CRD generation (`scripts/generate-crds/main.go`) and runtime conversions.
- **`conversion/`**: Handles bidirectional conversion between Kubernetes Resource Model (KRM) objects and DCL unstructured objects.
- **`extension/`**: Extracts and interprets custom DCL OpenAPI extensions (e.g., `x-dcl-id`, `x-dcl-labels`, container annotations, and reference configurations).
- **`kcclite/`**: Utilities to resolve references and convert KRM objects to a state usable without an active Kubernetes API server.
- **`livestate/`**: Manages reading and comparing live GCP state through DCL client calls.
- **`clientconfig/`**: Sets up Google client configurations, HTTP clients, and timeouts for DCL calls.

## Key Workflows

### Patching In-Tree DCL Code
Config Connector vendors and patches DCL directly in-tree under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/`. After modifying schemas, types, or conversion logic in `third_party/`:
```bash
make generate-including-dcl
make ready-pr
```

### Upstream Sync & Dependency Upgrades
In rare scenarios where syncing with the canonical upstream repository:
```bash
make upgrade-dcl
make generate-including-dcl
make ready-pr
```

For more details, refer to the [DCL Playbook](../../docs/develop-resources/scenarios/dcl-field-and-resource.md).
