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

package privateca

import (
	"context"
	"fmt"

	api "cloud.google.com/go/security/privateca/apiv1"
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

func (m *gcpClient) newCertificateAuthorityClient(ctx context.Context) (*api.CertificateAuthorityClient, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	service, err := api.NewCertificateAuthorityRESTClient(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("building service for certificate authority: %w", err)
	}

	return service, nil
}
