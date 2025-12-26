// Copyright 2025 Google LLC
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
// proto.service: google.cloud.iap.v1.IdentityAwareProxyAdminService
// proto.message: google.cloud.iap.v1.IapSettings

package mockiap

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

// GetIapSettings retrieves the IAP settings for a resource.
func (s *IdentityAwareProxyAdminService) GetIapSettings(ctx context.Context, req *pb.GetIapSettingsRequest) (*pb.IapSettings, error) {
	name, err := s.parseIapSettingsName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IapSettings{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// If settings don't exist, return a default empty object with the name.
			return &pb.IapSettings{Name: name.String()}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get iap settings: %v", err)
	}

	responseObj := ProtoClone(obj)
	responseObj.Name = name.String()

	return responseObj, nil
}

// UpdateIapSettings updates the IAP settings for a resource.
func (s *IdentityAwareProxyAdminService) UpdateIapSettings(ctx context.Context, req *pb.UpdateIapSettingsRequest) (*pb.IapSettings, error) {
	name, err := s.parseIapSettingsName(req.GetIapSettings().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	isAutoCreate := false
	obj := &pb.IapSettings{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, status.Errorf(codes.Internal, "failed to get IAP settings: %v", err)
		}
		// If not found, create a new empty object.
		obj.Name = name.String()
		isAutoCreate = true
	}

	// Apply the update mask.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If omitted, then all of the settings are updated. See
		paths = []string{
			"access_settings",
			"application_settings",
		}
	}

	source := req.GetIapSettings()

	for _, path := range paths {
		switch path {
		case "access_settings":
			obj.AccessSettings = source.GetAccessSettings()
		case "application_settings":
			obj.ApplicationSettings = source.GetApplicationSettings()
		default:
			return nil, status.Errorf(codes.Unimplemented, "update_mask path %q is not supported by this mock", path)
		}
	}

	if isAutoCreate {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create iap settings: %v", err)
		}
	} else {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update iap settings: %v", err)
		}
	}

	// The response `name` should use the project number.
	responseObj := proto.Clone(obj).(*pb.IapSettings)
	responseObj.Name = name.String()

	return responseObj, nil
}

// iapSettingsName is a parsed IAP Settings resource name.
type iapSettingsName struct {
	Project     *projects.ProjectData
	IAPResource string
}

// String returns the canonical name.
func (n *iapSettingsName) String() string {
	return fmt.Sprintf("projects/%d/%s", n.Project.Number, n.IAPResource)
}

// parseIapSettingsName parses a string into an iapSettingsName.
// The expected form is projects/{project}/{iap_resource_path}.
func (s *IdentityAwareProxyAdminService) parseIapSettingsName(name string) (*iapSettingsName, error) {
	tokens := strings.SplitN(name, "/", 3)
	if len(tokens) == 3 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "project %q not found", tokens[1])
		}

		iapResource := tokens[2]

		return &iapSettingsName{
			Project:     project,
			IAPResource: iapResource,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid iap settings name format: %q", name)
}
