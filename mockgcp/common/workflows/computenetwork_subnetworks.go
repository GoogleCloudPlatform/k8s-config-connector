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

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
)

func (e *Engine) subnetworksClient(ctx context.Context) (*compute.SubnetworksClient, error) {
	var options []option.ClientOption

	options = append(options, option.WithHTTPClient(e.httpClient))

	c, err := compute.NewSubnetworksRESTClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("error building compute subnetworks client: %w", err)
	}
	return c, nil
}

func (e *Engine) CreateComputeNetworkSubnetworks(ctx context.Context, projectID string, networkID string) error {
	subnetsClient, err := e.subnetworksClient(ctx)
	if err != nil {
		return err
	}

	regions := []string{
		"africa-south1",
		"asia-east1",
		"asia-east2",
		"asia-northeast1",
		"asia-northeast2",
		"asia-northeast3",
		"asia-south1",
		"asia-south2",
		"asia-southeast1",
		"asia-southeast2",
		"australia-southeast1",
		"australia-southeast2",
		"europe-central2",
		"europe-north1",
		"europe-southwest1",
		"europe-west1",
		"europe-west10",
		"europe-west12",
		"europe-west2",
		"europe-west3",
		"europe-west4",
		"europe-west6",
		"europe-west8",
		"europe-west9",
		"me-central1",
		"me-central2",
		"me-west1",
		"northamerica-northeast1",
		"northamerica-northeast2",
		"southamerica-east1",
		"southamerica-west1",
		"us-central1",
		"us-east1",
		"us-east4",
		"us-east5",
		"us-east7",
		"us-south1",
		"us-west1",
		"us-west2",
		"us-west3",
		"us-west4",
		"us-west8",
	}
	for _, region := range regions {
		subnet := &computepb.Subnetwork{
			Name:    PtrTo(networkID),
			Region:  PtrTo(region),
			Network: PtrTo(fmt.Sprintf("projects/%s/global/networks/%s", projectID, networkID)),
		}
		req := &computepb.InsertSubnetworkRequest{
			Project:            projectID,
			Region:             region,
			SubnetworkResource: subnet,
		}
		op, err := subnetsClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating automatic subnet %v: %w", req, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for subnet creation: %w", err)
		}
	}

	return nil
}
