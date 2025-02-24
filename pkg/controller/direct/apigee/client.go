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
	config  config.ControllerConfig
	service *api.Service
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient.service, err = api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building gcp service client: %w", err)
	}

	return gcpClient, nil
}

func (m *gcpClient) instancesClient() *api.OrganizationsInstancesService {
	return api.NewOrganizationsInstancesService(m.service)
}

func (m *gcpClient) envgroupsClient() *api.OrganizationsEnvgroupsService {
	return api.NewOrganizationsEnvgroupsService(m.service)
}

func (m *gcpClient) operationsClient() *api.OrganizationsOperationsService {
	return api.NewOrganizationsOperationsService(m.service)
}

func (m *gcpClient) endpointAttachmentClient() *api.OrganizationsEndpointAttachmentsService {
	return api.NewOrganizationsEndpointAttachmentsService(m.service)
}
