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

package mockdialogflow

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
)

type generatorsServer struct {
	*MockService
	pb.UnimplementedGeneratorsServer
}

type generatorName struct {
	Project   *projects.ProjectData
	Location  string
	Generator string
}

func (n *generatorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/generators/%s", n.Project.ID, n.Location, n.Generator)
}

func (s *MockService) parseGeneratorName(name string) (*generatorName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "generators" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &generatorName{
			Project:   project,
			Location:  tokens[3],
			Generator: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/generators/{generator}", name)
}

func (s *generatorsServer) GetGenerator(ctx context.Context, req *pb.GetGeneratorRequest) (*pb.Generator, error) {
	name, err := s.parseGeneratorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Generator{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Generator with ID %s not found in projects/%s.", name.Generator, name.Project.ID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *generatorsServer) CreateGenerator(ctx context.Context, req *pb.CreateGeneratorRequest) (*pb.Generator, error) {
	parent := req.GetParent() // e.g. "projects/{project}/locations/{location}"
	generatorID := req.GetGeneratorId()
	if generatorID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "generator_id must be provided")
	}

	name := fmt.Sprintf("%s/generators/%s", parent, generatorID)
	parsedName, err := s.parseGeneratorName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.CloneOf(req.GetGenerator())
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *generatorsServer) UpdateGenerator(ctx context.Context, req *pb.UpdateGeneratorRequest) (*pb.Generator, error) {
	reqName := req.GetGenerator().GetName()
	name, err := s.parseGeneratorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Generator{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Description = req.GetGenerator().GetDescription()
		obj.Context = req.GetGenerator().Context
		obj.InferenceParameter = req.GetGenerator().GetInferenceParameter()
		obj.TriggerEvent = req.GetGenerator().GetTriggerEvent()
		obj.FoundationModel = req.GetGenerator().FoundationModel
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetGenerator().GetDescription()
			case "free_form_context", "freeFormContext":
				obj.Context = req.GetGenerator().Context
			case "summarization_context", "summarizationContext":
				obj.Context = req.GetGenerator().Context
			case "inference_parameter", "inferenceParameter":
				obj.InferenceParameter = req.GetGenerator().GetInferenceParameter()
			case "trigger_event", "triggerEvent":
				obj.TriggerEvent = req.GetGenerator().GetTriggerEvent()
			case "published_model", "publishedModel":
				obj.FoundationModel = req.GetGenerator().FoundationModel
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *generatorsServer) DeleteGenerator(ctx context.Context, req *pb.DeleteGeneratorRequest) (*emptypb.Empty, error) {
	name, err := s.parseGeneratorName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.Generator{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
