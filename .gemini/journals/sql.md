### [2026-11-18] Prepopulating SQLBlueGreenDeployment
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `SQLBlueGreenDeployment` using `google.cloud.sql.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `BlueGreenDeployment` in `google.cloud.sql.v1` is defined in `google/cloud/sql/v1/cloud_sql_blue_green_deployments.proto`, added in googleapis commit `e7513c828b8db536fa12fc5493d541f323aedf90`.
  2. Following the pattern established in `apis/gkehub/generate.sh` and `apis/netapp/generate.sh`, `apis/sql/generate.sh` generates a versioned protobuf bundle `.build/googleapis-${PROTO_SHA}.pb` for `v1alpha1` while keeping `v1beta1` pinned to the stable default proto.
  3. Used Kind name `SQLBlueGreenDeployment` under API group `sql.cnrm.cloud.google.com/v1alpha1` following maintainer architectural decision.
  4. Added `cloud.google.com/go/sql v0.7.0` to `go.mod` for `cloud.google.com/go/sql/apiv1/sqlpb`.
  5. Configured `generate-mapper` with `--multiversion` and `--prepopulate-resource` across `google.cloud.sql.v1beta4,google.cloud.sql.v1`. Added backward compatibility forwarding wrappers for `v1beta1` helpers in `pkg/controller/direct/sql/mapper.go` so all existing `v1beta1` tests and fuzzers continue to pass cleanly.
  6. Added `SQLBlueGreenDeployment` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  7. Synchronized golden files `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `kind_name_violations.txt`.
