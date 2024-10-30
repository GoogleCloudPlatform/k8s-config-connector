# 5. Release

## 5.1 Turn on your Direct controller (TF/DCL Beta Only)

### For TF-based Beta resource

* Remove the `cnrm.cloud.google.com/tf2crd: "true"` label from the CRD will turn on SciFi controller. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/196a4b9a28b59b17936a443d5b36bb65f3c42fd9/apis/apikeys/v1alpha1/apikey_type.go#L44)

### For DCL-based Beta resource

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` label from the CRD will turn on SciFi controller.

## 5.2 Bump from v1alpha1 to v1beta1

* Copy-paste `./apis/service>/v1alpha1 `to` ./apis/service>/v1beta1 `using `git mv` (this is for PR review) and make sure the `./apis/service>/v1alpha1 `is still there.

* Update the new folder apiVersion to `v1beta1`

* Update all imports and update the MockGCP test suite to use `v1beta1`

* Add `kubebuilder:storageversion` tag to `v1beta1 `API. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/1b19153411653329177f4ba0991c982f36970707/apis/cloudbuild/v1beta1/workerpool_types.go#L155)

## 5.3 Turn on doc auto-generation (direct resource only)

* Add the direct resource to` config/servicemappings`, only specify the `name`, `kind` and <code>direct(true) </code>[example](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2182/files#diff-e463d47dab0190c35c12d64604451db84e0e7b6316ce33524a2a4eb29e0f2e47)

* For resource that is purely direct controller based, add the resource name to `pureDirectResources` in `IsPureDirectResource` function in [sets.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/test/resourcefixture/sets.go).

* If the resource is found in the resource-autogen [allowlist](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/6e9579347aadd08c07cfb0f1bd1747c4c9f4b197/scripts/resource-autogen/allowlist/allowlist.go#L31), remove it from the list.

## 5.4 Add samples 

1.  In [config/samples/resources](config/samples/resources), create a new
    directory for the resource's samples: `mkdir spannerinstance`.
1.  Follow the sample guidelines [here](../../../README.Samples.md) and create sample(s) for the resource.

### Run/Disable Sample Tests

Run the sample

```bash
go test -v -tags=integration ./config/tests/samples/create -test.run TestAll -run-tests <your_sample_test_name>
```

`your_sample_test_name` is the sample suite directory name, i.e `scheduler-job-pubsub`

Replace the environment variables to real values before running the tests.

### Cloud Code Snippets

We have a script that generates snippet files for Cloud Code using our samples.
If you created multiple samples (i.e. multiple samples subdirectories), you must
tell the script which sample to use for generating snippets. Update the
`preferredSampleForResource` map in
[pkg/snippet/snippetgeneration/snippetgeneration.go](../../../pkg/snippet/snippetgeneration/snippetgeneration.go).

## 5.5 Add reference document 

To add or update the Config Connector 
[reference doc](https://cloud.google.com/config-connector/docs/reference/overview), please do the following.

1.  Copy one of existing files under
    [scripts/generate-google3-docs/resource-reference/templates](../../../scripts/generate-google3-docs/resource-reference/templates),
    and name it like `spanner_spannerinstance.tmpl`, (i.e.,
    <service>_<kind>.tmpl).
1.  Add a paragraph to briefly introduce the resource at the very top of the file.
1.  If additional setup is necessary before the newly added resource can be applied successfully, 
    and the setup can't be done via the Config Connector resources or features, add a "Prerequisites" section to explain
    the prerequisites steps,
    E.g. [scripts/generate-google3-docs/resource-reference/generated/resource-docs/containerattached/containerattachedcluster.md](../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs/containerattached/containerattachedcluster.md)
    and
    [scripts/generate-google3-docs/resource-reference/generated/resource-docs/secretmanager/secretmanagersecret.md](../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs/secretmanager/secretmanagersecret.md)
1.  Update
    [scripts/generate-google3-docs/resource-reference/overview.md](../../../scripts/generate-google3-docs/resource-reference/overview.md)
    by adding a row for your resource to the resource table.
    
1.  Update
    [scripts/generate-google3-docs/resource-reference/_toc.yaml](../../../scripts/generate-google3-docs/resource-reference/_toc.yaml)
    by adding an entry and path to the corresponding API section. If there is no
    entry for the API, add one.

To generate the new resource doc, run `make resource-docs`. You should see your
resource generated to the
[scripts/generate-google3-docs/resource-reference/generated/resource-docs](../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs)
folder.
