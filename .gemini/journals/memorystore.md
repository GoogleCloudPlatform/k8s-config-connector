### [2026-11-13] Memorystore Google APIs Pin Update & Skip-Proto for AclPolicy
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `MemorystoreACLPolicy` (Greenfield Step 1a).
- **Problem**:
  1. The default googleapis commit pin in `apis/git.versions` does not include the `AclPolicy` protobuf message in `google/cloud/memorystore/v1/memorystore.proto`. Updating `apis/git.versions` globally breaks other APIs (such as `apigeeregistry`).
  2. The Go client library in `go.mod` (`cloud.google.com/go/memorystore v0.9.0`) does not yet include `AclPolicy` in `apiv1/memorystorepb`. Upgrading `memorystore` to `v1.9.0` would transitively pull in a newer `google.golang.org/genproto` that removes deprecated packages used by other controllers in the repository.
- **Solution**:
  1. Pinned `PROTO_SHA="97f6460c3558c1435d7fd001a9251f8df058ae2f"` in `apis/memorystore/generate.sh` where `AclPolicy` was introduced and passed `--proto-source-path ${PROTO_OUT}`, keeping the global `apis/git.versions` intact.
  2. Configured `--skip-proto AclPolicy,AclRule,AclPolicyRevision,AclPolicyRevisionStatus,InstanceAclPolicyAttachment` on `generate-mapper` in `apis/memorystore/generate.sh` to avoid generating `mapper.generated.go` references to types not present in `memorystore v0.9.0`.
- **Impact**: Enables mechanical prepopulation of `MemorystoreACLPolicy` types, identity, and CRD without causing compilation or dependency conflicts across the workspace.
