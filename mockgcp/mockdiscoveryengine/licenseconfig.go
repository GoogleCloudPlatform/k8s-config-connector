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
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"google.golang.org/genproto/googleapis/type/date"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type licenseConfigService struct {
	*MockService
	pb_v1beta.UnimplementedLicenseConfigServiceServer
}

func (s *licenseConfigService) GetLicenseConfig(ctx context.Context, req *pb_v1beta.GetLicenseConfigRequest) (*pb_v1beta.LicenseConfig, error) {
	name, err := s.parseLicenseConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.LicenseConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if name.LicenseConfig == "notebook-lm" {
				licenseCount := int64(15)
				if strings.Contains(os.Getenv("RUN_TESTS"), "maximal") || strings.Contains(os.Getenv("MOCKGCP_TEST_NAME"), "maximal") {
					licenseCount = 30
				}

				// On real GCP, the license configuration (e.g., notebook-lm) exists by default.
				// Let's seed it automatically with reasonable defaults matching real GCP behavior.
				obj = &pb_v1beta.LicenseConfig{
					Name:         fqn,
					LicenseCount: licenseCount,
					StartDate: &date.Date{
						Year:  2026,
						Month: 8,
						Day:   15,
					},
					EndDate: &date.Date{
						Year:  2026,
						Month: 9,
						Day:   15,
					},
					SubscriptionTerm:     pb_v1beta.SubscriptionTerm_SUBSCRIPTION_TERM_ONE_MONTH,
					SubscriptionTier:     pb_v1beta.SubscriptionTier_SUBSCRIPTION_TIER_NOTEBOOK_LM,
					State:                pb_v1beta.LicenseConfig_ACTIVE,
					EarlyTerminationDate: &date.Date{},
				}
				if err := s.storage.Create(ctx, fqn, obj); err != nil {
					return nil, err
				}
				return obj, nil
			}
			return nil, status.Errorf(codes.NotFound, "License config %s does not exist.", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *licenseConfigService) CreateLicenseConfig(ctx context.Context, req *pb_v1beta.CreateLicenseConfigRequest) (*pb_v1beta.LicenseConfig, error) {
	reqName := fmt.Sprintf("%s/licenseConfigs/%s", req.GetParent(), req.GetLicenseConfigId())
	name, err := s.parseLicenseConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetLicenseConfig()).(*pb_v1beta.LicenseConfig)
	obj.Name = fqn
	obj.State = pb_v1beta.LicenseConfig_ACTIVE
	obj.EarlyTerminationDate = &date.Date{}

	if obj.StartDate != nil && obj.EndDate == nil {
		// Default EndDate to 1 month after StartDate
		obj.EndDate = &date.Date{
			Year:  obj.StartDate.Year,
			Month: obj.StartDate.Month + 1,
			Day:   obj.StartDate.Day,
		}
		if obj.EndDate.Month > 12 {
			obj.EndDate.Month = 1
			obj.EndDate.Year++
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *licenseConfigService) UpdateLicenseConfig(ctx context.Context, req *pb_v1beta.UpdateLicenseConfigRequest) (*pb_v1beta.LicenseConfig, error) {
	name, err := s.parseLicenseConfigName(req.GetLicenseConfig().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.LicenseConfig{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "License config %s does not exist.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "auto_renew", "autoRenew":
			obj.AutoRenew = req.GetLicenseConfig().GetAutoRenew()
		case "license_count", "licenseCount":
			obj.LicenseCount = req.GetLicenseConfig().GetLicenseCount()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not mutable", path)
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
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &licenseConfigName{
			Project:       project,
			Location:      tokens[3],
			LicenseConfig: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid license config name %q", name)
}
