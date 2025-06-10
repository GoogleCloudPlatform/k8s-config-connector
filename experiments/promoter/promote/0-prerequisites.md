SERVICE is ${service}
KIND names are defined in `var {KIND}GVK = GroupVersion.WithKind("{KIND}")` in each `apis/<SERVICE>/v1alpha1/*_types.go`.

Controller can be found under `pkg/controller/direct/<SERVICE>/`, you can run `ls pkg/controller/direct/<SERVICE>/*_controller.go` to get all the controllers

Mappers can be found under `pkg/controller/direct/<SERVICE>/`, you can run `ls pkg/controller/direct/<SERVICE>/*map*.go` to get the mappers.
File `mapper.generated.go` can only have code removed, but not changed: if you want to change a function in `mapper.generated.go`, copy the function to another  `ls pkg/controller/direct/<SERVICE>/*_map*.go` file, change it there, and remove the orignal function from `mapper.generated.go`.

Fuzz test can be found under `pkg/controller/direct/<SERVICE>/`, you can run `ls pkg/controller/direct/<SERVICE>/*_fuzzer.go` to get the fuzz.

Fixture tests can be found under `pkg/test/resourcefixture/testdata/basic/<SERVICE>/`. 

Service and Kind name should be lower case when used in file path.
