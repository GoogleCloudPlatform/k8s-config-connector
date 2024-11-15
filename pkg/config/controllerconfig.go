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
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

type ControllerConfig struct {
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

	// GCPTokenSource mints OAuth2 tokens to be passed with GCP API calls,
	// allowing use of a non-default OAuth2 identity
	GCPTokenSource oauth2.TokenSource
}

func (c *ControllerConfig) RESTClientOptions() ([]option.ClientOption, error) {
	quotaProject := ""
	if c.UserProjectOverride && c.BillingProject != "" {
		quotaProject = c.BillingProject
	}

	var opts []option.ClientOption
	if c.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(c.UserAgent))
	}
	if c.HTTPClient != nil {
		httpClient := &http.Client{}
		*httpClient = *c.HTTPClient
		httpClient.Transport = &optionsRoundTripper{
			config:       *c,
			quotaProject: quotaProject,
			inner:        c.HTTPClient.Transport,
		}
		opts = append(opts, option.WithHTTPClient(httpClient))

		// quotaProject is incompatible with http client
		quotaProject = ""
	}
	if quotaProject != "" {
		opts = append(opts, option.WithQuotaProject(quotaProject))
	}
	if c.GCPTokenSource != nil {
		opts = append(opts, option.WithTokenSource(c.GCPTokenSource))
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
	if c.HTTPClient != nil {
		var dialOption grpc.DialOption
		if c.GCPTokenSource != nil {
			dialOption = grpc.WithPerRPCCredentials(oauth.TokenSource{TokenSource: c.GCPTokenSource})
		} else {
			dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
		}
		opts = append(opts, option.WithGRPCDialOption(dialOption))

	}
	if c.UserProjectOverride && c.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(c.BillingProject))
	}
	if c.GCPTokenSource != nil {
		opts = append(opts, option.WithTokenSource(c.GCPTokenSource))
	}
	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	return opts, nil
}

var unaryInterceptor = func() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
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
