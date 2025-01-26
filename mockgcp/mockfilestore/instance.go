// Copyright 2025 Google LLC
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
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.resource: Instance

package mockfilestore

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
	commonpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/filestore/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *FilestoreV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *FilestoreV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/instances/%s", req.GetParent(), req.GetInstanceId())
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.Instance_CREATING

	if obj.Tier == pb.Instance_TIER_UNSPECIFIED {
		return nil, status.Errorf(codes.InvalidArgument, "tier must be specified.")
	}

	if len(obj.Networks) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "must have exactly one network config.")
	}

	if len(obj.FileShares) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "expected exactly one file share.")
	}

	for _, fileShare := range obj.FileShares {
		if fileShare.CapacityGb < 1024 {
			return nil, status.Errorf(codes.InvalidArgument, "capacity must be at least 1024 GiB.")
		}
	}

	s.populateDefaultsForInstance(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	metadata := &commonpb.OperationMetadata{
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FilestoreV1) populateDefaultsForInstance(obj *pb.Instance) {

}

func (s *FilestoreV1) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.GetInstance().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetInstance().Description
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     obj.Name,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		return obj, nil
	})
}

func (s *FilestoreV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, prefix, nil, &emptypb.Empty{})

}

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/instances/%s", n.Project.Number, n.Location, n.Instance)
}

// parseInstanceName parses a string into an InstanceName.
// The expected form is `projects/*/locations/*/instances/*`.
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project:  project,
			Location: tokens[3],
			Instance: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
