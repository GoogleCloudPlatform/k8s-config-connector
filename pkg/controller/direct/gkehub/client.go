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

package gkehub

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	featureapi "google.golang.org/api/gkehub/v1beta"
	"google.golang.org/api/option"
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

type gkeHubClient struct {
	featureClient   *featureapi.ProjectsLocationsFeaturesService
	operationClient *featureapi.ProjectsLocationsOperationsService
}

func (m *gcpClient) newGkeHubClient(ctx context.Context) (*gkeHubClient, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	service, err := featureapi.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("building service for gkehub: %w", err)
	}
	return &gkeHubClient{
		featureClient:   featureapi.NewProjectsLocationsFeaturesService(service),
		operationClient: featureapi.NewProjectsLocationsOperationsService(service),
	}, nil
}
