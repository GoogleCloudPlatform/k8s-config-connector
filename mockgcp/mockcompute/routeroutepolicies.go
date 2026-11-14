// Copyright 2026 High Five Labs
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

// A Cloud Router's route
// policies, as the Compute Engine API documents them (routers.getRoutePolicy,
// updateRoutePolicy, patchRoutePolicy, deleteRoutePolicy, listRoutePolicies).
// Each is kept beside its router; the fingerprint changes with each write.

package mockcompute

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *RoutersV1) routerForPolicy(ctx context.Context, project, region, router string) (*routerName, *pb.Router, error) {
	name, err := s.parseRouterName("projects/" + project + "/regions/" + region + "/routers/" + router)
	if err != nil {
		return nil, nil, err
	}
	obj := &pb.Router{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name.String())
		}
		return nil, nil, err
	}
	return name, obj, nil
}

func routePolicyKey(router *routerName, policy string) string {
	return router.String() + "/routePolicies/" + policy
}

func (s *RoutersV1) GetRoutePolicy(ctx context.Context, req *pb.GetRoutePolicyRouterRequest) (*pb.RoutersGetRoutePolicyResponse, error) {
	name, _, err := s.routerForPolicy(ctx, req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}
	obj := &pb.RoutePolicy{}
	if err := s.storage.Get(ctx, routePolicyKey(name, req.GetPolicy()), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// The real API answers a missing policy with 400, not 404.
			return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'policy': '%s'. The policy does not exist", req.GetPolicy())
		}
		return nil, err
	}
	return &pb.RoutersGetRoutePolicyResponse{Resource: obj}, nil
}

func (s *RoutersV1) writeRoutePolicy(ctx context.Context, project, region, router string, in *pb.RoutePolicy, merge bool) (*pb.Operation, error) {
	name, r, err := s.routerForPolicy(ctx, project, region, router)
	if err != nil {
		return nil, err
	}
	if in.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "route policy name is required")
	}
	key := routePolicyKey(name, in.GetName())
	obj := proto.Clone(in).(*pb.RoutePolicy)
	existing := &pb.RoutePolicy{}
	err = s.storage.Get(ctx, key, existing)
	switch {
	case err == nil && merge:
		proto.Merge(existing, in)
		obj = existing
	case err != nil && status.Code(err) != codes.NotFound:
		return nil, err
	}
	obj.Fingerprint = PtrTo(fmt.Sprintf("%x", time.Now().UnixNano()))
	if err := s.storage.Update(ctx, key, obj); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		if err := s.storage.Create(ctx, key, obj); err != nil {
			return nil, err
		}
	}
	op := &pb.Operation{
		TargetId:      r.Id,
		TargetLink:    r.SelfLink,
		OperationType: PtrTo("compute.routers.updateRoutePolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return r, nil
	})
}

func (s *RoutersV1) UpdateRoutePolicy(ctx context.Context, req *pb.UpdateRoutePolicyRouterRequest) (*pb.Operation, error) {
	return s.writeRoutePolicy(ctx, req.GetProject(), req.GetRegion(), req.GetRouter(), req.GetRoutePolicyResource(), false)
}

func (s *RoutersV1) PatchRoutePolicy(ctx context.Context, req *pb.PatchRoutePolicyRouterRequest) (*pb.Operation, error) {
	return s.writeRoutePolicy(ctx, req.GetProject(), req.GetRegion(), req.GetRouter(), req.GetRoutePolicyResource(), true)
}

func (s *RoutersV1) DeleteRoutePolicy(ctx context.Context, req *pb.DeleteRoutePolicyRouterRequest) (*pb.Operation, error) {
	name, r, err := s.routerForPolicy(ctx, req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, routePolicyKey(name, req.GetPolicy()), &pb.RoutePolicy{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The route policy '%s' was not found", req.GetPolicy())
		}
		return nil, err
	}
	op := &pb.Operation{
		TargetId:      r.Id,
		TargetLink:    r.SelfLink,
		OperationType: PtrTo("compute.routers.deleteRoutePolicy"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return r, nil
	})
}

func (s *RoutersV1) ListRoutePolicies(ctx context.Context, req *pb.ListRoutePoliciesRoutersRequest) (*pb.RoutersListRoutePolicies, error) {
	name, _, err := s.routerForPolicy(ctx, req.GetProject(), req.GetRegion(), req.GetRouter())
	if err != nil {
		return nil, err
	}
	resp := &pb.RoutersListRoutePolicies{Kind: PtrTo("compute#routersListRoutePolicies")}
	kind := (&pb.RoutePolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: name.String() + "/routePolicies/"}, func(obj proto.Message) error {
		resp.Result = append(resp.Result, obj.(*pb.RoutePolicy))
		return nil
	}); err != nil {
		return nil, err
	}
	return resp, nil
}
