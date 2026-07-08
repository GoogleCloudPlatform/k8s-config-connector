// Copyright 2026 Google LLC
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

package apigeeregistry

import (
	"context"
	"fmt"

	api "cloud.google.com/go/apigeeregistry/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type gcpClient struct {
	config config.ControllerConfig
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	return gcpClient, nil
}

func (m *gcpClient) newProvisioningClient(ctx context.Context) (*api.ProvisioningClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewProvisioningClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building provisioning client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegistryClient(ctx context.Context) (*api.RegistryClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewRegistryClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building registry client: %w", err)
	}
	return client, err
}
