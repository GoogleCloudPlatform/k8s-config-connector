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
	"reflect"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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
			return nil, status.Errorf(codes.NotFound, "Instance %q not found.", fqn)
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

	s.populateDefaultsForInstance(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     obj.Name,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		return mutateObject(ctx, s.storage, fqn, func(obj *pb.Instance) error {
			obj.State = pb.Instance_READY
			for _, network := range obj.Networks {
				network.ReservedIpRange = "10.1.2.0/29"
				network.IpAddresses = []string{"10.1.2.1"}
			}
			return nil
		})
	})
}

// mutateObject updates the object; it gets the object by fqn, calls mutator, then updates the object
func mutateObject[T proto.Message](ctx context.Context, storage storage.Storage, fqn string, mutator func(obj T) error) (T, error) {
	var nilT T

	typeT := reflect.TypeOf(nilT)
	obj := reflect.New(typeT.Elem()).Interface().(T)
	if err := storage.Get(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	if err := mutator(obj); err != nil {
		return nilT, err
	}

	if err := storage.Update(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	return obj, nil
}

func (s *FilestoreV1) populateDefaultsForInstance(obj *pb.Instance) {
	if obj.Tier == pb.Instance_TIER_UNSPECIFIED {
		obj.Tier = pb.Instance_STANDARD
	}
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

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	proto.Merge(obj, req.GetInstance())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	updatedObj := proto.Clone(obj).(*pb.Instance)
	updatedObj.CreateTime = nil
	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	return s.operations.DoneLRO(ctx, prefix, nil, updatedObj)

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

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.Location, n.Instance)
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
