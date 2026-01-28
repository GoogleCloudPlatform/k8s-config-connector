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

package config

import (
	"context"
	"fmt"
	"net/http"
	"time"

	cloudresourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	metricstransport "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics/transport"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	ghttptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
)

type ControllerConfig struct {
	// UserAgent sets the User-Agent to pass in HTTP request headers
	UserAgent string

	// UserProjectOverride provides the option to use the resource project for preconditions, quota, and billing,
	// instead of the project the credentials belong to; false by default
	UserProjectOverride bool

	// BillingProject is the project used by the TF provider and DCL client to determine preconditions,
	// quota, and billing if UserProjectOverride is set to true. If this field is empty,
	// but UserProjectOverride is set to true, resource project will be used.
	BillingProject string

	// HTTPClient allows us to specify the HTTP client to use with DCL.
	// This is particularly useful in mocks/tests.
	HTTPClient *http.Client

	// GRPCUnaryClientInterceptor is the GRPC interceptor for use in tests.
	GRPCUnaryClientInterceptor grpc.UnaryClientInterceptor

	// GCPTokenSource mints OAuth2 tokens to be passed with GCP API calls,
	// allowing use of a non-default OAuth2 identity
	GCPTokenSource oauth2.TokenSource

	// EnableMetricsTransport enables automatic wrapping of HTTP clients with metrics transport
	EnableMetricsTransport bool

	// ProjectMapper maps between project ids and numbers
	ProjectMapper *projects.ProjectMapper
}

func (c *ControllerConfig) Init(ctx context.Context) error {
	if c.ProjectMapper == nil {
		opts, err := c.RESTClientOptions()
		if err != nil {
			return err
		}

		projectsClient, err := cloudresourcemanager.NewProjectsRESTClient(ctx, opts...)
		if err != nil {
			return fmt.Errorf("building cloudresourcemanager client: %w", err)
		}
		projectCache := projects.NewProjectCache(projectsClient, 4*time.Hour)
		c.ProjectMapper = projects.NewProjectMapper(projectCache)
	}
	return nil
}

type RESTClientOption func(o *restClientOptions)

type restClientOptions struct {
	defaultQuotaProject string
}

func WithDefaultQuotaProject(project string) RESTClientOption {
	return func(o *restClientOptions) {
		o.defaultQuotaProject = project
	}
}

func (c *ControllerConfig) RESTClientOptions(options ...RESTClientOption) ([]option.ClientOption, error) {
	var restClientOptions restClientOptions
	for _, option := range options {
		option(&restClientOptions)
	}

	quotaProject := ""
	if c.UserProjectOverride && c.BillingProject != "" {
		quotaProject = c.BillingProject
	}

	if restClientOptions.defaultQuotaProject != "" && quotaProject == "" {
		quotaProject = restClientOptions.defaultQuotaProject
	}

	var opts []option.ClientOption
	if c.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(c.UserAgent))
	}

	if c.HTTPClient != nil {
		httpClient := &http.Client{}
		*httpClient = *c.HTTPClient

		transport := c.HTTPClient.Transport
		if c.EnableMetricsTransport {
			transport = metricstransport.NewMetricsTransport(transport)
		}

		httpClient.Transport = &optionsRoundTripper{
			config:       *c,
			quotaProject: quotaProject,
			inner:        transport,
		}
		opts = append(opts, option.WithHTTPClient(httpClient))

		// quotaProject is incompatible with http client
		quotaProject = ""
	} else {
		// the default HTTP client is used and wired up with Google auth

		// we cannot pass both a custom http client and a token source to the Google transport
		if c.GCPTokenSource != nil {
			opts = append(opts, option.WithTokenSource(c.GCPTokenSource))
		}
	}

	if quotaProject != "" {
		opts = append(opts, option.WithQuotaProject(quotaProject))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	return opts, nil
}

func (c *ControllerConfig) GRPCClientOptions() ([]option.ClientOption, error) {
	var opts []option.ClientOption
	if c.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(c.UserAgent))
	}
	if c.UserProjectOverride && c.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(c.BillingProject))
	}
	if c.GCPTokenSource != nil {
		opts = append(opts, option.WithTokenSource(c.GCPTokenSource))
	}
	if c.GRPCUnaryClientInterceptor != nil {
		opts = append(opts, option.WithGRPCDialOption(grpc.WithUnaryInterceptor(c.GRPCUnaryClientInterceptor)))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	return opts, nil
}

type optionsRoundTripper struct {
	config       ControllerConfig
	quotaProject string
	inner        http.RoundTripper
}

func (m *optionsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.config.UserAgent != "" {
		req.Header.Set("User-Agent", m.config.UserAgent)
	}
	if m.quotaProject != "" {
		req.Header.Set("X-goog-user-project", m.quotaProject)
	}
	return m.inner.RoundTrip(req)
}

// NewAuthenticatedHTTPClient creates an HTTP client with proper authentication
// and optionally wraps it with metrics transport
func (c *ControllerConfig) NewAuthenticatedHTTPClient(ctx context.Context) (*http.Client, error) {
	opts, err := c.RESTClientOptions()
	if err != nil {
		return nil, fmt.Errorf("error creating REST client options: %w", err)
	}
	if c.HTTPClient != nil {
		c, _, err := ghttptransport.NewClient(ctx, opts...)
		return c, err
	}

	baseTransport := http.DefaultTransport
	if c.EnableMetricsTransport {
		baseTransport = metricstransport.NewMetricsTransport(baseTransport)
	}

	// Create an authenticated transport
	authTransport, err := ghttptransport.NewTransport(ctx, baseTransport, opts...)
	if err != nil {
		return nil, fmt.Errorf("error creating authenticated transport: %w", err)
	}

	return &http.Client{Transport: authTransport}, nil
}
