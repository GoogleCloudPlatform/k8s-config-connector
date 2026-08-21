// Copyright 2026 Google LLC
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

package objects

import (
	"bytes"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	sigsyaml "sigs.k8s.io/yaml"
)

// ParseObjectsFromStream parses KRM unstructured objects from a stream of YAML documents.
func ParseObjectsFromStream(r io.Reader) ([]*unstructured.Unstructured, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	docs, err := yaml.SplitYAML(data)
	if err != nil {
		return nil, err
	}
	var kccObjects []*unstructured.Unstructured
	for _, doc := range docs {
		if len(bytes.TrimSpace(doc)) == 0 {
			continue
		}
		var value map[string]interface{}
		if err := sigsyaml.Unmarshal(doc, &value); err != nil {
			return nil, fmt.Errorf("failed to parse YAML document: %w", err)
		}
		if len(value) == 0 {
			continue
		}
		if _, ok := value["apiVersion"]; !ok {
			continue
		}
		if _, ok := value["kind"]; !ok {
			continue
		}
		kccObjects = append(kccObjects, &unstructured.Unstructured{Object: value})
	}
	return kccObjects, nil
}
