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
	gkehubv1 "google.golang.org/api/gkehub/v1"
	featureapi "google.golang.org/api/gkehub/v1beta"
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
	scopeClient     *featureapi.ProjectsLocationsScopesService
	operationClient *featureapi.ProjectsLocationsOperationsService

	namespaceClient   *gkehubv1.ProjectsLocationsScopesNamespacesService
	v1ScopeClient     *gkehubv1.ProjectsLocationsScopesService
	v1OperationClient *gkehubv1.ProjectsLocationsOperationsService
}

func (m *gcpClient) newGkeHubClient(ctx context.Context) (*gkeHubClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	service, err := featureapi.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building service for gkehub v1beta: %w", err)
	}

	servicev1, err := gkehubv1.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building service for gkehub v1: %w", err)
	}

	return &gkeHubClient{
		featureClient:     featureapi.NewProjectsLocationsFeaturesService(service),
		scopeClient:       featureapi.NewProjectsLocationsScopesService(service),
		operationClient:   featureapi.NewProjectsLocationsOperationsService(service),
		namespaceClient:   gkehubv1.NewProjectsLocationsScopesNamespacesService(servicev1),
		v1ScopeClient:     gkehubv1.NewProjectsLocationsScopesService(servicev1),
		v1OperationClient: gkehubv1.NewProjectsLocationsOperationsService(servicev1),
	}, nil
}
