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

package sql

import (
	"context"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	api "google.golang.org/api/sqladmin/v1beta4"
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

func (m *gcpClient) sqlOperationsClient() *api.OperationsService {
	return api.NewOperationsService(m.service)
}

func (m *gcpClient) sqlInstancesClient() *api.InstancesService {
	return api.NewInstancesService(m.service)
}

func (m *gcpClient) sqlUsersClient() *api.UsersService {
	return api.NewUsersService(m.service)
}

// NewGCPOperationError builds an error from a GCP operation error.
func NewGCPOperationError(opErr *api.OperationErrors) error {
	var errs []error
	if opErr != nil {
		for _, err := range opErr.Errors {
			errs = append(errs, fmt.Errorf("code=%q, message=%q", err.Code, err.Message))
		}
	}
	if len(errs) == 0 {
		return fmt.Errorf("gcp operation failed with unknown error, raw error: %+v", opErr)
	}
	return fmt.Errorf("gcp operation failed: %w", errors.Join(errs...))
}
