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
	"log"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/logger"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/golang/glog"
	"golang.org/x/oauth2/google"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var nonretryable = dcl.Retryability{Retryable: false}

func New(enableUserProjectOverride bool, billingProject string) (*dcl.Config, error) {
	return newConfig(gcp.KCCUserAgent, enableUserProjectOverride, billingProject)
}

func newConfig(ua string, enableUserProjectOverride bool, billingProject string) (*dcl.Config, error) {
	httpClient, err := google.DefaultClient(context.Background(), gcp.ClientScopes...)
	if err != nil {
		log.Fatalf("error creating the http client to be used by DCL")
	}
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(httpClient),
		dcl.WithUserAgent(ua),
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
	if enableUserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
	}
	if billingProject != "" {
		configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
	}
	dclConfig := dcl.NewConfig(configOptions...)
	return dclConfig, nil
}

func NewForIntegrationTest() *dcl.Config {
	config, err := newConfig("kcc/dev", false, "")
	if err != nil {
		glog.Fatal(err)
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
