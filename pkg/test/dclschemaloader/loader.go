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

package testdclschemaloader

import (
	"fmt"
	"strings"
	"testing"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/nasa9084/go-openapi"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

type testSchemaLoader struct {
	schemas map[string]*openapi.Schema
}

func New(schemas map[string]*openapi.Schema) dclschemaloader.DCLSchemaLoader {
	return &testSchemaLoader{
		schemas: schemas,
	}
}

func (l *testSchemaLoader) GetDCLSchema(stv dclunstruct.ServiceTypeVersion) (*openapi.Schema, error) {
	key := fmt.Sprintf("%s_%s_%s", strings.ToLower(stv.Service), strings.ToLower(stv.Version), strings.ToLower(stv.Type))
	s, ok := l.schemas[key]
	if !ok {
		return nil, fmt.Errorf("couldn't find the dcl OpenAPI schema for %v", stv)
	}
	return s, nil
}

func DCLSchemaKeyForGVK(t *testing.T, gvk k8sschema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader) string {
	key, err := dclschemaloader.DCLSchemaKeyForGVK(gvk, smLoader)
	if err != nil {
		t.Fatalf("error resolving DCL schema key for GroupVersionKind %v: %v", gvk, err)
	}
	return key
}
