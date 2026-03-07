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
// proto.service: google.cloud.modelarmor.v1.ModelArmor
// proto.message: google.cloud.modelarmor.v1.FloorSetting

package mockmodelarmor

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/modelarmor/v1"
)

func (s *ModelArmorV1) GetFloorSetting(ctx context.Context, req *pb.GetFloorSettingRequest) (*pb.FloorSetting, error) {
	name, err := s.parseFloorSettingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ModelArmorV1) UpdateFloorSetting(ctx context.Context, req *pb.UpdateFloorSettingRequest) (*pb.FloorSetting, error) {
	reqName := req.GetFloorSetting().GetName()
	name, err := s.parseFloorSettingName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Create if not found
			obj = proto.Clone(req.FloorSetting).(*pb.FloorSetting)
			obj.Name = fqn
			now := time.Now()
			obj.CreateTime = timestamppb.New(now)
			obj.UpdateTime = timestamppb.New(now)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Simple merge for now, if no update mask
		proto.Merge(obj, req.FloorSetting)
	} else {
		if err := fields.UpdateByFieldMask(obj, req.FloorSetting, paths); err != nil {
			return nil, err
		}
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type floorSettingName struct {
	Project  string
	Location string
}

func (n *floorSettingName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/floorSetting", n.Project, n.Location)
}

func (s *MockService) parseFloorSettingName(name string) (*floorSettingName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "floorSetting" {
		name := &floorSettingName{
			Project:  tokens[1],
			Location: tokens[3],
		}
		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
