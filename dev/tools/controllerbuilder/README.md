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

A Kind is `<service prefix><resource type>`, with the resource type (from
`google.api.resource`) in the singular: `logging` + `logging.googleapis.com/Settings`
gives `LoggingSetting`. So only the service prefix needs deciding, once per API.

Approved service prefixes are in the hand-maintained
[`serviceprefixes/service-prefixes.yaml`](serviceprefixes/service-prefixes.yaml).
Proposed prefixes waiting for approval are in
[`serviceprefixes/service-prefix-exceptions.yaml`](serviceprefixes/service-prefix-exceptions.yaml),
each with a `reason:`. To approve one, move its entry to `service-prefixes.yaml`
(changing `servicePrefix` if needed) and record the decision in `reason:`. To
propose a prefix for a new API, add it to the exceptions file.

`generate-types --prepopulate-resource` checks every new resource against both
files: the Kind must be its API's service prefix followed by the singular
resource type, and the API must be in `service-prefixes.yaml` (resources of an
API that is only in the exceptions file are refused). A resource whose API is in
neither file is not checked. There is no override flag; change or approve the
prefix in the files instead.
