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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/logger"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"golang.org/x/oauth2/google"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

var nonretryable = dcl.Retryability{Retryable: false}

type Options struct {
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
}

func New(ctx context.Context, opt Options) (*dcl.Config, error) {
	if opt.UserAgent == "" {
		opt.UserAgent = gcp.KCCUserAgent
	}

	if opt.HTTPClient == nil {
		httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
		if err != nil {
			return nil, fmt.Errorf("error creating the http client to be used by DCL: %w", err)
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
	return dclConfig, nil
}

// NewForIntegrationTest creates a dcl.Config for use in integration tests.
// Deprecated: Prefer using a harness.
func NewForIntegrationTest() *dcl.Config {
	ctx := context.TODO()
	opt := Options{
		UserAgent: "kcc/dev",
	}

	config, err := New(ctx, opt)
	if err != nil {
		klog.Fatalf("error from NewForIntegrationTest: %v", err)
	}
	return config
}

func CopyAndModifyForKind(dclConfig *dcl.Config, kind string) *dcl.Config {
	// Configure a custom DCL timeout.
	if timeout, ok := kindToTimeout[kind]; ok {
		return dclConfig.Clone(dcl.WithTimeout(timeout))
	}
	return dclConfig.Clone()
}

// SetUserAgentWithBlueprintAttribution returns a new DCL Config with the user agent containing the blueprint attribution
// if the resource has the blueprint attribution annotation. Otherwise, the existing DCL Config is unmodified and returned.
func SetUserAgentWithBlueprintAttribution(dclConfig *dcl.Config, resource metav1.Object) *dcl.Config {
	bp, found := k8s.GetAnnotation(k8s.BlueprintAttributionAnnotation, resource)
	if !found {
		return dclConfig
	}
	userAgentWithBlueprintAttribution := fmt.Sprintf("%v blueprints/%v", gcp.KCCUserAgent, bp)
	newConfig := dclConfig.Clone(dcl.WithUserAgent(userAgentWithBlueprintAttribution))
	return newConfig
}
