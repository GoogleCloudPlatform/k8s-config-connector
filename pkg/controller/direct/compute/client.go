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

package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"google.golang.org/api/compute/v1"
	api "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type gcpClient struct {
	config controller.Config

	gcp *compute.Service
}

func newGCPClient(ctx context.Context, config *controller.Config) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	client, err := gcpClient.client(ctx)
	if err != nil {
		return nil, err
	}
	gcpClient.gcp = client
	return gcpClient, nil
}

func (a *gcpClient) waitForGlobalOperation(ctx context.Context, projectID string, opName string) (*compute.Operation, error) {
	// TODO: Use server-side wait
	// completed, err := a.gcp.GlobalOperations.Wait(a.projectID, op.Name).Context(ctx).Do()
	// if err != nil {
	// 	return fmt.Errorf("waiting for network creation: %w", err)
	// }
	// TODO: need to handle more Wait conditions?

	// TODO: Backoff?
	// TODO: Early-return and let the controller do it?

	timeout := 5 * time.Minute // TODO: Configurable?
	timeoutAt := time.Now().Add(timeout)
	for {
		op, err := a.gcp.GlobalOperations.Get(projectID, opName).Context(ctx).Do()
		if err != nil {
			return nil, fmt.Errorf("error getting operation: %w", err)
		}
		if op.Status == "DONE" {
			return op, nil
		}
		if time.Now().After(timeoutAt) {
			return nil, fmt.Errorf("timeout waiting for operation")
		}

		// TODO: Backoff?
		time.Sleep(2 * time.Second)
	}
}

func (m *gcpClient) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		// TODO: Set UserAgent in this scenario (error is: WithHTTPClient is incompatible with gRPC dial options)
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute client: %w", err)
	}
	return gcpClient, err
}
