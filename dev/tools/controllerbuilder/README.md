Examples:


```
go run . generate-types  --proto-source-path ../proto-to-mapper/build/googleapis.pb \
  --service google.bigtable.admin.v2 --version v1beta1 \
  --output-api ~/kcc/k8s-config-connector
  --kinds BigtableInstance


go run . generate-mapper --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
      --output-dir ~/kcc/k8s-config-connector/pkg/controller/direct/ \
      --proto-source-path ../proto-to-mapper/build/googleapis.pb \
      --service google.bigtable.admin.v2 \
      --api-dir ~/kcc/k8s-config-connector/apis/

```