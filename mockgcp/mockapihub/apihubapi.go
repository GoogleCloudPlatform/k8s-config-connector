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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ApiHubServer struct {
	*MockService
	pb.UnimplementedApiHubServer
}

func (s *ApiHubServer) GetApi(ctx context.Context, req *pb.GetApiRequest) (*pb.Api, error) {
	name, err := s.parseApiName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Api{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) CreateApi(ctx context.Context, req *pb.CreateApiRequest) (*pb.Api, error) {
	reqName := req.Parent + "/apis/" + req.ApiId
	name, err := s.parseApiName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Api).(*pb.Api)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) UpdateApi(ctx context.Context, req *pb.UpdateApiRequest) (*pb.Api, error) {
	apiName := req.GetApi().GetName()

	name, err := s.parseApiName(apiName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Api{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"display_name", "description", "owner", "documentation",
			"target_user", "team", "business_unit", "maturity_level", "attributes",
		}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetApi().GetDisplayName()
		case "description":
			obj.Description = req.GetApi().GetDescription()
		case "owner":
			obj.Owner = req.GetApi().GetOwner()
		case "documentation":
			obj.Documentation = req.GetApi().GetDocumentation()
		case "target_user", "targetUser":
			obj.TargetUser = req.GetApi().GetTargetUser()
		case "team":
			obj.Team = req.GetApi().GetTeam()
		case "business_unit", "businessUnit":
			obj.BusinessUnit = req.GetApi().GetBusinessUnit()
		case "maturity_level", "maturityLevel":
			obj.MaturityLevel = req.GetApi().GetMaturityLevel()
		case "attributes":
			obj.Attributes = req.GetApi().GetAttributes()
		case "api_requirements", "apiRequirements":
			obj.ApiRequirements = req.GetApi().GetApiRequirements()
		case "api_functional_requirements", "apiFunctionalRequirements":
			obj.ApiFunctionalRequirements = req.GetApi().GetApiFunctionalRequirements()
		case "api_technical_requirements", "apiTechnicalRequirements":
			obj.ApiTechnicalRequirements = req.GetApi().GetApiTechnicalRequirements()
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) DeleteApi(ctx context.Context, req *pb.DeleteApiRequest) (*emptypb.Empty, error) {
	name, err := s.parseApiName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Api{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
