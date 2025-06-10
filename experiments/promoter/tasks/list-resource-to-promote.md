SERVICE name is `apis/<SERVICE>`, KIND name is defined in `var APIGatewayAPIGVK = GroupVersion.WithKind` in each `apis/<SERVICE>/v1alpha1/*_types.go`. 

Iterate each sub-diretory under `apis`, to list the SERVICE/KIND pair, where KIND only shows in  `apis/<SERVICE>/v1alpha1/*_types.go`, but not  `apis/<SERVICE>/v1beta1/*_types.go`