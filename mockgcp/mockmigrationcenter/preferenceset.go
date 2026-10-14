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

package mockmigrationcenter

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/migrationcenter/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type preferenceSetName struct {
	Project         *projects.ProjectData
	Location        string
	PreferenceSetID string
}

func (n *preferenceSetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/preferenceSets/%s", n.Project.ID, n.Location, n.PreferenceSetID)
}

func (s *MockService) parsePreferenceSetName(name string) (*preferenceSetName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "preferenceSets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &preferenceSetName{
			Project:         project,
			Location:        tokens[3],
			PreferenceSetID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MigrationCenterV1) GetPreferenceSet(ctx context.Context, req *pb.GetPreferenceSetRequest) (*pb.PreferenceSet, error) {
	name, err := s.parsePreferenceSetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PreferenceSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MigrationCenterV1) CreatePreferenceSet(ctx context.Context, req *pb.CreatePreferenceSetRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/preferenceSets/%s", req.GetParent(), req.GetPreferenceSetId())
	name, err := s.parsePreferenceSetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetPreferenceSet()).(*pb.PreferenceSet)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

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
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MigrationCenterV1) UpdatePreferenceSet(ctx context.Context, req *pb.UpdatePreferenceSetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePreferenceSetName(req.GetPreferenceSet().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.PreferenceSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetPreferenceSet().GetDisplayName()
		obj.Description = req.GetPreferenceSet().GetDescription()
		obj.VirtualMachinePreferences = req.GetPreferenceSet().GetVirtualMachinePreferences()
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetPreferenceSet().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetPreferenceSet().GetDisplayName()
			case "virtual_machine_preferences", "virtualMachinePreferences":
				obj.VirtualMachinePreferences = req.GetPreferenceSet().GetVirtualMachinePreferences()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported in mock", path)
			}
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MigrationCenterV1) DeletePreferenceSet(ctx context.Context, req *pb.DeletePreferenceSetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePreferenceSetName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.PreferenceSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MigrationCenterV1) ListPreferenceSets(ctx context.Context, req *pb.ListPreferenceSetsRequest) (*pb.ListPreferenceSetsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is invalid", req.GetParent())
	}

	var preferenceSets []*pb.PreferenceSet
	listKind := (&pb.PreferenceSet{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		prefSet := obj.(*pb.PreferenceSet)
		name, err := s.parsePreferenceSetName(prefSet.GetName())
		if err == nil && name.Project.ID == tokens[1] && name.Location == tokens[3] {
			preferenceSets = append(preferenceSets, prefSet)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListPreferenceSetsResponse{
		PreferenceSets: preferenceSets,
	}, nil
}
