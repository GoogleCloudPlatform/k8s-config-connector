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
	"context"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/google/go-cmp/cmp"
)

const hclStreamFile = "testdata/expected-hcl-stream.golden.yaml"

func TestHCLStream(t *testing.T) {
	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	hclStream := stream.NewHCLStream(newTestUnstructuredResourceStreamFromAsset(t, newTestAssetStream(t)), smLoader, tfProvider)
	bytes := hclStreamToBytes(t, hclStream)
	validateHCLBytesMatchesExpectedFile(t, bytes)
}

func hclStreamToBytes(t *testing.T, stream *stream.HCLStream) []byte {
	ctx := context.TODO()

	results := make([]byte, 0)
	for bytes, _, err := stream.Next(ctx); !errors.Is(err, io.EOF); bytes, _, err = stream.Next(ctx) {
		if err != nil {
			t.Fatalf("error reading next terraform file: %v", err)
		}
		results = append(results, bytes...)
	}
	return results
}

func validateHCLBytesMatchesExpectedFile(t *testing.T, bytes []byte) {
	if *update {
		if err := ioutil.WriteFile(hclStreamFile, bytes, 0644); err != nil {
			t.Fatalf("error writing file '%v': %v", hclStreamFile, err)
		}
	}
	expectedBytes, err := ioutil.ReadFile(hclStreamFile)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}
	actualValue := string(bytes)
	// Trim the license header before comparison.
	expectedValue := test.TrimLicenseHeaderFromYaml(string(expectedBytes))
	if expectedValue != actualValue {
		diff := cmp.Diff(expectedValue, actualValue)
		t.Fatalf("mismatch between actual type and expected for type '%v', diff:\n%v", reflect.TypeOf(actualValue).Name(), diff)
	}
}
