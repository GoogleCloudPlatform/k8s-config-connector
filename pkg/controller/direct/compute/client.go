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
// proto.service: google.cloud.compute.v1.ExternalVpnGateways
// proto.service: google.cloud.compute.v1.NetworkEdgeSecurityServices
// proto.service: google.cloud.compute.v1.NetworkAttachments
// proto.service: google.cloud.compute.v1.HealthChecks
// proto.service: google.cloud.compute.v1.Images
// proto.service: google.cloud.compute.v1.SecurityPolicies
// proto.service: google.cloud.compute.v1.RegionSecurityPolicies
// proto.service: google.cloud.compute.v1.Addresses
// proto.service: google.cloud.compute.v1.GlobalAddresses
// proto.service: google.cloud.compute.v1.UrlMaps
// proto.service: google.cloud.compute.v1.RegionUrlMaps
// proto.service: google.cloud.compute.v1.Networks
// proto.service: google.cloud.compute.v1.Routes
// proto.service: google.cloud.compute.v1.Autoscalers
// proto.service: google.cloud.compute.v1.NodeTemplates
// proto.service: google.cloud.compute.v1.Firewalls

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

func (m *gcpClient) newBackendBucketsClient(ctx context.Context) (*compute.BackendBucketsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewBackendBucketsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute backendBuckets client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newFutureReservationsClient(ctx context.Context) (*compute.FutureReservationsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewFutureReservationsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute FutureReservations client: %w", err)

	}
	return client, err
}

func (m *gcpClient) newReservationsClient(ctx context.Context) (*compute.ReservationsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewReservationsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Reservations client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newBackendServicesClient(ctx context.Context) (*compute.BackendServicesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewBackendServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute BackendServices client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newSslPoliciesClient(ctx context.Context) (*compute.SslPoliciesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewSslPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute SslPolicies client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newHealthChecksClient(ctx context.Context) (*compute.HealthChecksClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewHealthChecksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute HealthChecksClient: %w", err)
	}
	return client, err
}

func (m *gcpClient) newInstancesClient(ctx context.Context) (*compute.InstancesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewInstancesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Instances client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newInstanceGroupsClient(ctx context.Context) (*compute.InstanceGroupsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewInstanceGroupsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute InstanceGroups client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRoutersClient(ctx context.Context) (*compute.RoutersClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRoutersRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Routers client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newImagesClient(ctx context.Context) (*compute.ImagesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewImagesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Images client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newSecurityPoliciesClient(ctx context.Context) (*compute.SecurityPoliciesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewSecurityPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute SecurityPolicies client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newAddressesClient(ctx context.Context) (*compute.AddressesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewAddressesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Addresses client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newUrlMapsClient(ctx context.Context) (*compute.UrlMapsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewUrlMapsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute UrlMaps client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegionSecurityPoliciesClient(ctx context.Context) (*compute.RegionSecurityPoliciesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionSecurityPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute RegionSecurityPolicies client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newGlobalAddressesClient(ctx context.Context) (*compute.GlobalAddressesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewGlobalAddressesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute GlobalAddresses client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegionUrlMapsClient(ctx context.Context) (*compute.RegionUrlMapsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionUrlMapsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute RegionUrlMaps client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newNetworksClient(ctx context.Context) (*compute.NetworksClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewNetworksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute networks client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRoutesClient(ctx context.Context) (*compute.RoutesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRoutesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute routes client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newDisksClient(ctx context.Context) (*compute.DisksClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewDisksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeDisks client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegionDisksClient(ctx context.Context) (*compute.RegionDisksClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionDisksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeRegionDisks client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newAutoscalersClient(ctx context.Context) (*compute.AutoscalersClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewAutoscalersRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeAutoscalers client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newFirewallsClient(ctx context.Context) (*compute.FirewallsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewFirewallsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeFirewalls client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newTargetHttpsProxiesClient(ctx context.Context) (*compute.TargetHttpsProxiesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewTargetHttpsProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute TargetHttpsProxies REST client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegionalTargetHttpsProxiesClient(ctx context.Context) (*compute.RegionTargetHttpsProxiesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionTargetHttpsProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute RegionTargetHttpsProxies REST client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newNodeTemplatesClient(ctx context.Context) (*compute.NodeTemplatesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewNodeTemplatesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeNodeTemplates client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newInstanceGroupManagersClient(ctx context.Context) (*compute.InstanceGroupManagersClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewInstanceGroupManagersRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeInstanceGroupManagers client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newRegionInstanceGroupManagersClient(ctx context.Context) (*compute.RegionInstanceGroupManagersClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewRegionInstanceGroupManagersRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeRegionInstanceGroupManagers client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newExternalVpnGatewaysClient(ctx context.Context) (*compute.ExternalVpnGatewaysClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := compute.NewExternalVpnGatewaysRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeExternalVpnGateways client: %w", err)
	}
	return client, err
}
