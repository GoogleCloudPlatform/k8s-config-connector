// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	"sort"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type organizationsServer struct {
	*MockService
	pb.UnimplementedOrganizationsServerServer
}

func (s *organizationsServer) CreateOrganization(ctx context.Context, req *pb.CreateOrganizationRequest) (*longrunningpb.Operation, error) {
	var name *OrganizationName

	projectID := ""

	parent := req.GetParent()
	tokens := strings.Split(parent, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		projectID = project.ID

		// Name is same as project ID
		name = &OrganizationName{
			ID: project.ID,
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.GetOrganization()).(*pb.GoogleCloudApigeeV1Organization)
	obj.Name = name.ID
	obj.CreatedAt = now.UnixMilli()
	obj.LastModifiedAt = now.UnixMilli()
	obj.ProjectId = projectID
	obj.State = "ACTIVE"

	obj.BillingType = "EVALUATION"
	obj.SubscriptionType = "TRIAL"

	expiresAt := now.Add(60 * 24 * time.Hour)
	obj.ExpiresAt = expiresAt.UnixMilli()

	obj.CaCertificate = []byte("LS0t...")

	if obj.AddonsConfig != nil {
		if obj.AddonsConfig.MonetizationConfig != nil {
			if !obj.AddonsConfig.MonetizationConfig.GetEnabled() {
				obj.AddonsConfig.MonetizationConfig = nil
			}
		}
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := "organizations/" + name.ID
	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = "FINISHED"
		return obj, nil
	})
}

func (s *organizationsServer) ListOrganizations(ctx context.Context, req *pb.ListOrganizationsRequest) (*pb.GoogleCloudApigeeV1ListOrganizationsResponse, error) {
	prefix := ""

	parent := req.GetParent()
	tokens := strings.Split(parent, "/")
	if len(tokens) == 1 && tokens[0] == "organizations" {
		prefix = "organizations/"
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	response := &pb.GoogleCloudApigeeV1ListOrganizationsResponse{}

	modelKind := (&pb.GoogleCloudApigeeV1Organization{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, modelKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		org := obj.(*pb.GoogleCloudApigeeV1Organization)
		mapping := &pb.GoogleCloudApigeeV1OrganizationProjectMapping{}
		mapping.Organization = org.Name
		mapping.ProjectId = org.ProjectId
		mapping.ProjectIds = []string{org.ProjectId}
		response.Organizations = append(response.Organizations, mapping)

		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil

}

func (s *organizationsServer) SetAddonsOrganization(ctx context.Context, req *pb.SetAddonsOrganizationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseOrganizationName(req.GetOrg())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.GoogleCloudApigeeV1Organization)
	updated.AddonsConfig = req.Organization.AddonsConfig

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	prefix := "organizations/" + name.ID
	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "UPDATE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = "FINISHED"
		return updated.AddonsConfig, nil
	})
}

func (s *organizationsServer) UpdateOrganization(ctx context.Context, req *pb.UpdateOrganizationRequest) (*pb.GoogleCloudApigeeV1Organization, error) {
	name, err := s.parseOrganizationName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// Updates the properties for an Apigee organization. No other fields in the organization profile will be updated.

	updated := proto.Clone(existing).(*pb.GoogleCloudApigeeV1Organization)
	props := req.Organization.Properties
	sort.Slice(props.Property, func(i, j int) bool {
		return props.Property[i].Name < props.Property[j].GetName()
	})
	updated.Properties = props

	updated.DisplayName = req.Organization.DisplayName
	updated.Description = req.Organization.Description

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *organizationsServer) GetOrganization(ctx context.Context, req *pb.GetOrganizationRequest) (*pb.GoogleCloudApigeeV1Organization, error) {
	name, err := s.parseOrganizationName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *organizationsServer) DeleteOrganization(ctx context.Context, req *pb.DeleteOrganizationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseOrganizationName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	prefix := "organizations/" + name.ID
	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = "FINISHED"
		return &emptypb.Empty{}, nil
	})
}

type OrganizationName struct {
	ID string
}

func (n *OrganizationName) String() string {
	return fmt.Sprintf("organizations/%s", n.ID)
}

// parseOrganizationName parses a string into a OrganizationName.
// The expected form is `organizations/*`.
func (s *MockService) parseOrganizationName(name string) (*OrganizationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "organizations" {
		name := &OrganizationName{
			ID: tokens[1],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
