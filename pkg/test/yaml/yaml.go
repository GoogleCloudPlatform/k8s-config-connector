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

package testyaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	cnrmyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"

	"github.com/ghodss/yaml" //nolint:depguard
	"github.com/google/go-cmp/cmp"
)

func SplitYAML(t *testing.T, yamlBytes []byte) [][]byte {
	t.Helper()
	result, err := cnrmyaml.SplitYAML(yamlBytes)
	if err != nil {
		t.Fatalf("error splitting YAML: %v", err)
	}
	return result
}

func WriteValueToFile(t *testing.T, value interface{}, filePath string) {
	t.Helper()
	bytes, err := yaml.Marshal(value)
	if err != nil {
		t.Fatalf("error marshalling value '%v' with kind '%v': %v", value, reflect.TypeOf(value).Name(), err)
	}
	WriteFile(t, bytes, filePath)
}

func WriteFile(t *testing.T, yamlBytes []byte, filePath string) {
	t.Helper()
	if err := ioutil.WriteFile(filePath, yamlBytes, 0644); err != nil {
		t.Fatalf("error writing file '%v': %v", filePath, err)
	}
}

func UnmarshalFile(t *testing.T, filePath string, value interface{}) {
	t.Helper()
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", filePath, err)
	}
	if err := yaml.Unmarshal(bytes, &value); err != nil {
		t.Fatalf("error unmarshalling bytes to '%v': %v\n\nstring value:\n\n%v", reflect.TypeOf(value).Name(), err, string(bytes))
	}
}

func AssertFileContentsMatchValue(t *testing.T, expectedFilePath string, actualValue interface{}) {
	t.Helper()
	newValue := reflect.New(reflect.TypeOf(actualValue)).Interface()
	UnmarshalFile(t, expectedFilePath, newValue)
	expectedValue := reflect.ValueOf(newValue).Elem().Interface()
	if !reflect.DeepEqual(expectedValue, actualValue) {
		diff := cmp.Diff(expectedValue, actualValue)
		t.Fatalf("mismatch between actual type and expected for type '%v', diff:\n%v", reflect.TypeOf(actualValue).Name(), diff)
	}
}
