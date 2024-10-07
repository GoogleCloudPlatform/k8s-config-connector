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

package firewallpolicyrule

import (
	"context"
	"fmt"
	"net/http"

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

func (m *gcpClient) options() ([]option.ClientOption, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		// TODO: Set UserAgent in this scenario (error is: WithHTTPClient is incompatible with gRPC dial options)

		httpClient := &http.Client{}
		*httpClient = *m.config.HTTPClient
		httpClient.Transport = &optionsRoundTripper{
			config: m.config,
			inner:  m.config.HTTPClient.Transport,
		}
		opts = append(opts, option.WithHTTPClient(httpClient))
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

type optionsRoundTripper struct {
	config config.ControllerConfig
	inner  http.RoundTripper
}

func (m *optionsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.config.UserAgent != "" {
		req.Header.Set("User-Agent", m.config.UserAgent)
	}
	return m.inner.RoundTrip(req)
}

func (m *gcpClient) firewallPoliciesClient(ctx context.Context) (*api.FirewallPoliciesClient, error) {
	opts, err := m.options()
	if err != nil {
		return nil, err
	}
	client, err := api.NewFirewallPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FirewallPolicy client: %w", err)
	}
	return client, err
}
