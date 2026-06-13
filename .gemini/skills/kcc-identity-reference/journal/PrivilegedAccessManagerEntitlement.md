# PrivilegedAccessManagerEntitlement

Journal of moving `PrivilegedAccessManagerEntitlement` to the canonical `IdentityV2` and `refs.Ref` pattern.

## Observations & Learnings

### 1. Multiple Optional Parents
`PrivilegedAccessManagerEntitlement` supports multiple parent resource types: `ProjectRef`, `FolderRef`, and `OrganizationRef`.
The custom parent resolution logic was originally implemented via a manual controller helper `oneOfContainer`.
We cleanly transitioned this logic to the canonical Identity format in `apis/privilegedaccessmanager/v1beta1/privilegedaccessmanagerentitlement_identity.go`.
Our `getIdentityFromPrivilegedAccessManagerEntitlementSpec` function matches this behavior, resolving:
- `ProjectRef` (or falling back to project ID from namespace if no parents are specified).
- `FolderRef`
- `OrganizationRef`

### 2. Missing from CAI (Cloud Asset Inventory)
The resource `privilegedaccessmanager.googleapis.com/Entitlement` is not tracked/defined in `docs/ai/metadata/cloudassetinventory_names.jsonl` (though `Grant` is).
Following the instruction for missing resources, we:
- Derived the URL templates from the direct controller's original implementation.
- Added exceptions in `pkg/gcpurls/registry_test.go` (`ignoredTemplates` map) to avoid `TestRegisteredTemplatesMatchCAI` failure.

### 3. Identity and Reference Struct Methods Compatibility
To keep the refactoring of `pkg/controller/direct/privilegedaccessmanager/entitlement_controller.go` minimal, surgical, and type-safe, we implemented helper methods directly on `PrivilegedAccessManagerEntitlementIdentity`:
- `Container()`: returning `projects/{project}`, `folders/{folder}`, or `organizations/{organization}`.
- `ParentString()`: returning `<container>/locations/{location}`.
- `FullyQualifiedName()`: matching `String()`.
- `AsExternalRef()`: returning `//privilegedaccessmanager.googleapis.com/<path>`.

This allowed us to simplify `AdapterForObject`, remove the redundant `oneOfContainer` and `checkExactlyOneOf` from the controller, and keep the rest of the controller extremely clean and compatible without rewriting all client/GCP calls.
