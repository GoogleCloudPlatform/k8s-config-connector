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
// proto.service: google.cloud.aiplatform.v1.SpecialistPoolService
// proto.message: google.cloud.aiplatform.v1.SpecialistPool

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type specialistPoolService struct {
	*MockService
	pb.UnimplementedSpecialistPoolServiceServer
}

func (s *specialistPoolService) GetSpecialistPool(ctx context.Context, req *pb.GetSpecialistPoolRequest) (*pb.SpecialistPool, error) {
	name, err := s.parseSpecialistPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.SpecialistPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.CloneOf(obj)
	result.Name = name.String()
	result.SpecialistManagersCount = int32(len(result.SpecialistManagerEmails))
	return result, nil
}

func (s *specialistPoolService) CreateSpecialistPool(ctx context.Context, req *pb.CreateSpecialistPoolRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/specialistPools/" + id
	name, err := s.parseSpecialistPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.CloneOf(req.SpecialistPool)
	obj.Name = name.String()
	obj.SpecialistManagersCount = int32(len(obj.SpecialistManagerEmails))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateSpecialistPoolOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.StringWithProjectNumber()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		result.Name = name.StringWithProjectNumber()
		result.SpecialistManagersCount = int32(len(result.SpecialistManagerEmails))
		return result, nil
	})
}

func (s *specialistPoolService) UpdateSpecialistPool(ctx context.Context, req *pb.UpdateSpecialistPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseSpecialistPoolName(req.GetSpecialistPool().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.SpecialistPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "displayName", "display_name":
			obj.DisplayName = req.GetSpecialistPool().GetDisplayName()
		case "specialistManagerEmails", "specialist_manager_emails":
			obj.SpecialistManagerEmails = req.GetSpecialistPool().GetSpecialistManagerEmails()
		case "specialistWorkerEmails", "specialist_worker_emails":
			obj.SpecialistWorkerEmails = req.GetSpecialistPool().GetSpecialistWorkerEmails()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.SpecialistManagersCount = int32(len(obj.SpecialistManagerEmails))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateSpecialistPoolOperationMetadata{
		SpecialistPool: name.StringWithProjectNumber(),
	}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.StringWithProjectNumber()
	result := proto.CloneOf(obj)
	result.Name = name.StringWithProjectNumber()
	result.SpecialistManagersCount = int32(len(result.SpecialistManagerEmails))
	return s.operations.DoneLRO(ctx, opPrefix, op, result)
}

func (s *specialistPoolService) DeleteSpecialistPool(ctx context.Context, req *pb.DeleteSpecialistPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseSpecialistPoolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	deleted := &pb.SpecialistPool{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *specialistPoolService) ListSpecialistPools(ctx context.Context, req *pb.ListSpecialistPoolsRequest) (*pb.ListSpecialistPoolsResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}
	projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/specialistPools/", project.Number, location)

	var specialistPools []*pb.SpecialistPool
	specialistPoolKind := (&pb.SpecialistPool{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, specialistPoolKind, storage.ListOptions{}, func(obj proto.Message) error {
		sp := obj.(*pb.SpecialistPool)
		if strings.HasPrefix(sp.GetName(), findPrefix) {
			specialistPools = append(specialistPools, sp)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListSpecialistPoolsResponse{
		SpecialistPools: specialistPools,
	}, nil
}

type SpecialistPoolName struct {
	Project        *projects.ProjectData
	Location       string
	SpecialistPool string
}

func (n *SpecialistPoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/specialistPools/%s", n.Project.ID, n.Location, n.SpecialistPool)
}

func (n *SpecialistPoolName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/specialistPools/%s", n.Project.Number, n.Location, n.SpecialistPool)
}

func (n *SpecialistPoolName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/specialistPools/%s", n.Project.Number, n.Location, n.SpecialistPool)
}

// parseSpecialistPoolName parses a string into a SpecialistPoolName.
// The expected form is projects/<projectID_or_number>/locations/<location>/specialistPools/<specialistPool>
func (s *MockService) parseSpecialistPoolName(name string) (*SpecialistPoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "specialistPools" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		specialistPoolID := tokens[5]
		if !isNumeric(specialistPoolID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid specialist pool id: %s", specialistPoolID)
		}

		return &SpecialistPoolName{
			Project:        project,
			Location:       tokens[3],
			SpecialistPool: specialistPoolID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
