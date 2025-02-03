gcloud_base: gcloud filestore instances
script: mockfilestore/testdata/instance/crud/script.yaml
http_host: file.googleapis.com
proto.service: google.cloud.filestore.v1.CloudFilestoreManager
proto.resource: Instance
proto.message: google.cloud.filestore.v1.Instance
proto.message: google.cloud.filestore.v1.Instance

```
mkdir -p mockworkflows

controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mockworkflows/service.go
// +tool:mockgcp-service
// http.host: workflows.googleapis.com
// proto.service: google.cloud.workflows.v1.Workflows
EOF
```

# For filestore instance
```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mockworkflows/workflow.go
// +tool:mockgcp-support
// proto.service: google.cloud.workflows.v1.Workflows
// proto.message: google.cloud.workflows.v1.Workflow
EOF
```

# For backups
```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > mockfilestore/backup.go
// +tool:mockgcp-support
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Backup
EOF
```


PROMPT:

I'm trying to implement the proto service google.cloud.workflows.v1.Workflows, and specifically the Workflow resource, in my workspace. I've taken a first pass, but it looks like I have some compilation errors.  Can you fix those problems for me?

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to change the code.
* You can use VerifyCode to see any remaining compilation errors as you are iterating towards a solution.

The main implementation file is mockworkflows/workflow.go.  

PROMPT:

I'm trying to implement the proto service google.cloud.filestore.v1.CloudFilestoreManager, and specifically the Backup resource, in my workspace. I've taken a first pass, but it looks like I have some compilation errors.  Can you fix those problems for me?

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to change the code.
* You can use VerifyCode to see any remaining compilation errors as you are iterating towards a solution.

The main implementation file is mockfilestore/backup.go.  

PROMPT:

Please add the services in `mockworkflows` to `mock_http_roundtrip.go`

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mockfilestore into the list of services.
* Don't forget to import the package!

