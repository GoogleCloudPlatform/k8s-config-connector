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

package mocknetworkservices

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
)

func (s *NetworkServicesServer) ListHttpRoutes(ctx context.Context, req *pb.ListHttpRoutesRequest) (*pb.ListHttpRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHttpRoutes not implemented")
}

func (s *NetworkServicesServer) GetHttpRoute(ctx context.Context, req *pb.GetHttpRouteRequest) (*pb.HttpRoute, error) {
	name, err := s.parseHttpRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.HttpRoute{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
func (s *NetworkServicesServer) CreateHttpRoute(ctx context.Context, req *pb.CreateHttpRouteRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/httpRoutes/" + req.HttpRouteId
	name, err := s.parseHttpRouteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.HttpRoute).(*pb.HttpRoute)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *NetworkServicesServer) UpdateHttpRoute(ctx context.Context, req *pb.UpdateHttpRouteRequest) (*longrunning.Operation, error) {
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

	// Field mask is used to specify the fields to be overwritten in the
	// HttpRoute resource by the update.
	// The fields specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the mask. If the
	// user does not provide a mask then all fields will be overwritten.
	paths := req.GetUpdateMask().GetPaths()
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetHttpRoute().GetDescription()
		case "labels":
			obj.Labels = req.GetHttpRoute().GetLabels()
		case "hostnames":
			obj.Hostnames = req.GetHttpRoute().GetHostnames()
		case "rules":
			obj.Rules = req.GetHttpRoute().GetRules()
		case "gateways":
			obj.Gateways = req.GetHttpRoute().GetGateways()
		case "meshes":
			obj.Meshes = req.GetHttpRoute().GetMeshes()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.NewLRO(ctx)
}

func (s *NetworkServicesServer) DeleteHttpRoute(ctx context.Context, req *pb.DeleteHttpRouteRequest) (*longrunning.Operation, error) {
	name, err := s.parseHttpRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.HttpRoute{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type httpRouteName struct {
	Project       *projects.ProjectData
	Location      string
	HttpRouteName string
}

func (n *httpRouteName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/httpRoutes/" + n.HttpRouteName
}

// parseHttpRouteName parses a string into a httpRouteName.
// The expected form is `projects/*/locations/*/httpRoutes/*`.
func (s *MockService) parseHttpRouteName(name string) (*httpRouteName, error) {
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
