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
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

const (
	deletedOrgSuffix = "@deleted"

	// defaultApigeeProjectID is the fallback tenant project ID matching authentic GCP baseline tests.
	defaultApigeeProjectID = "bba74dd9140b33c1e-tp"

	// defaultCACertificateBase64 is the root CA certificate used for mock Apigee organizations.
	defaultCACertificateBase64 = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURTVENDQWpHZ0F3SUJBZ0lRUTluaWl4akhCRTA0VHZMRDNTejAxVEFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRd01tUTBZVEJrTWkwNVlURmxMVFJpWXpBdFlUQXhZeTA1T0ROa05HSmtaalZsTVRndwpIaGNOTWpZd056SXpNRGN3T1RBeFdoY05Nell3TnpJd01EY3dPVEF4V2pBdk1TMHdLd1lEVlFRREV5UXdNbVEwCllUQmtNaTA1WVRGbExUUmlZekF0WVRBeFl5MDVPRE5rTkdKa1pqVmxNVGd3Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURHeUVVSDFrVE56dWxhbkVHL3ZicjZ4QlpTa3U0UTlXTmt1QWVoRFZ6ZQpnb1poVkk0bnhEMDYwMU1kS2NQWkNKK2k4b1d3OGpFTllOVThOK3BHSmlqTlM2WHErTjAzZjJvdnVFUjJJRlRNCk4zb1BuMXhKU0NZT0JBeVVhUGhEanZOaFBmYXFVWUJHRXBiSHRRbHpvZTZWd2lsQTkvU3ZqMFBwaWtNSjlFMWUKZGw5OHNtR05YOUZPRTk1dVhtZzNuVzRHOW5wTVY2cVh1K1FmRmZXUVdBRGJnS3VlK3FDWmlFaFRUMzVoM20yMApZQ3pydzF5elQ0b25HUzk1bHZub093L1dBcnc3bTdmZVRzWWhmNTVaR0Fyc2tCR3RtZ21zd3lnUU1ZZmRRb0hrClVQaHFXOC9KeDBUZzFZVzRaSHhlbmVqNGh6K1k0SjkwQUN3aDQwbGpQdUNMQWdNQkFBR2pZVEJmTUE0R0ExVWQKRHdFQi93UUVBd0lDcERBZEJnTlZIU1VFRmpBVUJnZ3JCZ0VGQlFjREFRWUlLd1lCQlFVSEF3SXdEd1lEVlIwVApBUUgvQkFVd0F3RUIvekFkQmdOVkhRNEVGZ1FVcHplQitIdHkrVlVGZHJrREdWZlNuM00rbDB3d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBSWtoMkx0VURBUXhOVW4vNkhYVVVEYXl4RUtQL29Ka3ZVcjZ3UXkwdGNnZ0JhRFoKVzlsUTFGeFU3NXQ1ME5iS3k0b2tSMUFpYUlQcjdoTVB4enZ1QytDR0VGcHYwcVJablc3aDVlNjlERUg5U0tMYwpEbFdXMGlFUUNSTUd4MGN1SVNsRWdGRDRlMEoyTVdadktsbFAwb2xZekwrQ2RzZzR0Q3BhNTlELzV1KzYyUm5XCjRnVDRvek5wa3JlZ0NBSElHMTFXZ2xUazNmc3RLa2xaUlY2Yjh6NFZ0SFo5UTRCa0ozVEtFVzZmTEhOQ2pmQysKVjBVNWJwdG9aVmRSd1ZBSStjaVpJZjlFc2ZhZ1RiQnNuMjVJdzBuU3V1ZXBMSGo0TmVZQWFOMTQyWnRwNXo1YwpLMlRkUStWbXRqMEpZdHVXbkZPWkQ5YWpCd1RNZ2RYYmRNbmNydHc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
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
	if err := s.clearDeleted(ctx, fqn); err != nil {
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

func (s *organizationsServer) SetAddonsOrganization(ctx context.Context, req *pb.SetAddonsOrganizationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseOrganizationName(req.GetOrg())
	if err != nil {
		return nil, err
	}

	existing, err := s.getOrCreateDefaultOrganization(ctx, name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

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
	return s.getOrCreateDefaultOrganization(ctx, name)
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
	if err := s.markDeleted(ctx, fqn, name); err != nil {
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

// getOrCreateDefaultOrganization retrieves an existing organization from storage,
// or auto-populates defaults if unseeded and not previously deleted.
// Handles concurrent auto-population race conditions safely.
func (s *organizationsServer) getOrCreateDefaultOrganization(ctx context.Context, name *OrganizationName) (*pb.GoogleCloudApigeeV1Organization, error) {
	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if s.isDeleted(ctx, fqn) {
				return nil, status.Errorf(codes.NotFound, "googleCloudApigeeV1Organization %q not found", fqn)
			}
			defaultObj := s.populateDefaultOrganization(ctx, name)
			if err := s.storage.Create(ctx, fqn, defaultObj); err != nil {
				if status.Code(err) == codes.AlreadyExists {
					// Concurrently created by another request; retrieve the stored object
					if err := s.storage.Get(ctx, fqn, obj); err != nil {
						return nil, err
					}
					return obj, nil
				}
				return nil, err
			}
			return defaultObj, nil
		}
		return nil, err
	}
	return obj, nil
}

// populateDefaultOrganization initializes default configuration for unseeded test organizations.
func (s *organizationsServer) populateDefaultOrganization(ctx context.Context, name *OrganizationName) *pb.GoogleCloudApigeeV1Organization {
	now := time.Now()
	certBytes := s.getCaCertificate(name.ID)
	apigeeProjectID := s.deriveApigeeProjectID(name.ID)

	return &pb.GoogleCloudApigeeV1Organization{
		Name:              name.ID,
		ProjectId:         name.ID,
		ApigeeProjectId:   apigeeProjectID,
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

// deriveApigeeProjectID returns the tenant project ID for an Apigee organization.
// It defaults to the authentic tenant project ID recorded in golden logs.
func (s *organizationsServer) deriveApigeeProjectID(projectID string) string {
	return defaultApigeeProjectID
}

// getCaCertificate returns the decoded CA certificate bytes for an organization.
func (s *organizationsServer) getCaCertificate(projectID string) []byte {
	certBytes, _ := base64.StdEncoding.DecodeString(defaultCACertificateBase64)
	return certBytes
}

func (s *organizationsServer) markDeleted(ctx context.Context, fqn string, name *OrganizationName) error {
	deletedMarker := &pb.GoogleCloudApigeeV1Organization{
		Name:  name.ID,
		State: "DELETED",
	}
	deletedFQN := fqn + deletedOrgSuffix
	if err := s.storage.Create(ctx, deletedFQN, deletedMarker); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return s.storage.Update(ctx, deletedFQN, deletedMarker)
		}
		return err
	}
	return nil
}

func (s *organizationsServer) clearDeleted(ctx context.Context, fqn string) error {
	deletedFQN := fqn + deletedOrgSuffix
	deletedMarker := &pb.GoogleCloudApigeeV1Organization{}
	if err := s.storage.Delete(ctx, deletedFQN, deletedMarker); err != nil {
		if status.Code(err) != codes.NotFound {
			return err
		}
	}
	return nil
}

func (s *organizationsServer) isDeleted(ctx context.Context, fqn string) bool {
	deletedFQN := fqn + deletedOrgSuffix
	deletedMarker := &pb.GoogleCloudApigeeV1Organization{}
	err := s.storage.Get(ctx, deletedFQN, deletedMarker)
	return err == nil
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
