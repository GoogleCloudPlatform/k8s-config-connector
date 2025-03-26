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
// proto.service: google.cloud.compute.v1.Interconnects
// proto.message: google.cloud.compute.v1.Interconnect

package mockcompute

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type interconnectsV1 struct {
	*MockService
	pb.UnimplementedInterconnectsServer
}

func (s *interconnectsV1) Get(ctx context.Context, req *pb.GetInterconnectRequest) (*pb.Interconnect, error) {
	name, err := s.parseInterconnectName(req.Project, req.Interconnect)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Interconnect{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name.String())
		}
		return nil, err
	}

	return obj, nil
}

func (s *interconnectsV1) Insert(ctx context.Context, req *pb.InsertInterconnectRequest) (*pb.Operation, error) {
	name, err := s.parseInterconnectName(req.Project, req.GetInterconnectResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetInterconnectResource()).(*pb.Interconnect)
	obj.SelfLink = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/%s", name.String()))
	obj.Kind = PtrTo("compute#interconnect")
	obj.Id = proto.Uint64(s.generateID())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetLink:    obj.SelfLink,
		Status:        PtrTo(pb.Operation_RUNNING),
		User:          PtrTo("user@example.com"),
		Progress:      PtrTo(int32(0)),
		TargetId:      obj.Id,
	}

	return s.startGlobalLRO(ctx, req.Project, op, func() (proto.Message, error) {
		obj, err := s.Get(ctx, &pb.GetInterconnectRequest{Project: name.Project, Interconnect: name.Name})
		if err != nil {
			return nil, fmt.Errorf("getting object: %w", err)
		}

		obj.State = PtrTo("UNPROVISIONED")
		obj.InterconnectType = PtrTo("IT_PRIVATE")
		obj.AdminEnabled = PtrTo(true)
		obj.CreationTimestamp = PtrTo(timestamppb.New(time.Now()).String())
		obj.GoogleReferenceId = PtrTo("1234567890")
		obj.LabelFingerprint = PtrTo("abcdef0123A=")
		obj.MacsecEnabled = PtrTo(false)
		obj.OperationalStatus = PtrTo("OS_UNPROVISIONED")

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "error updating object: %v", err)
		}
		return obj, nil
	})
}

func (s *interconnectsV1) Delete(ctx context.Context, req *pb.DeleteInterconnectRequest) (*pb.Operation, error) {
	name, err := s.parseInterconnectName(req.Project, req.Interconnect)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Interconnect{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%v not found", req.Interconnect)
		}
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetLink:    obj.SelfLink,
		Status:        PtrTo(pb.Operation_RUNNING),
		User:          PtrTo("user@example.com"),
		TargetId:      obj.Id,
		Progress:      PtrTo(int32(0)),
	}

	return s.startGlobalLRO(ctx, req.Project, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

// todo: more fidelity for Patch
func (s *interconnectsV1) Patch(ctx context.Context, req *pb.PatchInterconnectRequest) (*pb.Operation, error) {
	name, err := s.parseInterconnectName(req.GetProject(), req.GetInterconnect())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Interconnect{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetInterconnectResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetLink:    obj.SelfLink,
		Status:        PtrTo(pb.Operation_RUNNING),
		User:          PtrTo("user@example.com"),
		TargetId:      obj.Id,
		Progress:      PtrTo(int32(0)),
	}
	return s.startGlobalLRO(ctx, req.Project, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *MockService) parseInterconnectName(project, name string) (*interconnectName, error) {
	if project == "" {
		return nil, fmt.Errorf("project cannot be empty")
	}
	return &interconnectName{
		Project: project,
		Name:    name,
	}, nil
}

type interconnectName struct {
	Project string
	Name    string
}

func (n *interconnectName) String() string {
	return "projects/" + n.Project + "/global/interconnects/" + n.Name
}
