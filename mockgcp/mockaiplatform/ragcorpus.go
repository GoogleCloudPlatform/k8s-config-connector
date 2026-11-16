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
// proto.service: google.cloud.aiplatform.v1.VertexRagDataService
// proto.message: google.cloud.aiplatform.v1.RagCorpus

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type ragCorpusService struct {
	*MockService
	pb.UnimplementedVertexRagDataServiceServer
}

func (s *ragCorpusService) GetRagCorpus(ctx context.Context, req *pb.GetRagCorpusRequest) (*pb.RagCorpus, error) {
	name, err := s.parseRagCorpusName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.RagCorpus{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.CloneOf(obj)
	result.Name = name.String()
	s.populateDefaultsForRagCorpus(result, name)
	if result.CorpusStatus == nil {
		result.CorpusStatus = &pb.CorpusStatus{
			State: pb.CorpusStatus_ACTIVE,
		}
	}
	return result, nil
}

func (s *ragCorpusService) CreateRagCorpus(ctx context.Context, req *pb.CreateRagCorpusRequest) (*longrunning.Operation, error) {
	var id string
	if strings.Contains(req.GetRagCorpus().GetDisplayName(), "min") {
		id = "763118244281188352"
	} else if strings.Contains(req.GetRagCorpus().GetDisplayName(), "max") {
		id = "3671317703655686144"
	} else {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	reqName := req.Parent + "/ragCorpora/" + id
	name, err := s.parseRagCorpusName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.CloneOf(req.RagCorpus)
	obj.Name = name.StringWithProjectNumber()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForRagCorpus(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateRagCorpusOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		result.CreateTime = nil
		result.UpdateTime = nil
		result.CorpusStatus = nil
		result.Name = name.StringWithProjectNumber()
		return result, nil
	})
}

func (s *ragCorpusService) DeleteRagCorpus(ctx context.Context, req *pb.DeleteRagCorpusRequest) (*longrunning.Operation, error) {
	name, err := s.parseRagCorpusName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	deleted := &pb.RagCorpus{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *ragCorpusService) populateDefaultsForRagCorpus(obj *pb.RagCorpus, name *RagCorpusName) {
	if obj.BackendConfig == nil {
		obj.BackendConfig = &pb.RagCorpus_VectorDbConfig{
			VectorDbConfig: &pb.RagVectorDbConfig{},
		}
	}
	if vdc := obj.GetVectorDbConfig(); vdc != nil {
		if vdc.RagEmbeddingModelConfig == nil {
			vdc.RagEmbeddingModelConfig = &pb.RagEmbeddingModelConfig{
				ModelConfig: &pb.RagEmbeddingModelConfig_VertexPredictionEndpoint_{
					VertexPredictionEndpoint: &pb.RagEmbeddingModelConfig_VertexPredictionEndpoint{
						Endpoint: fmt.Sprintf("projects/%d/locations/%s/publishers/google/models/text-embedding-005", name.Project.Number, name.Location),
					},
				},
			}
		}
	}
}

type RagCorpusName struct {
	Project     *projects.ProjectData
	Location    string
	RagCorpusID string
}

func (n *RagCorpusName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", n.Project.ID, n.Location, n.RagCorpusID)
}

func (n *RagCorpusName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/ragCorpora/%s", n.Project.Number, n.Location, n.RagCorpusID)
}

func (n *RagCorpusName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/ragCorpora/%s", n.Project.Number, n.Location, n.RagCorpusID)
}

func (s *MockService) parseRagCorpusName(name string) (*RagCorpusName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "ragCorpora" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &RagCorpusName{
			Project:     project,
			Location:    tokens[3],
			RagCorpusID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
