# ComputeInstanceGroupNamedPort Direct Types Migration Journal

## Observations & Lessons Learned

### 1. Handling Alpha Project References without `kind`
For `v1alpha1` resources like `ComputeInstanceGroupNamedPort`, the baseline CRD of `projectRef` might lack a `kind` field. In such cases, importing the canonical `ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` (rather than `apis/refs/v1beta1`) matches this schema exactly and avoids schema diffs.

### 2. Omitted `observedState` in Baseline CRD
To ensure 100% strict schema compatibility with the baseline CRD, we must not include `observedState` in `ComputeInstanceGroupNamedPortStatus`. However, the KRM fuzzer framework (`NewKRMTypedFuzzer`) expects status mapper functions. 

We solved this beautifully by defining the `ComputeInstanceGroupNamedPortObservedState` struct as an unreferenced/dangling type in Go. Since it is never referenced as a field in `Status`, it has no impact on the CRD schema (diff-crds remains empty), but allows the fuzzer, deepcopy, and mappers to compile and run perfectly.

### 3. Sub-resources as Standalone KRM Types
In Google Cloud, named ports are sub-resources (nested array of name/port under `InstanceGroup` or configured via `setNamedPorts` API). In KCC, they are represented as standalone KRM types. The parent `groupRef` and `zone` are not part of the `NamedPort` message in GCP's API but are used in the controller adapter to set/locate the parent `InstanceGroup`.

### 4. Fast Mapper Generation
Running `./generate-proto.sh` via `generate.sh` can be slow due to network clones. We can run `go run . generate-mapper` directly in `dev/tools/controllerbuilder` to regenerate mappers instantly during development.
