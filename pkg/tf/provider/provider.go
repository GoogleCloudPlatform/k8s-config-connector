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

package tfprovider

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"k8s.io/klog/v2"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Config holds additional configuration for the google TF provider
type Config struct {
	// AccessToken is the access_token to be passed to the TF provider (if non-empty),
	// allowing use of a non-default OAuth2 identity
	AccessToken string

	// Scopes is the list of OAuth2 scopes to be passed to the TF provider,
	// allowing use of non-default OAuth2 scopes. If none are specified, then
	// Terraform has a default list of scopes that it will use.
	Scopes []string

	// Controls the quota project used in requests to GCP APIs for the purpose of preconditions,
	// quota, and billing. If false, the quota project is determined by the API and may be the project
	// associated with your credentials, or the resource project. If true, most resources in
	// the provider will explicitly supply their resource project, as described in their documentation.
	// Otherwise, a billing_project value must be supplied.
	// https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override
	UserProjectOverride bool

	// BillingProject is the project used by the TF provider for preconditions,
	// quota, and billing if UserProjectOverride is set to true. If this field is empty,
	// but UserProjectOverride is set to true, then the TF provider uses the resource's project.
	// https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#billing_project
	BillingProject string
}

var DefaultConfig = NewConfig()

func UnitTestConfig() Config {
	return Config{
		Scopes: append(deepcopy.StringSlice(transport_tpg.DefaultClientScopes),

			// Needed by the KCC controller to be able to create resources that
			// read Google Drive files.
			"https://www.googleapis.com/auth/drive.readonly",
		),
		AccessToken: "dummyToken",
	}
}

func NewConfig() Config {
	return Config{
		Scopes: append(deepcopy.StringSlice(transport_tpg.DefaultClientScopes),

			// Needed by the KCC controller to be able to create resources that
			// read Google Drive files.
			"https://www.googleapis.com/auth/drive.readonly",
		),
	}
}

// New builds a new tfschema.Provider for the google provider.
func New(ctx context.Context, config Config) (*tfschema.Provider, error) {
	googleProvider := google.Provider()
	cfgMap := map[string]interface{}{}
	if config.AccessToken != "" {
		cfgMap["access_token"] = config.AccessToken
	}

	cfgMap["scopes"] = config.Scopes
	cfgMap["user_project_override"] = config.UserProjectOverride
	cfgMap["billing_project"] = config.BillingProject

	schema := tfschema.InternalMap(googleProvider.Schema).CoreConfigSchema()
	cfg := terraform.NewResourceConfigShimmed(krmtotf.MapToCtyVal(cfgMap, schema.ImpliedType()), schema)
	if err := googleProvider.Configure(ctx, cfg); err != nil {
		return nil, fmt.Errorf("error configuring provider: %v", err)
	}
	return googleProvider, nil
}

// NewOrLogFatal calls New and panics on error
// deprecated: Prefer New and handle the error
func NewOrLogFatal(config Config) *tfschema.Provider {
	ctx := context.TODO()
	p, err := New(ctx, config)
	if err != nil {
		klog.Fatal(err)
	}
	return p
}
