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
// proto.service: google.cloud.dataplex.v1.ContentService
// proto.service: google.cloud.dataplex.v1.DataplexService

package dataplex

import (
	"context"
	"fmt"

	api "cloud.google.com/go/dataplex/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"google.golang.org/api/option"
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

func (m *gcpClient) options() ([]option.ClientOption, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	return opts, nil
}

func (m *gcpClient) client(ctx context.Context) (*api.Client, error) {
	opts, err := m.options()
	if err != nil {
		return nil, err
	}

	grpcClient, err := api.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataplex client: %w", err)
	}

	return grpcClient, err
}

func (m *gcpClient) newContentClient(ctx context.Context) (*api.ContentClient, error) {
	opts, err := m.options()
	if err != nil {
		return nil, err
	}

	grpcClient, err := api.NewContentClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataplex content client: %w", err)
	}

	return grpcClient, err
}
