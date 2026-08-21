1. Make sure the directory apis/${SERVICE}/v1beta1 exist. If not, create one.
2. Copy the go files from apis/${SERVICE}/v1alpha1 to apis/${SERVICE}/v1beta1.
3. Make sure the go files under apis/${SERVICE}/v1beta1 have package name `v1beta1`. For example, you can run run `grep -rl "package v1alpha1" apis/<SERVICE>/v1beta1/ | xargs -i sed "s/package v1alpha1/package v1beta1/g"`
4. Iterate the new files and change the code `v1alpha1` to `v1beta1` if it represents `apis/${SERVICE}/v1alpha1`
3. Remove file apis/${SERVICE}/v1beta1/zz_generated.deepcopy.go
4. Add `// +kubebuilder:storageversion` as a comment to `type {KIND} struct` in file *_types.go 
5. Make sure all the go file under apis/${SERVICE}/v1beta1 have the right go package name.  
6. Finally, run `dev/tasks/generate-crds`. 