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