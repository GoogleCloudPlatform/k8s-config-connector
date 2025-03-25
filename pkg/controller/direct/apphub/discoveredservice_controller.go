// Copyright 2024 Google LLC
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

// +tool:controller
// proto.service: google.cloud.apphub.v1.AppHub
// proto.message: google.cloud.apphub.v1.DiscoveredService
// crd.type: AppHubDiscoveredService
// crd.version: v1alpha1

package apphub

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/apphub/apiv1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func init() {
	registry.RegisterModel(krm.AppHubDiscoveredServiceGVK, NewDiscoveredServiceModel)
}

func NewDiscoveredServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &discoveredServiceModel{config: *config}, nil
}

var _ directbase.Model = &discoveredServiceModel{}

type discoveredServiceModel struct {
	config config.ControllerConfig
}

func (m *discoveredServiceModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building apphub discovered service client: %w", err)
	}
	return gcpClient, nil
}

func (m *discoveredServiceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	return nil, fmt.Errorf("AdapterForObject not implemented")
}

func (m *discoveredServiceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, fmt.Errorf("AdapterForURL not implemented")
}
