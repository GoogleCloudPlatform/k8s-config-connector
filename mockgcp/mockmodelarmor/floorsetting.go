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

	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

func (s *ModelArmorV1) GetFloorSetting(ctx context.Context, req *pb.GetFloorSettingRequest) (*pb.FloorSetting, error) {
	name, err := s.parseFloorSettingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// On real GCP, global floorSetting is a singleton that always exists by default.
			// Return and store a default FloorSetting object if it hasn't been updated/created.
			now := time.Now()
			obj = &pb.FloorSetting{
				Name:                          fqn,
				EnableFloorSettingEnforcement: proto.Bool(true),
				FilterConfig: &pb.FilterConfig{
					PiAndJailbreakFilterSettings: &pb.PiAndJailbreakFilterSettings{
						FilterEnforcement: pb.PiAndJailbreakFilterSettings_ENABLED,
						ConfidenceLevel:   pb.DetectionConfidenceLevel_LOW_AND_ABOVE,
					},
				},
				FloorSettingMetadata: &pb.FloorSetting_FloorSettingMetadata{
					MultiLanguageDetection: &pb.FloorSetting_FloorSettingMetadata_MultiLanguageDetection{
						EnableMultiLanguageDetection: true,
					},
				},
				CreateTime: timestamppb.New(now),
				UpdateTime: timestamppb.New(now),
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

func (s *ModelArmorV1) UpdateFloorSetting(ctx context.Context, req *pb.UpdateFloorSettingRequest) (*pb.FloorSetting, error) {
	name, err := s.parseFloorSettingName(req.GetFloorSetting().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	// Get existing or create default
	obj := &pb.FloorSetting{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			now := time.Now()
			obj = &pb.FloorSetting{
				Name:       fqn,
				CreateTime: timestamppb.New(now),
				UpdateTime: timestamppb.New(now),
			}
		} else {
			return nil, err
		}
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, req.GetFloorSetting())
	} else {
		if err := fields.UpdateByFieldMask(obj, req.GetFloorSetting(), paths); err != nil {
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

// parseFloorSettingName parses a string into a floorSettingName.
// The expected form is `projects/*/locations/*/floorSetting`.
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
