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

package workflows

import (
	"context"
	"fmt"
	"net/http"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
)

type Engine struct {
	httpClient *http.Client
}

func NewEngine(httpClient *http.Client) (*Engine, error) {
	return &Engine{httpClient: httpClient}, nil
}

func (e *Engine) computeClient(ctx context.Context) (*compute.NetworksClient, error) {
	var options []option.ClientOption

	options = append(options, option.WithHTTPClient(e.httpClient))

	c, err := compute.NewNetworksRESTClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("error building compute networks client: %w", err)
	}
	return c, nil
}

func (e *Engine) PopulateNewProject(ctx context.Context, projectID string) error {
	networksClient, err := e.computeClient(ctx)
	if err != nil {
		return err
	}
	network := &computepb.Network{
		Name:        PtrTo("default"),
		Description: PtrTo("Default network for the project"),
	}
	req := &computepb.InsertNetworkRequest{
		Project:         projectID,
		NetworkResource: network,
	}
	op, err := networksClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating network: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for network creation: %w", err)
	}

	return nil
}
