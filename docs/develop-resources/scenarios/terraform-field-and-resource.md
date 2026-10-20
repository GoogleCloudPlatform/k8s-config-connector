# Terraform-Based Field and Resource Addition Playbook

This playbook provides end-to-end guidance for contributing fields or resources managed by Config Connector's legacy Terraform-based reconciler.

Config Connector historically manages approximately 70% of resources via a dynamic controller backed by a locally vendored copy of the [Terraform Provider Google (TPG) Beta](https://github.com/hashicorp/terraform-provider-google-beta). While new resources and active migrations prioritize the [Direct Reconciliation approach](./new-resource.md), maintaining existing brownfield resources and backporting newly added GCP features requires patching the vendored Terraform provider and updating ServiceMapping configurations.

---

## 1. Controller Archetype Triage

Before making code changes, verify whether your target resource is managed by the Terraform controller:

1. **Inspect Static Config**:
   Check `pkg/controller/resourceconfig/static_config.go`. A Terraform-managed resource has `DefaultController` set to `k8s.ReconcilerTypeTerraform` (or has `k8s.ReconcilerTypeTerraform` listed in `SupportedControllers`).
2. **Inspect CRD Schema**:
   Check the resource CRD under `config/crds/resources/`. A Terraform-managed CRD contains the label:
   ```yaml
   metadata:
     labels:
       cnrm.cloud.google.com/tf2crd: "true"
   ```
3. **Inspect ServiceMapping**:
   Check `config/servicemappings/<service>.yaml`. A corresponding `ResourceConfig` entry will exist for the resource.

> **Note**: If `DefaultController` is `k8s.ReconcilerTypeDirect`, do **not** use this playbook; follow the [Direct Resource Playbook](./new-field.md) instead. If the resource is dual-controlled (supports both Direct and Terraform), updates must be coordinated across both engines.

---

## 2. TPG Beta Git Subtree Workflow

Config Connector maintains a local copy of Terraform Provider Google Beta as a vendored directory tree at:
```
third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/
```

### 2.1 Modifying Schema Definitions

The schema definition for a resource is located in `resource_<service>_<resource>.go` (or related helper files such as `node_config.go`).

When adding or modifying a field in `Schema: map[string]*schema.Schema`:

```go
"my_new_field": {
    Type:        schema.TypeString,
    Description: `Description of the field matching GCP API documentation.`,
    Optional:    true,
    ForceNew:    true, // Set to true if the GCP API requires recreation when changed
},
```

#### Schema Attributes & Invariants:
* **`Type`**: Use standard Terraform schema types (`schema.TypeString`, `schema.TypeBool`, `schema.TypeInt`, `schema.TypeFloat`, `schema.TypeList`, `schema.TypeSet`, `schema.TypeMap`).
* **`Optional` vs `Required`**:
  * Set `Optional: true` for non-mandatory fields.
  * Set `Required: true` only if the GCP API strictly requires the field upon resource creation.
* **`Computed`**:
  * Set `Computed: true` if the field is populated or defaulted by the server.
  * For output-only fields, set `Computed: true` and omit `Optional`/`Required`.
* **`ForceNew` (Immutability)**:
  * Setting `ForceNew: true` informs Terraform that modifying this field requires destroying and recreating the resource.
  * In Config Connector, `ForceNew: true` causes the CRD generator to prefix the field description with `Immutable. ` and marks the field as immutable in KRM.
  * **Rule**: If the field can be modified in-place via the GCP API, do **not** set `ForceNew: true`. Instead, ensure the resource's `Update` function handles the modification via `d.HasChange(...)`.

---

### 2.2 Implementing Expanders

An expander converts Terraform schema state (`*schema.ResourceData` or `map[string]interface{}`) into GCP client SDK structs for API requests.

```go
func expandMyResourceMyNewField(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) (interface{}, error) {
    if v == nil {
        return nil, nil
    }
    s, ok := v.(string)
    if !ok || s == "" {
        return nil, nil
    }
    return s, nil
}
```

For complex structs or nested objects:
```go
func expandMyResourceNestedConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) (*service.NestedConfig, error) {
    l, ok := v.([]interface{})
    if !ok || len(l) == 0 || l[0] == nil {
        return nil, nil
    }
    raw := l[0].(map[string]interface{})
    req := &service.NestedConfig{
        Attribute: raw["attribute"].(string),
    }
    return req, nil
}
```

---

### 2.3 Implementing Flatteners

A flattener converts GCP client SDK response structs back into Terraform schema data structures (`[]interface{}` or primitive types) for state storage.

```go
func flattenMyResourceNestedConfig(c *service.NestedConfig) []interface{} {
    if c == nil {
        return nil
    }
    transformed := make(map[string]interface{})
    transformed["attribute"] = c.Attribute
    return []interface{}{transformed}
}
```

---

### 2.4 The Empty Block Flattener Invariant (`ForceNew: true`)

> [!CAUTION]
> **CRITICAL PITFALL: Empty Block Flatteners on `ForceNew` Fields**
>
> For optional nested blocks marked `ForceNew: true`, if all attributes within the block are empty in GCP's API response, the flattener **MUST return `nil` or an empty slice (`[]interface{}{}`)**, never a map containing zero-values.
>
> **Why this matters**:
> If a user omits an optional block from their KRM YAML, their desired state contains 0 blocks. If the flattener returns a 1-element slice containing a map with zero-values (e.g. `[{"attribute": ""}]`), Terraform computes a state diff between 0 blocks (desired) and 1 block (live state). Because the block is marked `ForceNew: true`, Terraform generates a replacement plan, triggering **endless resource recreation cycles on every reconciliation loop** in Config Connector!

#### Incorrect Pattern:
```go
// INCORRECT: Always returns a slice with a zero-value map
func flattenSubConfig(c *service.SubConfig) []interface{} {
    if c == nil {
        return nil
    }
    // BUG: If c is non-nil but has empty fields, this returns a 1-item slice with zero-values!
    return []interface{}{
        map[string]interface{}{
            "enabled": c.Enabled,
            "mode":    c.Mode,
        },
    }
}
```

#### Correct Pattern:
```go
// CORRECT: Returns nil if all sub-attributes are empty or zero-value
func flattenSubConfig(c *service.SubConfig) []interface{} {
    if c == nil {
        return nil
    }
    if !c.Enabled && c.Mode == "" {
        return nil // Or return []interface{}{}
    }
    return []interface{}{
        map[string]interface{}{
            "enabled": c.Enabled,
            "mode":    c.Mode,
        },
    }
}
```

---

### 2.5 Boolean and Zero-Value Handling

Handling optional boolean fields (and zero-values like `false`, `0`, `""`) in Terraform Provider SDK v2 requires special care due to zero-value collapse.

For full architectural details, see [docs/design-decisions/boolean-handling-in-tf-based-resources.md](../../design-decisions/boolean-handling-in-tf-based-resources.md).

#### The "Non-Zero Default" Trap in Terraform Schemas:
* **Never add `Default: true` to optional boolean fields in Terraform provider schemas.**
* In Terraform Plugin SDK v2, `false` is collapsed into the zero-value/unset category. If `Default: true` is configured:
  1. A user specifies `spec.myBoolField: false`.
  2. TF SDK v2 evaluates `false` as unset/omitted.
  3. TF automatically replaces the value with `Default: true`.
  4. The user's explicit intent is permanently overwritten, making it impossible to opt out or set `false` on GCP!
  5. Furthermore, existing live resources with `false` will be unexpectedly mutated to `true` on subsequent reconciliations.

#### `d.Get(...)` vs `d.GetOk(...)` for Expanders:
* Use `d.Get(...)` for boolean fields instead of `d.GetOk(...)`.
* `d.GetOk("bool_field")` returns `ok = false` when the value is `false` (since `isZero(false) == true`), silently dropping explicit `false` inputs and preventing users from sending `false` to GCP APIs.
* `d.Get("bool_field")` returns the actual boolean value, correctly preserving explicit `false` declarations.

---

### 2.6 Compiling the Vendored Provider

To verify that modifications in `third_party/` compile correctly:

1. Temporarily uncomment the vendored module in `go.work` if needed:
   ```go
   // go.work
   ./third_party/github.com/hashicorp/terraform-provider-google-beta
   ```
2. Run `go vet` scoped to the modified service package:
   ```bash
   go vet ./third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/...
   ```
3. Revert any local `go.work` changes:
   ```bash
   git checkout go.work
   ```

---

## 3. ServiceMapping Configuration (`config/servicemappings/<service>.yaml`)

The `ServiceMapping` resource defines how Terraform provider schemas map to Kubernetes Custom Resource Definitions (CRDs).

### 3.1 Structure of a ServiceMapping File

```yaml
apiVersion: core.cnrm.cloud.google.com/v1alpha1
kind: ServiceMapping
metadata:
  name: compute.cnrm.cloud.google.com
  namespace: cnrm-system
spec:
  name: Compute
  version: v1beta1
  serviceHostName: "compute.googleapis.com"
  resources:
    - name: google_compute_subnetwork
      kind: ComputeSubnetwork
      # Resource configuration...
```

---

### 3.2 Field Whitelisting

When adding new fields to an existing resource:
1. Check if the field is present in `ignoredFields` or `ignoredOutputOnlySpecFields`.
2. **Remove the field from `ignoredFields`** to expose it in the generated CRD spec.
3. If the field is a subfield under a writable parent block but is output-only according to the GCP API, add it to `ignoredOutputOnlySpecFields`:
   ```yaml
   ignoredOutputOnlySpecFields:
     - parent_block.generated_id
   ```

---

### 3.3 Structured Resource References

To allow KRM objects to reference other Config Connector resources using Kubernetes references (e.g. `subnetworkRef` instead of raw URLs or self links), declare them under `resourceReferences`:

```yaml
resourceReferences:
  - tfField: network_config.subnetwork
    key: subnetworkRef
    description: |-
      The ComputeSubnetwork that this resource is attached to.
    gvk:
      kind: ComputeSubnetwork
      version: v1beta1
      group: compute.cnrm.cloud.google.com
    targetField: self_link
```

#### Fields:
* **`tfField`**: The path to the field in the Terraform schema (using dot notation for nested fields).
* **`key`**: The name of the reference field in the KRM spec (must end with the `Ref` suffix, e.g. `subnetworkRef`, `kmsKeyRef`).
* **`gvk`**: The Group, Version, and Kind of the referenced resource.
* **`targetField`**: The field of the referenced resource to supply to Terraform (typically `self_link` or `name`).

---

### 3.4 Hierarchical References

If the resource resides within a GCP resource hierarchy (project, folder, organization, or billing account), declare both `hierarchicalReferences` and the corresponding `resourceReferences`:

```yaml
hierarchicalReferences:
  - type: project
    key: projectRef

resourceReferences:
  - tfField: project
    key: projectRef
    description: |-
      The project that this resource belongs to.
    gvk:
      kind: Project
      version: v1beta1
      group: resourcemanager.cnrm.cloud.google.com
```

---

### 3.5 Output-Only and Observed Fields

If a computed or server-generated field should be exposed under `status.observedState` (or status conditions) for consumers to inspect:

```yaml
observedFields:
  - master_auth.client_certificate
  - endpoint
```

---

### 3.6 Directives and Mutable-But-Unreadable Fields

* **`directives`**: Terraform-specific behavioral flags configured via resource annotations (e.g. `cnrm.cloud.google.com/force-destroy`):
  ```yaml
  directives:
    - force_destroy
    - remove_default_node_pool
  ```
* **`mutableButUnreadableFields`**: Fields that are mutable via API requests but not returned in read/GET responses (such as sensitive credentials or passwords):
  ```yaml
  mutableButUnreadableFields:
    - password
  ```

---

## 4. Adding Brownfield Resources via ServiceMapping

To add a new resource that already exists in Terraform Provider Google Beta:

### 4.1 Define the ResourceConfig

Add an entry under `spec.resources` in `config/servicemappings/<service>.yaml`:

```yaml
resources:
  - name: google_spanner_instance
    kind: SpannerInstance
    metadataMapping:
      name: name
      labels: labels
    resourceID:
      targetField: name
    idTemplate: "projects/{{project}}/instances/{{name}}"
    idTemplateCanBeUsedToMatchResourceName: true
```

### 4.2 Using `scripts/resource-autogen/` & Magic Modules

Config Connector maintains autogenerated service mapping templates under `scripts/resource-autogen/generated/servicemappings/`:

1. Check `scripts/resource-autogen/generated/servicemappings/<service>.yaml` for pre-generated resource configurations.
2. If available, copy the resource configuration into `config/servicemappings/<service>.yaml`.
3. Remove or unset `autoGenerated: true` (which should only remain true for automated generators).
4. Review and enrich `resourceReferences`, `metadataMapping`, and `idTemplate`.
5. If the resource is generated using [Magic Modules](https://github.com/GoogleCloudPlatform/magic-modules), `scripts/resource-autogen/generate-resources.sh` can be executed to refresh templates from upstream compiler specifications.

### 4.3 Enable GCP Service in Test Config

If introducing a new GCP service, ensure the API is registered in `SUPPORTED_SERVICES` in `scripts/shared-vars-public.sh`.

---

## 5. Manifest Generation and Tooling

Once TPG Beta and ServiceMappings are configured, regenerate CRDs and documentation:

### 5.1 CRD Manifest Generation
```bash
make manifests
```
Verify the generated CRD under `config/crds/resources/` (e.g. check `git diff config/crds/resources/`).

### 5.2 Resource Documentation Generation
```bash
make resource-docs
```
Verify generated documentation under `scripts/generate-google3-docs/resource-reference/generated/resource-docs/`.

### 5.3 If the Resource Uses Go Types in `apis/`
If the resource has explicit Go types defined under `apis/<service>/<version>/<kind>_types.go`:
1. Add new fields to the `<Kind>Spec` Go struct with appropriate JSON tags and `+optional` markers.
2. Run the types and mappers generator:
   ```bash
   dev/tasks/generate-types-and-mappers
   ```
3. Run `make manifests` and `make resource-docs`.

---

## 6. Testing & Dynamic Controller Verification

### 6.1 Creating Test Fixtures

Add test fixtures under `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testcase>/`:

* **`create.yaml`**: Minimal configuration for resource creation. Use `${uniqueId}` and `${projectId}` template variables:
  ```yaml
  apiVersion: spanner.cnrm.cloud.google.com/v1beta1
  kind: SpannerInstance
  metadata:
    name: spannerinstance-${uniqueId}
  spec:
    config: regional-us-west1
    displayName: Spanner Instance Sample
    numNodes: 1
  ```
* **`update.yaml`**: Modifies mutable fields to verify in-place updates succeed without recreating the resource.
* **`dependencies.yaml`**: Declares prerequisite resources in topological order (e.g., a `ComputeNetwork` needed by a subnetwork).

### 6.2 Test Context Overrides

If a resource does not support updates or acquisition, configure `pkg/test/resourcefixture/contexts/<service>_context.go`:

```go
func init() {
    resourceContextMap["spannerinstance"] = ResourceContext{
        ResourceKind: "SpannerInstance",
        SkipUpdate:   false,
        SkipAcquire:  false,
    }
}
```

### 6.3 Executing Tests

Run integration fixture tests for the service:
```bash
go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
```

Run the dynamic controller end-to-end CRUD suite directly:
```bash
go test -v -tags=integration ./pkg/controller/dynamic/ -test.run TestCreateNoChangeUpdateDelete -run-tests <test-target> -timeout 900s
```

Capture HTTP logs for debugging if needed:
```bash
ARTIFACTS=/tmp/kcc-logs go test -v -tags=integration ./pkg/controller/dynamic/ -test.run TestCreateNoChangeUpdateDelete -run-tests <test-target> -timeout 900s
```

### 6.4 Resource Skeleton Tests (Export & Asset Inventory)

* **Asset Skeleton**: If supported by Cloud Asset Inventory, add an entry to `pkg/resourceskeleton/testdata/asset-skeleton.yaml`.
* **URI Skeleton**: Verify `config-connector export` support by adding an entry to `pkg/resourceskeleton/testdata/uri-skeleton.yaml`.

---

## 7. Pre-PR Checklist & Verification

Before submitting your pull request, run the repository validation suite:

```bash
make ready-pr
```

Ensure all generated files, CRDs, and docs are committed cleanly.
