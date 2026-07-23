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

package networkconnectivity

import (
	"context"
	"fmt"

	gcpapi "cloud.google.com/go/networkconnectivity/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	api "google.golang.org/api/networkconnectivity/v1"
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

func (m *gcpClient) newNetworkConnectivityClient(ctx context.Context) (*api.Service, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networkconnectivity client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newInternalRangeClient(ctx context.Context) (*gcpapi.InternalRangeClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := gcpapi.NewInternalRangeClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networkconnectivity internalrange client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newCrossNetworkAutomationClient(ctx context.Context) (*gcpapi.CrossNetworkAutomationClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := gcpapi.NewCrossNetworkAutomationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networkconnectivity crossnetworkautomation client: %w", err)
	}
	return client, err
}

