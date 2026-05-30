---
name: update-terraform-fields
description: Guides modifying, patching, and updating resources and fields managed by the Terraform/legacy controller, including patching the local/vendored Terraform Google Beta provider.
---

# KCC Update Terraform Fields

This skill guides an automated agent or developer through adding/updating fields or resources managed by the Terraform/legacy controller (marked by `cnrm.cloud.google.com/tf2crd: "true"`). It also covers patching the local/vendored copy of the Terraform Google Beta Provider (TPG Beta).

## Pre-requisites
1. Identify the resource Kind and API Group (e.g., `SpannerDatabase` under `spanner.cnrm.cloud.google.com`).
2. Identify the field to add/update, and its matching field in the Terraform Google Provider (e.g., `deletion_protection` in `google_spanner_database`).

---

## 1. Modifying the Vendored Terraform Provider (TPG Beta)

If the field/fix is not yet present in the vendored version of TPG Beta under `third_party/github.com/hashicorp/terraform-provider-google-beta/`, you may need to patch the provider locally.

### Steps to Patch TPG Beta:
1. **Implement Modifications**:
   - Make code changes directly inside the files under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/`.
   - Example: Adding a new field to `resource_spanner_database.go`'s schema.
2. **Verify Compilation**:
   - Run `make vet` (or `go vet ./...`) to verify that the modified provider code compiles successfully within the workspace.
3. **Commit TPG Beta Changes Separately**:
   - Group all changes under `third_party/github.com/hashicorp/terraform-provider-google-beta/` into a **single, dedicated git commit**.
   - Use a clear, descriptive commit message explaining the patch (e.g., `deletion_protection_default.patch` or `add_missing_field_to_resource.patch`).

---

## 2. Updating KCC CRD Definitions

Once the Terraform provider schema supports the desired field, propagate it to the KCC CRDs.

### Case A: The resource is NOT yet migrated to Go types in `apis/`
If the resource does NOT have a `<kind>_types.go` file under `apis/<service>/<version>/`, KCC dynamically generates the CRD from the Terraform schema:

1. **Configure ServiceMapping**:
   - Open `config/servicemappings/<service>.yaml`.
   - Ensure the field is not ignored (check `ignoredFields` and `ignoredOutputOnlySpecFields`).
   - If the field is a resource reference, map it under `resourceReferences`:
     ```yaml
     resourceReferences:
       - tfField: encryption.default_kms_key_name
         key: kmsKeyRef
         gvk:
           kind: KMSCryptoKey
           version: v1beta1
           group: kms.cnrm.cloud.google.com
         targetField: self_link
     ```
   - If the field is output-only and should be in the KCC observed state, add it under `observedFields`.
2. **Regenerate manifests**:
   - Run:
     ```bash
     make manifests
     ```
   - Verify the generated files under `config/crds/resources/` using `git diff`.

### Case B: The resource HAS Go types in `apis/` (but uses the Terraform controller)
If the resource has a `<kind>_types.go` file with `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"`:

1. **Update Go Types**:
   - Open `apis/<service>/<version>/<kind>_types.go`.
   - Add the field with the proper Go type and json tag matching the camelCase name of the Terraform field.
   - Example:
     ```go
     DeletionProtection *bool `json:"deletionProtection,omitempty"`
     ```
2. **Regenerate CRD and Go client**:
   - Run:
     ```bash
     dev/tasks/generate-types-and-mappers
     ```
   - Run `make manifests` if needed to ensure list of GVKs and general configurations are updated.
   - Run `make generate-go-client` to update Go clients.

---

## 3. Verify and Test

1. **Add Test Coverage**:
   - Add the new field / test case under `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/`.
2. **Update Golden Output**:
   - Generate golden HTTP logs and object files:
     ```bash
     WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/[testname]/
     ```
3. **Verify Compliance**:
   - Run `make ready-pr` to run formatters, linters, and verify the PR is ready for review.
