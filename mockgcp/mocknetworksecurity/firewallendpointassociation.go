// Copyright 2026 Google LLC
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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FirewallActivationServer struct {
	*MockService
	pbv1.UnimplementedFirewallActivationServer
}

func (s *FirewallActivationServer) CreateFirewallEndpointAssociation(ctx context.Context, req *pbv1.CreateFirewallEndpointAssociationRequest) (*longrunningpb.Operation, error) {
	name := req.Parent + "/firewallEndpointAssociations/" + req.FirewallEndpointAssociationId

	fqn := name

	obj := proto.CloneOf(req.FirewallEndpointAssociation)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pbv1.FirewallEndpointAssociation_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *FirewallActivationServer) GetFirewallEndpointAssociation(ctx context.Context, req *pbv1.GetFirewallEndpointAssociationRequest) (*pbv1.FirewallEndpointAssociation, error) {
	name, err := s.parseFirewallEndpointAssociationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.FirewallEndpointAssociation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *FirewallActivationServer) UpdateFirewallEndpointAssociation(ctx context.Context, req *pbv1.UpdateFirewallEndpointAssociationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFirewallEndpointAssociationName(req.GetFirewallEndpointAssociation().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pbv1.FirewallEndpointAssociation{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(obj)
	updated.UpdateTime = timestamppb.New(time.Now())

	// Apply field mask updates
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Default to all updateable spec fields if empty
		paths = []string{"labels", "disabled", "tls_inspection_policy"}
	}

	for _, path := range paths {
		switch path {
		case "labels":
			updated.Labels = req.GetFirewallEndpointAssociation().GetLabels()
		case "disabled":
			updated.Disabled = req.GetFirewallEndpointAssociation().GetDisabled()
		case "tlsInspectionPolicy", "tls_inspection_policy":
			updated.TlsInspectionPolicy = req.GetFirewallEndpointAssociation().GetTlsInspectionPolicy()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updateable", path)
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *FirewallActivationServer) DeleteFirewallEndpointAssociation(ctx context.Context, req *pbv1.DeleteFirewallEndpointAssociationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFirewallEndpointAssociationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.FirewallEndpointAssociation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

type firewallEndpointAssociationName struct {
	Project                     *projects.ProjectData
	Location                    string
	FirewallEndpointAssociation string
}

func (n *firewallEndpointAssociationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", n.Project.ID, n.Location, n.FirewallEndpointAssociation)
}

func (s *MockService) parseFirewallEndpointAssociationName(name string) (*firewallEndpointAssociationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "firewallEndpointAssociations" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		return &firewallEndpointAssociationName{
			Project:                     project,
			Location:                    tokens[3],
			FirewallEndpointAssociation: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
