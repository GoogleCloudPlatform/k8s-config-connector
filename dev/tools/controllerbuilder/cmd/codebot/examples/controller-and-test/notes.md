
1. Create client.go

Will be:

```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > pkg/controller/direct/filestore/client.go
// +tool:controller-client
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
EOF
```


1. Create instance_controller.go

Will be:

```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > pkg/controller/direct/filestore/client.go
// +tool:controller
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
// crd.type: FilestoreInstance
// crd.version: v1alpha1
EOF
```

1. Update instance_identity.go

Will be:

```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > apis/filestore/v1alpha1/instance_identity.go
// +tool:krm-identity
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
// crd.type: FilestoreInstance
// crd.version: v1alpha1
EOF
```

(Maybe the URL path would be more useful than the full service, or maybe the CAIS path)

1. Iterate until `go build ./pkg/controller/direct/filestore/...` works

1. Create a minimal test

```
mkdir -p pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal

cat <<EOF > pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/create.yaml
apiVersion: filestore.cnrm.cloud.google.com/v1alpha1
kind: FilestoreInstance
metadata:
  name: filestoreinstance-minimal-\${uniqueId}
spec:
  projectRef:
    external: \${projectId}
  locationID: us-west2
  description: "Initial description"
EOF

cat <<EOF > pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/update.yaml
apiVersion: filestore.cnrm.cloud.google.com/v1alpha1
kind: FilestoreInstance
metadata:
  name: filestoreinstance-minimal-\${uniqueId}
spec:
  projectRef:
    external: \${projectId}
  locationID: us-west2
  description: "Updated description"
EOF
```

1. Update the MaybeSkip function in `config/tests/samples/create/harness.go` to add our type


```
case schema.GroupKind{Group: "filestore.cnrm.cloud.google.com", Kind: "FilestoreInstance"}:
```


1. Test the controller against mockgcp.  (We do this before testing against realgcp just to catch any obvious problems):

```
hack/compare-mock fixtures/filestoreinstance-minimal
```

The test will always fail the first time because it is writing the golden output.  It may also fail before that, of course!  Once you have reasonable golden output against the mocks, proceed to the next step and commit the work.  Later we'll verify against real gcp.


1. Commit to git

```
git add pkg/controller/direct/filestore/client.go
git add pkg/controller/direct/filestore/instance_controller.go
git commit -m "FilestoreInstance: create controller"

git add config/tests/samples/create/harness.go
git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/create.yaml
git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/update.yaml
git commit -m "FilestoreInstance: create minimal test"

git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/_generated*
git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/_http.log
git commit -m "FilestoreInstance: golden output for minimal test"

1. Verify that you can run against the mocks and get repeatable output:

```
hack/compare-mock fixtures/filestoreinstance-minimal
```


1. Now run against real gcp and capture the _real_ output:

```
hack/record-gcp fixtures/filestoreinstance-minimal
```

I like to stage (but not commit) the output, this makes the differences very visible during iteration:

```
hack/record-gcp fixtures/filestoreinstance-minimal
git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/
hack/compare-mock fixtures/filestoreinstance-minimal
git diff
# <make changes to mock or controller or normalization>
hack/compare-mock fixtures/filestoreinstance-minimal
git diff
# <repeat until no diffs or only acceptable diffs>
git add pkg/test/resourcefixture/testdata/basic/filestore/v1alpha1/filestoreinstance/filestoreinstance-minimal/
# Add any code changes also
git absorb
```

Sometimes hack/record-gcp will fail:

* Where the mock differs from GCP (for example the test case doesn't pass some required parameters), I like to update the mock to add the validation at the same time as updating the test case (then re-run `compare-mock`, and git add and git absorb)


TODO: Add export test

SCRATCH

gcloud_base: gcloud filestore instances
script: mockfilestore/testdata/instance/crud/script.yaml
http_host: file.googleapis.com
http.host: file.googleapis.com
proto.service: google.cloud.filestore.v1.CloudFilestoreManager
proto.resource: Instance
proto.message: google.cloud.filestore.v1.Instance
crd.type: FilestoreInstance
crd.version: v1alpha1
