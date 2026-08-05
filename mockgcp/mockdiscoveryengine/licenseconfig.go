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
// proto.service: google.cloud.discoveryengine.v1beta.LicenseConfigService
// proto.message: google.cloud.discoveryengine.v1beta.LicenseConfig

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type licenseConfigService struct {
	*MockService
	discoveryenginepb.UnimplementedLicenseConfigServiceServer
}

func (s *licenseConfigService) CreateLicenseConfig(ctx context.Context, req *discoveryenginepb.CreateLicenseConfigRequest) (*discoveryenginepb.LicenseConfig, error) {
	reqName := fmt.Sprintf("%s/licenseConfigs/%s", req.GetParent(), req.GetLicenseConfigId())
	name, err := s.parseLicenseConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetLicenseConfig()).(*discoveryenginepb.LicenseConfig)
	obj.Name = fqn
	obj.State = discoveryenginepb.LicenseConfig_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *licenseConfigService) GetLicenseConfig(ctx context.Context, req *discoveryenginepb.GetLicenseConfigRequest) (*discoveryenginepb.LicenseConfig, error) {
	name, err := s.parseLicenseConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &discoveryenginepb.LicenseConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "LicenseConfig %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *licenseConfigService) UpdateLicenseConfig(ctx context.Context, req *discoveryenginepb.UpdateLicenseConfigRequest) (*discoveryenginepb.LicenseConfig, error) {
	name, err := s.parseLicenseConfigName(req.LicenseConfig.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &discoveryenginepb.LicenseConfig{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "LicenseConfig %q not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "autoRenew", "auto_renew":
			obj.AutoRenew = req.GetLicenseConfig().GetAutoRenew()
		case "licenseCount", "license_count":
			obj.LicenseCount = req.GetLicenseConfig().GetLicenseCount()
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type licenseConfigName struct {
	Project       *projects.ProjectData
	Location      string
	LicenseConfig string
}

func (n *licenseConfigName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/licenseConfigs/%s", n.Project.Number, n.Location, n.LicenseConfig)
}

func (s *MockService) parseLicenseConfigName(name string) (*licenseConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "licenseConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &licenseConfigName{
			Project:       project,
			Location:      tokens[3],
			LicenseConfig: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
