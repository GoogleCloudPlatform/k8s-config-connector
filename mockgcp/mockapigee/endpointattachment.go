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

package mockapigee

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

type organizationsEndpointAttachmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsEndpointAttachmentsServerServer
}

func (s *organizationsEndpointAttachmentsServer) CreateOrganizationsEndpointAttachment(ctx context.Context, req *pb.CreateOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEndpointAttachmentName(req.GetParent() + "/endpointAttachments/" + req.GetOrganizationsEndpointAttachment().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetOrganizationsEndpointAttachment()).(*pb.GoogleCloudApigeeV1EndpointAttachment)
	obj.Name = fqn
	obj.State = "ACTIVE"
	//obj.ConnectionState = "CONNECTED"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}

	// TODO: StartLRO
	return s.operations.DoneLRO(ctx, fqn, opMetadata, func() *pb.GoogleCloudApigeeV1EndpointAttachment {
		obj.Name = name.EndpointAttachmentID
		obj.Location = name.Location
		return obj
	}())
}

func (s *organizationsEndpointAttachmentsServer) GetOrganizationsEndpointAttachment(ctx context.Context, req *pb.GetOrganizationsEndpointAttachmentRequest) (*pb.GoogleCloudApigeeV1EndpointAttachment, error) {
	name, err := s.parseEndpointAttachmentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *organizationsEndpointAttachmentsServer) DeleteOrganizationsEndpointAttachment(ctx context.Context, req *pb.DeleteOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEndpointAttachmentName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.OrganizationID)
	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, &emptypb.Empty{})
}

type EndpointAttachmentName struct {
	OrganizationID       string
	EndpointAttachmentID string
	Location             string
}

func (n *EndpointAttachmentName) String() string {
	return fmt.Sprintf("organizations/%s/endpointAttachments/%s", n.OrganizationID, n.EndpointAttachmentID)
}

// parseEndpointAttachmentName parses a string into a EndpointAttachmentName.
// The expected form is `organizations/*/endpointAttachments/*`.
func (s *MockService) parseEndpointAttachmentName(name string) (*EndpointAttachmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "endpointAttachments" {
		name := &EndpointAttachmentName{
			OrganizationID:       tokens[1],
			EndpointAttachmentID: tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}