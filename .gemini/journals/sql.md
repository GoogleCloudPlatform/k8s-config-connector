### [2026-11-19] Prepopulating SQLBlueGreenDeployment
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `SQLBlueGreenDeployment` using `google.cloud.sql.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `BlueGreenDeployment` in `google.cloud.sql.v1` is defined in `google/cloud/sql/v1/cloud_sql_blue_green_deployments.proto`, added in googleapis commit `e7513c828b8db536fa12fc5493d541f323aedf90`.
  2. Following the pattern established in `apis/gkehub/generate.sh` and `apis/netapp/generate.sh`, `apis/sql/generate.sh` generates a versioned protobuf bundle `.build/googleapis-${PROTO_SHA}.pb` for `v1alpha1` while keeping `v1beta1` pinned to the stable default proto.
  3. Used Kind name `SQLBlueGreenDeployment` under API group `sql.cnrm.cloud.google.com/v1alpha1` following maintainer architectural decision.
  4. Added `cloud.google.com/go/sql v0.7.0` to `go.mod` for `cloud.google.com/go/sql/apiv1/sqlpb`.
  5. Configured `generate-mapper` with `--multiversion` and `--prepopulate-resource` across `google.cloud.sql.v1beta4,google.cloud.sql.v1`. Added backward compatibility forwarding wrappers for `v1beta1` helpers in `pkg/controller/direct/sql/mapper.go` so all existing `v1beta1` tests and fuzzers continue to pass cleanly.
  6. Added `SQLBlueGreenDeployment` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  7. Synchronized golden files `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `kind_name_violations.txt`.

### [2026-11-19] Triaging Judgement Queue for SQLBlueGreenDeployment
- **Context**: Triaged and resolved all entries in `apis/sql/judgement_queue.yaml` for `SQLBlueGreenDeployment`.
- **Learnings & Observations**:
  1. `BlueGreenDeployment` is a regional project-parented resource (`projects/{project}/locations/{location}/blueGreenDeployments/{blue_green_deployment}`), resolved using inlined `*parent.ProjectAndLocationRef`.
  2. `source_instance` is an authoritative proto reference to `sqladmin.googleapis.com/Instance`, modeled as `SourceInstanceRef *refsv1beta1.SQLInstanceRef` (`json:"sourceInstanceRef"`) and mapped automatically by `generate-mapper`.
  3. `CreateBlueGreenDeploymentRequest` has a required `blue_green_deployment_id` client parameter, confirming user-specified ID and conformance with `identity.IdentityV2`.
  4. Selected canonical v1 pattern `projects/{project}/locations/{location}/blueGreenDeployments/{blueGreenDeployment}` to resolve `identity-multi-pattern`.
  5. Removed unverified `resource-docs` link guess as public REST documentation page is not yet live.

### [2026-11-19] Direct Controller, Fuzzer, and E2E Fixtures for SQLBlueGreenDeployment
- **Context**: Implemented direct controller, fuzzer, and recorded E2E fixtures against real GCP for `SQLBlueGreenDeployment`.
- **Learnings & Observations**:
  1. **LRO Polling with SqlOperationsClient**: `BlueGreenDeploymentsClient.CreateBlueGreenDeployment` and `DeleteBlueGreenDeployment` return Cloud SQL Admin operations (`*sqlpb.Operation`). The controller configures `SqlOperationsClient` (`NewSqlOperationsRESTClient`) and polls operation completion (`pollOperation`) until `Operation_DONE`, checking `op.GetError()`.
  2. **Immutable Resource Reconciliation**: `SQLBlueGreenDeployment` has no `Update` RPC on GCP. In `Update`, the controller validates drift via `common.CompareGreenfieldSpec` and updates status in-place.
  3. **source_instance Normalization**: GCP returns `sourceInstance` in `GetBlueGreenDeployment` as the short instance ID (`sqlinstance-...`) instead of the full path (`projects/{project}/instances/{instance}`) provided on Create. In `Update`, a normalizer strips the resource prefix so both desired and actual compare as the short instance ID, avoiding false immutable field modification diffs during re-reconciliation.
  4. **Reference Resolution**: `SourceInstanceRef` (`*refsv1beta1.SQLInstanceRef`) is resolved in `AdapterForObject` via `refsv1beta1.ResolveSQLInstanceRef` to format `External` as `projects/{project}/instances/{instance}`. In `Export`, any unqualified instance name is prefixed with `projects/{project}/instances/` so exported manifests can be re-resolved.
  5. **Real GCP Fixtures & Prerequisites**: Cloud SQL Blue-Green deployment requires the source `SQLInstance` to have automated backups and binary logging enabled (`settings.backupConfiguration.enabled: true`, `binaryLogEnabled: true`). Minimal and maximal fixtures (`MYSQL_8_0` source, maximal specifying `MYSQL_8_4` target database version) were recorded and verified against real GCP with all validation checks passing.
