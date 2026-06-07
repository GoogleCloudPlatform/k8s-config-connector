# Journal: ServiceDirectoryService Transition to Direct KRM types

## Learnings & Observations

### 1. Handling of missing ObservedState in CRD schema
When a baseline CRD's `status` doesn't have an `observedState` field (e.g. `ServiceDirectoryService`), we must not add it to the `ServiceDirectoryServiceStatus` Go struct to preserve strict schema compatibility.
However, because the GCP proto contains output-only fields (e.g., `create_time`, `update_time`, `uid`), the `generate-mapper` tool expects a `ServiceDirectoryServiceObservedState` struct to map these fields.
**Solution:** Define an empty `ServiceDirectoryServiceObservedState` struct in `service_types.go`, but do NOT reference it as a field inside `ServiceDirectoryServiceStatus`. This allows the generated mapper to compile successfully while ensuring `controller-gen` does not generate an `observedState` property in the final CRD.

### 2. Strict Schema Compatibility
Running `dev/tasks/diff-crds` returned an empty output, confirming that our handcoded `service_types.go` conforms perfectly and maintains 100% strict compatibility with the baseline custom resource definition.
