# ComputeFirewallPolicy Direct Type Implementation Journal

## Observations & Implementation Details

- **Brownfield-like Migration using Greenfield Skill**: Although `ComputeFirewallPolicy` was previously an existing DCL-based resource, the task instructed using the greenfield workflow skill.
- **Strict Schema Compatibility**:
  - **ID Type**: In the previous DCL-based CRD, the `id` field was defined as `type: string`. In the Compute proto, it is naturally a `uint64`. To avoid breaking backwards-compatibility, `ID` in the Go type `ComputeFirewallPolicyStatus` was modeled as `*string` rather than `*uint64`, which successfully keeps `id` as `type: string` in the generated CRD schema.
  - **Rule Tuple Count Type**: In the previous CRD, `ruleTupleCount` was `format: int64`. We used `*int64` instead of `*int32` to preserve the exact format.
- **Resource Reference Types**: Standardized on using `FolderRef` and `OrganizationRef` from the `refsv1beta1` package.
- **OneOf Validation**:
  - Used `x-kubernetes-validations` to ensure exactly one of `folderRef` or `organizationRef` is specified, which is the clean, declarative, modern kubebuilder/K8s OpenAPI v3 way.
