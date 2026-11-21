// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Config Connector patch: per-identity copies of the provider Config, so one
// controller process can make Terraform calls with different credentials.

package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

// CopyWithTokenSource returns a copy of c that authenticates with tokenSource.
// The copy gets its own HTTP client, built like the one LoadAndValidate builds,
// and its own request batchers, so batched calls never mix identities. c must
// already be loaded with LoadAndValidate. The batchers stop when c.Context is
// done.
func (c *Config) CopyWithTokenSource(tokenSource oauth2.TokenSource) (*Config, error) {
	if c.Context == nil || c.Client == nil {
		return nil, fmt.Errorf("config is not loaded; call LoadAndValidate first")
	}
	cp := *c
	cp.tokenSource = tokenSource
	client, err := cp.newHTTPClient(c.Context, tokenSource)
	if err != nil {
		return nil, err
	}
	cp.Client = client
	cp.RequestBatcherServiceUsage = NewRequestBatcher("Service Usage", c.Context, c.BatchingConfig)
	cp.RequestBatcherIam = NewRequestBatcher("IAM", c.Context, c.BatchingConfig)
	return &cp, nil
}

// newHTTPClient builds the client chain used for GCP API calls: mTLS and
// authentication, then request logging, retries and extra headers.
func (c *Config) newHTTPClient(ctx context.Context, tokenSource oauth2.TokenSource) (*http.Client, error) {
	oauth2Client := cleanhttp.DefaultClient()
	if OAuth2HTTPClientTransformer != nil {
		oauth2Client = OAuth2HTTPClientTransformer(ctx, oauth2Client)
	}
	cleanCtx := context.WithValue(ctx, oauth2.HTTPClient, oauth2Client)

	// 1. MTLS TRANSPORT/CLIENT - sets up proper auth headers
	client, _, err := transport.NewHTTPClient(cleanCtx, option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, err
	}
	if DefaultHTTPClientTransformer != nil {
		client = DefaultHTTPClientTransformer(ctx, client)
	}

	// 2. Logging Transport - ensure we log HTTP requests to GCP APIs.
	loggingTransport := logging.NewTransport("Google", client.Transport)

	// 3. Retry Transport - retries common temporary errors
	// Keep order for wrapping logging so we log each retried request as well.
	// This value should be used if needed to create shallow copies with additional retry predicates.
	// See ClientWithAdditionalRetries
	retryTransport := NewTransportWithDefaultRetries(loggingTransport)

	// 4. Header Transport - outer wrapper to inject additional headers we want to apply
	// before making requests
	headerTransport := NewTransportWithHeaders(retryTransport)
	if c.RequestReason != "" {
		headerTransport.Set("X-Goog-Request-Reason", c.RequestReason)
	}

	// Ensure $userProject is set for all HTTP requests using the client if specified by the provider config
	// See https://cloud.google.com/apis/docs/system-parameters
	if c.UserProjectOverride && c.BillingProject != "" {
		headerTransport.Set("X-Goog-User-Project", c.BillingProject)
	}

	// Set final transport value.
	client.Transport = headerTransport

	// This timeout is a timeout per HTTP request, not per logical operation.
	client.Timeout = c.synchronousTimeout()
	return client, nil
}
