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
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.ExternalAddress

package mockvmwareengine

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vmwareengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *VMwareEngineV1) GetExternalAddress(ctx context.Context, req *pb.GetExternalAddressRequest) (*pb.ExternalAddress, error) {
	name, err := s.parseExternalAddressName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExternalAddress{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VMwareEngineV1) CreateExternalAddress(ctx context.Context, req *pb.CreateExternalAddressRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/externalAddresses/" + req.ExternalAddressId
	name, err := s.parseExternalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetExternalAddress()).(*pb.ExternalAddress)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.ExternalAddress_ACTIVE // Or CREATING, then update in LRO
	obj.Uid = fmt.Sprintf("%d", rand.Int63())
	obj.ExternalIp = fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256))

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
		metadata.EndTime = timestamppb.Now()
		// If state was CREATING, update it to ACTIVE here
		// obj.State = pb.ExternalAddress_ACTIVE
		// obj.UpdateTime = timestamppb.Now()
		// if err := s.storage.Update(ctx, fqn, obj); err != nil {
		// 	return nil, err
		// }
		return obj, nil
	})
}

func (s *VMwareEngineV1) UpdateExternalAddress(ctx context.Context, req *pb.UpdateExternalAddressRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseExternalAddressName(req.GetExternalAddress().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.ExternalAddress{}
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
			obj.Description = req.GetExternalAddress().Description
		case "internal_ip":
			obj.InternalIp = req.GetExternalAddress().InternalIp
		// Add other updatable fields if any
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for update", path)
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
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *VMwareEngineV1) DeleteExternalAddress(ctx context.Context, req *pb.DeleteExternalAddressRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseExternalAddressName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ExternalAddress{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent resource should succeed according to API spec
		} else {
			return nil, err
		}
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type externalAddressName struct {
	Project           *projects.ProjectData
	Location          string
	PrivateCloudID    string
	ExternalAddressID string
}

func (n *externalAddressName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", n.Project.ID, n.Location, n.PrivateCloudID, n.ExternalAddressID)
}

// parseExternalAddressName parses a string into a externalAddressName.
// The expected form is `projects/*/locations/*/privateClouds/*/externalAddresses/*`.
func (s *MockService) parseExternalAddressName(name string) (*externalAddressName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "privateClouds" && tokens[6] == "externalAddresses" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &externalAddressName{
			Project:           project,
			Location:          tokens[3],
			PrivateCloudID:    tokens[5],
			ExternalAddressID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
