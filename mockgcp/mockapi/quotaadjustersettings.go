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
// proto.service: google.api.cloudquotas.v1beta.QuotaAdjusterSettingsManager
// proto.message: google.api.cloudquotas.v1beta.QuotaAdjusterSettings

package mockcloudquotas

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/cloudquotas/v1beta"
)

func (s *QuotaAdjusterSettingsManagerV1Beta) GetQuotaAdjusterSettings(ctx context.Context, req *pb.GetQuotaAdjusterSettingsRequest) (*pb.QuotaAdjusterSettings, error) {
	name, err := s.parseQuotaAdjusterSettingsName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QuotaAdjusterSettings{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj = s.buildDefaultQuotaAdjusterSettings(name)
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

func (s *QuotaAdjusterSettingsManagerV1Beta) UpdateQuotaAdjusterSettings(ctx context.Context, req *pb.UpdateQuotaAdjusterSettingsRequest) (*pb.QuotaAdjusterSettings, error) {
	reqName := req.GetQuotaAdjusterSettings().GetName()
	name, err := s.parseQuotaAdjusterSettingsName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.QuotaAdjusterSettings{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		// If there is no existing QuotaAdjusterSettings, create a default one.
		existing = s.buildDefaultQuotaAdjusterSettings(name)
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	updated := proto.Clone(existing).(*pb.QuotaAdjusterSettings)
	updated.UpdateTime = timestamppb.Now()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "enablement":
			updated.Enablement = req.GetQuotaAdjusterSettings().GetEnablement()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	// The etag is missing or empty, or if provided, does not match the stored etag.
	if req.GetQuotaAdjusterSettings().GetEtag() != "" && req.GetQuotaAdjusterSettings().GetEtag() != existing.GetEtag() {
		return nil, status.Errorf(codes.Aborted, "etag mismatch: update of QuotaAdjusterSettings blocked")
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *QuotaAdjusterSettingsManagerV1Beta) buildDefaultQuotaAdjusterSettings(n *quotaAdjusterSettingsName) *pb.QuotaAdjusterSettings {
	return &pb.QuotaAdjusterSettings{
		Name:       n.String(),
		Enablement: pb.QuotaAdjusterSettings_DISABLED,
		UpdateTime: timestamppb.New(time.Now()),
	}
}

type quotaAdjusterSettingsName struct {
	Project *projects.ProjectData
}

func (q *quotaAdjusterSettingsName) String() string {
	return fmt.Sprintf("projects/%s/locations/global/quotaAdjusterSettings", q.Project.ID)
}

func (s *MockService) parseQuotaAdjusterSettingsName(name string) (*quotaAdjusterSettingsName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "quotaAdjusterSettings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &quotaAdjusterSettingsName{
			Project: project,
		}

		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
