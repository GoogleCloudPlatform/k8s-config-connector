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
	"fmt"
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

type saasName struct {
	Project  *projects.ProjectData
	Location string
	Saas     string
}

func (n *saasName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/saas/" + n.Saas
}

func (s *MockService) parseSaasName(name string) (*saasName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "saas" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &saasName{
			Project:  project,
			Location: tokens[3],
			Saas:     tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasDeploymentsServer) GetSaas(ctx context.Context, req *pb.GetSaasRequest) (*pb.Saas, error) {
	name, err := s.parseSaasName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Saas{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateSaas(ctx context.Context, req *pb.CreateSaasRequest) (*pb.Saas, error) {
	reqName := req.Parent + "/saas/" + req.SaasId
	name, err := s.parseSaasName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Saas).(*pb.Saas)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Saas)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) UpdateSaas(ctx context.Context, req *pb.UpdateSaasRequest) (*pb.Saas, error) {
	reqName := req.GetSaas().GetName()
	name, err := s.parseSaasName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Saas{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.Saas, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update Saas fields: %v", err)
		}
	} else {
		desired := req.GetSaas()
		obj.Locations = desired.Locations
		obj.Labels = desired.Labels
		obj.Annotations = desired.Annotations
	}
	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.Etag = "mock-etag-updated"

	// In real GCP, background provisioning completes after creation, populating blueprintRepo, state, and conditions.
	if obj.BlueprintRepo == "" {
		obj.BlueprintRepo = fmt.Sprintf("%s-docker.pkg.dev/%s/%s-00000000", name.Location, name.Project.ID, name.Saas)
	}
	if obj.State == pb.Saas_STATE_TYPE_UNSPECIFIED {
		obj.State = pb.Saas_STATE_ACTIVE
	}
	if len(obj.Conditions) == 0 {
		obj.Conditions = []*pb.SaasCondition{
			{
				Type:               pb.SaasCondition_TYPE_READY,
				Status:             pb.SaasCondition_STATUS_TRUE,
				LastTransitionTime: now,
			},
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Saas)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	response.BlueprintRepo = ""
	response.State = pb.Saas_STATE_TYPE_UNSPECIFIED
	response.Conditions = nil
	return response, nil
}

func (s *SaasDeploymentsServer) DeleteSaas(ctx context.Context, req *pb.DeleteSaasRequest) (*emptypb.Empty, error) {
	name, err := s.parseSaasName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Saas{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SaasDeploymentsServer) ListSaas(ctx context.Context, req *pb.ListSaasRequest) (*pb.ListSaasResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/locations/" + tokens[3] + "/saas/"
	response := &pb.ListSaasResponse{}

	saasKind := (&pb.Saas{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, saasKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		obj := msg.(*pb.Saas)
		response.Saas = append(response.Saas, obj)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
