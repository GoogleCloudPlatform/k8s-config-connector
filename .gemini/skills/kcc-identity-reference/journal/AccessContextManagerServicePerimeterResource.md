# Journal - AccessContextManagerServicePerimeterResource Migration

## Context
Migrating `AccessContextManagerServicePerimeterResource` to the modern identity and reference pattern (`IdentityV2` and `refs.Ref` with `gcpurls.Template`).

## Key Observations and Learnings

### 1. Name Collisions in Packages
The Kind/CRD name was `AccessContextManagerServicePerimeterResource`, which was already defined as a nested struct inside `apis/accesscontextmanager/v1beta1/serviceperimeter_types.go` for the `Resources` field of the service perimeter. 
Declaring the top-level KRM type `AccessContextManagerServicePerimeterResource` led to a redeclaration compilation error and caused `controller-gen` to panic during the CRD generation task.
* **Resolution:** Renamed the nested struct in `serviceperimeter_types.go` to `AccessContextManagerServicePerimeterResourceNested` (and updated its slice elements in `Resources []AccessContextManagerServicePerimeterResourceNested`), allowing the top-level KRM type to cleanly compile and generate deepcopy/CRD definitions.

### 2. Virtual/Non-Protobuf Resources
Some resources are virtual (created by Terraform to manage lists/nested properties as standalone resources) and do not correspond to any official Cloud Asset Inventory (CAI) name or GCP Protobuf message. 
* **Resolution:** 
  1. Manually specify the parent reference (`perimeterNameRef` referencing `AccessContextManagerServicePerimeter`) and standard resource references (`resourceRef` referencing a project).
  2. Implement the template manually mapping to the composite GCP URL structure: `accessPolicies/{accessPolicy}/servicePerimeters/{servicePerimeter}/projects/{project}`.
  3. Add the normalized template `//accesscontextmanager.googleapis.com/accessPolicies/{}/servicePerimeters/{}/projects/{}` to the `ignoredTemplates` list in `pkg/gcpurls/registry_test.go` to bypass the CAI definition check.
