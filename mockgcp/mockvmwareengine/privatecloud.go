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

//go:build mockgcp
// +build mockgcp

// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.PrivateCloud
package mockvmwareengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetPrivateCloud(ctx context.Context, req *pb.GetPrivateCloudRequest) (*pb.PrivateCloud, error) {
	name, err := s.parsePrivateCloudName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateCloud{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreatePrivateCloud(ctx context.Context, req *pb.CreatePrivateCloudRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/privateClouds/" + req.PrivateCloudId
	name, err := s.parsePrivateCloudName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetPrivateCloud()).(*pb.PrivateCloud)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.PrivateCloud_ACTIVE
	obj.Uid = "111111111111111111111"

	if obj.ManagementCluster != nil && obj.ManagementCluster.StretchedClusterConfig != nil {
		obj.Type = pb.PrivateCloud_STRETCHED
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		obj.UpdateTime = timestamppb.New(now)
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdatePrivateCloud(ctx context.Context, req *pb.UpdatePrivateCloudRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateCloudName(req.GetPrivateCloud().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateCloud{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetPrivateCloud().Description
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeletePrivateCloud(ctx context.Context, req *pb.DeletePrivateCloudRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateCloudName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateCloud{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.State = pb.PrivateCloud_DELETED
	obj.DeleteTime = timestamppb.New(now)
	// ExpireTime uses the same default of 3 hours in prod.
	obj.ExpireTime = timestamppb.New(now.Add(time.Hour * 3))
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		// TODO: actually delete the resource after waiting until the expireTime.
		metadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *VMwareEngineV1) UndeletePrivateCloud(ctx context.Context, req *pb.UndeletePrivateCloudRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePrivateCloudName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PrivateCloud{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Can only undelete if the resource has not expired yet.
	if time.Now().After(obj.GetExpireTime().AsTime()) {
		return nil, status.Errorf(codes.FailedPrecondition, "private cloud %q has expired and cannot undeleted", fqn)
	}

	now := time.Now()
	obj.State = pb.PrivateCloud_ACTIVE
	obj.DeleteTime = nil
	obj.ExpireTime = nil
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "undelete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

type privateCloudName struct {
	Project        *projects.ProjectData
	Location       string
	PrivateCloudID string
}

func (n *privateCloudName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", n.Project.ID, n.Location, n.PrivateCloudID)
}

// parsePrivateCloudName parses a string into a privateCloudName.
// The expected form is `projects/*/locations/*/privateClouds/*`.
func (s *MockService) parsePrivateCloudName(name string) (*privateCloudName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "privateClouds" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &privateCloudName{
			Project:        project,
			Location:       tokens[3],
			PrivateCloudID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
