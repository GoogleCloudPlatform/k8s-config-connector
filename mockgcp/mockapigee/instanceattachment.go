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

type instanceAttachmentName struct {
	Organization string
	Instance     string
	Attachment   string
}

func (n *instanceAttachmentName) Parent() string {
	return fmt.Sprintf("organizations/%v/instances/%v", n.Organization, n.Instance)
}

func (n *instanceAttachmentName) String() string {
	return fmt.Sprintf("organizations/%v/instances/%v/attachments/%v", n.Organization, n.Instance, n.Attachment)
}

// parseInstanceAttachmentName parses a string into a instanceName.
// The expected form is organizations/{organization}/instances/{instance}.
func (s *instancesAttachmentsServer) parseInstanceAttachmentName(name string) (*instanceAttachmentName, error) {
	expectedFormat := "organizations/{organization}/instances/{instance}/attachments/{attachment}"
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "organizations" || parts[2] != "instances" || parts[4] != "attachments" {
		return nil, fmt.Errorf("name '%s' is not of the form %s", name, expectedFormat)
	}
	return &instanceAttachmentName{
		Organization: parts[1],
		Instance:     parts[3],
		Attachment:   parts[5],
	}, nil
}

type instancesAttachmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsInstancesAttachmentsServerServer
}

func (s *instancesAttachmentsServer) GetOrganizationsInstancesAttachment(ctx context.Context, req *pb.GetOrganizationsInstancesAttachmentRequest) (*pb.GoogleCloudApigeeV1InstanceAttachment, error) {
	name, err := s.parseInstanceAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1InstanceAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *instancesAttachmentsServer) CreateOrganizationsInstancesAttachment(ctx context.Context, req *pb.CreateOrganizationsInstancesAttachmentRequest) (*longrunningpb.Operation, error) {
	obj := proto.Clone(req.OrganizationsInstancesAttachment).(*pb.GoogleCloudApigeeV1InstanceAttachment)
	populateDefaultsForOrganizationsInstancesAttachment(obj)

	reqName := req.Parent + "/attachments/" + obj.Name
	name, err := s.parseInstanceAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instanceName, err := ParseInstanceName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := instanceName.Parent()

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
		result := proto.Clone(obj).(*pb.GoogleCloudApigeeV1InstanceAttachment)
		populateOutputsForOrganizationsInstancesAttachment(result)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *instancesAttachmentsServer) DeleteOrganizationsInstancesAttachment(ctx context.Context, req *pb.DeleteOrganizationsInstancesAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instanceName, err := ParseInstanceName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := instanceName.Parent()

	oldObj := &pb.GoogleCloudApigeeV1InstanceAttachment{}
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
		return &pb.GoogleCloudApigeeV1InstanceAttachment{}, nil
	})
	return op, err
}

func populateDefaultsForOrganizationsInstancesAttachment(obj *pb.GoogleCloudApigeeV1InstanceAttachment) {
	obj.Name = "${attachmentId}"
}

func populateOutputsForOrganizationsInstancesAttachment(obj *pb.GoogleCloudApigeeV1InstanceAttachment) {
	obj.CreatedAt = 1740434641
}
