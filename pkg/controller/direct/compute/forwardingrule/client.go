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

package forwardingrule

import (
	"context"
	"fmt"

	api "cloud.google.com/go/compute/apiv1"
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

func (m *gcpClient) newClientOptions(ctx context.Context) ([]option.ClientOption, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	opts := []option.ClientOption{option.WithHTTPClient(httpClient)}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}
	return opts, nil
}

func (m *gcpClient) globalForwardingRuleClient(ctx context.Context) (*api.GlobalForwardingRulesClient, error) {
	opts, err := m.newClientOptions(ctx)
	if err != nil {
		return nil, err
	}
	client, err := api.NewGlobalForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building global ComputeForwardingRule client: %w", err)
	}
	return client, err
}

func (m *gcpClient) forwardingRuleClient(ctx context.Context) (*api.ForwardingRulesClient, error) {
	opts, err := m.newClientOptions(ctx)
	if err != nil {
		return nil, err
	}
	client, err := api.NewForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeForwardingRule client: %w", err)
	}
	return client, err
}
