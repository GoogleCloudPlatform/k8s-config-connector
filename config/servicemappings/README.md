# ServiceMappings Configuration

This directory contains `ServiceMapping` manifests for Google Cloud Platform (GCP) services managed by Config Connector's Terraform-based reconciler.

A `ServiceMapping` defines the translation layer between Config Connector Kubernetes Resource Model (KRM) Custom Resource Definitions (CRDs) and the underlying resources from the locally vendored [Terraform Provider Google Beta (TPG Beta)](../../third_party/github.com/hashicorp/terraform-provider-google-beta).

The metadata defined in these files is consumed by:
1. **CRD Generation** (`pkg/crd/crdgeneration`): Automatically generates OpenAPI v3 schema definitions for KCC CRDs from Terraform provider schemas.
2. **Dynamic Controller** (`pkg/controller/dynamic`): Guides runtime reconciliation, field mapping, reference resolution, and state synchronization between Kubernetes and GCP APIs.

The Go schema definitions for ServiceMappings are maintained in [`pkg/apis/core/v1alpha1/servicemapping_types.go`](../../pkg/apis/core/v1alpha1/servicemapping_types.go).

---

## ServiceMapping File Structure

Each GCP service has an associated file named `<service>.yaml` (in lowercase, e.g. `spanner.yaml`, `compute.yaml`, `pubsub.yaml`).

```yaml
apiVersion: core.cnrm.cloud.google.com/v1alpha1
kind: ServiceMapping
metadata:
  name: spanner.cnrm.cloud.google.com
  namespace: cnrm-system
spec:
  name: Spanner
  version: v1beta1
  serviceHostName: "spanner.googleapis.com"
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

### Top-Level Fields

| Field | Type | Description |
| :--- | :--- | :--- |
| `metadata.name` | `string` | The API group domain for the service (e.g. `spanner.cnrm.cloud.google.com`). |
| `spec.name` | `string` | The PascalCase name of the service (e.g. `Spanner`, `Compute`). |
| `spec.version` | `string` | Default API version for resource CRDs generated under this service (e.g. `v1beta1`). |
| `spec.serviceHostName` | `string` | The GCP service endpoint hostname (e.g. `spanner.googleapis.com`). |
| `spec.resources` | `[]ResourceConfig` | List of resource mapping definitions. |

---

## `ResourceConfig` Specification

Each entry under `spec.resources` configures a single Kubernetes Kind backed by a Terraform resource:

### Core Identity & Metadata

* **`name`** (`string`): The underlying Terraform resource name (e.g. `google_compute_subnetwork`).
* **`kind`** (`string`): The Kubernetes Kind name in PascalCase (e.g. `ComputeSubnetwork`).
* **`version`** (`*string`, optional): Overrides the default API version for this specific resource.
* **`metadataMapping`**:
  * **`name`**: Terraform field mapped from KRM `metadata.name` (typically `name`).
  * **`labels`**: Terraform field mapped from KRM `metadata.labels` (typically `labels`).
* **`resourceID`**:
  * **`targetField`**: Terraform field that receives either the user-specified resource ID (`metadata.name` / `spec.resourceID`) or server-generated ID.
* **`idTemplate`** (`string`): Format string used to construct resource IDs for Terraform importers and live state lookups (e.g. `projects/{{project}}/regions/{{region}}/subnetworks/{{name}}`).
* **`idTemplateCanBeUsedToMatchResourceName`** (`bool`): Set to `true` if the template matches the One Platform / GCP resource URI format for use in `config-connector export`.

### Field Filtering & Control

* **`ignoredFields`** (`[]string`):
  List of Terraform fields excluded from the generated CRD spec. To expose a newly supported or backported field in KCC, **remove it from this list**.
* **`ignoredOutputOnlySpecFields`** (`*[]string`):
  List of output-only subfields nested within writable parent objects. Prevents server-computed attributes from appearing in `spec`.
* **`observedFields`** (`*[]string`):
  List of server-computed attributes exposed under `status.observedState` (e.g. `master_auth.client_certificate`).
* **`mutableButUnreadableFields`** (`[]string`):
  Fields that are mutable in update requests but not returned by GCP read APIs (such as passwords or credentials).
* **`directives`** (`[]string`):
  Terraform provider behavioral flags configured on the resource via annotations (e.g. `cnrm.cloud.google.com/force-destroy`).

### References

* **`resourceReferences`** (`[]ReferenceConfig`):
  Maps implicit string/URL fields in Terraform to explicit Kubernetes object references:
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
* **`hierarchicalReferences`** (`[]HierarchicalReference`):
  Configures parent scoping within the GCP resource hierarchy (`project`, `folder`, `organization`, `billingAccount`):
  ```yaml
  hierarchicalReferences:
    - type: project
      key: projectRef
  ```

### IAM Configuration

* **`iamConfig`**:
  Configures IAM policy and member bindings associated with the resource:
  ```yaml
  iamConfig:
    policyName: google_spanner_instance_iam_policy
    policyMemberName: google_spanner_instance_iam_member
    supportsConditions: true
    referenceField:
      name: instance
      type: name
  ```

---

## Common Workflows

### 1. Whitelisting a Field
If a field exists in the vendored Terraform provider schema but is absent from the CRD:
1. Locate the resource in `config/servicemappings/<service>.yaml`.
2. Remove the field from `ignoredFields` (or `ignoredOutputOnlySpecFields`).
3. Run `make manifests` to regenerate the CRD.
4. Run `make resource-docs` to regenerate reference documentation.

### 2. Adding Resource References
When a field refers to another GCP resource managed by KCC:
1. Add an entry to `resourceReferences` with `tfField`, `key` (ending in `Ref`), `gvk`, and `targetField`.
2. Run `make manifests`.

### 3. Adding a Brownfield Resource
1. Check `scripts/resource-autogen/generated/servicemappings/<service>.yaml` for existing templates generated from Magic Modules.
2. Add the `ResourceConfig` block under `spec.resources`.
3. Run `make manifests` and `make resource-docs`.

---

## Detailed Playbook

For comprehensive instructions on patching the vendored Terraform provider, expander/flattener invariants, boolean zero-value mechanics, and integration testing, see:

👉 **[Terraform-Based Field and Resource Playbook](../../docs/develop-resources/scenarios/terraform-field-and-resource.md)**
