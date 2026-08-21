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
// proto.message: google.cloud.securitycenter.v1.MuteConfig

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

type SecurityCenterServer struct {
	*MockService
	pb.UnimplementedSecurityCenterServer
}

type muteConfigName struct {
	Organization string
	ID           string
}

func (s *SecurityCenterServer) parseMuteConfigName(name string) (*muteConfigName, error) {
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimPrefix(name, "v1/")
	tokens := strings.Split(name, "/")
	// Expected formats:
	// organizations/{organization}/muteConfigs/{mute_config}
	// organizations/{organization}/locations/{location}/muteConfigs/{mute_config}
	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "muteConfigs" {
		return &muteConfigName{
			Organization: tokens[1],
			ID:           tokens[3],
		}, nil
	}
	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "muteConfigs" {
		return &muteConfigName{
			Organization: tokens[1],
			ID:           tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid mute config name", name)
}

func (n *muteConfigName) String() string {
	return fmt.Sprintf("organizations/%s/muteConfigs/%s", n.Organization, n.ID)
}

func (s *SecurityCenterServer) GetMuteConfig(ctx context.Context, req *pb.GetMuteConfigRequest) (*pb.MuteConfig, error) {
	name, err := s.parseMuteConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MuteConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) CreateMuteConfig(ctx context.Context, req *pb.CreateMuteConfigRequest) (*pb.MuteConfig, error) {
	parent := req.GetParent() // e.g. organizations/{organization} or organizations/{organization}/locations/{location}
	// Check if there is location/global
	var muteConfigID string
	if req.GetMuteConfigId() != "" {
		muteConfigID = req.GetMuteConfigId()
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "mute_config_id is required")
	}

	reqName := fmt.Sprintf("%s/muteConfigs/%s", parent, muteConfigID)
	name, err := s.parseMuteConfigName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetMuteConfig()).(*pb.MuteConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.MostRecentEditor = "mock-editor@gcp-mock.iam.gserviceaccount.com"

	// STATIC is 1, DYNAMIC is 2. STATIC by default.
	if obj.Type == pb.MuteConfig_MUTE_CONFIG_TYPE_UNSPECIFIED {
		obj.Type = pb.MuteConfig_STATIC
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) UpdateMuteConfig(ctx context.Context, req *pb.UpdateMuteConfigRequest) (*pb.MuteConfig, error) {
	reqName := req.GetMuteConfig().GetName()

	name, err := s.parseMuteConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MuteConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Description = req.GetMuteConfig().GetDescription()
		obj.Filter = req.GetMuteConfig().GetFilter()
		obj.DisplayName = req.GetMuteConfig().GetDisplayName()
		obj.Type = req.GetMuteConfig().GetType()
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetMuteConfig().GetDescription()
			case "filter":
				obj.Filter = req.GetMuteConfig().GetFilter()
			case "display_name", "displayName":
				obj.DisplayName = req.GetMuteConfig().GetDisplayName()
			case "type":
				obj.Type = req.GetMuteConfig().GetType()
			case "expiry_time", "expiryTime":
				obj.ExpiryTime = req.GetMuteConfig().GetExpiryTime()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) DeleteMuteConfig(ctx context.Context, req *pb.DeleteMuteConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseMuteConfigName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.MuteConfig{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
