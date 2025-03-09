// Copyright 2024 Google LLC
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
// proto.service: google.api.cloudquotas.v1.CloudQuotas
// proto.message: google.api.cloudquotas.v1.QuotaPreference

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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/cloudquotas/v1"
)

type CloudQuotasV1 struct {
	*MockService
	pb.UnimplementedCloudQuotasServer
}

func (s *CloudQuotasV1) GetQuotaPreference(ctx context.Context, req *pb.GetQuotaPreferenceRequest) (*pb.QuotaPreference, error) {
	name, err := s.parseQuotaPreferenceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.QuotaPreference{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CloudQuotasV1) CreateQuotaPreference(ctx context.Context, req *pb.CreateQuotaPreferenceRequest) (*pb.QuotaPreference, error) {
	reqName := fmt.Sprintf("%s/quotaPreferences/%s", req.GetParent(), req.GetQuotaPreferenceId())
	name, err := s.parseQuotaPreferenceName(reqName)
	if err != nil {
		return nil, err
		}

	fqn := name.String()

	obj := proto.Clone(req.GetQuotaPreference()).(*pb.QuotaPreference)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CloudQuotasV1) UpdateQuotaPreference(ctx context.Context, req *pb.UpdateQuotaPreferenceRequest) (*pb.QuotaPreference, error) {
	reqName := req.GetQuotaPreference().GetName()

	name, err := s.parseQuotaPreferenceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.QuotaPreference{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	now := time.Now()

	obj.UpdateTime = timestamppb.New(now)

	proto.Merge(obj, req.QuotaPreference)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type quotaPreferenceName struct {
	Parent              string
	QuotaPreferenceName string
}

func (n *quotaPreferenceName) String() string {
	return n.Parent + "/quotaPreferences/" + n.QuotaPreferenceName
}

// parseQuotaPreferenceName parses a string into a quotaPreferenceName.
// The expected form is `projects/*/locations/*/quotaPreferences/*`.
func (s *MockService) parseQuotaPreferenceName(name string) (*quotaPreferenceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[2] == "locations" && tokens[4] == "quotaPreferences" {
		if tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations" {
			name := &quotaPreferenceName{
				Parent:              strings.Join(tokens[0:4], "/"),
				QuotaPreferenceName: tokens[5],
			}

			return name, nil
		}
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}



