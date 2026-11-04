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

package mockbigqueryanalyticshub

import (
	"context"
	"fmt"
	"strings"

	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *analyticsHubServer) GetQueryTemplate(ctx context.Context, request *pb.GetQueryTemplateRequest) (*pb.QueryTemplate, error) {
	name, err := a.parseQueryTemplateID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.QueryTemplate{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Failed to find query template: %v", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) CreateQueryTemplate(ctx context.Context, request *pb.CreateQueryTemplateRequest) (*pb.QueryTemplate, error) {
	reqName := request.Parent + "/queryTemplates/" + request.QueryTemplateId
	name, err := a.parseQueryTemplateID(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.CloneOf(request.QueryTemplate)
	obj.Name = fqn

	a.populateQueryTemplateDefaults(obj, name)

	if err := a.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) populateQueryTemplateDefaults(obj *pb.QueryTemplate, name *queryTemplateName) {
	now := timestamppb.Now()
	if obj.CreateTime == nil {
		obj.CreateTime = now
	}
	if obj.UpdateTime == nil {
		obj.UpdateTime = now
	}
	if obj.State == pb.QueryTemplate_STATE_UNSPECIFIED {
		obj.State = pb.QueryTemplate_DRAFTED
	}
}

func (a *analyticsHubServer) UpdateQueryTemplate(ctx context.Context, request *pb.UpdateQueryTemplateRequest) (*pb.QueryTemplate, error) {
	name, err := a.parseQueryTemplateID(request.GetQueryTemplate().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.QueryTemplate{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range request.GetUpdateMask().Paths {
		switch path {
		case "displayName":
			obj.DisplayName = request.QueryTemplate.GetDisplayName()
		case "description":
			obj.Description = request.QueryTemplate.GetDescription()
		case "proposer":
			obj.Proposer = request.QueryTemplate.GetProposer()
		case "primaryContact":
			obj.PrimaryContact = request.QueryTemplate.GetPrimaryContact()
		case "documentation":
			obj.Documentation = request.QueryTemplate.GetDocumentation()
		case "routine":
			obj.Routine = request.QueryTemplate.GetRoutine()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := a.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) DeleteQueryTemplate(ctx context.Context, request *pb.DeleteQueryTemplateRequest) (*empty.Empty, error) {
	name, err := a.parseQueryTemplateID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.QueryTemplate{}
	if err := a.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type queryTemplateName struct {
	Project         *projects.ProjectData
	Location        string
	DataExchangeID  string
	QueryTemplateID string
}

func (n *queryTemplateName) String() string {
	location := strings.ToLower(n.Location)
	return fmt.Sprintf("projects/%v/locations/%v/dataExchanges/%v/queryTemplates/%v", n.Project.Number, location, n.DataExchangeID, n.QueryTemplateID)
}

func (s *MockService) parseQueryTemplateID(name string) (*queryTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataExchanges" && tokens[6] == "queryTemplates" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &queryTemplateName{
			Project:         project,
			Location:        tokens[3],
			DataExchangeID:  tokens[5],
			QueryTemplateID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
