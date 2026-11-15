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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
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

	fqn := name.String()

	obj := proto.CloneOf(req.Asset)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "111111111111111111111"
	obj.State = pb.State_ACTIVE
	if obj.DiscoveryStatus == nil {
		obj.DiscoveryStatus = &pb.Asset_DiscoveryStatus{}
	}
	if obj.ResourceStatus == nil {
		obj.ResourceStatus = &pb.Asset_ResourceStatus{
			State: pb.Asset_ResourceStatus_READY,
		}
	}
	if obj.SecurityStatus == nil {
		obj.SecurityStatus = &pb.Asset_SecurityStatus{
			State: pb.Asset_SecurityStatus_READY,
		}
	}
	if obj.ResourceSpec != nil {
		if obj.ResourceSpec.ReadAccessMode == pb.Asset_ResourceSpec_ACCESS_MODE_UNSPECIFIED {
			obj.ResourceSpec.ReadAccessMode = pb.Asset_ResourceSpec_DIRECT
		}
		if obj.ResourceSpec.Type == pb.Asset_ResourceSpec_STORAGE_BUCKET {
			tokens := strings.Split(obj.ResourceSpec.Name, "/buckets/")
			if len(tokens) == 2 {
				obj.ResourceSpec.Name = tokens[1]
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:                name.String(),
		Verb:                  "create",
		CreateTime:            timestamppb.New(time.Now()),
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataplexService) UpdateAsset(ctx context.Context, req *pb.UpdateAssetRequest) (*longrunning.Operation, error) {
	name, err := s.parseAssetName(req.Asset.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Asset{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	updated := proto.CloneOf(existing)
	proto.Merge(updated, req.Asset)
	updated.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:                name.String(),
		Verb:                  "update",
		CreateTime:            timestamppb.New(time.Now()),
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
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
		Target:                name.String(),
		Verb:                  "delete",
		CreateTime:            timestamppb.New(time.Now()),
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
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
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &assetName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
			ZoneID:   tokens[7],
			AssetID:  tokens[9],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/locations/*/lakes/*/zones/*/assets/*)", name)
}
