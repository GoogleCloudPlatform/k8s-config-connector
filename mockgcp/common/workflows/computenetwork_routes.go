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
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
)

func (e *Engine) routesClient(ctx context.Context) (*compute.RoutesClient, error) {
	var options []option.ClientOption

	options = append(options, option.WithHTTPClient(e.httpClient))

	c, err := compute.NewRoutesRESTClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("error building compute routes client: %w", err)
	}
	return c, nil
}

func computeHashSuffix(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func (e *Engine) CreateComputeNetworkRoutes(ctx context.Context, projectID string, networkID string, networkSelfLink string) error {
	subnetsClient, err := e.subnetworksClient(ctx)
	if err != nil {
		return err
	}

	routesClient, err := e.routesClient(ctx)
	if err != nil {
		return err
	}

	// Add default route
	{
		hash := computeHashSuffix(projectID + "::" + networkID)
		route := &computepb.Route{
			Description:    PtrTo("Default route to the Internet."),
			DestRange:      PtrTo("0.0.0.0/0"),
			Name:           PtrTo("default-route-" + hash),
			Network:        PtrTo(networkSelfLink),
			NextHopGateway: PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/gateways/%s", projectID, "default-internet-gateway")),
			Priority:       PtrTo(uint32(1000)),
		}
		req := &computepb.InsertRouteRequest{
			Project:       projectID,
			RouteResource: route,
		}
		op, err := routesClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating automatic default route %v: %w", req, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for route creation: %w", err)
		}
	}

	// Add subnet routes
	for scope, err := range subnetsClient.AggregatedList(ctx, &computepb.AggregatedListSubnetworksRequest{
		Project: projectID,
	}).All() {
		if err != nil {
			return fmt.Errorf("listing subnets: %w", err)
		}

		for _, subnet := range scope.Value.Subnetworks {
			if subnet.GetNetwork() != networkSelfLink {
				continue
			}

			hash := computeHashSuffix(ValueOf(subnet.SelfLink))
			route := &computepb.Route{
				Description:    PtrTo(fmt.Sprintf("Default local route to the subnetwork %s.", ValueOf(subnet.IpCidrRange))),
				DestRange:      subnet.IpCidrRange,
				Name:           PtrTo("default-route-r-" + hash),
				Network:        subnet.Network,
				NextHopNetwork: subnet.Network,
				Priority:       PtrTo(uint32(0)),
			}
			req := &computepb.InsertRouteRequest{
				Project:       projectID,
				RouteResource: route,
			}
			op, err := routesClient.Insert(ctx, req)
			if err != nil {
				return fmt.Errorf("creating automatic route for subnet %v %v: %w", subnet, req, err)
			}
			if err := op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting for route creation: %w", err)
			}
		}
	}

	return nil
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
