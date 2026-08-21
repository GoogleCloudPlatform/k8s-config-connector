# ComputeFirewallPolicyAssociation Direct Type Transition Journal

## Learnings & Observations
- **Resource Scope:** `ComputeFirewallPolicyAssociation` is not a project-scoped resource but a folder- or organization-level resource. Therefore, `projectRef` and `location` should not be added to its spec, and we need to verify and match the original CRD schema.
- **Reference Types:** Custom references like `attachmentTargetRef` (which can point to Folder or Organization resources) must be hand-coded locally in `firewallpolicyassociation_types.go` because they differ from canonical project/location schemes.
- **Strict Schema Compatibility:** By ensuring `Kind` in `attachmentTargetRef` remains optional at the OpenAPI property level (but required dynamically in the `oneOf` constraint added by `scripts/add-validation-to-crds`), `dev/tasks/diff-crds` was kept 100% empty.
- **Mapper Hand-coding:** Handwritten `FromProto` and `ToProto` functions were implemented in a dedicated `pkg/controller/direct/compute/firewallpolicyassociation_mapper.go` file to elegantly translate the structured KRM reference types to/from flat string fields in the proto message.
