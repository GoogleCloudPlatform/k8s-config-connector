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

package mocklivestream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type assetName struct {
	Project  *projects.ProjectData
	Location string
	Asset    string
}

func (n *assetName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/assets/" + n.Asset
}

func (s *MockService) parseAssetName(name string) (*assetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "assets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &assetName{
			Project:  project,
			Location: tokens[3],
			Asset:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *LivestreamServer) GetAsset(ctx context.Context, req *pb.GetAssetRequest) (*pb.Asset, error) {
	name, err := s.parseAssetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Asset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *LivestreamServer) CreateAsset(ctx context.Context, req *pb.CreateAssetRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/assets/" + req.AssetId
	name, err := s.parseAssetName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.Asset).(*pb.Asset)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Asset_ACTIVE
	if obj.Crc32C == "" {
		obj.Crc32C = "hq+43Q=="
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "create",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *LivestreamServer) DeleteAsset(ctx context.Context, req *pb.DeleteAssetRequest) (*longrunning.Operation, error) {
	name, err := s.parseAssetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Asset{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "delete",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *LivestreamServer) ListAssets(ctx context.Context, req *pb.ListAssetsRequest) (*pb.ListAssetsResponse, error) {
	response := &pb.ListAssetsResponse{}
	findPrefix := req.Parent + "/assets/"

	assetKind := (&pb.Asset{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, assetKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		response.Assets = append(response.Assets, obj.(*pb.Asset))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
