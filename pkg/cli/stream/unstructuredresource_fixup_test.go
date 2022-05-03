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
