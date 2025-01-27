gcloud_base: gcloud filestore instances
script: mockfilestore/testdata/instance/crud/script.yaml
http_host: file.googleapis.com
http.host: file.googleapis.com
proto.service: google.cloud.filestore.v1.CloudFilestoreManager
proto.resource: Instance
proto.message: google.cloud.filestore.v1.Instance
proto.message: google.cloud.filestore.v1.Instance

```
mkdir -p mockfilestore

controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mockfilestore/service.go
// +tool:mockgcp-service
// http.host: file.googleapis.com
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
EOF
```


```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mockfilestore/instance.go
// +tool:mockgcp-support
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
EOF
```

PROMPT:

I'm trying to implement the proto service google.cloud.filestore.v1.CloudFilestoreManager, and specifically the Instance resource, in my workspace. I've taken a first pass, but it looks like I have some compilation errors.  Can you fix those problems for me?

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to change the code.
* You can use VerifyCode to see any remaining compilation errors as you are iterating towards a solution.

The main implementation file is mockfilestore/instance.go.  

PROMPT:

Please add the services in `mockfilestore` to `mock_http_roundtrip.go`

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mockfilestore into the list of services.
* Don't forget to import the package!

