// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dclschemaloader

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/embed"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/nasa9084/go-openapi"
	"gopkg.in/yaml.v2"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"

	// import dcl schema files
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apigee/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/billingbudgets/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/binaryauthorization"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudfunctions/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudidentity/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudkms"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudresourcemanager"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudscheduler"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/compute/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/configcontroller/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeranalysis"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/datafusion/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/dataproc"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/dlp/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/eventarc/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/filestore/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gkehub/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iap"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/identitytoolkit"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/logging"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/monitoring"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkconnectivity"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networksecurity/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkservices"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkservices/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkservices/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/osconfig/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/privateca"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/pubsub"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/recaptchaenterprise"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/run/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/storage"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/vpcaccess/beta"
)

// DCLSchemaLoader is the DCL schema loader interface.
type DCLSchemaLoader interface {
	GetDCLSchema(stv dclunstruct.ServiceTypeVersion) (*openapi.Schema, error)
}

type EmbedDCLSchemaLoader struct {
	dclSchemaMap map[string]*openapi.Schema
}

func New() (*EmbedDCLSchemaLoader, error) {
	loader := EmbedDCLSchemaLoader{}
	schemaMap, err := constructSchemaMap()
	if err != nil {
		return nil, err
	}
	loader.dclSchemaMap = schemaMap
	return &loader, nil
}

func constructSchemaMap() (map[string]*openapi.Schema, error) {
	schemaMap := make(map[string]*openapi.Schema)
	baseDir := "/"
	var queue []string
	queue = append(queue, baseDir)
	for len(queue) != 0 {
		dirPath := queue[0]
		queue = queue[1:]
		curDir, err := embed.Assets.Open(dirPath)
		if err != nil {
			return nil, fmt.Errorf("error opening directory %v: %w", dirPath, err)
		}
		fileInfos, err := curDir.Readdir(0)
		if err != nil {
			return nil, fmt.Errorf("error reading files in directory %v: %w", curDir, err)
		}
		for _, f := range fileInfos {
			if f.IsDir() {
				queue = append(queue, path.Join(dirPath, f.Name()))
				continue
			}
			if strings.HasSuffix(f.Name(), ".yaml") {
				filePath := path.Join(dirPath, f.Name())
				schema, err := parseOpenAPISchema(filePath)
				if err != nil {
					return nil, err
				}
				key, err := getKey(filePath)
				if err != nil {
					return nil, err
				}
				schemaMap[key] = schema
			}
		}
		curDir.Close()
	}
	patchDCLSchemas(schemaMap)
	return schemaMap, nil
}

func getKey(filePath string) (string, error) {
	p := strings.TrimPrefix(filePath, "/")
	components := strings.Split(strings.ReplaceAll(p, ".yaml", ""), "/")
	if len(components) == 2 {
		kind := strings.ReplaceAll(components[1], "_", "")
		return fmt.Sprintf("%s_ga_%s", components[0], kind), nil
	}
	if len(components) == 3 {
		kind := strings.ReplaceAll(components[2], "_", "")
		return fmt.Sprintf("%s_%s_%s", components[0], components[1], kind), nil
	}
	return "", fmt.Errorf("path to the dcl schema yaml file has invalid format: %v", filePath)
}

func parseOpenAPISchema(filePath string) (*openapi.Schema, error) {
	document := &openapi.Document{}
	file, err := embed.Assets.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if err := yaml.Unmarshal(b, document); err != nil {
		return nil, err
	}
	var schema *openapi.Schema
	for k, v := range document.Components.Schemas {
		components := strings.Split(document.Info.Title, "/")
		if len(components) != 2 {
			return nil, fmt.Errorf("invalid format for title %v", document.Info.Title)
		}
		if k == components[1] {
			schema = v
			schema.Title = k
		}
	}
	schema, err = CheckAndResolveRefs(schema, document)
	if err != nil {
		return nil, err
	}
	if err := schema.Validate(); err != nil {
		return nil, err
	}
	return schema, nil
}

func CheckAndResolveRefs(schema *openapi.Schema, doc *openapi.Document) (*openapi.Schema, error) {
	if schema.Ref != "" {
		newSchema, err := ResolveSchemaForRef(schema, doc)
		if err != nil {
			return nil, err
		}
		return newSchema, nil
	}
	if schema.Type == "object" {
		// The field AdditionalProperties is mutually exclusive with Properties
		// in Schema.
		if schema.AdditionalProperties != nil {
			newVal, err := CheckAndResolveRefs(schema.AdditionalProperties, doc)
			if err != nil {
				return nil, err
			}
			schema.AdditionalProperties = newVal
			return schema, nil
		}
		for k, v := range schema.Properties {
			newVal, err := CheckAndResolveRefs(v, doc)
			if err != nil {
				return nil, err
			}
			schema.Properties[k] = newVal
		}
	}
	if schema.Type == "array" {
		newItems, err := CheckAndResolveRefs(schema.Items, doc)
		if err != nil {
			return nil, err
		}
		schema.Items = newItems
	}
	return schema, nil
}

func ResolveSchemaForRef(schema *openapi.Schema, doc *openapi.Document) (*openapi.Schema, error) {
	var ret *openapi.Schema
	ret, err := openapi.ResolveSchema(doc, schema.Ref)
	if err != nil {
		return nil, fmt.Errorf("error resolving reference schema: %w", err)
	}
	// Check for resource refs within the referenced schema
	if ret, err = CheckAndResolveRefs(ret, doc); err != nil {
		return nil, err
	}
	return ret, nil
}

func (l *EmbedDCLSchemaLoader) GetDCLSchema(stv dclunstruct.ServiceTypeVersion) (*openapi.Schema, error) {
	key := dclSchemaKeyForSTV(stv)
	s, ok := l.dclSchemaMap[key]
	if !ok {
		return nil, fmt.Errorf("couldn't find the dcl OpenAPI schema for %v", stv)
	}
	return s, nil
}

func GetDCLSchemaForGVK(gvk k8sschema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader DCLSchemaLoader) (*openapi.Schema, error) {
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error resolving the DCL ServiceTypeVersion for GroupVersionKind %v: %w", gvk, err)
	}
	dclSchema, err := schemaLoader.GetDCLSchema(stv)
	if err != nil {
		return nil, err
	}
	return dclSchema, nil
}

func dclSchemaKeyForSTV(stv dclunstruct.ServiceTypeVersion) string {
	return fmt.Sprintf("%s_%s_%s", strings.ToLower(stv.Service), strings.ToLower(stv.Version), strings.ToLower(stv.Type))
}

func DCLSchemaKeyForGVK(gvk k8sschema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader) (string, error) {
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, smLoader)
	if err != nil {
		return "", fmt.Errorf("error resolving the DCL ServiceTypeVersion for GroupVersionKind %v: %w", gvk, err)
	}
	return dclSchemaKeyForSTV(stv), nil
}
