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

package snippetgeneration

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func TestSnippifyResourceConfig(t *testing.T) {
	tests := []struct {
		testName string
		config   string
		expected string
	}{
		{
			testName: "basicConfig",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  field: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  field: \${2:value}`,
		},
		{
			testName: "basicConfigWithLabel",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
  labels:
    labelA: valueA
spec:
  field: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  labels:
    \${1:labelA}: \${2:valueA}
  name: \${3:myresourcekind-name}
spec:
  field: \${4:value}`,
		},

		{
			testName: "basicConfigWithLabels",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
  labels:
    labelA: valueA
    labelB: valueB
spec:
  field: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  labels:
    \${1:labelA}: \${2:valueA}
    \${3:labelB}: \${4:valueB}
  name: \${5:myresourcekind-name}
spec:
  field: \${6:value}`,
		},
		{
			testName: "basicConfigWithComments",
			config: `# Top-level comment
apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind #In-line comment
metadata:
  name: myresourcekind-sample
spec:
  # Comment
  field: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  field: \${2:value}`,
		},
		{
			testName: "basicConfigWithBraces",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  field: ${uniqueId}`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  field: \${2:[uniqueId]}`,
		},
		{
			testName: "configWithVariousFieldTypes",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  numVal: 9000
  numWithScientificNotation: 1e+12
  floatVal: 3.14
  boolVal: true
  stringValWithQuotes: "Hello world"
  stringVal: Hello world`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  numVal: \${2:9000}
  numWithScientificNotation: \${3:1e+12}
  floatVal: \${4:3.14}
  boolVal: \${5:true}
  stringValWithQuotes: \${6:Hello world}
  stringVal: \${7:Hello world}`,
		},
		{
			testName: "configWithNestedMaps",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  fieldA:
    field: value
  fieldB:
    field:
      field: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  fieldA:
    field: \${2:value}
  fieldB:
    field:
      field: \${3:value}`,
		},
		{
			testName: "configWithSequences",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  field:
  - valueOne
  - valueTwo
  - valueThree`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  field:
  - \${2:valueOne}
  - \${3:valueTwo}
  - \${4:valueThree}`,
		},
		{
			testName: "configWithNestedMapsAndSequences",
			config: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: myresourcekind-sample
spec:
  field:
    fieldA:
    - fieldA: value
      fieldB:
        fieldA:
        - valueOne
        - valueTwo
        - valueThree
    - fieldA: value
      fieldB:
        fieldA:
        - value
    fieldB: value`,
			expected: `apiVersion: myservice.cnrm.cloud.google.com/v1alpha2
kind: MyResourceKind
metadata:
  name: \${1:myresourcekind-name}
spec:
  field:
    fieldA:
    - fieldA: \${2:value}
      fieldB:
        fieldA:
        - \${3:valueOne}
        - \${4:valueTwo}
        - \${5:valueThree}
    - fieldA: \${6:value}
      fieldB:
        fieldA:
        - \${7:value}
    fieldB: \${8:value}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			output, err := SnippifyResourceConfig([]byte(tc.config))
			if err != nil {
				t.Fatal(err)
			}
			if tc.expected != strings.TrimSuffix(output.InsertText, "\n") {
				t.Errorf("\n# Expected:\n%v\n---\n# Actual:\n%v\n", tc.expected, output.InsertText)
			}
		})
	}
}

func TestAllResourcesHaveSampleFileUsedForSnippets(t *testing.T) {
	samplesPath := repo.GetResourcesSamplesPath()
	resources, err := fileutil.SubdirsIn(samplesPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, resource := range resources {
		_, err := PathToSampleFileUsedForSnippets(resource)
		if err != nil {
			t.Fatalf("Failed to get sample file to use for snippet generation "+
				"for resource samples directory '%v': %v\n\n"+

				"Did you forget to update the 'preferredSampleForResource' map "+
				"in snippetgeneration.go ?", resource, err)
		}
	}
}
