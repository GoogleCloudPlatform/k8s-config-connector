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
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
)

type NetworkServicesServer struct {
	*MockService
	pb.UnimplementedNetworkServicesServer
}

func (s *NetworkServicesServer) ListMeshes(ctx context.Context, req *pb.ListMeshesRequest) (*pb.ListMeshesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeshes not implemented")
}

func (s *NetworkServicesServer) GetMesh(ctx context.Context, req *pb.GetMeshRequest) (*pb.Mesh, error) {
	name, err := s.parseMeshName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Mesh{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}
func (s *NetworkServicesServer) CreateMesh(ctx context.Context, req *pb.CreateMeshRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/meshes/" + req.MeshId
	name, err := s.parseMeshName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := ProtoClone(req.Mesh)
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
		result := ProtoClone(obj)
		result.SelfLink = "" // Not populated here
		return result, nil
	})
}

func (s *NetworkServicesServer) UpdateMesh(ctx context.Context, req *pb.UpdateMeshRequest) (*longrunning.Operation, error) {
	reqName := req.GetMesh().GetName()

	name, err := s.parseMeshName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Mesh{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()

	// Field mask is used to specify the fields to be overwritten in the
	// Mesh resource by the update.
	// The fields specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the mask. If the
	// user does not provide a mask then all fields will be overwritten.
	paths := req.GetUpdateMask().GetPaths()
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetMesh().GetDescription()
		case "interceptionPort":
			obj.InterceptionPort = req.GetMesh().GetInterceptionPort()
		case "labels":
			obj.Labels = req.GetMesh().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.NewLRO(ctx)
}

func (s *NetworkServicesServer) DeleteMesh(ctx context.Context, req *pb.DeleteMeshRequest) (*longrunning.Operation, error) {
	name, err := s.parseMeshName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	deletedObj := &pb.Mesh{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

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

type meshName struct {
	Project  *projects.ProjectData
	Location string
	MeshName string
}

func (n *meshName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/meshes/" + n.MeshName
}

// parseMeshName parses a string into a meshName.
// The expected form is `projects/*/locations/global/meshes/*`.
func (s *MockService) parseMeshName(name string) (*meshName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[3] == "global" && tokens[4] == "meshes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &meshName{
			Project:  project,
			Location: "global",
			MeshName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
