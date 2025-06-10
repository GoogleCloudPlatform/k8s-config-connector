modify the fixture tests under pkg/test/resourcefixture/testdata/basic/${SERVICE}

1. Move pkg/test/resourcefixture/testdata/basic/${SERVICE}/ directory to pkg/test/resourcefixture/testdata/basic/${SERVICE}/
2. Change the `apiVersion` in create.yaml and update.yaml from `v1alpha1` to `v1beta1`
3. 