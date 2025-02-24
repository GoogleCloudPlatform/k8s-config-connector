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

// +tool:mockgcp-support-apigee
// proto.service: mockgcp.cloud.apigee.v1.OrganizationsEndpointAttachmentsServer
// proto.message: mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment

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

type endpointAttachmentName struct {
	Organization       string
	EndpointAttachment string
}

func (n *endpointAttachmentName) String() string {
	return fmt.Sprintf("organizations/%v/endpointattachments/%v", n.Organization, n.EndpointAttachment)
}

// parseEndpointAttachmentName parses a string into an endpointAttachmentName.
// The expected form is organizations/{org}/endpointattachments/{endpoint_attachment}
func (s *MockService) parseEndpointAttachmentName(name string) (*endpointAttachmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "endpointattachments" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid endpoint attachment name %q: must be of the form organizations/org/endpointattachments/endpoint_attachment", name)
	}

	return &endpointAttachmentName{
		Organization:       tokens[1],
		EndpointAttachment: tokens[3],
	}, nil
}

type endpointAttachmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsEndpointAttachmentsServerServer
}

func (s *endpointAttachmentsServer) GetOrganizationsEndpointAttachment(ctx context.Context, req *pb.GetOrganizationsEndpointAttachmentRequest) (*pb.GoogleCloudApigeeV1EndpointAttachment, error) {
	name, err := s.parseEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *endpointAttachmentsServer) CreateOrganizationsEndpointAttachment(ctx context.Context, req *pb.CreateOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/endpointattachments/" + req.OrganizationsEndpointAttachment.Name
	name, err := s.parseEndpointAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsEndpointAttachment).(*pb.GoogleCloudApigeeV1EndpointAttachment)
	obj.Name = req.OrganizationsEndpointAttachment.Name
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "CREATE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)

	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, func() *pb.GoogleCloudApigeeV1EndpointAttachment {
		obj.Name = name.EndpointAttachment
		return obj
	}())
}

func (s *endpointAttachmentsServer) DeleteOrganizationsEndpointAttachment(ctx context.Context, req *pb.DeleteOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)
	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, &emptypb.Empty{})
}
