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
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
	"sync"
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

	deletedOrgs sync.Map
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

	obj := proto.CloneOf(req.GetOrganization())
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
	s.deletedOrgs.Delete(fqn)

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

func populateDefaultOrganization(name *OrganizationName) *pb.GoogleCloudApigeeV1Organization {
	now := time.Now()
	caCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURTVENDQWpHZ0F3SUJBZ0lRUTluaWl4akhCRTA0VHZMRDNTejAxVEFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRd01tUTBZVEJrTWkwNVlURmxMVFJpWXpBdFlUQXhZeTA1T0ROa05HSmtaalZsTVRndwpIaGNOTWpZd056SXpNRGN3T1RBeFdoY05Nell3TnpJd01EY3dPVEF4V2pBdk1TMHdLd1lEVlFRREV5UXdNbVEwCllUQmtNaTA1WVRGbExUUmlZekF0WVRBeFl5MDVPRE5rTkdKa1pqVmxNVGd3Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURHeUVVSDFrVE56dWxhbkVHL3ZicjZ4QlpTa3U0UTlXTmt1QWVoRFZ6ZQpnb1poVkk0bnhEMDYwMU1kS2NQWkNKK2k4b1d3OGpFTllOVThOK3BHSmlqTlM2WHErTjAzZjJvdnVFUjJJRlRNCk4zb1BuMXhKU0NZT0JBeVVhUGhEanZOaFBmYXFVWUJHRXBiSHRRbHpvZTZWd2lsQTkvU3ZqMFBwaWtNSjlFMWUKZGw5OHNtR05YOUZPRTk1dVhtZzNuVzRHOW5wTVY2cVh1K1FmRmZXUVdBRGJnS3VlK3FDWmlFaFRUMzVoM20yMApZQ3pydzF5elQ0b25HUzk1bHZub093L1dBcnc3bTdmZVRzWWhmNTVaR0Fyc2tCR3RtZ21zd3lnUU1ZZmRRb0hrClVQaHFXOC9KeDBUZzFZVzRaSHhlbmVqNGh6K1k0SjkwQUN3aDQwbGpQdUNMQWdNQkFBR2pZVEJmTUE0R0ExVWQKRHdFQi93UUVBd0lDcERBZEJnTlZIU1VFRmpBVUJnZ3JCZ0VGQlFjREFRWUlLd1lCQlFVSEF3SXdEd1lEVlIwVApBUUgvQkFVd0F3RUIvekFkQmdOVkhRNEVGZ1FVcHplQitIdHkrVlVGZHJrREdWZlNuM00rbDB3d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBSWtoMkx0VURBUXhOVW4vNkhYVVVEYXl4RUtQL29Ka3ZVcjZ3UXkwdGNnZ0JhRFoKVzlsUTFGeFU3NXQ1ME5iS3k0b2tSMUFpYUlQcjdoTVB4enZ1QytDR0VGcHYwcVJablc3aDVlNjlERUg5U0tMYwpEbFdXMGlFUUNSTUd4MGN1SVNsRWdGRDRlMEoyTVdadktsbFAwb2xZekwrQ2RzZzR0Q3BhNTlELzV1KzYyUm5XCjRnVDRvek5wa3JlZ0NBSElHMTFXZ2xUazNmc3RLa2xaUlY2Yjh6NFZ0SFo5UTRCa0ozVEtFVzZmTEhOQ2pmQysKVjBVNWJwdG9aVmRSd1ZBSStjaVpJZjlFc2ZhZ1RiQnNuMjVJdzBuU3V1ZXBMSGo0TmVZQWFOMTQyWnRwNXo1YwpLMlRkUStWbXRqMEpZdHVXbkZPWkQ5YWpCd1RNZ2RYYmRNbmNydHc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
	certBytes, _ := base64.StdEncoding.DecodeString(caCert)
	return &pb.GoogleCloudApigeeV1Organization{
		Name:              name.ID,
		ProjectId:         name.ID,
		ApigeeProjectId:   "bba74dd9140b33c1e-tp",
		CreatedAt:         now.UnixMilli(),
		LastModifiedAt:    now.UnixMilli(),
		State:             "ACTIVE",
		BillingType:       "PAYG",
		SubscriptionType:  "PAID",
		RuntimeType:       "CLOUD",
		AnalyticsRegion:   "us-central1",
		DisableVpcPeering: true,
		CaCertificate:     certBytes,
		CaCertificates:    [][]byte{certBytes},
		Properties: &pb.GoogleCloudApigeeV1Properties{
			Property: []*pb.GoogleCloudApigeeV1Property{
				{
					Name:  "features.hybrid.enabled",
					Value: "true",
				},
				{
					Name:  "features.mart.connect.enabled",
					Value: "true",
				},
			},
		},
	}
}

func (s *organizationsServer) SetAddonsOrganization(ctx context.Context, req *pb.SetAddonsOrganizationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseOrganizationName(req.GetOrg())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			if _, deleted := s.deletedOrgs.Load(fqn); !deleted {
				existing = populateDefaultOrganization(name)
				if err := s.storage.Create(ctx, fqn, existing); err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	var addonsConfig *pb.GoogleCloudApigeeV1AddonsConfig
	if req.GetOrganization() != nil {
		addonsConfig = req.GetOrganization().GetAddonsConfig()
	}

	updated := proto.CloneOf(existing)
	updated.AddonsConfig = addonsConfig
	if updated.AddonsConfig != nil {
		if updated.AddonsConfig.MonetizationConfig != nil {
			if !updated.AddonsConfig.MonetizationConfig.GetEnabled() {
				updated.AddonsConfig.MonetizationConfig = nil
			}
		}
	}

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
		if updated.AddonsConfig == nil {
			return &pb.GoogleCloudApigeeV1AddonsConfig{}, nil
		}
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

	updated := proto.CloneOf(existing)
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
		if status.Code(err) == codes.NotFound {
			if _, deleted := s.deletedOrgs.Load(fqn); !deleted {
				obj = populateDefaultOrganization(name)
				if err := s.storage.Create(ctx, fqn, obj); err != nil {
					return nil, err
				}
				return obj, nil
			}
		}
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
	s.deletedOrgs.Store(fqn, true)
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
