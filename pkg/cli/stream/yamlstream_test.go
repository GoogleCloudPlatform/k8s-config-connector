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
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const yamlStreamFile = "testdata/expected-yaml-stream.golden.yaml"

func TestYAMLStream(t *testing.T) {
	yamlStream := stream.NewYAMLStream(newTestUnstructuredResourceStreamFromAsset(t, newTestAssetStream(t)))
	bytes := yamlStreamToBytes(t, yamlStream)
	validateYAMLBytesMatchesExpectedFile(t, bytes)
}

func TestEmptyStream(t *testing.T) {
	testFile, err := ioutil.TempFile("", "empty-stream")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	assetStream, err := asset.NewStreamFromFile(testFile.Name())
	if err != nil {
		t.Fatalf("error creating asset stream from file '%v': %v", testFile.Name(), err)
	}
	yamlStream := stream.NewYAMLStream(newTestUnstructuredResourceStreamFromAsset(t, assetStream))
	bytes := yamlStreamToBytes(t, yamlStream)
	expectedLength := 0
	if len(bytes) > expectedLength {
		t.Fatalf("unexpected stream length: got '%v', want '%v'", len(bytes), expectedLength)
	}
}

func TestErrorInAssetStream(t *testing.T) {
	errorAsset := asset.Asset{
		Name:      "//compute.googleapis.com/projects/cli-test-project/global/backendServices/computebackendservice-8034p172pirpg3bg31te",
		AssetType: "compute.googleapis.com/InvalidResourceType",
		Ancestors: nil,
	}
	realAssetStream := newTestAssetStream(t)
	allAssets := assetStreamToSlice(t, realAssetStream)
	if len(allAssets) < 2 {
		t.Fatalf("test requires input asset stream has at least two assets to be valid")
	}
	// create a stream of assets where every odd indexed asset results in an error
	nextResults := make([]NextAssetResult, 0, 2*len(allAssets))
	for i := 0; i < len(allAssets); i++ {
		successResult := NextAssetResult{Asset: &allAssets[i]}
		nextResults = append(nextResults, successResult)
		errorResult := NextAssetResult{Asset: &errorAsset}
		nextResults = append(nextResults, errorResult)
	}
	assetStream := newMockAssetStream(nextResults)
	yamlStream := stream.NewYAMLStream(newTestUnstructuredResourceStreamFromAsset(t, assetStream))
	bytes := yamlStreamToBytesIgnoreErrors(yamlStream)
	validateYAMLBytesMatchesExpectedFile(t, bytes)
	// create a stream of assets where every even indexed asset results in an error
	nextResults = make([]NextAssetResult, 0, 2*len(allAssets))
	for i := 0; i < len(allAssets); i++ {
		errorResult := NextAssetResult{Asset: &errorAsset}
		nextResults = append(nextResults, errorResult)
		successResult := NextAssetResult{Asset: &allAssets[i]}
		nextResults = append(nextResults, successResult)
	}
	assetStream = newMockAssetStream(nextResults)
	yamlStream = stream.NewYAMLStream(newTestUnstructuredResourceStreamFromAsset(t, assetStream))
	bytes = yamlStreamToBytesIgnoreErrors(yamlStream)
	validateYAMLBytesMatchesExpectedFile(t, bytes)
}

func TestStatusIsRemoved(t *testing.T) {
	ctx := context.TODO()

	results := []NextUnstructuredResult{
		{
			Unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec":   "specValue",
					"status": "statusValue",
				},
			},
			Err: nil,
		},
	}
	yamlStream := stream.NewYAMLStream(newMockUnstructuredStream(results))
	bytes, _, err := yamlStream.Next(ctx)
	if err != nil {
		t.Fatalf("unexpected error: got '%v', want 'nil'", err)
	}
	yaml := string(bytes)
	if strings.Contains(yaml, "status") {
		t.Fatalf("yaml contains status field when it should have been removed")
	}
}

type NextUnstructuredResult struct {
	Unstructured *unstructured.Unstructured
	Err          error
}

type MockUnstructuredStream struct {
	idx     int
	results []NextUnstructuredResult
}

func newMockUnstructuredStream(results []NextUnstructuredResult) stream.UnstructuredStream {
	return &MockUnstructuredStream{
		results: results,
	}
}

func (m *MockUnstructuredStream) Next(_ context.Context) (*unstructured.Unstructured, error) {
	if m.idx == len(m.results) {
		return nil, io.EOF
	}
	result := m.results[m.idx]
	m.idx++
	return result.Unstructured, result.Err
}

type NextAssetResult struct {
	Asset *asset.Asset
	Err   error
}

type MockAssetStream struct {
	idx     int
	results []NextAssetResult
}

func newMockAssetStream(results []NextAssetResult) stream.AssetStream {
	return &MockAssetStream{
		results: results,
	}
}

func (m *MockAssetStream) Next() (*asset.Asset, error) {
	if m.idx == len(m.results) {
		return nil, io.EOF
	}
	result := m.results[m.idx]
	m.idx++
	return result.Asset, result.Err
}

func (m *MockAssetStream) Close() error {
	return nil
}

func yamlStreamToBytes(t *testing.T, stream *stream.YAMLStream) []byte {
	ctx := context.TODO()

	results := make([]byte, 0)
	for bytes, _, err := stream.Next(ctx); !errors.Is(err, io.EOF); bytes, _, err = stream.Next(ctx) {
		if err != nil {
			t.Fatalf("error reading next yaml: %v", err)
		}
		results = append(results, bytes...)
	}
	return results
}

func yamlStreamToBytesIgnoreErrors(stream *stream.YAMLStream) []byte {
	ctx := context.TODO()

	results := make([]byte, 0)
	for bytes, _, err := stream.Next(ctx); !errors.Is(err, io.EOF); bytes, _, err = stream.Next(ctx) {
		if err != nil {
			continue
		}
		results = append(results, bytes...)
	}
	return results
}

func validateYAMLBytesMatchesExpectedFile(t *testing.T, bytes []byte) {
	if *update {
		testyaml.WriteFile(t, bytes, yamlStreamFile)
	}
	expectedBytes, err := ioutil.ReadFile(yamlStreamFile)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}
	// Trim the license header before comparison.
	expectedValue := test.TrimLicenseHeaderFromYaml(string(expectedBytes))
	actualValue := string(bytes)
	if expectedValue != actualValue {
		diff := cmp.Diff(expectedValue, actualValue)
		t.Fatalf("mismatch between actual type and expected for type '%v', diff:\n%v", reflect.TypeOf(actualValue).Name(), diff)
	}
}
