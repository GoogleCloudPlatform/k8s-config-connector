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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Asset

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *DataplexService) GetAsset(ctx context.Context, req *pb.GetAssetRequest) (*pb.Asset, error) {
	name, err := s.parseAssetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Asset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataplexService) CreateAsset(ctx context.Context, req *pb.CreateAssetRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/assets/" + req.AssetId
	name, err := s.parseAssetName(reqName)
	if err != nil {
		return nil, err
	}

	parentZone := &pb.Zone{}
	parentFQN := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", name.Project.ID, name.Location, name.LakeID, name.ZoneID)
	if err := s.storage.Get(ctx, parentFQN, parentZone); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentFQN)
		}
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Asset).(*pb.Asset)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "111111111111111111111"
	obj.State = pb.State_ACTIVE

	s.populateAssetDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		Target:     name.String(),
		Verb:       "create",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataplexService) UpdateAsset(ctx context.Context, req *pb.UpdateAssetRequest) (*longrunning.Operation, error) {
	name, err := s.parseAssetName(req.GetAsset().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Asset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	updateMask := req.GetUpdateMask()

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			obj.Description = req.GetAsset().GetDescription()
		case "display_name", "displayName":
			obj.DisplayName = req.GetAsset().GetDisplayName()
		case "labels":
			obj.Labels = req.GetAsset().GetLabels()
		case "discovery_spec", "discoverySpec":
			obj.DiscoverySpec = req.GetAsset().GetDiscoverySpec()
			if obj.DiscoveryStatus == nil {
				obj.DiscoveryStatus = &pb.Asset_DiscoveryStatus{
					UpdateTime: timestamppb.New(time.Now()),
				}
			}
			if obj.DiscoverySpec != nil && obj.DiscoverySpec.Enabled {
				obj.DiscoveryStatus.State = pb.Asset_DiscoveryStatus_IN_PROGRESS
				if obj.DiscoveryStatus.LastRunTime == nil {
					obj.DiscoveryStatus.LastRunTime = timestamppb.New(time.Now())
				}
			} else {
				obj.DiscoveryStatus.State = pb.Asset_DiscoveryStatus_DISABLED
			}
		case "resource_spec", "resourceSpec":
			if req.GetAsset().GetResourceSpec() != nil {
				obj.ResourceSpec = proto.Clone(req.GetAsset().GetResourceSpec()).(*pb.Asset_ResourceSpec)
				s.normalizeResourceSpec(obj.ResourceSpec)
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of field %q for Asset", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataplexService) ListAssets(ctx context.Context, req *pb.ListAssetsRequest) (*pb.ListAssetsResponse, error) {
	_, err := s.parseZoneName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListAssetsResponse{}

	assetKind := (&pb.Asset{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, assetKind, storage.ListOptions{}, func(obj proto.Message) error {
		asset := obj.(*pb.Asset)
		if strings.HasPrefix(asset.GetName(), req.Parent+"/assets/") {
			response.Assets = append(response.Assets, asset)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DataplexService) DeleteAsset(ctx context.Context, req *pb.DeleteAssetRequest) (*longrunning.Operation, error) {
	name, err := s.parseAssetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Asset{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		Target:     name.String(),
		Verb:       "delete",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *DataplexService) normalizeResourceSpec(spec *pb.Asset_ResourceSpec) {
	if spec == nil {
		return
	}
	if spec.Name != "" {
		tokens := strings.Split(spec.Name, "/")
		spec.Name = tokens[len(tokens)-1]
	}
	if spec.ReadAccessMode == pb.Asset_ResourceSpec_ACCESS_MODE_UNSPECIFIED {
		spec.ReadAccessMode = pb.Asset_ResourceSpec_DIRECT
	}
}

func (s *DataplexService) populateAssetDefaults(obj *pb.Asset) {
	s.normalizeResourceSpec(obj.ResourceSpec)

	if obj.ResourceStatus == nil {
		obj.ResourceStatus = &pb.Asset_ResourceStatus{
			State:      pb.Asset_ResourceStatus_READY,
			UpdateTime: timestamppb.New(time.Now()),
		}
	}
	if obj.SecurityStatus == nil {
		obj.SecurityStatus = &pb.Asset_SecurityStatus{
			State:      pb.Asset_SecurityStatus_READY,
			UpdateTime: timestamppb.New(time.Now()),
		}
	}
	if obj.DiscoveryStatus == nil {
		obj.DiscoveryStatus = &pb.Asset_DiscoveryStatus{
			UpdateTime: timestamppb.New(time.Now()),
		}
		if obj.DiscoverySpec != nil && obj.DiscoverySpec.Enabled {
			obj.DiscoveryStatus.State = pb.Asset_DiscoveryStatus_IN_PROGRESS
			obj.DiscoveryStatus.LastRunTime = timestamppb.New(time.Now())
		} else {
			obj.DiscoveryStatus.State = pb.Asset_DiscoveryStatus_DISABLED
		}
	}
}

type assetName struct {
	Project  *projects.ProjectData
	Location string
	LakeID   string
	ZoneID   string
	AssetID  string
}

func (n *assetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", n.Project.ID, n.Location, n.LakeID, n.ZoneID, n.AssetID)
}

func (s *MockService) parseAssetName(name string) (*assetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "zones" && tokens[8] == "assets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &assetName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
			ZoneID:   tokens[7],
			AssetID:  tokens[9],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/locations/*/lakes/*/zones/*/assets/*)", name)
}
