# Terraform-Based Field and Resource Additions Playbook

This playbook provides comprehensive guidance for adding new fields and introducing brownfield resources managed by the Terraform-based controller in Config Connector.

---

## 1. Controller Archetype Overview

Approximately 70% of legacy Config Connector resources are backed by the vendored Terraform Provider Google Beta (TPG Beta) library and the dynamic controller.

A resource is Terraform-based if:
- `pkg/controller/resourceconfig/static_config.go` specifies `DefaultController: Terraform`, or
- The CRD schema under `config/crds/resources/` contains the label `cnrm.cloud.google.com/tf2crd: "true"`.

---

## 2. TPG Beta Git Subtree Workflow

Config Connector maintains a local copy of TPG Beta under:
```
third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/
```

### 2.1 Modifying Schema Definitions
Locate `resource_<service>_<resource>.go` in the service directory. Schema definitions define the attributes supported by Terraform:

```go
"my_field": {
    Type:        schema.TypeString,
    Optional:    true,
    Description: `Description of the field.`,
},
```

Schema attributes can specify:
- `Optional: true`: The user can configure this field in KRM `spec`.
- `Required: true`: The field must be specified by the user.
- `Computed: true`: The field is computed by the server or defaults. If only `Computed: true` (not `Optional`), it maps to KRM `status.observedState`.
- `ForceNew: true`: Changing this field requires deleting and recreating the GCP resource.

### 2.2 Implementing Expanders
Expanders convert Terraform state/map representation into the GCP SDK client struct:
```go
func expand<Resource><Field>(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
    // Convert v into GCP client struct
}
```

### 2.3 Implementing Flatteners
Flatteners convert the GCP SDK response into the Terraform map representation:
```go
func flatten<Resource><Field>(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
    // Convert GCP client struct into map or primitive
}
```

### 2.4 The Empty Block Flattener Invariant (`ForceNew: true`)

> **CRITICAL PITFALL**:
> For optional nested blocks marked `ForceNew: true`, if all attributes within the block are empty in GCP's API response, the flattener **must return `nil` or an empty slice**, **never** a map with zero-values!
> 
> Returning a map with zero-values causes Terraform to compute a permanent non-empty diff against the empty user configuration. Because `ForceNew` is `true`, this permanent diff causes Config Connector to endlessly recreate the resource on every reconciliation loop!

Example safe flattener:
```go
func flattenNestedBlock(v interface{}) interface{} {
    if v == nil {
        return nil
    }
    obj := v.(*clientsdk.NestedBlock)
    if obj.FieldA == "" && obj.FieldB == 0 {
        return nil // Return nil when all inner fields are zero-value
    }
    return []interface{}{
        map[string]interface{}{
            "field_a": obj.FieldA,
            "field_b": obj.FieldB,
        },
    }
}
```

### 2.5 Boolean Handling
In Terraform-based resources, boolean fields should follow pointer semantics to distinguish between an explicit `false` and an unset/omitted value. Refer to `docs/design-decisions/boolean-handling-in-tf-based-resources.md` for full details.

---

## 3. ServiceMapping Configuration (`config/servicemappings/<service>.yaml`)

The `ServiceMapping` resource instructs Config Connector how to map Terraform schemas into Kubernetes Custom Resource Definitions (CRDs).

### 3.1 Field Whitelisting
By default, new fields in TPG Beta may be ignored if listed under `ignoredFields`. To expose a new field in KCC:
1. Locate the resource in `config/servicemappings/<service>.yaml`.
2. If the field is in `ignoredFields`, remove it.

### 3.2 Declaring Structured References (`resourceReferences`)
If a field references another GCP resource, map it to a KCC reference field using `resourceReferences`:
```yaml
resourceReferences:
  - tfField: network_config.subnetwork
    key: subnetworkRef
    gvk:
      kind: ComputeSubnetwork
      version: v1beta1
      group: compute.cnrm.cloud.google.com
    targetField: self_link
```

### 3.3 Server-Computed Attributes (`observedFields`)
Server-generated or output-only fields that should appear in `status.observedState` are declared under `observedFields`:
```yaml
observedFields:
  - selfLink
  - createTime
```

---

## 4. Adding Brownfield Resources via ServiceMapping

To expose a previously unsupported TPG Beta resource in Config Connector:
1. Add a new resource definition under `config/servicemappings/<service>.yaml`:
```yaml
- name: google_compute_address
  kind: ComputeAddress
  version: v1beta1
  idTemplate: "projects/{{project}}/regions/{{region}}/addresses/{{name}}"
```
2. Leverage `scripts/resource-autogen/` or Magic Modules definitions where applicable.

---

## 5. Manifest Generation and Testing

After updating TPG Beta and ServiceMapping:

1. **Verify TPG compilation**:
   ```bash
   make vet
   ```
2. **Generate CRDs and documentation**:
   ```bash
   make manifests
   make resource-docs
   ```
3. **Run integration test fixtures**:
   ```bash
   go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
   ```
4. **Pre-PR Verification**:
   ```bash
   make ready-pr
   ```

---

## 6. Public Issue/PR Hygiene (SOUL.md Invariant 7)

Never mention Buganizer IDs (`b/<id>`), internal Buganizer URLs, or customer names in issue descriptions, PR titles, PR bodies, commit messages, code comments, test fixtures, or release notes.
