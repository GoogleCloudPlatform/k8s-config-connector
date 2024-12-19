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

package apigee

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	api "google.golang.org/api/apigee/v1"
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

type apigeeEnvgroupClient struct {
	featureClient   *api.OrganizationsEnvgroupsService
	operationClient *api.OrganizationsOperationsService
}

func (m *gcpClient) newApigeeEnvgroupClient(ctx context.Context) (*apigeeEnvgroupClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	service, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ApigeeEnvgroup client: %w", err)
	}
	return &apigeeEnvgroupClient{
		featureClient:   api.NewOrganizationsEnvgroupsService(service),
		operationClient: api.NewOrganizationsOperationsService(service),
	}, err
}
