.PHONY: update-schema
update-schema:
	curl https://raw.githubusercontent.com/kubernetes/kubernetes/master/staging/src/k8s.io/client-go/applyconfigurations/internal/internal.go | \
	awk '/typed.YAMLObject/,/`)/'  | sed 's/`)//g' | sed 's/var schemaYAML = typed.YAMLObject(`//g' > kubernetes_builtin_schema.yaml
