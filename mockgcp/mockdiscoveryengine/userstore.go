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

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type userStoreService struct {
	*MockService
	pb_v1beta.UnimplementedUserStoreServiceServer
}

func (s *userStoreService) GetUserStore(ctx context.Context, req *pb_v1beta.GetUserStoreRequest) (*pb_v1beta.UserStore, error) {
	name, err := s.parseUserStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.UserStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Real GCP has userStore default_user_store implicitly, let's seed it when requested.
			obj = &pb_v1beta.UserStore{
				Name: fqn,
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func (s *userStoreService) UpdateUserStore(ctx context.Context, req *pb_v1beta.UpdateUserStoreRequest) (*pb_v1beta.UserStore, error) {
	name, err := s.parseUserStoreName(req.GetUserStore().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.UserStore{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Create it if it doesn't exist
			obj = proto.Clone(req.GetUserStore()).(*pb_v1beta.UserStore)
			obj.Name = fqn
			if obj.DefaultLicenseConfig != "" {
				licName, err := s.parseLicenseConfigName(obj.DefaultLicenseConfig)
				if err == nil {
					obj.DefaultLicenseConfig = licName.String()
				}
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"display_name", "default_license_config", "enable_license_auto_register", "enable_expired_license_auto_update"}
	}

	if err := fields.UpdateByFieldMask(obj, req.GetUserStore(), paths); err != nil {
		return nil, err
	}

	if obj.DefaultLicenseConfig != "" {
		licName, err := s.parseLicenseConfigName(obj.DefaultLicenseConfig)
		if err == nil {
			obj.DefaultLicenseConfig = licName.String()
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type userStoreName struct {
	Project   *projects.ProjectData
	Location  string
	UserStore string
}

func (n *userStoreName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/userStores/%s", n.Project.Number, n.Location, n.UserStore)
}

func (s *MockService) parseUserStoreName(name string) (*userStoreName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "userStores" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &userStoreName{
			Project:   project,
			Location:  tokens[3],
			UserStore: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid user store name %q", name)
}
