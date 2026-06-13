# ComputeOrganizationSecurityPolicyAssociation Journal

- **Date:** June 13, 2026
- **Author:** codebot-robot (Gemini CLI)

## Observations & Learnings

### 1. Protobuf Schema Mismatch in `v1` vs `v1beta`
During initial generation of the direct types using `--service google.cloud.compute.v1`, the generator reported:
`failed to find the proto message google.cloud.compute.v1.SecurityPolicyAssociation: proto: not found`

We investigated and confirmed that the public `google/cloud/compute/v1/compute.proto` does NOT define the `SecurityPolicyAssociation` message nor any of the related organization security policy messages. They are completely absent in `v1` (likely because organization-level security features are private or not published in the `v1` gRPC/proto public repository), even though they are available via the REST Discovery API (which MockGCP uses).

However, `SecurityPolicyAssociation` IS defined in the official `google/cloud/compute/v1beta/compute.proto`. 

**Solution:**
We mapped the resource to the versioned proto package by specifying:
`--resource ComputeOrganizationSecurityPolicyAssociation:google.cloud.compute.v1beta.SecurityPolicyAssociation` in `generate.sh`. This mirrors the pattern used for other resources like `ComputeFutureReservation` which is also mapped to the `v1beta` package.

### 2. Strict Schema Compatibility for Migrated Types
Because the type generator assumes a standard namespaced GCP resource by default (generating boilerplate `projectRef` and `location` in the Spec), we had to manually strip out the boilerplate and restore the original fields (`attachmentId` and `policyId`) to keep the schema strictly backward-compatible with the old TF-based CRD.

Comparing with the original TF-based `ComputeOrganizationSecurityPolicyAssociationSpec` helped us achieve perfect 100% backward compatibility (verified via `dev/tasks/diff-crds`).
