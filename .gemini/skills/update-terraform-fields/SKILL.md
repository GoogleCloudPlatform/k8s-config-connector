---
name: update-terraform-fields
description: Guides modifying, patching, and updating resources and fields managed by the Terraform/legacy controller, including patching the local/vendored Terraform Google Beta provider.
---

# KCC Update Terraform Fields (Agentic-Friendly Guide)

This skill provides step-by-step instructions for an AI agent to add or update fields or resources managed by the Terraform/legacy controller (distinguished by the label `cnrm.cloud.google.com/tf2crd: "true"` or the static mappings). It also covers patching the local/vendored copy of the Terraform Google Beta Provider (TPG Beta).

---

## 1. Pre-requisites & Codebase Discovery

Before making any changes, the agent must inspect the repository to locate the resource files:

1. **Identify Kind and API Group**:
   - Determine the Kubernetes resource Kind (e.g., `SpannerDatabase`) and the api group (e.g., `spanner.cnrm.cloud.google.com`).
2. **Find the Terraform Resource Name**:
   - Locate the service mapping file in `config/servicemappings/<service>.yaml` (e.g., `config/servicemappings/spanner.yaml`).
   - Find the mapping block matching the Kind to extract the underlying Terraform resource name (e.g., `google_spanner_database`).
3. **Verify the Controller Mode**:
   - Check if the resource is managed by the Terraform controller by verifying that `config/servicemappings/<service>.yaml` contains the kind and does NOT override the controller to `direct` in `pkg/controller/resourceconfig/static_config.go`.
4. **Identify Go Types Existence**:
   - Run a search (e.g., `find_by_name`) for `apis/<service>/<version>/<kind>_types.go`.
   - **If it exists**: Go to **Section 3 (Case B)**.
   - **If it does NOT exist**: Go to **Section 3 (Case A)**.

---

## 2. Modifying the Vendored Terraform Provider (TPG Beta)

If the field or bug fix is not yet present in the vendored copy of TPG Beta under `third_party/github.com/hashicorp/terraform-provider-google-beta/`, patch the provider locally.

### Steps to Patch TPG Beta:
1. **Locate the Provider Code**:
   - Locate the schema and flattener/expander files under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/`.
   - Common files are named `resource_<terraform_resource_name>.go` or `node_config.go`.
2. **Implement Changes**:
   - Add/update the field in the schema map (`Schema: map[string]*schema.Schema`).
   - Ensure the field type (`TypeSchema`), description, optional/required attributes, and `ForceNew` properties match the GCP API behavior.
   - Update the corresponding **expander** (which maps KRM/Terraform values to GCP SDK structures) and **flattener** (which maps GCP responses back to Terraform schema maps).
3. **Handle Optional/ForceNew Block Flattening (Critical Pitfall)**:
   - For optional blocks/structs that are immutable (`ForceNew: true`), the flattener MUST explicitly return an empty slice or `nil` if all attributes inside the block are empty. Otherwise, Terraform calculates a diff (from 0 blocks to 1 block containing default values) and fails the update.
   - Example:
     ```go
     func flattenMyConfig(c *service.MyConfig) []map[string]interface{} {
         if c == nil || c.SomeAttribute == "" {
             return nil // Or return []map[string]interface{}{}
         }
         ...
     }
     ```
4. **Verify Compilation**:
   - Run the compiler check locally:
     ```bash
     go vet ./third_party/github.com/hashicorp/terraform-provider-google-beta/...
     ```
5. **Commit Provider Changes Separately (Repository Rule)**:
   - Run `git status` and verify only files under `third_party/` are modified.
   - Stage and commit the third-party changes in a **dedicated, separate git commit** before editing other KCC configurations:
     ```bash
     git add third_party/
     git commit -m "Patch vendored Terraform provider: add support for <field_name> on <resource>"
     ```

---

## 3. Updating KCC CRD Definitions

### Case A: The resource is NOT yet migrated to Go types in `apis/`
KCC dynamically generates the CRD directly from the Terraform schema:

1. **Modify Service Mapping**:
   - Open `config/servicemappings/<service>.yaml`.
   - Locate the resource definition block.
   - Ensure the field is not listed in `ignoredFields` or `ignoredOutputOnlySpecFields`.
   - **Resource References**: If the field points to another KCC resource, define it under `resourceReferences`:
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
   - **Observed Fields**: If the field is output-only and should be updated in the resource status, specify it in `observedFields`.
2. **Regenerate CRD Manifests**:
   - Run:
     ```bash
     make manifests
     ```
   - Run `git diff config/crds/resources/` to verify that the new field is properly generated in the CRD schema.

### Case B: The resource HAS Go types in `apis/`
If a `<kind>_types.go` file exists under `apis/<service>/<version>/`:

1. **Update Go Types**:
   - Open `apis/<service>/<version>/<kind>_types.go`.
   - Locate the `<Kind>Spec` struct (or the nested block struct).
   - Add the field using the correct Go type and tag. Use camelCase for the json tag:
     ```go
     // +optional
     MyNewField *string `json:"myNewField,omitempty"`
     ```
   - Ensure you use appropriate kubebuilder annotations (e.g. `// +optional`, `// +kubebuilder:validation:Enum`).
2. **Regenerate CRD & Mappers**:
   - Run:
     ```bash
     dev/tasks/generate-types-and-mappers
     ```
   - Verify the generated files (`_types.go` tags, auto-generated mappers, and `config/crds/resources/`) via `git diff`.
3. **Update Go Client**:
   - Run:
     ```bash
     make generate-go-client
     ```

---

## 4. Verification & E2E Testing

Once definitions are updated, the agent must verify correctness of the field using the fixture testing framework.

1. **Check Copyright Headers**:
   - If you created any new files, make sure they have the copyright header:
     ```go
     // Copyright 2026 Google LLC
     //
     // Licensed under the Apache License, Version 2.0 (the "License");
     // you may not use this file except in compliance with the License.
     ...
     ```
2. **Proceed to Test Skill**:
   - Refer to [.gemini/skills/test-terraform-fields/SKILL.md](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/test-terraform-fields/SKILL.md) to set up E2E tests, record golden files, check against MockGCP, and run linters.
