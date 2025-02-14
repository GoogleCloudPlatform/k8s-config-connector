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
	"google.golang.org/protobuf/types/known/emptypb"
)

type envgroupAttachmentServer struct {
	*MockService
	pb.UnimplementedOrganizationsEnvgroupsAttachmentsServerServer
}

func (s *envgroupAttachmentServer) CreateOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.CreateOrganizationsEnvgroupsAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvgroupAttachmentName(req.GetParent() + "/attachments/" + req.GetOrganizationsEnvgroupsAttachment().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetOrganizationsEnvgroupsAttachment()).(*pb.GoogleCloudApigeeV1EnvironmentGroupAttachment)

	// The name field in the request body is ignored by Apigee.
	// Set it to the fully qualified name for consistency with GCP
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}

	opPrefix := fmt.Sprintf("organizations/%s/envgroups/%s", name.Organization, name.Envgroup)
	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, obj)
}

func (s *envgroupAttachmentServer) GetOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.GetOrganizationsEnvgroupsAttachmentRequest) (*pb.GoogleCloudApigeeV1EnvironmentGroupAttachment, error) {
	name, err := s.parseEnvgroupAttachmentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *envgroupAttachmentServer) DeleteOrganizationsEnvgroupsAttachment(ctx context.Context, req *pb.DeleteOrganizationsEnvgroupsAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvgroupAttachmentName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s/envgroups/%s", name.Organization, name.Envgroup)
	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, &emptypb.Empty{})
}

// There is no UPDATE func for this API based on the generated proto

// EnvgroupAttachmentName represents a "fully qualified name" for an EnvgroupAttachment resource.
type envgroupAttachmentName struct {
	Organization string
	Envgroup     string
	AttachmentID string
}

func (n *envgroupAttachmentName) String() string {
	return fmt.Sprintf("organizations/%s/envgroups/%s/attachments/%s", n.Organization, n.Envgroup, n.AttachmentID)
}

// parseEnvgroupAttachmentName parses the given name string into a envgroupAttachmentName struct.
// The expected format is: organizations/<organization>/envgroups/<envgroup>/attachments/<attachmentID>.
func (s *MockService) parseEnvgroupAttachmentName(name string) (*envgroupAttachmentName, error) {
	split := strings.Split(name, "/")
	if len(split) != 6 || split[0] != "organizations" || split[2] != "envgroups" || split[4] != "attachments" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid envgroup attachment name: %q", name)
	}

	result := &envgroupAttachmentName{
		Organization: split[1],
		Envgroup:     split[3],
		AttachmentID: split[5],
	}
	return result, nil
}
