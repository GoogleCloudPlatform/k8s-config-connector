### 2026-06-13 Scaffolding ComputeInstanceGroupNamedPort Direct Types
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeInstanceGroupNamedPort` ([Issue #9839](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9839)).
- **Problem**: The old KCC implementation of `ComputeInstanceGroupNamedPort` was managed via Terraform, and we are scaffolding the direct KRM types to enable migrating it to a direct controller. The GCP `NamedPort` proto message is part of an `InstanceGroup` resources' `named_ports` array, rather than a standalone GCP API resource.
- **Solution**:
  - Registered `ComputeInstanceGroupNamedPort:NamedPort` inside `apis/compute/v1alpha1/generate.sh`.
  - Created `ComputeInstanceGroupIdentity` to parse and represent the format `projects/{project}/zones/{zone}/instanceGroups/{instanceGroup}` for referencing `ComputeInstanceGroup` resources.
  - Implemented `ComputeInstanceGroupRef` in `apis/compute/v1alpha1/computeinstancegroup_reference.go` as a proper KCC reference implementing `refs.Ref`.
  - Scaffolded and refined `apis/compute/v1alpha1/namedport_types.go` with exact fields matching the existing CRD definition, ensuring that `groupRef` uses our new `ComputeInstanceGroupRef`, `port` uses `int32`, and GVK/stability labels are correctly set.
  - Successfully regenerated CRD files by executing `generate.sh` and formatted all files with `make fmt`.
