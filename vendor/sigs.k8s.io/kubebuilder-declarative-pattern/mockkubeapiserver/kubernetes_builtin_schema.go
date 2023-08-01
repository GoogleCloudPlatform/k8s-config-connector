package mockkubeapiserver

import (
	"fmt"
	"sync"

	"sigs.k8s.io/structured-merge-diff/v4/typed"

	_ "embed"
)

//go:embed kubernetes_builtin_schema.yaml
var kubernetesBuiltinSchemaYAML string

type Schema struct {
	Parser *typed.Parser
}

var cacheKubernetesBuiltinSchema *Schema

var mutexKubernetesBuiltinSchema sync.Mutex

func KubernetesBuiltInSchema() (*Schema, error) {
	mutexKubernetesBuiltinSchema.Lock()
	defer mutexKubernetesBuiltinSchema.Unlock()

	if cacheKubernetesBuiltinSchema != nil {
		return cacheKubernetesBuiltinSchema, nil
	}

	schemaYAML := kubernetesBuiltinSchemaYAML

	parser, err := typed.NewParser(typed.YAMLObject(schemaYAML))
	if err != nil {
		return nil, fmt.Errorf("error parsing schema: %w", err)
	}
	cacheKubernetesBuiltinSchema = &Schema{Parser: parser}
	return cacheKubernetesBuiltinSchema, nil
}
