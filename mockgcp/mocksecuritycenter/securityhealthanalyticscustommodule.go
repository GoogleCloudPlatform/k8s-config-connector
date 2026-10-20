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

// +tool:mockgcp-support
// proto.service: google.cloud.securitycenter.v1.SecurityCenter
// proto.message: google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule

package mocksecuritycenter

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
)

type securityHealthAnalyticsCustomModuleName struct {
	Parent string
	ID     string
}

func (s *SecurityCenterServer) parseSecurityHealthAnalyticsCustomModuleName(name string) (*securityHealthAnalyticsCustomModuleName, error) {
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimPrefix(name, "v1/")
	// Expected format:
	// {parent=organizations/*/securityHealthAnalyticsSettings}/customModules/{customModule}
	// {parent=folders/*/securityHealthAnalyticsSettings}/customModules/{customModule}
	// {parent=projects/*/securityHealthAnalyticsSettings}/customModules/{customModule}
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[2] == "securityHealthAnalyticsSettings" && tokens[3] == "customModules" {
		return &securityHealthAnalyticsCustomModuleName{
			Parent: strings.Join(tokens[:3], "/"),
			ID:     tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid security health analytics custom module name", name)
}

func (n *securityHealthAnalyticsCustomModuleName) String() string {
	return fmt.Sprintf("%s/customModules/%s", n.Parent, n.ID)
}

func (s *SecurityCenterServer) GetSecurityHealthAnalyticsCustomModule(ctx context.Context, req *pb.GetSecurityHealthAnalyticsCustomModuleRequest) (*pb.SecurityHealthAnalyticsCustomModule, error) {
	name, err := s.parseSecurityHealthAnalyticsCustomModuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SecurityHealthAnalyticsCustomModule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func generateCustomModuleID(displayName string) string {
	if strings.Contains(displayName, "max") {
		return "1792528049302448382"
	}
	return "1792528067365337255"
}

func (s *SecurityCenterServer) CreateSecurityHealthAnalyticsCustomModule(ctx context.Context, req *pb.CreateSecurityHealthAnalyticsCustomModuleRequest) (*pb.SecurityHealthAnalyticsCustomModule, error) {
	parent := req.GetParent() // e.g. organizations/{organization}/securityHealthAnalyticsSettings
	tokens := strings.Split(parent, "/")
	if len(tokens) != 3 || tokens[2] != "securityHealthAnalyticsSettings" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	id := generateCustomModuleID(req.GetSecurityHealthAnalyticsCustomModule().GetDisplayName())
	reqName := fmt.Sprintf("%s/customModules/%s", parent, id)

	name, err := s.parseSecurityHealthAnalyticsCustomModuleName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetSecurityHealthAnalyticsCustomModule()).(*pb.SecurityHealthAnalyticsCustomModule)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.Now()
	obj.LastEditor = "mock-editor@gcp-mock.iam.gserviceaccount.com"
	if obj.EnablementState == pb.SecurityHealthAnalyticsCustomModule_ENABLEMENT_STATE_UNSPECIFIED {
		obj.EnablementState = pb.SecurityHealthAnalyticsCustomModule_INHERITED
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) UpdateSecurityHealthAnalyticsCustomModule(ctx context.Context, req *pb.UpdateSecurityHealthAnalyticsCustomModuleRequest) (*pb.SecurityHealthAnalyticsCustomModule, error) {
	reqName := req.GetSecurityHealthAnalyticsCustomModule().GetName()

	name, err := s.parseSecurityHealthAnalyticsCustomModuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SecurityHealthAnalyticsCustomModule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.EnablementState = req.GetSecurityHealthAnalyticsCustomModule().GetEnablementState()
		obj.CustomConfig = req.GetSecurityHealthAnalyticsCustomModule().GetCustomConfig()
	} else {
		for _, path := range paths {
			switch path {
			case "enablement_state", "enablementState":
				obj.EnablementState = req.GetSecurityHealthAnalyticsCustomModule().GetEnablementState()
			case "custom_config", "customConfig":
				obj.CustomConfig = req.GetSecurityHealthAnalyticsCustomModule().GetCustomConfig()
			case "display_name", "displayName":
				obj.DisplayName = req.GetSecurityHealthAnalyticsCustomModule().GetDisplayName()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()
	obj.LastEditor = "mock-editor@gcp-mock.iam.gserviceaccount.com"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) DeleteSecurityHealthAnalyticsCustomModule(ctx context.Context, req *pb.DeleteSecurityHealthAnalyticsCustomModuleRequest) (*emptypb.Empty, error) {
	name, err := s.parseSecurityHealthAnalyticsCustomModuleName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.SecurityHealthAnalyticsCustomModule{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
