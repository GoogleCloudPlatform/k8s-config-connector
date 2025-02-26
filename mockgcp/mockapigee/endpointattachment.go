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

package mockapigee

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type endpointAttachmentName struct {
	Organization       string
	EndpointAttachment string
}

func (n *endpointAttachmentName) Parent() string {
	return fmt.Sprintf("organizations/%v", n.Organization)
}

func (n *endpointAttachmentName) String() string {
	return fmt.Sprintf("organizations/%v/endpointAttachments/%v", n.Organization, n.EndpointAttachment)
}

// ParseEndpointAttachmentName parses a string into a endpointAttachmentName.
// The expected form is organizations/{organization}/endpointAttachments/{attachment}.
func ParseEndpointAttachmentName(name string) (*endpointAttachmentName, error) {
	expectedFormat := "organizations/{organization}/endpointAttachments/{attachment}"
	parts := strings.Split(name, "/")
	if len(parts) != 4 || parts[0] != "organizations" || parts[2] != "endpointAttachments" {
		return nil, fmt.Errorf("name '%s' is not of the form %s", name, expectedFormat)
	}
	return &endpointAttachmentName{
		Organization:       parts[1],
		EndpointAttachment: parts[3],
	}, nil
}

type endpointAttachmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsEndpointAttachmentsServerServer
}

func (s *endpointAttachmentsServer) GetOrganizationsEndpointAttachment(ctx context.Context, req *pb.GetOrganizationsEndpointAttachmentRequest) (*pb.GoogleCloudApigeeV1EndpointAttachment, error) {
	name, err := ParseEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *endpointAttachmentsServer) CreateOrganizationsEndpointAttachment(ctx context.Context, req *pb.CreateOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/endpointAttachments/" + req.GetEndpointAttachmentId()
	name, err := ParseEndpointAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsEndpointAttachment).(*pb.GoogleCloudApigeeV1EndpointAttachment)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, req.GetParent(), metadata, func() (proto.Message, error) {
		metadata.Progress = &pb.GoogleCloudApigeeV1OperationMetadataProgress{
			Description: "Succeeded",
			PercentDone: 100,
		}
		metadata.State = "FINISHED"
		result := proto.Clone(obj).(*pb.GoogleCloudApigeeV1EndpointAttachment)
		populateOutputsForOrganizationsEndpointAttachment(result)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *endpointAttachmentsServer) DeleteOrganizationsEndpointAttachment(ctx context.Context, req *pb.DeleteOrganizationsEndpointAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := ParseEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		return &pb.GoogleCloudApigeeV1EndpointAttachment{}, nil
	})
	return op, err
}

func populateOutputsForOrganizationsEndpointAttachment(obj *pb.GoogleCloudApigeeV1EndpointAttachment) {
	obj.ConnectionState = "ACCEPTED"
	obj.Host = "10.1.2.3"
	obj.State = "ACTIVE"
}
