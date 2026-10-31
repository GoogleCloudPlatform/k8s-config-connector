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

package stream_test

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestFolderNameShouldGetFilled(t *testing.T) {
	// an exported folder does not have a 'name' so one should get filled in with the resource id value
	testFixupFile(t, "testdata/fixup/folder.yaml")
}

func TestBigQueryJobUppercaseNameShouldGetChanged(t *testing.T) {
	// the name the job has an uppercase character which should get replaced with a lowercase character
	// and an underscore which should get replaced by a hyphen
	testFixupFile(t, "testdata/fixup/bigqueryjob.yaml")
}

func TestFilterSensitiveFields(t *testing.T) {
	newInput := func() *unstructured.Unstructured {
		return &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
				"kind":       "SQLUser",
				"metadata": map[string]interface{}{
					"name": "test-sql-user",
				},
				"spec": map[string]interface{}{
					"host": "%",
					"password": map[string]interface{}{
						"value": "super-secret-password",
					},
					"nestedConfig": map[string]interface{}{
						"clientSecret": map[string]interface{}{
							"value": "oauth-secret-val",
						},
						"displayName": "public-name",
					},
					"flags": []interface{}{
						map[string]interface{}{
							"name":  "max_connections",
							"value": "100",
						},
					},
				},
			},
		}
	}

	// Default behavior: sensitive field values are retained
	defaultStream := stream.NewUnstructuredResourceFixupStream(newMockUnstructuredStream([]NextUnstructuredResult{
		{Unstructured: newInput()},
	}))
	defaultRes := unstructuredStreamToSlice(t, defaultStream)
	if len(defaultRes) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(defaultRes))
	}
	pw, found, err := unstructured.NestedString(defaultRes[0].Object, "spec", "password", "value")
	if err != nil || !found || pw != "super-secret-password" {
		t.Fatalf("expected default stream to retain spec.password.value, got %q (found=%v, err=%v)", pw, found, err)
	}

	// Filtered behavior: sensitive field values are omitted while keeping the field skeleton
	filteredStream := stream.NewUnstructuredResourceFixupStreamWithOptions(newMockUnstructuredStream([]NextUnstructuredResult{
		{Unstructured: newInput()},
	}), stream.FixupOptions{FilterSensitiveFields: true})
	filteredRes := unstructuredStreamToSlice(t, filteredStream)
	if len(filteredRes) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(filteredRes))
	}
	pwMap, found, err := unstructured.NestedMap(filteredRes[0].Object, "spec", "password")
	if err != nil || !found {
		t.Fatalf("expected spec.password skeleton map to remain, found=%v, err=%v", found, err)
	}
	if _, hasVal := pwMap["value"]; hasVal {
		t.Fatalf("expected spec.password.value to be omitted when FilterSensitiveFields is true, got %v", pwMap)
	}
	nestedSecretMap, found, err := unstructured.NestedMap(filteredRes[0].Object, "spec", "nestedConfig", "clientSecret")
	if err != nil || !found {
		t.Fatalf("expected spec.nestedConfig.clientSecret skeleton map to remain, found=%v, err=%v", found, err)
	}
	if _, hasVal := nestedSecretMap["value"]; hasVal {
		t.Fatalf("expected spec.nestedConfig.clientSecret.value to be omitted, got %v", nestedSecretMap)
	}
	// Non-sensitive slice item with {name, value} must remain intact
	flags, found, err := unstructured.NestedSlice(filteredRes[0].Object, "spec", "flags")
	if err != nil || !found || len(flags) != 1 {
		t.Fatalf("expected spec.flags to remain intact, found=%v, err=%v", found, err)
	}
	flagMap := flags[0].(map[string]interface{})
	if flagMap["value"] != "100" {
		t.Fatalf("expected non-sensitive flag value '100' to be preserved, got %v", flagMap["value"])
	}
}

func testFixupFile(t *testing.T, inputFile string) {
	goldenFile := strings.Replace(inputFile, ".yaml", ".golden.yaml", 1)
	unstructuredStream := newMockUnstructuredStreamFromFile(t, inputFile)
	fixupStream := stream.NewUnstructuredResourceFixupStream(unstructuredStream)
	unstructs := unstructuredStreamToSlice(t, fixupStream)
	if *update {
		testyaml.WriteValueToFile(t, unstructs, goldenFile)
	}
	testyaml.AssertFileContentsMatchValue(t, goldenFile, unstructs)
}

func newMockUnstructuredStreamFromFile(t *testing.T, filePath string) stream.UnstructuredStream {
	var unstruct *unstructured.Unstructured
	testyaml.UnmarshalFile(t, filePath, &unstruct)
	results := make([]NextUnstructuredResult, 1)
	results[0] = NextUnstructuredResult{
		Unstructured: unstruct,
	}
	return newMockUnstructuredStream(results)
}
