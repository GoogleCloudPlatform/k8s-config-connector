package apis

//go:generate go run k8s.io/code-generator/cmd/deepcopy-gen@v0.29.0 -O zz_generated.deepcopy -h ../hack/boilerplate.go.txt -i  github.com/GoogleCloudPlatform/k8s-config-connector/apis/... --trim-path-prefix github.com/GoogleCloudPlatform/k8s-config-connector/apis
