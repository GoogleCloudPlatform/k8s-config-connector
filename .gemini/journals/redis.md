### [2026-11-19] Prepopulating RedisACLPolicy
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `RedisACLPolicy` using `google.cloud.redis.cluster.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `AclPolicy` in `google.cloud.redis.cluster.v1` is defined in `google/cloud/redis/cluster/v1/cloud_redis_cluster.proto`, added in googleapis commit `e7513c828b8db536fa12fc5493d541f323aedf90`.
  2. Following the pattern established in `apis/sql/generate.sh`, `apis/redis/generate.sh` generates a versioned protobuf bundle `.build/googleapis-${PROTO_SHA}.pb` for `v1alpha1` and `generate-mapper`.
  3. Upgraded `cloud.google.com/go/redis` from `v1.25.0` to `v1.29.0` in `go.mod` to pull in `AclPolicy` types in `clusterpb`.
  4. Configured `generate-mapper` with `--multiversion` and `--prepopulate-resource` across `google.cloud.redis.cluster.v1,google.cloud.redis.v1`. Added backward compatibility forwarding wrappers for `v1beta1` helpers in `pkg/controller/direct/redis/mapper.go` so all existing `v1beta1` tests pass cleanly.
  5. Added `RedisACLPolicy` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  6. Synchronized golden file `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-23] Finalizing RedisACLPolicy types and triaging judgement queue
- **Context**: Triaged `apis/redis/judgement_queue.yaml` entries for `RedisACLPolicy` and finalized KRM types and identity.
- **Learnings & Observations**:
  1. Replaced `projectRef` and `location` with inlined `*parent.ProjectAndLocationRef` in `RedisACLPolicySpec`.
  2. Verified `CreateAclPolicyRequest` in `google.cloud.redis.cluster.v1`: client specifies `acl_policy_id` (supported via `ResourceID`), confirming client-specified resource identity (`IdentityV2`).
  3. Fixed resource-docs link from guessed `https://docs.cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.aclPolicies` to `https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.aclPolicies`.
  4. Confirmed `AclRule` in `RedisACLPolicySpec` does not hold GCP control-plane resource references (`username` specifies Redis OSS username/IAM principal name, and `rule` specifies the Redis OSS ACL rule string).

