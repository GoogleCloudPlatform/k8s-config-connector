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

// +tool:controller-client
// proto.service: google.cloud.recaptchaenterprise.v1.RecaptchaEnterpriseService

package recaptchaenterprise

import (
	"context"
	"fmt"

	api "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	"google.golang.org/api/option"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type gcpClient struct {
	config config.ControllerConfig
}

func newReCAPTCHAEnterpriseClient(ctx context.Context, config *config.ControllerConfig) (*api.Client, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	opts, err := gcpClient.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	opts = append(opts, option.WithEndpoint("public-preview-recaptchaenterprise.googleapis.com:443"))

	grpcClient, err := api.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReCAPTCHAEnterprise client: %w", err)
	}

	return grpcClient, err
}
