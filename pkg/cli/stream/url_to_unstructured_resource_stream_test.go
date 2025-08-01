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
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"golang.org/x/oauth2"
)

const (
	spannerDatabaseURLGoldenFile = "testdata/expected-spanner-datbase-url-to-unstructured-resource-stream.golden.yaml"
	spannerDatabaseURL           = "https://spanner.googleapis.com/v1/projects/test-project/instances/spannerdatabase-dep/databases/asdf"
)

func TestURLToUnstructuredStream(t *testing.T) {
	unstructuredStream := newTestUnstructuredResourceStreamFromURL(t, spannerDatabaseURL)
	unstructs := unstructuredStreamToSlice(t, unstructuredStream)
	if *update {
		testyaml.WriteValueToFile(t, unstructs, spannerDatabaseURLGoldenFile)
	}
	testyaml.AssertFileContentsMatchValue(t, spannerDatabaseURLGoldenFile, unstructs)
}

func newTestUnstructuredResourceStreamFromURL(t *testing.T, url string) *stream.URLToUnstructuredResourceStream {
	ctx := context.TODO()

	mockClient := newMockGCPClient(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	smLoader := testservicemappingloader.New(t)
	httpClient := http.DefaultClient

	controllerConfig := &config.ControllerConfig{
		GCPTokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
	}

	if err := controllerConfig.Init(ctx); err != nil {
		t.Fatalf("error initializing controller config: %v", err)
	}

	return stream.NewUnstructuredResourceStreamFromURL(url, tfProvider, smLoader, mockClient, httpClient, controllerConfig)
}
