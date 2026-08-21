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

package printer_test

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/printer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/resourcedescription"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"github.com/google/go-cmp/cmp"
)

// set this flag to true to update the expected test output
var update = flag.Bool("update", false, "update .golden files")

func TestPrintYAML(t *testing.T) {
	testPrinter(t, printer.PrintYAML, "testdata/print-yaml.golden.yaml")
}

func TestPrintJSON(t *testing.T) {
	testPrinter(t, printer.PrintJSON, "testdata/print-json.golden.json")
}

func TestPrintTable(t *testing.T) {
	testPrinter(t, printer.PrintTable, "testdata/print-table.golden.txt")
}

type PrintFunc func([]resourcedescription.ResourceDescription, io.Writer) error

func testPrinter(t *testing.T, printFunc PrintFunc, expectedOutputFile string) {
	resourceDescs := newResourceDescriptionsFixture(t)
	buf := bytes.Buffer{}
	if err := printFunc(resourceDescs, &buf); err != nil {
		t.Fatalf("unexpected error printing: %v", err)
	}
	bytes := buf.Bytes()
	if *update {
		if err := ioutil.WriteFile(expectedOutputFile, bytes, 0644); err != nil {
			t.Fatalf("error writing file '%v': %v", expectedOutputFile, err)
		}
	}
	expectedBytes, err := ioutil.ReadFile(expectedOutputFile)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}
	actualValue := string(bytes)
	// Trim the license header before comparison.
	expectedValue := test.TrimLicenseHeaderFromYaml(string(expectedBytes))
	if expectedValue != actualValue {
		diff := cmp.Diff(expectedValue, actualValue)
		t.Fatalf("mismatch between actual output and expected for output '%v', diff:\n%v", reflect.TypeOf(actualValue).Name(), diff)
	}
}

func newResourceDescriptionsFixture(t *testing.T) []resourcedescription.ResourceDescription {
	var resourceDescs []resourcedescription.ResourceDescription
	testyaml.UnmarshalFile(t, "testdata/resource-descriptions.yaml", &resourceDescs)
	return resourceDescs
}
