// Copyright 2022 Google LLC
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

package clientconfig

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/logger"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"golang.org/x/oauth2/google"
	"k8s.io/klog/v2"
)

var nonretryable = dcl.Retryability{Retryable: false}

type Options struct {
	config.ControllerConfig
}

func newConfigAndClient(ctx context.Context, opt Options) (*dcl.Config, *http.Client, error) {
	if opt.UserAgent == "" {
		opt.UserAgent = gcp.KCCUserAgent()
	}

	if opt.HTTPClient == nil {
		httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating the http client to be used by DCL: %w", err)
		}
		opt.HTTPClient = httpClient
	}

	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(opt.HTTPClient),
		dcl.WithUserAgent(opt.UserAgent),
		dcl.WithLogger(logger.SimpleDCLLogger()),
		// ConfigControllerInstance takes ~17 minutes to be created.
		dcl.WithTimeout(30 * time.Minute),
		dcl.WithCodeRetryability(
			map[int]dcl.Retryability{
				// ComputeFirewallPolicy uses error code 403 for quota exceeded
				// errors. Quota exceeded errors should not be retryable.
				403: nonretryable,
				// Quota exceeded errors are usually surfaced by a 429 error
				// code in GCP. Quota exceeded errors should not be retryable.
				429: nonretryable,
			},
		),
	}
	if opt.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
	}
	if opt.BillingProject != "" {
		configOptions = append(configOptions, dcl.WithBillingProject(opt.BillingProject))
	}
	dclConfig := dcl.NewConfig(configOptions...)
	return dclConfig, opt.HTTPClient, nil
}

func New(ctx context.Context, opt Options) (*dcl.Config, error) {
	dclConfig, _, err := newConfigAndClient(ctx, opt)
	if err != nil {
		return nil, err
	}

	return dclConfig, nil
}

// NewForIntegrationTest creates a dcl.Config for use in integration tests.
// Deprecated: Prefer using a harness.
func NewForIntegrationTest() *dcl.Config {
	dclConfig, _ := NewConfigAndClientForIntegrationTest()
	return dclConfig
}

func NewConfigAndClientForIntegrationTest() (*dcl.Config, *http.Client) {
	ctx := context.TODO()
	eventSinks := test.EventSinksFromContext(ctx)

	if artifacts := os.Getenv("ARTIFACTS"); artifacts != "" {
		outputDir := filepath.Join(artifacts, "http-logs")

		eventSinks = append(eventSinks, test.NewDirectoryEventSink(outputDir))
	}

	opt := Options{}
	opt.UserAgent = "kcc/dev"

	// Log DCL requests
	if len(eventSinks) != 0 {
		if opt.HTTPClient == nil {
			httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
			if err != nil {
				klog.Fatalf("error creating the http client to be used by DCL: %v", err)
			}
			opt.HTTPClient = httpClient
		}
		t := test.NewHTTPRecorder(opt.HTTPClient.Transport, eventSinks...)
		opt.HTTPClient = &http.Client{Transport: t}
	}

	config, httpClient, err := newConfigAndClient(ctx, opt)
	if err != nil {
		klog.Fatalf("error from NewForIntegrationTest: %v", err)
	}
	return config, httpClient
}

func CopyAndModifyForKind(dclConfig *dcl.Config, kind string) *dcl.Config {
	// Configure a custom DCL timeout.
	if timeout, ok := kindToTimeout[kind]; ok {
		return dclConfig.Clone(dcl.WithTimeout(timeout))
	}
	return dclConfig.Clone()
}
