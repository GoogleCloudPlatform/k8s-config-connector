# 3. Add the KRM and API mapper

## 3.1 Generate the API and proto mapper (repeat-run safe) 

Run the following command

```
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd $REPO_ROOT/dev/tools/controllerbuilder

go run . generate-mapper \
   --proto-source-path ../proto-to-mapper/build/googleapis.pb \
   --service google.storage.v1  \
   --api-version "storage.cnrm.cloud.google.com/v1beta1" \
   --api-go-package-path  $REPO_ROOT/apis \
   --output-dir $REPO_ROOT/pkg/controller/direct/ \
   --api-dir $REPO_ROOT/apis 
```

**Note**: We suggest using the same proto for your mock GCP and for you type-generation tool to generate the Config Connector API and mapper to avoid mismatch in schema definitions. There are some exceptions where you need to [replace the proto go package](https://github.com/xiaoweim/k8s-config-connector/blob/master/dev/tools/controllerbuilder/pkg/codegen/mappergenerator.go#L132).

### Simple path

If no comments marked `MISSING` and all the mapper functions look good, you are done. You can move to “Add the controller” (step 4).

### Custom path

If you need manual changes to some methods, remove the method or/and the `MISSING` comment from the generated code and move to the next step (step 3.2).   


## 3.2 Add manual mapper code

Create a` <resource>_mappings.go` under` pkg/controller/direct/<service> `only if the auto-generated code cannot fulfill the mappings.

If you have a lot of fields that need to be manually written, you can split the work into several PRs, just keep the `MISSING `comments for reference. Also, comment out the code you plan to update with `/*NOTYET*/ `comments


### To change a method

* Move the code that needs manual change from `<resource>_generated.mappings.go `to `<resource>_mappings.go` 

### To add a missing method

* Follow the naming convention as the auto-generated code. You shall have two methods `<field>_ToProto` and` <field>_FromProto`
* Remove the `MISSING` comment


## Exit Criteria

* No `MISSING` comments left in the code
* No `/*NOTYET*/` comments left in the code.
* Each mapper method shall reflect in the `_http.log` as the value from `create.yaml` and `update.yaml` recorded in the `_http.log` POST and PUT/PATCH method.