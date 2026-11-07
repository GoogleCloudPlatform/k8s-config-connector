Examples:


```
# To generate Go types for BigtableInstance
go run . generate-types \
  --service google.bigtable.admin.v2 \
  --api-version bigtable.cnrm.cloud.google.com/v1beta1 \
  --resource BigtableInstance:Instance

# To generate mapping function between KRM and proto for BigtableInstance
go run . generate-mapper \
  --output-dir ~/kcc/k8s-config-connector/pkg/controller/direct/ \
  --service google.bigtable.admin.v2

# To scaffold generate the SecretManagerSecretVersion controller
go run . generate-direct-reconciler \
  --resource SecretManagerSecretVersion:SecretVersion \
  --api-version  "secretmanager.cnrm.cloud.google.com/v1beta1" \
  --service "google.cloud.secretmanager.v1"

```
## Kind names for new resources

A Kind is `<service prefix><proto message>`, with the acronyms in
[`serviceprefixes/acronyms.yaml`](serviceprefixes/acronyms.yaml) applied to the message: `sqladmin`
+ `BackupRun` gives `SQLAdminBackupRun`. So only the service prefix needs
deciding, once per API.

Approved service prefixes are in the hand-maintained
[`serviceprefixes/service-prefixes.yaml`](serviceprefixes/service-prefixes.yaml).
Proposed prefixes waiting for approval are in
[`serviceprefixes/service-prefix-pending.yaml`](serviceprefixes/service-prefix-pending.yaml),
each with a `reason:`. To approve one, move its entry to `service-prefixes.yaml`
(changing `servicePrefix` if needed) and record the decision in `reason:`. To
propose a prefix for a new API, add it to the pending file.

Each entry also lists, in `protoPackages`, the proto packages of the API that
`apis/*/generate.sh` passes to `--service` (e.g. `google.cloud.sql.v1` for
`sqladmin`). When a `generate.sh` starts using a new package (for example a new
API version), add it to the API's `protoPackages`.

`TestKindNamesFollowServicePrefixes` in `tests/apichecks` checks every
`generate-types --resource <Kind>:<Message>` in `apis/*/generate.sh`: it finds
the API from the `--service` package, the API's prefix must be approved, and the
Kind must be the service prefix followed by the message. Kinds that keep an
older name (for example from Terraform) are listed in `knownKindNames` in the
test, and existing violations in
`tests/apichecks/testdata/exceptions/kind_name_violations.txt`.
