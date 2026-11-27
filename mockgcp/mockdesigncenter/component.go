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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DesignCenterServer) GetComponent(ctx context.Context, req *pb.GetComponentRequest) (*pb.Component, error) {
	name, err := s.parseComponentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Component{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateComponent(obj), nil
}

func (s *DesignCenterServer) CreateComponent(ctx context.Context, req *pb.CreateComponentRequest) (*pb.Component, error) {
	parent, err := s.parseApplicationTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent application template does not exist
	parentAppTemplate := &pb.ApplicationTemplate{}
	if err := s.storage.Get(ctx, parent.String(), parentAppTemplate); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.ComponentId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "componentId is required")
	}

	if req.Component == nil {
		return nil, status.Errorf(codes.InvalidArgument, "component is required")
	}

	componentID := req.ComponentId
	fqn := fmt.Sprintf("%s/components/%s", parent.String(), componentID)

	name, err := s.parseComponentName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Component).(*pb.Component)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj = s.populateComponent(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdateComponent(ctx context.Context, req *pb.UpdateComponentRequest) (*pb.Component, error) {
	if req.Component == nil {
		return nil, status.Errorf(codes.InvalidArgument, "component is required")
	}

	name, err := s.parseComponentName(req.Component.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Component{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Component)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Component, paths); err != nil {
		return nil, err
	}

	updated = s.populateComponent(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeleteComponent(ctx context.Context, req *pb.DeleteComponentRequest) (*emptypb.Empty, error) {
	name, err := s.parseComponentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Component{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DesignCenterServer) ListComponents(ctx context.Context, req *pb.ListComponentsRequest) (*pb.ListComponentsResponse, error) {
	parent, err := s.parseApplicationTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/components/"
	response := &pb.ListComponentsResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	componentKind := (&pb.Component{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, componentKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.Component)
		response.Components = append(response.Components, s.populateComponent(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateComponent(obj *pb.Component) *pb.Component {
	if len(obj.Apis) == 0 {
		obj.Apis = []string{
			"cloudbilling.googleapis.com",
			"servicenetworking.googleapis.com",
			"notebooks.googleapis.com",
			"aiplatform.googleapis.com",
			"iap.googleapis.com",
			"iamcredentials.googleapis.com",
			"cloudkms.googleapis.com",
			"secretmanager.googleapis.com",
			"modelarmor.googleapis.com",
			"dlp.googleapis.com",
			"cloudresourcemanager.googleapis.com",
			"serviceusage.googleapis.com",
			"iam.googleapis.com",
		}
	}
	if len(obj.Roles) == 0 {
		obj.Roles = []string{
			"roles/modelarmor.admin",
			"roles/modelarmor.floorSettingsAdmin",
			"roles/artifactregistry.reader",
			"roles/compute.admin",
			"roles/compute.networkAdmin",
			"roles/iam.securityAdmin",
			"roles/iam.serviceAccountAdmin",
			"roles/cloudkms.admin",
			"roles/storage.admin",
			"roles/cloudkms.cryptoKeyEncrypterDecrypter",
			"roles/aiplatform.admin",
			"roles/notebooks.admin",
			"roles/iam.serviceAccountUser",
			"roles/iam.roleAdmin",
			"roles/serviceusage.serviceUsageAdmin",
		}
	}
	if len(obj.ComponentParameterSchema) == 0 {
		obj.ComponentParameterSchema = []*pb.ComponentParameterSchema{
			{
				IsRequired: true,
				Key:        "agent-runtime-1_display_name",
				Type:       "string",
			},
			{
				IsRequired: true,
				Key:        "agent-runtime-1_project_id",
				Type:       "string",
			},
			{
				IsRequired: true,
				Key:        "agent-runtime-1_region",
				Type:       "string",
			},
		}
	}
	return obj
}
