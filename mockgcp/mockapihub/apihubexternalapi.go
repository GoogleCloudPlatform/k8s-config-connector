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

package mockapihub

import (
	"context"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func populateExternalApiAttributes(obj *pb.ExternalApi) {
	if obj == nil {
		return
	}
	for k, v := range obj.Attributes {
		if v != nil {
			v.Attribute = k
		}
	}
}

func (s *ApiHubServer) GetExternalApi(ctx context.Context, req *pb.GetExternalApiRequest) (*pb.ExternalApi, error) {
	name, err := s.parseExternalApiName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExternalApi{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	populateExternalApiAttributes(obj)
	return obj, nil
}

func (s *ApiHubServer) CreateExternalApi(ctx context.Context, req *pb.CreateExternalApiRequest) (*pb.ExternalApi, error) {
	reqName := req.Parent + "/externalApis/" + req.ExternalApiId
	name, err := s.parseExternalApiName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ExternalApi).(*pb.ExternalApi)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	populateExternalApiAttributes(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) UpdateExternalApi(ctx context.Context, req *pb.UpdateExternalApiRequest) (*pb.ExternalApi, error) {
	extApiName := req.GetExternalApi().GetName()

	name, err := s.parseExternalApiName(extApiName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExternalApi{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"display_name", "description", "endpoints", "paths", "documentation", "attributes",
		}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetExternalApi().GetDisplayName()
		case "description":
			obj.Description = req.GetExternalApi().GetDescription()
		case "endpoints":
			obj.Endpoints = req.GetExternalApi().GetEndpoints()
		case "paths":
			obj.Paths = req.GetExternalApi().GetPaths()
		case "documentation":
			obj.Documentation = req.GetExternalApi().GetDocumentation()
		case "attributes":
			obj.Attributes = req.GetExternalApi().GetAttributes()
		}
	}

	obj.UpdateTime = timestamppb.Now()

	populateExternalApiAttributes(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) DeleteExternalApi(ctx context.Context, req *pb.DeleteExternalApiRequest) (*emptypb.Empty, error) {
	name, err := s.parseExternalApiName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExternalApi{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
