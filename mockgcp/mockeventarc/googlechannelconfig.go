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
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.GoogleChannelConfig

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/eventarc/v1"
)

func (s *EventarcV1) GetGoogleChannelConfig(ctx context.Context, req *pb.GetGoogleChannelConfigRequest) (*pb.GoogleChannelConfig, error) {
	reqName := req.GetName()
	id, err := s.parseGoogleChannelConfigName(reqName)
	if err != nil {
		return nil, err
	}
	existing := &pb.GoogleChannelConfig{}
	if err := s.storage.Get(ctx, reqName, existing); err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}
	if existing.Name != "" {
		return existing, nil
	}

	minimalConfig := &pb.GoogleChannelConfig{
		Name:          reqName,
		UpdateTime:    timestamppb.New(time.Now()),
		CryptoKeyName: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", id.Project.ID, id.Location, "test-key-ring", "test-crypto-key"),
	}

	return minimalConfig, nil
}

func (s *EventarcV1) UpdateGoogleChannelConfig(ctx context.Context, req *pb.UpdateGoogleChannelConfigRequest) (*pb.GoogleChannelConfig, error) {
	reqName := req.GetGoogleChannelConfig().GetName()
	name, err := s.parseGoogleChannelConfigName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.GoogleChannelConfig{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	updated := &pb.GoogleChannelConfig{}
	if existing.Name == "" {
		updated.Name = reqName
	} else {
		updated = proto.Clone(existing).(*pb.GoogleChannelConfig)
	}

	if req.GetUpdateMask() == nil || len(req.GetUpdateMask().GetPaths()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "cryptoKeyName":
			updated.CryptoKeyName = req.GetGoogleChannelConfig().GetCryptoKeyName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field mask path %q not supported", path)
		}
	}
	updated.UpdateTime = timestamppb.New(time.Now())
	if existing.Name == "" {
		if err := s.storage.Create(ctx, fqn, updated); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Update(ctx, fqn, updated); err != nil {
			return nil, err
		}
	}
	return updated, nil
}

func applyUpdateMask(mask *fieldmaskpb.FieldMask, src *pb.GoogleChannelConfig, dest *pb.GoogleChannelConfig) error {
	for _, path := range mask.GetPaths() {
		switch path {
		case "cryptoKeyName":
			dest.CryptoKeyName = src.CryptoKeyName
		default:
			return fmt.Errorf("unexpected field path: %s", path)
		}
	}
	return nil
}

type googleChannelConfigName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *googleChannelConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/googleChannelConfig", n.Project.ID, n.Location)
}

// parseGoogleChannelConfigName parses a string into an googleChannelConfigName.
// The expected form is `projects/*/locations/*/googleChannelConfig`.
func (s *EventarcV1) parseGoogleChannelConfigName(name string) (*googleChannelConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "googleChannelConfig" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &googleChannelConfigName{
			Project:  project,
			Location: tokens[3],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
