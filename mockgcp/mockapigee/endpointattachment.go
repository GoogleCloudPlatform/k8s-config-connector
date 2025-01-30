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
// proto.service: mockgcp.cloud.apigee.v1.OrganizationsEndpointAttachmentsServer
// proto.message: mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment

package mockapigee

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	lro "google.golang.org/genproto/googleapis/longrunning"
)

type organizationsEndpointAttachments struct {
	*MockService
	pb.UnimplementedOrganizationsEndpointAttachmentsServer
}

func (s *organizationsEndpointAttachments) GetOrganizationsEndpointAttachment(ctx context.Context, req *pb.GetOrganizationsEndpointAttachmentRequest) (*pb.GoogleCloudApigeeV1EndpointAttachment, error) {
	name, err := s.parseOrganizationsEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "endpoint attachment %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *organizationsEndpointAttachments) CreateOrganizationsEndpointAttachment(ctx context.Context, req *pb.CreateOrganizationsEndpointAttachmentRequest) (*lro.Operation, error) {
	reqName := req.GetParent() + "/endpointAttachments/" + req.GetOrganizationsEndpointAttachmentId()
	name, err := s.parseOrganizationsEndpointAttachmentName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetOrganizationsEndpointAttachment()).(*pb.GoogleCloudApigeeV1EndpointAttachment)
	obj.Name = fqn
	obj.State = "ACTIVE"
	obj.ConnectionState = "CONNECTED"
	obj.Host = "mock-host"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.StartLRO(ctx, name.String(), nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *organizationsEndpointAttachments) DeleteOrganizationsEndpointAttachment(ctx context.Context, req *pb.DeleteOrganizationsEndpointAttachmentRequest) (*lro.Operation, error) {
	name, err := s.parseOrganizationsEndpointAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.GoogleCloudApigeeV1EndpointAttachment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return s.operations.StartLRO(ctx, name.String(), nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *organizationsEndpointAttachments) ListOrganizationsEndpointAttachments(ctx context.Context, req *pb.ListOrganizationsEndpointAttachmentsRequest) (*pb.GoogleCloudApigeeV1ListEndpointAttachmentsResponse, error) {
	parent, err := s.parseOrganizationsEndpointAttachmentParentName(req.GetParent())
	if err != nil {
		return nil, err
	}

	namePrefix := parent.String() + "/endpointAttachments/"
	response := &pb.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}

	err = s.storage.List(ctx, (&pb.GoogleCloudApigeeV1EndpointAttachment{}).ProtoReflect().Descriptor(), func(obj proto.Message) error {
		ep := obj.(*pb.GoogleCloudApigeeV1EndpointAttachment)
		if strings.HasPrefix(ep.GetName(), namePrefix) {
			response.EndpointAttachments = append(response.EndpointAttachments, ep)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

type organizationsEndpointAttachmentName struct {
	Organization         *projects.ProjectData
	EndpointAttachmentID string
}

func (a *organizationsEndpointAttachmentName) String() string {
	return fmt.Sprintf("organizations/%d/endpointAttachments/%s", a.Organization.Number, a.EndpointAttachmentID)
}

func (s *organizationsEndpointAttachments) parseOrganizationsEndpointAttachmentName(name string) (*organizationsEndpointAttachmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "endpointAttachments" {
		org, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &organizationsEndpointAttachmentName{
			Organization:         org,
			EndpointAttachmentID: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type organizationsEndpointAttachmentParentName struct {
	Organization *projects.ProjectData
}

func (a *organizationsEndpointAttachmentParentName) String() string {
	return fmt.Sprintf("organizations/%d", a.Organization.Number)
}

func (s *organizationsEndpointAttachments) parseOrganizationsEndpointAttachmentParentName(name string) (*organizationsEndpointAttachmentParentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "organizations" {
		org, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &organizationsEndpointAttachmentParentName{
			Organization: org,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
