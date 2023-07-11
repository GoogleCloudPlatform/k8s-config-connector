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

	pb "google.golang.org/genproto/googleapis/cloud/networkservices/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "mesh %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading mesh: %v", err)
		}
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

	obj := proto.Clone(req.Mesh).(*pb.Mesh)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating mesh: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *NetworkServicesServer) UpdateMesh(ctx context.Context, req *pb.UpdateMeshRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMesh not implemented")
}

func (s *NetworkServicesServer) DeleteMesh(ctx context.Context, req *pb.DeleteMeshRequest) (*longrunning.Operation, error) {
	name, err := s.parseMeshName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	meshKind := (&pb.Mesh{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, meshKind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "mesh %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting mesh: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
