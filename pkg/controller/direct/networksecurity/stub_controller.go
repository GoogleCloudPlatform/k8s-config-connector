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

package networksecurity

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityBackendAuthenticationConfigGVK, NewNetworkSecurityBackendAuthenticationConfigModel)
	registry.RegisterModel(krm.NetworkSecurityInterceptEndpointGroupAssociationGVK, NewNetworkSecurityInterceptEndpointGroupAssociationModel)
	registry.RegisterModel(krm.NetworkSecurityInterceptEndpointGroupGVK, NewNetworkSecurityInterceptEndpointGroupModel)
	registry.RegisterModel(krm.NetworkSecurityInterceptDeploymentGVK, NewNetworkSecurityInterceptDeploymentModel)
}

func NewNetworkSecurityBackendAuthenticationConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkSecurityBackendAuthenticationConfigModel{config: config}, nil
}

type networkSecurityBackendAuthenticationConfigModel struct {
	config *config.ControllerConfig
}

func (m *networkSecurityBackendAuthenticationConfigModel) client(ctx context.Context) (client.Client, error) {
	return nil, nil
}

func (m *networkSecurityBackendAuthenticationConfigModel) AdapterForObject(ctx context.Context, opts *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return nil, nil
}

func (m *networkSecurityBackendAuthenticationConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func NewNetworkSecurityInterceptEndpointGroupAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkSecurityInterceptEndpointGroupAssociationModel{config: config}, nil
}

type networkSecurityInterceptEndpointGroupAssociationModel struct {
	config *config.ControllerConfig
}

func (m *networkSecurityInterceptEndpointGroupAssociationModel) client(ctx context.Context) (client.Client, error) {
	return nil, nil
}

func (m *networkSecurityInterceptEndpointGroupAssociationModel) AdapterForObject(ctx context.Context, opts *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return nil, nil
}

func (m *networkSecurityInterceptEndpointGroupAssociationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func NewNetworkSecurityInterceptEndpointGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkSecurityInterceptEndpointGroupModel{config: config}, nil
}

type networkSecurityInterceptEndpointGroupModel struct {
	config *config.ControllerConfig
}

func (m *networkSecurityInterceptEndpointGroupModel) client(ctx context.Context) (client.Client, error) {
	return nil, nil
}

func (m *networkSecurityInterceptEndpointGroupModel) AdapterForObject(ctx context.Context, opts *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return nil, nil
}

func (m *networkSecurityInterceptEndpointGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func NewNetworkSecurityInterceptDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkSecurityInterceptDeploymentModel{config: config}, nil
}

type networkSecurityInterceptDeploymentModel struct {
	config *config.ControllerConfig
}

func (m *networkSecurityInterceptDeploymentModel) client(ctx context.Context) (client.Client, error) {
	return nil, nil
}

func (m *networkSecurityInterceptDeploymentModel) AdapterForObject(ctx context.Context, opts *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return nil, nil
}

func (m *networkSecurityInterceptDeploymentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}
