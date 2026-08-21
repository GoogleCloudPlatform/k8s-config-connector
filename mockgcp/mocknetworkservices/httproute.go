// Copyright 2026 Google LLC
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

// +tool:mockgcp-support
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.HttpRoute

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetHttpRoute(ctx context.Context, req *pb.GetHttpRouteRequest) (*pb.HttpRoute, error) {
	name, err := s.parseHttpRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HttpRoute{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListHttpRoutes(ctx context.Context, req *pb.ListHttpRoutesRequest) (*pb.ListHttpRoutesResponse, error) {
	response := &pb.ListHttpRoutesResponse{}

	prefixName, err := s.parseHttpRouteName(req.Parent + "/httpRoutes/placeholder-name")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder-name")

	findKind := (&pb.HttpRoute{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		httpRoute := obj.(*pb.HttpRoute)
		response.HttpRoutes = append(response.HttpRoutes, httpRoute)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *NetworkServicesServer) CreateHttpRoute(ctx context.Context, req *pb.CreateHttpRouteRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/httpRoutes/" + req.HttpRouteId
	name, err := s.parseHttpRouteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.HttpRoute)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.SelfLink = buildSelfLink(ctx, fqn)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		result.SelfLink = "" // Not populated here
		return result, nil
	})
}

func (s *NetworkServicesServer) UpdateHttpRoute(ctx context.Context, req *pb.UpdateHttpRouteRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetHttpRoute().GetName()

	name, err := s.parseHttpRouteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.HttpRoute{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		now := time.Now()
		req.HttpRoute.CreateTime = obj.CreateTime
		req.HttpRoute.UpdateTime = timestamppb.New(now)
		req.HttpRoute.Name = obj.Name
		obj = req.HttpRoute
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetHttpRoute().GetLabels()
			case "description":
				obj.Description = req.GetHttpRoute().GetDescription()
			case "hostnames":
				obj.Hostnames = req.GetHttpRoute().GetHostnames()
			case "gateways":
				obj.Gateways = req.GetHttpRoute().GetGateways()
			case "meshes":
				obj.Meshes = req.GetHttpRoute().GetMeshes()
			case "rules":
				obj.Rules = req.GetHttpRoute().GetRules()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(time.Now())
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *NetworkServicesServer) DeleteHttpRoute(ctx context.Context, req *pb.DeleteHttpRouteRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseHttpRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.HttpRoute{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := &emptypb.Empty{}
		return result, nil
	})
}

type httpRouteName struct {
	Project       *projects.ProjectData
	Location      string
	HttpRouteName string
}

func (n *httpRouteName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/httpRoutes/" + n.HttpRouteName
}

// parseHttpRouteName parses a string into an httpRouteName.
// The expected form is `projects/*/locations/*/httpRoutes/*`.
func (s *NetworkServicesServer) parseHttpRouteName(name string) (*httpRouteName, error) {
	// DCL has a bug where it sometimes sends names like "projects/projects/my-project/..."
	if strings.HasPrefix(name, "projects/projects/") {
		name = strings.TrimPrefix(name, "projects/")
	}

	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "httpRoutes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &httpRouteName{
			Project:       project,
			Location:      tokens[3],
			HttpRouteName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
