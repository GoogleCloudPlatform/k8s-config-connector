# ComputePerInstanceConfig Greenfield Types Implementation Journal

## Observations & Implementation Details

1. **Handling of Message-Valued Maps**:
   - The GCP Proto for `google.cloud.compute.v1.PreservedState` utilizes map fields with complex message values (e.g., `map<string, PreservedStatePreservedDisk> disks`).
   - The KCC `generate-types` scaffolding tool marked these maps as `unsupported map type with key string and value message` and commented them out in `types.generated.go`.
   - To preserve 100% backward compatibility with the existing legacy/Terraform KRM schema for `ComputePerInstanceConfig`, we hand-coded the types (`PerinstanceconfigPreservedState`, `PerinstanceconfigDisk`, `PerinstanceconfigExternalIp`, `PerinstanceconfigInternalIp`, `PerinstanceconfigIpAddress`) as a list of structs/slices exactly matching the established schema.

2. **Scaffolding Missing References**:
   - The `ComputePerInstanceConfig` spec references `ComputeInstanceGroupManager` via `instanceGroupManagerRef`.
   - Because `ComputeInstanceGroupManager` is still a legacy resource and did not have its reference type defined under `apis/`, we created a new reference file at `apis/compute/v1beta1/instancegroupmanager_reference.go`.
   - This new type implements the `refsv1beta1.Ref` and `refsv1beta1.ExternalNormalizer` interfaces, fully conforming to KCC's strict reference guidelines.

3. **Validation and Code Generation**:
   - Successfully ran the generation script which compiles, validates, and automatically outputs deepcopy functions (`zz_generated.deepcopy.go`) and the updated CRD manifest `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_computeperinstanceconfigs.compute.cnrm.cloud.google.com.yaml` containing the complete and correct schema.
   - Performed formatting checks (`make fmt`) and static analysis (`go vet ./apis/compute/...`) confirming zero compilation issues.
