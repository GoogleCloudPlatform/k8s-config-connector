# Journal: ComputeRegionDiskResourcePolicyAttachment Implementation

## Observations & Architectural Choices

1. **No Proto Message Mapping**: 
   - `ComputeRegionDiskResourcePolicyAttachment` represents a relationship/action in the GCP Compute Engine API (attaching a resource policy to a regional disk via the `AddResourcePolicies`/`RemoveResourcePolicies` endpoints) rather than a single proto resource message.
   - For this reason, the resource types (`computeregiondiskresourcepolicyattachment_types.go`) were implemented directly/by-hand under `apis/compute/v1alpha1` and registered via the `SchemeBuilder`.

2. **Required Legacy Resource Reference**:
   - The resource spec requires referencing a `ComputeDisk`.
   - Since `ComputeDisk` is currently a legacy (Terraform-reconciled) resource in KCC and has not yet been migrated to the direct controller, there was no existing `ComputeDiskRef` or `ComputeDiskIdentity` in `apis/compute/v1beta1`.
   - We implemented:
     - `apis/compute/v1beta1/computedisk_identity.go`: To format/parse external zonal (`projects/{project}/zones/{zone}/disks/{name}`) and regional (`projects/{project}/regions/{region}/disks/{name}`) disk URLs.
     - `apis/compute/v1beta1/computedisk_reference.go`: To validate and normalize references pointing to `ComputeDisk` KRM objects, falling back to reading `status.selfLink` since `status.externalRef` is not yet populated for legacy resources.

3. **Validation & Generation Integrity**:
   - Running `./apis/compute/v1alpha1/generate.sh` correctly compiled and outputted updated deepcopy and CRD files.
   - `cnrm.cloud.google.com/tf2crd: "true"` label was removed from the CRD as it is now migrated from a pure TF definition to direct types.
   - All validation tags (`+required`, `+optional`, `stability-level=alpha`, and `+kubebuilder:subresource:status`) have been correctly applied, resulting in a valid generated CRD matching Config Connector's direct controllers standard.
