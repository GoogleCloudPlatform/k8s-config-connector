1. Make sure the file `pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND in lower>/<KIND in lower>-full/create.yaml` exists. If not, create one
2. The created create.yaml is a Kubernetes YAML CustomResource of the CustomResourceDefinition (CRD) config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<KIND>s.<SERVICE>.cnrm.cloud.google.com.yaml. If you cannot find the CRD, it may because the KIND uses a different plural form. Please fill in the content of create.yaml. Remember that (described in JSON path), 
  - the `.metadata.name` should be <KIND>-${uniqueId} with all lower case.
  - the `.spec.projectRef` (If exist) should be `spec.projectRef.external: ${projectId}` 
  - try to understand the meaning of each field and give it a valid value, you can run `hack/record-gcp fixtures/<KIND in lower>-full to see if the value is correct. If not, try to fix it based on the result, and try 10 times or until the test passes. You can not only change this directory but the pkg/controller/direct/<service>/ directory as well. 
  - try to configure as many CR fields as possible. 

3. If the create.yaml succeeds, add a update.yaml file under the same directory. The content of the update.yaml should be the same as create.yaml, except the mutable field in `.spec` is assigned to another valid value. You should try to modify as many fields as possible. You can run `hack/record-gcp fixtures/<KIND in lower>-full` to see if the value is correct. If not, try to fix it based on the result. You should only change either the update.yaml (if it is a bad value) or the `Update` method in golang file pkg/controller/direct/<service>/*_controller.go.

<!--  
TODO: Two problems. 
    1.) If Create/Update fails, the fixture test cannot always give meaningful error or the LLM consider the timeout the root cause. We need to process the error to let LLM triage (the controller and mock). 
    2.) The fixture test log is very long. If use LLM to analyze it, it would exceed the token limit easily. 
4. If the update.yaml succeeds, we want to verify the mockgcp/mock<service>.
  - Store the real GCP GOOD log by running `git add pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND in lowercase>/<KIND in lowercase>-full/`. 
  - Compare with the Mock GCP log by running `hack/compare-mock fixtures/<KIND in lower>-full` to verify the mockgcp/mock<service>. If the test fail, you can see the diff between real GCP (hack/record-gcp) and mock GCP (hack/compare-mock) from `git diff pkg/test/resourcefixture/testdata/basic/<SERVICE>/v1beta1/<KIND in lowercase>/<KIND in lowercase>-full/_http.log`. Try to modify the functions in mockgcp/mock<service> to make minimum diff. For example, the mockgcp/mocksql/sqlinstance.go `Insert` `Update` `Delete` maps to the `a.sqlInstancesClient.Insert`, `sqlInstancesClient.Update` and `sqlInstancesClient.Delete` in pkg/controller/direct/sql/sqlinstance_controller.go -->