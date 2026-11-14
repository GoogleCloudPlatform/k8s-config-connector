### [2026-11-13] Memorystore Google APIs Pin Update & Skip-Proto for AclPolicy
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `MemorystoreACLPolicy` (Greenfield Step 1a).
- **Problem**:
  1. The default googleapis commit pin in `apis/git.versions` does not include the `AclPolicy` protobuf message in `google/cloud/memorystore/v1/memorystore.proto`. Updating `apis/git.versions` globally breaks other APIs (such as `apigeeregistry`).
  2. The Go client library in `go.mod` (`cloud.google.com/go/memorystore v0.9.0`) does not yet include `AclPolicy` in `apiv1/memorystorepb`. Upgrading `memorystore` to `v1.9.0` would transitively pull in a newer `google.golang.org/genproto` that removes deprecated packages used by other controllers in the repository.
- **Solution**:
  1. Pinned `PROTO_SHA="97f6460c3558c1435d7fd001a9251f8df058ae2f"` in `apis/memorystore/generate.sh` where `AclPolicy` was introduced and passed `--proto-source-path ${PROTO_OUT}`, keeping the global `apis/git.versions` intact.
  2. Configured `--skip-proto AclPolicy,AclRule,AclPolicyRevision,AclPolicyRevisionStatus,InstanceAclPolicyAttachment` on `generate-mapper` in `apis/memorystore/generate.sh` to avoid generating `mapper.generated.go` references to types not present in `memorystore v0.9.0`.
- **Impact**: Enables mechanical prepopulation of `MemorystoreACLPolicy` types, identity, and CRD without causing compilation or dependency conflicts across the workspace.

### [2026-11-13] MemorystoreACLPolicy Judgement Queue Triage (Step 1b)
- **Context**: Triaging judgement queue and finalizing direct KRM types, identity, and reference for `MemorystoreACLPolicy` under `memorystore.cnrm.cloud.google.com/v1alpha1`.
- **Triage Decisions**:
  1. `location-or-parent-ref`: `MemorystoreACLPolicy` is a regional project-parented resource with pattern `projects/{project}/locations/{location}/aclPolicies/{acl_policy}`. Replaced `spec.projectRef` and `spec.location` with inlined `*parent.ProjectAndLocationRef`.
  2. `identity-not-in-cai`: The resource URL format matches the proto `(google.api.resource)` pattern `projects/{project}/locations/{location}/aclPolicies/{aclPolicy}`. Accepted as valid even though not indexed in CAIS.
  3. `untriaged-bulk-generation`: Resolved after verifying all spec fields (`rules` with Valkey ACL username and rules) and output-only observed state fields (`state`, `etag`, `instanceAclPolicyAttachments`, `createTime`, `updateTime`).
- **Impact**: All findings for `MemorystoreACLPolicy` in `apis/memorystore/judgement_queue.yaml` resolved cleanly, and CRD validation updated.

