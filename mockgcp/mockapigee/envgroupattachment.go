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

type envgroupAttachmentName struct {
	Organization string
	Envgroup     string
	Attachment   string
}

func (n *envgroupAttachmentName) Parent() string {
	return fmt.Sprintf("organizations/%v/envgroups/%v", n.Organization, n.Envgroup)
}

func (n *envgroupAttachmentName) String() string {
	return fmt.Sprintf("organizations/%v/envgroups/%v/attachments/%v", n.Organization, n.Envgroup, n.Attachment)
}

// parseEnvgroupAttachmentName parses a string into a envgroupAttachmentName.
// The expected form is organizations/{organization}/envgroups/{envgroup}/attachments/{attachment}.
func (s *envgroupsAttachmentsServer) parseEnvgroupAttachmentName(name string) (*envgroupAttachmentName, error) {
	expectedFormat := "organizations/{organization}/envgroups/{envgroup}/attachments/{attachment}"
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "organizations" || parts[2] != "envgroups" || parts[4] != "attachments" {
		return nil, fmt.Errorf("name '%s' is not of the form %s", name, expectedFormat)
	}
	return &envgroupAttachmentName{
		Organization: parts[1],
		Envgroup:     parts[3],
		Attachment:   parts[5],
	}, nil
}

type envgroupsAttachmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsEnvgroupsAttachmentsServerServer
}

func (s *envgroupsAttachmentsServer) GetOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.GetOrganizationsEnvgroupsAttachmentRequest) (*pb.GoogleCloudApigeeV1EnvironmentGroupAttachment, error) {
	name, err := s.parseEnvgroupAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *envgroupsAttachmentsServer) CreateOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.CreateOrganizationsEnvgroupsAttachmentRequest) (*longrunningpb.Operation, error) {
	obj := proto.Clone(req.OrganizationsEnvgroupsAttachment).(*pb.GoogleCloudApigeeV1EnvironmentGroupAttachment)
	populateDefaultsForOrganizationsEnvgroupsAttachment(obj)

	reqName := req.Parent + "/attachments/" + obj.Name
	name, err := s.parseEnvgroupAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	envgroupName, err := ParseEnvgroupName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := envgroupName.Parent()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, orgID, metadata, func() (proto.Message, error) {
		metadata.Progress = &pb.GoogleCloudApigeeV1OperationMetadataProgress{
			Description: "Succeeded",
			PercentDone: 100,
		}
		metadata.State = "FINISHED"
		result := proto.Clone(obj).(*pb.GoogleCloudApigeeV1EnvironmentGroupAttachment)
		populateOutputsForOrganizationsEnvgroupsAttachment(result)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *envgroupsAttachmentsServer) DeleteOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.DeleteOrganizationsEnvgroupsAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvgroupAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	envgroupName, err := ParseEnvgroupName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := envgroupName.Parent()

	oldObj := &pb.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, orgID, metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		return &pb.GoogleCloudApigeeV1EnvironmentGroupAttachment{}, nil
	})
	return op, err
}

func populateDefaultsForOrganizationsEnvgroupsAttachment(obj *pb.GoogleCloudApigeeV1EnvironmentGroupAttachment) {
	obj.Name = "${attachmentId}"
}

func populateOutputsForOrganizationsEnvgroupsAttachment(obj *pb.GoogleCloudApigeeV1EnvironmentGroupAttachment) {
	obj.CreatedAt = 1740434641
}
