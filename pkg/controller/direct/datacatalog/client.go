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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.service: google.cloud.datacatalog.v1.PolicyTagManager
// proto.service: google.cloud.datacatalog.v1.PolicyTagManagerSerialization

package datacatalog

import (
	"context"
	"fmt"

	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
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

func (m *gcpClient) newDataCatalogClient(ctx context.Context) (*datacatalog.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := datacatalog.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newPolicyTagManagerClient(ctx context.Context) (*datacatalog.PolicyTagManagerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := datacatalog.NewPolicyTagManagerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog policyTagManager client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newPolicyTagManagerSerializationClient(ctx context.Context) (*datacatalog.PolicyTagManagerSerializationClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := datacatalog.NewPolicyTagManagerSerializationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog policyTagManagerSerialization client: %w", err)
	}
	return client, err
}
