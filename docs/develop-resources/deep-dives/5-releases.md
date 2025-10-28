# 5. Release

## 5.1 Make Direct Controller the Default for Beta

When a resource is promoted to Beta, the `direct` controller should typically become the default. This is achieved by ensuring the `v1beta1` CRD does not contain the `cnrm.cloud.google.com/tf2crd` or `cnrm.cloud.google.com/dcl2crd` labels, and then regenerating the static controller configuration.

For detailed instructions, please see [4-add-controller.md](./4-add-controller.md), under "Scenario 3: Promoting an alpha direct resource to beta".

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

1.  Make sure a <kind> directory exist under [config/samples/resources](config/samples/resources).
    For example, if the kind is `spannerinstance`, the directory should be `config/samples/resources/spannerinstance`.
1.  Make sure a a sample YAML file named `<service>-<version>-<kind>.yaml` exist under the directory. 
1.  Look at example in pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<kind>-minimal/create.yaml to fill in the content for the sample YAML.
1.  Fulfill the following requirements for the sample YAML file (the path is in JSON format).
    - The `.metadata.name` is in the format `<kind>-sample` (lower case).
    - If the resource has a `.spec.projectRef` field, use `projects/${PROJECT_ID?}` literally as `.spec.projectRef.external` value. 
1.  Enable your service API in [SUPPORTED_SERVICES](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/scripts/shared-vars-public.sh#L32) if it's not there.
1.  If the test takes more than 10 minutes to finish, add it to LONG_RUNNING_SAMPLES_TESTS_REGEX in [./scripts/shared-vars-public.sh](./scripts/shared-vars-public.sh).

### Run/Disable Sample Tests

Run the sample

```bash
go test -v -tags=integration ./config/tests/samples/create -test.run TestAll -run-tests <your_sample_test_name>
```

`your_sample_test_name` is the sample suite directory name, i.e `spannerinstance`

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
    by adding a row for your resource to the resource table. Note that {{spanner_name}} is the template variable for Cloud Spanner. 
    The template variables for gcp product names can be found [here](https://source.corp.google.com/piper///depot/google3/third_party/devsite/cloud/en/_shared/_product_names.html).
    
1.  Update
    [scripts/generate-google3-docs/resource-reference/_toc.yaml](./../../../scripts/generate-google3-docs/resource-reference/_toc.yaml)
    by adding an entry and path to the corresponding API section. If there is no
    entry for the API, add one.

1.  Run `make resource-docs`. You should see your resource generated to the [scripts/generate-google3-docs/resource-reference/generated/resource-docs](./../../../scripts/generate-google3-docs/resource-reference/generated/resource-docs) folder.

1. Wait for the next Config Connector release, the release sheriff will publish the reference doc in Google based on the auto-generated doc folder. 

## 5.6 Add reference document 

Once you finish adding an Alpha resource, a Beta resource, or new fields to existing resources, you can add a [releasenote](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/docs/releasenotes) for the next release. The release sheriff will publish the note.
