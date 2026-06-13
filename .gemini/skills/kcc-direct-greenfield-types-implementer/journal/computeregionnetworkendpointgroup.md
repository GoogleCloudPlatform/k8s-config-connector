# Greenfield Types Implementation Journal - ComputeRegionNetworkEndpointGroup

## Key Observations and Quirks

### 1. Commented-out "Unreachable" Types in `types.generated.go`
When generating KRM types from a complex proto service like Google Cloud Compute Engine, types that are not directly referenced by any generated `Spec` types will be pruned and commented out under `types.generated.go` as `/* unreachable type ... */`.
To support these nested structures in hand-written KRM types, we can manually copy/uncomment the required types or define them directly inside our handcrafted `_types.go` file (e.g. `NetworkEndpointGroupCloudFunction` and `NetworkEndpointGroupCloudRun` under `networkendpointgroup_types.go`).

### 2. Breaking Circular Package Import Cycles
In direct KRM resources, references to other direct resources are typically typed using their specific direct KRM type package (e.g. `runv1beta1.RunServiceRef`). However, if package `run` already imports `compute`, having `compute` import `run` creates a circular dependency.
To resolve Go import cycles without losing KRM schema-level compatibility, we can define the required reference types (such as `RunServiceRef` or `CloudFunctionsFunctionRef`) under the shared and decoupled reference package `apis/refs/v1beta1/`. Both `run` and `compute` can import `refs` without any package dependency loops.
