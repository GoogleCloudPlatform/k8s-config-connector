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

package mocksaasservicemgmt

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type tenantName struct {
	Project  *projects.ProjectData
	Location string
	Tenant   string
}

func (n *tenantName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/tenants/" + n.Tenant
}

func (s *MockService) parseTenantName(name string) (*tenantName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "tenants" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &tenantName{
			Project:  project,
			Location: tokens[3],
			Tenant:   tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasDeploymentsServer) GetTenant(ctx context.Context, req *pb.GetTenantRequest) (*pb.Tenant, error) {
	name, err := s.parseTenantName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tenant{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateTenant(ctx context.Context, req *pb.CreateTenantRequest) (*pb.Tenant, error) {
	reqName := req.Parent + "/tenants/" + req.TenantId
	name, err := s.parseTenantName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Tenant).(*pb.Tenant)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Tenant)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) UpdateTenant(ctx context.Context, req *pb.UpdateTenantRequest) (*pb.Tenant, error) {
	reqName := req.GetTenant().GetName()
	name, err := s.parseTenantName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tenant{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.Tenant, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update Tenant fields: %v", err)
		}
	} else {
		desired := req.GetTenant()
		obj.ConsumerResource = desired.ConsumerResource
		obj.Saas = desired.Saas
		obj.Labels = desired.Labels
		obj.Annotations = desired.Annotations
	}
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Tenant)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) DeleteTenant(ctx context.Context, req *pb.DeleteTenantRequest) (*emptypb.Empty, error) {
	name, err := s.parseTenantName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tenant{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SaasDeploymentsServer) ListTenants(ctx context.Context, req *pb.ListTenantsRequest) (*pb.ListTenantsResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/locations/" + tokens[3] + "/tenants/"
	response := &pb.ListTenantsResponse{}

	tenantKind := (&pb.Tenant{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tenantKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		obj := msg.(*pb.Tenant)
		response.Tenants = append(response.Tenants, obj)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
