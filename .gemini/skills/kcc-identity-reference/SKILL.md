---
name: kcc-identity-reference
description: Creates or updates the _identity.go and _reference.go files for a Config Connector resource, ensuring they follow the canonical gcpurls.Template pattern. Use this when you need to make sure the identity and reference is up to date for a KCC resource or when implementing IdentityV2 and refs.Ref for a resource.
---

# Kcc Identity Reference

## Overview

This skill guides you through creating or updating the `_identity.go` and `_reference.go` files for a given Config Connector resource (e.g. `MemoryStoreInstance`). The goal is to implement `identity.IdentityV2` and `refs.Ref` using the canonical `gcpurls.Template` pattern, identical to the pattern used in `apis/artifactregistry/v1beta1/artifactregistryrepository_identity.go`.

## Hints

* If the identity or reference files already exist, do not update the copyright year.

## Workflow

When asked to update or create the identity and reference for a "resource of the day" (e.g., `group: memorystore.cnrm.cloud.google.com`, `kind: MemoryStoreInstance`), follow these steps:

### Step 1: Locate the target files

1. Identify the `groupPrefix` (the group without the `.cnrm.cloud.google.com` suffix). Example: `memorystore`
2. Look for the resource version directory under `apis/<groupPrefix>/`. Example: `apis/memorystore/v1beta1/` or `apis/memorystore/v1alpha1/`. (Check the filesystem to see which versions exist).
3. The files to edit/create are `apis/<groupPrefix>/<version>/<kind>_identity.go` and `apis/<groupPrefix>/<version>/<kind>_reference.go`.
   - Note: `<kind>` is typically the full lowercase Kind, e.g. `memorystoreinstance`. Sometimes the `groupPrefix` is dropped, so check if the files already exist.

### Step 2: Determine the Identity Template

1. Read the corresponding line in `docs/ai/metadata/cloudassetinventory_names.jsonl` using grep. Search for the resource kind to find its URL format.
   - Example: `grep -i memorystore docs/ai/metadata/cloudassetinventory_names.jsonl`
   - Output might be: `{"resourceType": "memorystore.googleapis.com/Instance", "nameFormats": ["//memorystore.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/instances/{{INSTANCE}}"]}`
   - **Note:** If the resource is missing from `cloudassetinventory_names.jsonl` (e.g. not handled by CAIS), check the existing `_identity.go` or direct controller to infer the URL format. Pay attention to camelCase path segments (e.g. `entryGroups` instead of `entrygroups`), as GCP URLs are case-sensitive. Additionally, if the resource is missing from `cloudassetinventory_names.jsonl`, you will also need to add its URL format as an exception in `pkg/gcpurls/registry_test.go` to prevent the `TestRegisteredTemplatesMatchCAI` test from failing.
2. Map the format to the `gcpurls.Template` format: `"projects/{project}/locations/{location}/instances/{instance}"`.
3. Read the canonical `apis/artifactregistry/v1beta1/artifactregistryrepository_identity.go` to refresh your understanding of the implementation details.

### Step 3: Implement the Identity (`<kind>_identity.go`)

Create or update the file to match the canonical example. Key requirements:
- Use the standard copyright header (Year 2026).
- Declare interface implementations: `_ identity.IdentityV2 = &<Kind>Identity{}` and `_ identity.Resource = &<Kind>{}`
- Define the template var: `var <Kind>IdentityFormat = gcpurls.Template[<Kind>Identity]("api.googleapis.com", "projects/{project}/...")`
- The struct must map exactly to the template fields (e.g., `Project string`, `Location string`, `Instance string`) and have `// +k8s:deepcopy-gen=false`.
  - **Important:** The variables in your `gcpurls.Template` (e.g. `{instance}`) MUST match the struct fields when both are lowercased (e.g. `{deploymentresourcepool}` matches `DeploymentResourcePool`). Do not use underscores in the template variables (e.g. `{deployment_resource_pool}`) if your struct field is CamelCased, as `gcpurls.Template` will panic at initialization.
  - **Note:** If an existing deepcopy method was previously generated for this identity struct, run `dev/tasks/generate-types-and-mappers` to regenerate the types and remove the obsolete code.
- Implement `String()`, `FromExternal(ref string)`, and `Host()` by delegating to the format var.
- Implement `getIdentityFrom<Kind>Spec(...)` to extract fields from the spec/obj (often using `refs.ResolveProjectID`, `refs.GetLocation`, etc.).
- Implement `GetIdentity(ctx, reader)` on the Resource struct, including cross-checking `externalRef` or `status.Name`. (Look at `artifactregistryrepository_identity.go`'s `GetIdentity` implementation for exactly how to do this cross-check).
  - **Note:** If you are updating an existing resource's Identity struct to the IdentityV2 pattern, be sure to check for existing usages of the struct and its old methods (e.g. `.Parent()`, `.ID()`) in dependent identity files and direct controllers, and update them to use the new fields (e.g. `.Project`, `.Location`, etc.).  The compiler is your friend: remove the functions, then run `go vet ./...` or `go build ./...` to look for references to functions that no longer exist.

### Step 4: Implement the Reference (`<kind>_reference.go`)

Read the canonical `apis/artifactregistry/v1beta1/artifactregistryrepository_reference.go` to refresh your understanding.

Create or update the file to match the canonical example. Key requirements:
- Use the standard copyright header (Year 2026).
- Implement `_ refs.Ref = &<Kind>Ref{}`.
- Define the GVK variable: `var <Kind>GVK = schema.GroupVersionKind{...}` (It is also acceptable if this is defined in `<kind>_types.go`).
- Define the `<Kind>Ref` struct with exactly 3 fields: `External`, `Name`, and `Namespace`.
  - The `External` field MUST have specific godoc: `"A reference to an externally managed <Kind> resource. Should be in the format \"projects/{{projectID}}/...\""`. Do not use generic docstrings.
  - The `Name` and `Namespace` fields should have godocs: `"The name of a <Kind> resource."` and `"The namespace of a <Kind> resource."`.
- Include `func init() { refs.Register(&<Kind>Ref{}) }`.
- Implement boilerplate methods: `GetGVK`, `GetNamespacedName`, `GetExternal`, `SetExternal`, `ValidateExternal`, `ParseExternalToIdentity`.
- Implement `Normalize` delegating to `refs.NormalizeWithFallback`. In the fallback function `func(u *unstructured.Unstructured) string`, simply pass `u` directly to `getIdentityFrom<Kind>Spec` since `*unstructured.Unstructured` implements `client.Object`.

### Step 5: Verify

Ensure the code compiles and there are no lint errors. You MUST always run `go vet ./...` and `go build ./...` before sending the PR to verify that your changes have not introduced any compilation errors across the entire project.
