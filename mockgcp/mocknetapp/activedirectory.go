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
// proto.service: google.cloud.netapp.v1.NetApp
// proto.message: google.cloud.netapp.v1.ActiveDirectory

package mocknetapp

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

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *backupVaultsService) GetActiveDirectory(ctx context.Context, req *pb.GetActiveDirectoryRequest) (*pb.ActiveDirectory, error) {
	name, err := s.parseActiveDirectoryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ActiveDirectory{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", fqn)
		}
		return nil, err
	}

	s.populateDefaultsForActiveDirectory(obj)

	return obj, nil
}

func (s *backupVaultsService) ListActiveDirectories(ctx context.Context, req *pb.ListActiveDirectoriesRequest) (*pb.ListActiveDirectoriesResponse, error) {
	parent := req.GetParent()
	response := &pb.ListActiveDirectoriesResponse{}

	findPrefix := fmt.Sprintf("%s/activeDirectories/", parent)

	activeDirectoryKind := (&pb.ActiveDirectory{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, activeDirectoryKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cloned := proto.Clone(obj).(*pb.ActiveDirectory)
		s.populateDefaultsForActiveDirectory(cloned)
		response.ActiveDirectories = append(response.ActiveDirectories, cloned)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *backupVaultsService) CreateActiveDirectory(ctx context.Context, req *pb.CreateActiveDirectoryRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/activeDirectories/%s", req.GetParent(), req.GetActiveDirectoryId())
	name, err := s.parseActiveDirectoryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetActiveDirectory()).(*pb.ActiveDirectory)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.ActiveDirectory_READY
	obj.StateDetails = "Credentials saved and available for use"
	obj.Password = "******************"

	s.populateDefaultsForActiveDirectory(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) UpdateActiveDirectory(ctx context.Context, req *pb.UpdateActiveDirectoryRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseActiveDirectoryName(req.GetActiveDirectory().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ActiveDirectory{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	if err := fields.UpdateByFieldMask(obj, req.GetActiveDirectory(), req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	obj.Password = "******************"
	s.populateDefaultsForActiveDirectory(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) populateDefaultsForActiveDirectory(obj *pb.ActiveDirectory) {
	if obj.OrganizationalUnit == "" {
		obj.OrganizationalUnit = "CN=Computers"
	}
}

func (s *backupVaultsService) DeleteActiveDirectory(ctx context.Context, req *pb.DeleteActiveDirectoryRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseActiveDirectoryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ActiveDirectory{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type activeDirectoryName struct {
	Project           *projects.ProjectData
	Location          string
	ActiveDirectoryID string
}

func (n *activeDirectoryName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", n.Project.ID, n.Location, n.ActiveDirectoryID)
}

// parseActiveDirectoryName parses a string into a activeDirectoryName.
// The expected form is `projects/*/locations/*/activeDirectories/*`.
func (s *MockService) parseActiveDirectoryName(name string) (*activeDirectoryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "activeDirectories" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &activeDirectoryName{
			Project:           project,
			Location:          tokens[3],
			ActiveDirectoryID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
