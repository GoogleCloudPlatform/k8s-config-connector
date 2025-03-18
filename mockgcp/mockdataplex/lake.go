// Copyright 2024 Google LLC
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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Lake

package mockdataplex

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DataplexV1) GetLake(ctx context.Context, req *pb.GetLakeRequest) (*pb.Lake, error) {
	name, err := s.parseLakeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Lake{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataplexV1) CreateLake(ctx context.Context, req *pb.CreateLakeRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/lakes/" + req.LakeId
	name, err := s.parseLakeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Lake).(*pb.Lake)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "lake-" + name.LakeID // TODO: maybe a proper random value?
	obj.State = pb.State_ACTIVE
	s.populateDefaultsForLake(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) populateDefaultsForLake(obj *pb.Lake) {
}

func (s *DataplexV1) UpdateLake(ctx context.Context, req *pb.UpdateLakeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLakeName(req.GetLake().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Lake{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
		EndTime:    timestamppb.Now(),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) DeleteLake(ctx context.Context, req *pb.DeleteLakeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLakeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Lake{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type lakeName struct {
	Project    *projects.ProjectData
	Location   string
	LakeID     string
	ResourceID string // Needed for nested resources
}

func (n *lakeName) String() string {
	base := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", n.Project.ID, n.Location, n.LakeID)
	if n.ResourceID != "" {
		return base + "/" + n.ResourceID
	}
	return base
}

// parseLakeName parses a string into a lakeName.
// The expected form is `projects/*/locations/*/lakes/*`.
func (s *MockService) parseLakeName(name string) (*lakeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &lakeName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


