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
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
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
}

func NewUnstructuredResourceStreamFromURL(url string, provider *schema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, gcpClient gcpclient.Client) *URLToUnstructuredResourceStream {
	stream := URLToUnstructuredResourceStream{
		url:        url,
		smLoader:   smLoader,
		tfProvider: provider,
		gcpClient:  gcpClient,
	}
	return &stream
}

func (s *URLToUnstructuredResourceStream) Next() (*unstructured.Unstructured, error) {
	if s.done {
		return nil, io.EOF
	}
	skel, err := resourceskeleton.NewFromURI(s.url, s.smLoader, s.tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error converting url '%v' to skeleton: %v", s.url, err)
	}
	u, err := s.gcpClient.Get(skel)
	if err != nil {
		return nil, fmt.Errorf("error getting '%v': %v", s.url, err)
	}
	s.done = true
	return u, nil
}
