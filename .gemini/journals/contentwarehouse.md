### [2026-05-26] ContentWarehouseSchema Type Generation Quirks
- **Context**: Implementing Greenfield types for `ContentWarehouseSchema` mapped to `DocumentSchema`.
- **Problem**: `contentwarehouse.cnrm.cloud.google.com/v1alpha1/ContentWarehouseSchema`'s proto spec uses embedded definitions like `PropertyDefinition`, which the generator initially places in `types.generated.go` without required validations. Since the generator marks them as `unreachable` when not directly referenced, they get commented out or pruned if not explicitly copied.
- **Solution**: Copied the unreferenced struct definitions from `types.generated.go` directly into `documentschema_types.go` and added `+kubebuilder:validation:Required` tags manually. Rerunning `generate.sh` correctly pruned the duplicate definitions from `types.generated.go`. Also added `contentwarehouse.googleapis.com` to `pkg/gcpurls/registry_test.go` exception list.
- **Impact**: When scaffolding new resources where fields are deeply nested or not automatically mapped to the root `Spec` object, agents must manually extract those types to the `*_types.go` file to allow `prunetypes.go` to clean up the `types.generated.go` and to safely apply `kubebuilder` validation markers without them getting overwritten on subsequent generations.

### [2026-06-05] ContentWarehouse RuleSet Initial Types and Identity
- **Context**: Implementing ContentWarehouseRuleSet KRM types, CRD, and IdentityV2 under `v1alpha1`.
- **Problem**: The proto generator skipped compiling all nested types (such as `Rule`, `Action`, etc.) inside `types.generated.go`, labeling them as "unreachable" because they were not explicitly referenced by Go types during the first generation.
- **Solution**: Copied the nested structures explicitly into `ruleset_types.go` and mapped them to the `ContentWarehouseRuleSetSpec` type with `+optional` tags. This ensured all nested structs are fully typed and accessible, and subsequent generation runs cleanly marked them as reachable and fully compiled.
- **Impact**: Ensures schema compatibility and complete coverage for nested array structures under direct types.
