# ComputeInstanceGroupNamedPort Direct Types Journal

## 1. Context and Problem
We migrated `ComputeInstanceGroupNamedPort` KRM Go types from the legacy auto-generated types to the hand-scaffolded/direct types. The resource is an existing `v1alpha1` KRM type that utilizes `google.cloud.compute.v1.NamedPort` as its proto but also contains `groupRef`, `projectRef`, and `zone`.

## 2. Design Choices
- **CRD Schema Preservation**: The original baseline CRD schema is extremely clean and matches exactly. To maintain strict schema compatibility, we removed the generator-default fields like `externalRef` and `observedState` in `Status` since they were not part of the baseline CRD.
- **Reference Types Mapping**:
  - We used `*computev1beta1.ComputeInstanceGroupRef` for `groupRef` to leverage the existing canonical `ComputeInstanceGroupRef` reference type.
  - We used `*refs.ProjectRef` for `projectRef` since the project reference does not contain a `kind` field.
- **Field Mappings**:
  - `port`: Left as `int64` to match the baseline schema integer format, even though the GCP proto for NamedPort defines it as an `int32`.

## 3. Results
Running `generate.sh` and validating schemas via `dev/tasks/diff-crds` produced a completely identical KRM schema definition. This ensures 100% backward compatibility.
