package mockkubeapiserver

import (
	"fmt"
	"sync"

	"sigs.k8s.io/structured-merge-diff/v4/typed"
	"sigs.k8s.io/yaml"

	_ "embed"
)

//go:embed kubernetes_builtin_schema.yaml
var kubernetesBuiltinSchemaYAML string

//go:embed kubernetes_builtin_schema.meta.yaml
var kubernetesBuiltinSchemaMetaYAML string

const BuiltinKey = "kubernetes_builtin_schema"

type Schema struct {
	Parser *typed.Parser
	Meta   *SchemaMeta
}

type SchemaMeta struct {
	Resources []SchemaMetaResource `json:"resources"`
}

type SchemaMetaResource struct {
	Key      string `json:"key"`
	Group    string `json:"group"`
	Version  string `json:"version"`
	Kind     string `json:"kind"`
	Resource string `json:"resource"`
	Scope    string `json:"scope"`
}

type schemaCache struct {
	mutex   sync.Mutex
	schemas map[string]*Schema
}

var globalSchemaCache schemaCache

func KubernetesBuiltInSchema() (*Schema, error) {
	return globalSchemaCache.Get(BuiltinKey)
}

func (c *schemaCache) Get(key string) (*Schema, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if schema := c.schemas[key]; schema != nil {
		return schema, nil
	}

	if key != BuiltinKey {
		return nil, fmt.Errorf("schema %q not known", key)
	}

	schemaYAML := kubernetesBuiltinSchemaYAML
	parser, err := typed.NewParser(typed.YAMLObject(schemaYAML))
	if err != nil {
		return nil, fmt.Errorf("error parsing schema: %w", err)
	}

	metaYAML := kubernetesBuiltinSchemaMetaYAML
	meta := &SchemaMeta{}
	if err := yaml.Unmarshal([]byte(metaYAML), meta); err != nil {
		return nil, fmt.Errorf("error parsing schema metadata: %w", err)
	}

	schema := &Schema{
		Parser: parser,
		Meta:   meta,
	}
	if c.schemas == nil {
		c.schemas = make(map[string]*Schema)
	}
	c.schemas[key] = schema
	return schema, nil
}
