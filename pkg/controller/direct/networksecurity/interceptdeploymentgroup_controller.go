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
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityInterceptDeploymentGroupGVK, dummyModel)
	registry.RegisterModel(krm.NetworkSecurityInterceptDeploymentGVK, dummyModel)
	registry.RegisterModel(krm.NetworkSecurityInterceptEndpointGroupGVK, dummyModel)
	registry.RegisterModel(krm.NetworkSecurityBackendAuthenticationConfigGVK, dummyModel)
}

type model struct{}

func dummyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{}, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}
