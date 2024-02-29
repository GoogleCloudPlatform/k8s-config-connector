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
	"errors"
	"io"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
)

const (
	assetFile                     = "testdata/assets.json"
	expectedFilteredAssetYAMLFile = "testdata/expected-filtered-asset-stream.golden.yaml"
)

func TestFilteredAssetStream(t *testing.T) {
	// The function below returns on each invocation in order:
	// 1. true
	// 2. false
	// 3. true
	// 4. false
	// ...
	// the purpose is to filter every other item in the stream, i.e. return values at indices 0, 2, 4, 6, etc from the stream.
	shouldInclude := false
	includeEvenIndexes := func(a *asset.Asset) bool {
		shouldInclude = !shouldInclude
		return shouldInclude
	}
	filteredStream := stream.NewFilteredAssetStream(newTestAssetStream(t), includeEvenIndexes)
	filteredAssets := assetStreamToSlice(t, filteredStream)
	if *update {
		testyaml.WriteValueToFile(t, filteredAssets, expectedFilteredAssetYAMLFile)
	}
	testyaml.AssertFileContentsMatchValue(t, expectedFilteredAssetYAMLFile, filteredAssets)
}

func assetStreamToSlice(t *testing.T, stream stream.AssetStream) []asset.Asset {
	results := make([]asset.Asset, 0)
	for asset, err := stream.Next(); !errors.Is(err, io.EOF); asset, err = stream.Next() {
		if err != nil {
			t.Fatalf("error reading asset: %v", err)
		}
		results = append(results, *asset)
	}
	return results
}

func newTestAssetStream(t *testing.T) *asset.Stream {
	t.Helper()
	assetStream, err := asset.NewStreamFromFile(assetFile)
	if err != nil {
		t.Fatalf("error creating asset stream: %v", err)
	}
	return assetStream
}
