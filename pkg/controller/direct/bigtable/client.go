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

package bigtable

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	api "google.golang.org/api/bigtableadmin/v2"
	"google.golang.org/api/option"
)

type gcpClient struct {
	config controller.Config
}

func newGCPClient(ctx context.Context, config *controller.Config) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	return gcpClient, nil
}

func (m *gcpClient) options() ([]option.ClientOption, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		// TODO: Set UserAgent in this scenario (error is: WithHTTPClient is incompatible with gRPC dial options)
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	return opts, nil
}

func (m *gcpClient) newAdminClient(ctx context.Context) (*api.Service, error) {
	opts, err := m.options()
	if err != nil {
		return nil, err
	}
	client, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building bigtable admin client: %w", err)
	}
	return client, err
}
