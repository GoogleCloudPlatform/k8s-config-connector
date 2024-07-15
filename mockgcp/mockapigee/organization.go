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
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
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
	obj.Name = PtrTo(name.ID)
	obj.CreatedAt = PtrTo(now.UnixMilli())
	obj.LastModifiedAt = PtrTo(now.UnixMilli())
	obj.ProjectId = PtrTo(projectID)
	obj.State = PtrTo("ACTIVE")

	obj.BillingType = PtrTo("EVALUATION")
	obj.SubscriptionType = PtrTo("TRIAL")

	expiresAt := now.Add(60 * 24 * time.Hour)
	obj.ExpiresAt = PtrTo(expiresAt.UnixMilli())

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
		OperationType:      PtrTo("INSERT"),
		State:              PtrTo("IN_PROGRESS"),
		TargetResourceName: PtrTo(fqn),
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = PtrTo("FINISHED")
		return obj, nil
	})
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
		OperationType:      PtrTo("UPDATE"),
		State:              PtrTo("IN_PROGRESS"),
		TargetResourceName: PtrTo(fqn),
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = PtrTo("FINISHED")
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
	updated.Properties = req.Organization.Properties
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
		OperationType:      PtrTo("DELETE"),
		State:              PtrTo("IN_PROGRESS"),
		TargetResourceName: PtrTo(fqn),
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = PtrTo("FINISHED")
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
