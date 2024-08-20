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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const assetToUnstructuredResourceStreamYAMLFile = "testdata/expected-asset-to-unstructured-resource-stream.golden.yaml"

func TestAssetToUnstructuredStream(t *testing.T) {
	unstructuredStream := newTestUnstructuredResourceStreamFromAsset(t, newTestAssetStream(t))
	unstructs := unstructuredStreamToSlice(t, unstructuredStream)
	if *update {
		testyaml.WriteValueToFile(t, unstructs, assetToUnstructuredResourceStreamYAMLFile)
	}
	testyaml.AssertFileContentsMatchValue(t, assetToUnstructuredResourceStreamYAMLFile, unstructs)
}

func newTestUnstructuredResourceStreamFromAsset(t *testing.T, assetStream stream.AssetStream) *stream.AssetToUnstructuredResourceStream {
	mockClient := newMockGCPClient(t)
	serviceClient := serviceclient.NewMockServiceClient(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	config := &config.ControllerConfig{}
	unstructuredStream, err := stream.NewUnstructuredResourceStreamFromAssetStream(assetStream, mockClient, tfProvider, &serviceClient, config)
	if err != nil {
		t.Fatalf("error creating unstructured stream: %v", err)
	}
	return unstructuredStream
}

func unstructuredStreamToSlice(t *testing.T, stream stream.UnstructuredStream) []*unstructured.Unstructured {
	ctx := context.TODO()

	results := make([]*unstructured.Unstructured, 0)
	for u, err := stream.Next(ctx); !errors.Is(err, io.EOF); u, err = stream.Next(ctx) {
		if err != nil {
			t.Fatalf("error reading asset: %v", err)
		}
		results = append(results, u)
	}
	return results
}

type mockGCPClient struct {
	t *testing.T
}

func newMockGCPClient(t *testing.T) gcpclient.Client {
	return &mockGCPClient{
		t: t,
	}
}

func (m *mockGCPClient) Get(_ context.Context, u *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	newUnstruct := unstructured.Unstructured{
		Object: deepcopy.DeepCopy(u.Object).(map[string]interface{}),
	}
	value := "this value verifies I was created by the mock client"
	if err := unstructured.SetNestedField(newUnstruct.Object, value, "spec", "testKey"); err != nil {
		return nil, err
	}

	return &newUnstruct, nil
}

func (m *mockGCPClient) Apply(_ *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	m.t.Fatalf("unimplemented")
	return nil, nil
}

func (m *mockGCPClient) Delete(_ *unstructured.Unstructured) error {
	m.t.Fatalf("unimplemented")
	return nil
}

func (m *mockGCPClient) IsSupported(_ string) bool {
	return true
}
