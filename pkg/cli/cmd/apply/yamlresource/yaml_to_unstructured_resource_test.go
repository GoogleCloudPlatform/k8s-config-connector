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
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	PathPrefix         = "testdata"
	InvalidYamlFile    = "invalid_resource.yaml"
	GoldenPathYamlFile = "golden_path_resource.yaml"
	JSONOutputFile     = "render_output.json"
)

func TestReadEmptyFile(t *testing.T) {
	_, err := UnstructuredFromYamlFile("")
	if err == nil {
		t.Fatalf("Should have Failed for empty file path.")
	}
}

func TestNotYamlFile(t *testing.T) {
	testPath := fmt.Sprintf("%v/%v", PathPrefix, InvalidYamlFile)
	_, err := UnstructuredFromYamlFile(testPath)
	if err == nil {
		t.Fatalf("Should have Failed for file %v", testPath)
	}
}

func TestReadSuccess(t *testing.T) {
	testPath := fmt.Sprintf("%v/%v", PathPrefix, GoldenPathYamlFile)
	unstructured, err := UnstructuredFromYamlFile(testPath)
	if err != nil {
		t.Fatalf("error reading file %v: %v", testPath, err)
	}

	if unstructured.GetKind() != "ComputeInstance" {
		t.Fatalf("Error parsing resource file %v: %+v", testPath, unstructured)
	}

}

func TestRenderJSONSuccess(t *testing.T) {
	testPath := fmt.Sprintf("%v/%v", PathPrefix, GoldenPathYamlFile)
	unstructured, err := UnstructuredFromYamlFile(testPath)
	if err != nil {
		t.Fatalf("Error loading test_file %v, %v", testPath, err)
	}
	buf := bytes.Buffer{}
	if err := RenderJSON(unstructured, &buf); err != nil {
		t.Fatal(err)
	}
	renderOutput, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", PathPrefix, JSONOutputFile))
	if err != nil {
		t.Fatalf("Error loading test_file %v, %v", JSONOutputFile, err)
	}
	expectedValue := string(renderOutput)
	actualValue := buf.String()
	if expectedValue != actualValue {
		diff := cmp.Diff(expectedValue, actualValue)
		t.Fatalf("mismatch between actual output and expected for output diff:\n%v", diff)
	}
}
