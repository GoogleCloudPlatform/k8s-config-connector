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

package yamlresource

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml" //nolint:depguard
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// unstructuredFromYamlFile returns a unstructured.Unstructured
// struct from a yaml file.
func UnstructuredFromYamlFile(filePath string) (*unstructured.Unstructured, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var value map[string]interface{}
	if err = yaml.Unmarshal(bytes, &value); err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: value}, nil
}

func RenderJSON(res *unstructured.Unstructured, output io.Writer) error {
	bytes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling unstructured value to json: %w", err)
	}
	_, err = output.Write(bytes)
	return err
}

func RenderYAML(res *unstructured.Unstructured, output io.Writer) error {
	bytes, err := yaml.Marshal(res)
	if err != nil {
		return fmt.Errorf("error marshalling unstructured to yaml: %w", err)
	}
	_, err = output.Write(bytes)
	return err
}
