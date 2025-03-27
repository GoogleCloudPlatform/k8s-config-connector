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
// proto.service: google.cloud.compute.v1.NetworkAttachments
// proto.message: google.cloud.compute.v1.NetworkAttachment

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type networkAttachments struct {
	*MockService
	pb.UnimplementedNetworkAttachmentsServer
}

func (s *networkAttachments) Get(ctx context.Context, req *pb.GetNetworkAttachmentRequest) (*pb.NetworkAttachment, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", req.GetProject(), req.GetRegion(), req.GetNetworkAttachment())
	name, err := s.parseNetworkAttachmentName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.NetworkAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "NetworkAttachment %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *networkAttachments) Insert(ctx context.Context, req *pb.InsertNetworkAttachmentRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", req.GetProject(), req.GetRegion(), req.GetNetworkAttachmentResource().GetName())
	name, err := s.parseNetworkAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetNetworkAttachmentResource()).(*pb.NetworkAttachment)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#networkAttachment")
	obj.CreationTimestamp = PtrTo(s.nowString())

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	obj.ConnectionPreference = PtrTo("ACCEPT_AUTOMATIC")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *networkAttachments) Delete(ctx context.Context, req *pb.DeleteNetworkAttachmentRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", req.GetProject(), req.GetRegion(), req.GetNetworkAttachment())
	name, err := s.parseNetworkAttachmentName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.NetworkAttachment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type networkAttachmentName struct {
	Project             string
	Region              string
	NetworkAttachmentID string
}

func (n *networkAttachmentName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", n.Project, n.Region, n.NetworkAttachmentID)
}

// parseNetworkAttachmentName parses a string into a networkAttachmentName.
// The expected form is `projects/*/regions/*/networkAttachments/*`.
func (s *MockService) parseNetworkAttachmentName(name string) (*networkAttachmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "networkAttachments" {
		name := &networkAttachmentName{
			Project:             tokens[1],
			Region:              tokens[3],
			NetworkAttachmentID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
