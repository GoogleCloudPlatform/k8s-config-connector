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

package mockiap

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/iap/apiv1/iappb"
)

func (s *MockService) parseIapSettingsName(name string) (string, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return "", status.Errorf(codes.NotFound, "project %q not found", tokens[1])
		}
		tokens[1] = strconv.FormatInt(project.Number, 10)
		return strings.Join(tokens, "/"), nil
	}
	return name, nil
}

// GetIapSettings gets the IAP settings on a resource.
func (s *IdentityAwareProxyAdminService) GetIapSettings(ctx context.Context, req *pb.GetIapSettingsRequest) (*pb.IapSettings, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be specified")
	}

	fqn, err := s.parseIapSettingsName(name)
	if err != nil {
		return nil, err
	}

	obj := &pb.IapSettings{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// In GCP, IAP settings always exist on any valid resource, even if they've never been modified.
			// Return a default empty settings object.
			defaultObj := &pb.IapSettings{
				Name: fqn,
			}
			return defaultObj, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get iap settings: %v", err)
	}

	return obj, nil
}

// UpdateIapSettings updates the IAP settings on a resource.
func (s *IdentityAwareProxyAdminService) UpdateIapSettings(ctx context.Context, req *pb.UpdateIapSettingsRequest) (*pb.IapSettings, error) {
	desired := req.GetIapSettings()
	if desired == nil {
		return nil, status.Errorf(codes.InvalidArgument, "iap_settings must be specified")
	}

	name := desired.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be specified")
	}

	fqn, err := s.parseIapSettingsName(name)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing settings or use a new default one.
	existing := &pb.IapSettings{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			existing = &pb.IapSettings{
				Name: fqn,
			}
		} else {
			return nil, status.Errorf(codes.Internal, "failed to check for existing iap settings: %v", err)
		}
	}

	// For mockgcp, we can just replace the stored settings with the desired ones.
	updated := proto.Clone(desired).(*pb.IapSettings)
	updated.Name = fqn

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		if status.Code(err) == codes.NotFound {
			if err := s.storage.Create(ctx, fqn, updated); err != nil {
				return nil, status.Errorf(codes.Internal, "failed to create iap settings: %v", err)
			}
		} else {
			return nil, status.Errorf(codes.Internal, "failed to update iap settings: %v", err)
		}
	}

	return updated, nil
}
