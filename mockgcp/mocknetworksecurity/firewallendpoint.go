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

func (s *FirewallActivationServer) CreateProjectFirewallEndpoint(ctx context.Context, req *pbv1.CreateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.createFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) CreateFirewallEndpoint(ctx context.Context, req *pbv1.CreateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.createFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) createFirewallEndpoint(ctx context.Context, req *pbv1.CreateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	name := req.Parent + "/firewallEndpoints/" + req.FirewallEndpointId

	fqn := name

	obj := proto.Clone(req.FirewallEndpoint).(*pbv1.FirewallEndpoint)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pbv1.FirewallEndpoint_ACTIVE

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
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *FirewallActivationServer) GetProjectFirewallEndpoint(ctx context.Context, req *pbv1.GetFirewallEndpointRequest) (*pbv1.FirewallEndpoint, error) {
	return s.getFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) GetFirewallEndpoint(ctx context.Context, req *pbv1.GetFirewallEndpointRequest) (*pbv1.FirewallEndpoint, error) {
	return s.getFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) getFirewallEndpoint(ctx context.Context, req *pbv1.GetFirewallEndpointRequest) (*pbv1.FirewallEndpoint, error) {
	name, err := s.parseFirewallEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.FirewallEndpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *FirewallActivationServer) UpdateProjectFirewallEndpoint(ctx context.Context, req *pbv1.UpdateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.updateFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) UpdateFirewallEndpoint(ctx context.Context, req *pbv1.UpdateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.updateFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) updateFirewallEndpoint(ctx context.Context, req *pbv1.UpdateFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFirewallEndpointName(req.GetFirewallEndpoint().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pbv1.FirewallEndpoint{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pbv1.FirewallEndpoint)
	updated.UpdateTime = timestamppb.New(time.Now())

	// Apply field mask updates
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"labels", "description", "billing_project_id"}
	}

	for _, path := range paths {
		switch path {
		case "labels":
			updated.Labels = req.GetFirewallEndpoint().GetLabels()
		case "description":
			updated.Description = req.GetFirewallEndpoint().GetDescription()
		case "billingProjectId", "billing_project_id":
			updated.BillingProjectId = req.GetFirewallEndpoint().GetBillingProjectId()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updateable", path)
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}

	var lroPrefix string
	if name.Project != nil {
		lroPrefix = fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	} else {
		lroPrefix = fmt.Sprintf("organizations/%s/locations/%s", name.Organization, name.Location)
	}

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

func (s *FirewallActivationServer) DeleteProjectFirewallEndpoint(ctx context.Context, req *pbv1.DeleteFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.deleteFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) DeleteFirewallEndpoint(ctx context.Context, req *pbv1.DeleteFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	return s.deleteFirewallEndpoint(ctx, req)
}

func (s *FirewallActivationServer) deleteFirewallEndpoint(ctx context.Context, req *pbv1.DeleteFirewallEndpointRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFirewallEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.FirewallEndpoint{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	var lroPrefix string
	if name.Project != nil {
		lroPrefix = fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	} else {
		lroPrefix = fmt.Sprintf("organizations/%s/locations/%s", name.Organization, name.Location)
	}

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

type firewallEndpointName struct {
	Project          *projects.ProjectData
	Organization     string
	Location         string
	FirewallEndpoint string
}

func (n *firewallEndpointName) String() string {
	if n.Project != nil {
		return fmt.Sprintf("projects/%s/locations/%s/firewallEndpoints/%s", n.Project.ID, n.Location, n.FirewallEndpoint)
	}
	return fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", n.Organization, n.Location, n.FirewallEndpoint)
}

func (s *MockService) parseFirewallEndpointName(name string) (*firewallEndpointName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[2] == "locations" && tokens[4] == "firewallEndpoints" {
		if tokens[0] == "projects" {
			project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
			if err != nil {
				return nil, err
			}
			return &firewallEndpointName{
				Project:          project,
				Location:         tokens[3],
				FirewallEndpoint: tokens[5],
			}, nil
		}
		if tokens[0] == "organizations" {
			return &firewallEndpointName{
				Organization:     tokens[1],
				Location:         tokens[3],
				FirewallEndpoint: tokens[5],
			}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
