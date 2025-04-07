# 5. Release

## 5.1 Turn on your Direct controller (TF/DCL Beta Only)

### For TF-based Beta resource

* Remove the `cnrm.cloud.google.com/tf2crd: "true"` label from the CRD will turn on SciFi controller. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/196a4b9a28b59b17936a443d5b36bb65f3c42fd9/apis/apikeys/v1alpha1/apikey_type.go#L44)

### For DCL-based Beta resource

* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` label from the CRD will turn on SciFi controller.

## 5.2 Bump from v1alpha1 to v1beta1

* Copy-paste `./apis/service>/v1alpha1 `to` ./apis/service>/v1beta1 `using `git mv` (this is for PR review) and make sure the `./apis/service>/v1alpha1 `is still there.
  * NOTE: if you have more than one resource in your service, only copy the files relevant to your resource for now. You will need to re-run `dev/tasks/generate-crds` for the `deepcopy` file. 

* Update the new package version to `v1beta1`

* Update all direct controller imports to use  the `v1beta1` api of your resource

* For all the fixtures tests under [pkg/test/resourcefixture/testdata](pkg/test/resourcefixture/testdata) make sure your resources use the `v1beta1` version.

* Add `kubebuilder:storageversion` tag to `v1beta1 `API. [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/1b19153411653329177f4ba0991c982f36970707/apis/cloudbuild/v1beta1/workerpool_types.go#L155)

* Run `dev/tasks/generate-crds` to patch your resource's CRD with the `v1beta1` version.

## 5.3 Turn on doc auto-generation (direct resource only)

* For resource that is purely direct controller based, add the resource name to `pureDirectResources` in `IsPureDirectResource` function in [sets.go](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/test/resourcefixture/sets.go).

* If the resource is found in the resource-autogen [allowlist](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/6e9579347aadd08c07cfb0f1bd1747c4c9f4b197/scripts/resource-autogen/allowlist/allowlist.go#L31), remove it from the list.

## 5.4 Add samples 

1.  In [config/samples/resources](config/samples/resources), create a new
    directory for the resource's samples: `mkdir spannerinstance`.
1.  Follow the [sample guidelines](./../../../README.Samples.md) and create sample(s) for the resource.

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
[pkg/snippet/snippetgeneration/snippetgeneration.go](./../../../pkg/snippet/snippetgeneration/snippetgeneration.go).

## 5.5 Add reference document 

All beta resource can have a Google Reference documentation Connector 
[reference doc](https://cloud.google.com/config-connector/docs/reference/overview). Please follow the steps to add the reference doc.

1.  Copy one of existing files under
    [scripts/generate-google3-docs/resource-reference/templates](./../../../scripts/generate-google3-docs/resource-reference/templates),
    and name it like `spanner_spannerinstance.tmpl`, (i.e.,
    <service>_<kind>.tmpl).
1.  Add a paragraph to briefly introduce the resource at the very top of the file.
1.  If additional setup is necessary before the newly added resource can be applied successfully, 
    and the setup can't be done via the Config Connector resources or features, add a "Prerequisites" section to explain
    the prerequisites steps,
    E.g. [scripts/generate-google3-docs/resource-reference/generated/resource-docs/containerattached/containerattachedcluster.md](./../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs/containerattached/containerattachedcluster.md)
    and
    [scripts/generate-google3-docs/resource-reference/generated/resource-docs/secretmanager/secretmanagersecret.md](./../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs/secretmanager/secretmanagersecret.md)
1.  Update
    [scripts/generate-google3-docs/resource-reference/overview.md](./../../../scripts/generate-google3-docs/resource-reference/overview.md)
    by adding a row for your resource to the resource table.
    
1.  Update
    [scripts/generate-google3-docs/resource-reference/_toc.yaml](./../../../scripts/generate-google3-docs/resource-reference/_toc.yaml)
    by adding an entry and path to the corresponding API section. If there is no
    entry for the API, add one.

1.  Run `make resource-docs`. You should see your resource generated to the [scripts/generate-google3-docs/resource-reference/generated/resource-docs](./../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs) folder.

1. Wait for the next Config Connector release, the release sheriff will publish the reference doc in Google based on the auto-generated doc folder. 

## 5.6 Add reference document 

Once you finish adding an Alpha resource, a Beta resource, or new fields to existing resources, you can add a [releasenote](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/releasenotes) for the next release. The release sheriff will publish the note.
