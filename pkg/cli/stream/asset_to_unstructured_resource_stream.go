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
	"errors"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type AssetToUnstructuredResourceStream struct {
	assetStream   AssetStream
	gcpClient     gcpclient.Client
	serviceClient serviceclient.ServiceClient
	smLoader      *servicemappingloader.ServiceMappingLoader
	tfProvider    *schema.Provider
	config        config.ControllerConfig
}

// NewUnstructuredResourceStreamFromAssetStream returns an unstructured stream. The stream converts each asset in the 'assetStream' to
// a KCC resource and does a GET request to GCP finally returning the current value of the resource in KCC format
// as an unstructured
func NewUnstructuredResourceStreamFromAssetStream(assetStream AssetStream, client gcpclient.Client, tfProvider *schema.Provider, serviceClient serviceclient.ServiceClient, config *config.ControllerConfig) (*AssetToUnstructuredResourceStream, error) {
	stream, err := newUnstructuredResourceStreamFromAssetStream(assetStream, tfProvider, serviceClient, config)
	if err != nil {
		return nil, err
	}
	stream.gcpClient = client
	return stream, nil
}

func newUnstructuredResourceStreamFromAssetStream(assetStream AssetStream, tfProvider *schema.Provider, serviceClient serviceclient.ServiceClient, config *config.ControllerConfig) (*AssetToUnstructuredResourceStream, error) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error creating service mapping loader: %w", err)
	}
	stream := AssetToUnstructuredResourceStream{
		assetStream:   assetStream,
		serviceClient: serviceClient,
		smLoader:      smLoader,
		tfProvider:    tfProvider,
		config:        *config,
	}
	return &stream, nil
}

func (s *AssetToUnstructuredResourceStream) Next(ctx context.Context) (*unstructured.Unstructured, error) {
	asset, err := s.assetStream.Next()
	if err != nil {
		if !errors.Is(err, io.EOF) {
			err = fmt.Errorf("error getting next asset: %w", err)
		}
		return nil, err
	}

	// First check if this resource uses our direct-reconciliation model
	exported, err := direct.Export(ctx, asset.Name, &s.config)
	if err != nil {
		return nil, err
	}
	if exported != nil {
		return exported, nil
	}

	skel, err := resourceskeleton.NewFromAsset(asset, s.smLoader, s.tfProvider, s.serviceClient)
	if err != nil {
		return nil, fmt.Errorf("error converting asset '%v' with kind '%v' to skeleton: %w", asset.Name, asset.AssetType, err)
	}
	u, err := s.gcpClient.Get(ctx, skel)
	if err != nil {
		return nil, fmt.Errorf("error getting '%v': %w", asset.Name, err)
	}
	return u, nil
}
