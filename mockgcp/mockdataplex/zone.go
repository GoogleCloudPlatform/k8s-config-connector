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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Zone

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunning "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *DataplexService) GetZone(ctx context.Context, req *pb.GetZoneRequest) (*pb.Zone, error) {
	name, err := s.parseZoneName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Zone{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
	}

	return obj, nil
}

func (s *DataplexService) CreateZone(ctx context.Context, req *pb.CreateZoneRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/zones/" + req.ZoneId
	name, err := s.parseZoneName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Zone).(*pb.Zone)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "zone-" + name.ZoneID // TODO: maybe a proper random value?
	obj.State = pb.State_ACTIVE
	if obj.AssetStatus == nil {
		obj.AssetStatus = &pb.AssetStatus{UpdateTime: timestamppb.New(time.Now())}
	}
	if obj.DiscoverySpec == nil {
		obj.DiscoverySpec = &pb.Zone_DiscoverySpec{
			CsvOptions:  &pb.Zone_DiscoverySpec_CsvOptions{},
			JsonOptions: &pb.Zone_DiscoverySpec_JsonOptions{},
			Trigger:     &pb.Zone_DiscoverySpec_Schedule{Schedule: ""},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		obj.AssetStatus = &pb.AssetStatus{}
		// Update the object in storage after LRO completion simulation
		return obj, nil
	})
}

func (s *DataplexService) UpdateZone(ctx context.Context, req *pb.UpdateZoneRequest) (*longrunning.Operation, error) {
	name, err := s.parseZoneName(req.GetZone().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Zone{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	updateMask := req.GetUpdateMask()

	// Field behavior immutable means we cannot update these fields.
	// We only allow updates for mutable fields.
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			obj.Description = req.GetZone().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetZone().GetDisplayName()
		case "labels":
			obj.Labels = req.GetZone().GetLabels()
		case "discovery_spec":
			obj.DiscoverySpec = req.GetZone().GetDiscoverySpec()
		// Immutable fields 'type' and 'resource_spec' are not updatable
		// case "type":
		// case "resource_spec":
		default:
			// For other unhandled fields, return a generic error or log a warning.
			return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of field %q for Zone", path)

		}
	}
	if obj.DiscoverySpec == nil {
		obj.DiscoverySpec = &pb.Zone_DiscoverySpec{
			CsvOptions:  &pb.Zone_DiscoverySpec_CsvOptions{},
			JsonOptions: &pb.Zone_DiscoverySpec_JsonOptions{},
			Trigger:     &pb.Zone_DiscoverySpec_Schedule{Schedule: ""},
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		obj.AssetStatus = &pb.AssetStatus{UpdateTime: timestamppb.New(time.Now())}
		return obj, nil
	})
}

func (s *DataplexService) ListZones(ctx context.Context, req *pb.ListZonesRequest) (*pb.ListZonesResponse, error) {
	_, err := s.parseLakeName(req.Parent) // Validate parent format
	if err != nil {
		return nil, err
	}

	response := &pb.ListZonesResponse{}

	zoneKind := (&pb.Zone{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, zoneKind, storage.ListOptions{}, func(obj proto.Message) error {
		zone := obj.(*pb.Zone)
		// Check if the zone's name starts with the correct parent prefix
		// (e.g., projects/p/locations/l/lakes/lk)
		if strings.HasPrefix(zone.GetName(), req.Parent+"/") {
			response.Zones = append(response.Zones, zone)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DataplexService) DeleteZone(ctx context.Context, req *pb.DeleteZoneRequest) (*longrunning.Operation, error) {
	name, err := s.parseZoneName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Zone{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err // Return other storage errors
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type zoneName struct {
	Project  *projects.ProjectData
	Location string
	LakeID   string
	ZoneID   string
}

func (n *zoneName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", n.Project.ID, n.Location, n.LakeID, n.ZoneID)
}

// parseZoneName parses a string into a zoneName.
// The expected form is `projects/*/locations/*/lakes/*/zones/*`.
func (s *MockService) parseZoneName(name string) (*zoneName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "zones" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zoneName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
			ZoneID:   tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/locations/*/lakes/*/zones/*)", name)
}
