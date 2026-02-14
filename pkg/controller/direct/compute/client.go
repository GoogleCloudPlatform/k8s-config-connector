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

// +tool:controller-client
// proto.service: google.cloud.compute.v1.NetworkEdgeSecurityServices
// proto.service: google.cloud.compute.v1.NetworkAttachments

package compute

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type gcpClient struct {
	config config.ControllerConfig
}

func newGCPClient(config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	return gcpClient, nil
}

func (m *gcpClient) newGlobalForwardingRuleClient(ctx context.Context) (*compute.GlobalForwardingRulesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	client, err := compute.NewGlobalForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building global compute ForwardingRule client: %w", err)

	}
	return client, err
}

func (m *gcpClient) forwardingRuleClient(ctx context.Context) (*compute.ForwardingRulesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeForwardingRule client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newNetworkEdgeSecurityServicesClient(ctx context.Context) (*compute.NetworkEdgeSecurityServicesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewNetworkEdgeSecurityServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute networkEdgeSecurityServices client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newNetworkAttachmentsClient(ctx context.Context) (*compute.NetworkAttachmentsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewNetworkAttachmentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute networkEdgeSecurityServices client: %w", err)

	}
	return client, err
}

func (m *gcpClient) newTargetTcpProxiesClient(ctx context.Context) (*compute.TargetTcpProxiesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	client, err := compute.NewTargetTcpProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute TargetTcpProxiesClient client: %w", err)

	}
	return client, err
}

func (m *gcpClient) newRegionalTargetTcpProxiesClient(ctx context.Context) (*compute.RegionTargetTcpProxiesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionTargetTcpProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute RegionalTargetTcpProxiesClient client: %w", err)

	}
	return client, err
}

func (m *gcpClient) urlMapsClient(ctx context.Context) (*compute.UrlMapsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewUrlMapsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute UrlMapsClient client: %w", err)
	}
	return client, err
}

func (m *gcpClient) regionUrlMapsClient(ctx context.Context) (*compute.RegionUrlMapsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionUrlMapsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute RegionUrlMapsClient client: %w", err)
	}
	return client, err
}
