gcloud_base: gcloud filestore instances
script: mockfilestore/testdata/instance/crud/script.yaml
http_host: file.googleapis.com
http.host: file.googleapis.com
proto.service: google.cloud.filestore.v1.CloudFilestoreManager
proto.resource: Instance

```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF
// +tool:mockgcp-service
// http.host: file.googleapis.com
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
EOF
```


```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF
// +tool:mockgcp-support
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.resource: Instance
EOF
```


PROMPT:

Please add the services in `mockfilestore` to `mock_http_roundtrip.go`

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mockfilestore into the list of services.
* Don't forget to import the package!
