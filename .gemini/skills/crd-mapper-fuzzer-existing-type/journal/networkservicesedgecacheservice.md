# NetworkServicesEdgeCacheService Journal

## Observations & Learnings

### Handling Missing Public Proto Definitions
Some GCP services (such as Media CDN and its `EdgeCacheService`) are proprietary or restrict public API access, meaning their protocol buffer definitions (`.proto` files) are not published in the canonical public `googleapis` repository.

Because `generate-types` requires a corresponding `.proto` message to generate Go structs, migrating these resources using automatic generation is not possible out-of-the-box. Attempting to add `--resource` to `generate-types` results in a compilation/lookup error (`failed to find the proto message ...: proto: not found`).

### Best Practices for Migration

1. **Leverage Clients Generated Go Types:**
   Use the existing Go structs located under `pkg/clients/generated/apis/<service>/<version>/<kind>_types.go`. These types already represent 100% of the CRD schema accurately. Copy, adapt, and place them under `apis/<service>/<version>/<kind>_types.go`.

2. **Standardize Project Reference:**
   Ensure the `ProjectRef` field uses the standard `refs.ProjectRef` type from the `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` package, rather than any client-scoped `ResourceRef` definition, as long as it matches the schema structure of `external`, `name`, `namespace` without a `kind` field.

3. **Align Spec Annotations:**
   Ensure `// +required` is added to the `Spec` field to guarantee that `spec` is generated as required in the OpenAPI validation of the CRD, maintaining exact structural schema compatibility with the baseline.

4. **Document Transition in generate.sh:**
   Instead of omitting documentation, add a clear comment block in `apis/<service>/<version>/generate.sh` noting that the resource's Go types are handcoded due to the missing public proto definitions.

5. **Exhaustive Validation:**
   Run `dev/tasks/generate-crds` and verify that `dev/tasks/diff-crds` returns empty. Run `make generate-go-client ensure fmt` to synchronize downstream clients and format.
