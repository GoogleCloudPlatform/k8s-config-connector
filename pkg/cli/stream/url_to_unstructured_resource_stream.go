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

package stream

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type URLToUnstructuredResourceStream struct {
	url        string
	done       bool
	gcpClient  gcpclient.Client
	smLoader   *servicemappingloader.ServiceMappingLoader
	tfProvider *schema.Provider
	httpClient *http.Client

	controllerConfig *config.ControllerConfig
}

func NewUnstructuredResourceStreamFromURL(url string, provider *schema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, gcpClient gcpclient.Client, httpClient *http.Client, controllerConfig *config.ControllerConfig) *URLToUnstructuredResourceStream {
	stream := URLToUnstructuredResourceStream{
		url:              url,
		smLoader:         smLoader,
		tfProvider:       provider,
		gcpClient:        gcpClient,
		httpClient:       httpClient,
		controllerConfig: controllerConfig,
	}
	return &stream
}

func (s *URLToUnstructuredResourceStream) Next(ctx context.Context) (*unstructured.Unstructured, error) {
	if s.done {
		return nil, io.EOF
	}

	// First check if this resource uses our direct-reconciliation model
	exported, err := direct.Export(ctx, s.url, s.controllerConfig)
	if err != nil {
		return nil, err
	}
	if exported != nil {
		s.done = true
		return exported, nil
	}

	skel, err := resourceskeleton.NewFromURI(s.url, s.smLoader, s.tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error converting url '%v' to skeleton: %w", s.url, err)
	}
	u, err := s.gcpClient.Get(ctx, skel)
	if err != nil {
		return nil, fmt.Errorf("error getting '%v': %w", s.url, err)
	}
	s.done = true
	return u, nil
}
