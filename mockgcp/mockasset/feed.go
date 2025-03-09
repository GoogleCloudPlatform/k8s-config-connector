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
// proto.service: google.cloud.asset.v1.AssetService
// proto.message: google.cloud.asset.v1.Feed

package mockasset

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/asset/v1"
)

type assetService struct {
	*MockService
	pb.UnimplementedAssetServiceServer
}

func (s *assetService) GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.Feed, error) {
	name, err := s.parseFeedName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Feed{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "feed %q not found in %q", name.feedID, name.parent)
		}
		return nil, err
	}

	return obj, nil
}

func (s *assetService) CreateFeed(ctx context.Context, req *pb.CreateFeedRequest) (*pb.Feed, error) {
	reqName := req.Parent + "/feeds/" + req.FeedId
	name, err := s.parseFeedName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	if err := s.storage.Create(ctx, fqn, req.Feed); err != nil {
		return nil, err
	}
	return req.Feed, nil
}

func (s *assetService) UpdateFeed(ctx context.Context, req *pb.UpdateFeedRequest) (*pb.Feed, error) {
	name, err := s.parseFeedName(req.GetFeed().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Feed{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if req.UpdateMask != nil {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "asset_names":
				obj.AssetNames = req.Feed.AssetNames
			case "asset_types":
				obj.AssetTypes = req.Feed.AssetTypes
			case "content_type":
				obj.ContentType = req.Feed.ContentType
			case "feed_output_config":
				obj.FeedOutputConfig = req.Feed.FeedOutputConfig
			case "condition":
				obj.Condition = req.Feed.Condition
			case "relationship_types":
				obj.RelationshipTypes = req.Feed.RelationshipTypes
			default:
				return nil, fmt.Errorf("unexpected field mask path: %q", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *assetService) DeleteFeed(ctx context.Context, req *pb.DeleteFeedRequest) (*emptypb.Empty, error) {
	name, err := s.parseFeedName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Feed{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type feedName struct {
	parent string
	feedID string
}

func (n *feedName) String() string {
	return fmt.Sprintf("%s/feeds/%s", n.parent, n.feedID)
}

// parseFeedName parses a string into an feedName.
func (s *MockService) parseFeedName(name string) (*feedName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] != "" && tokens[2] == "feeds" {
		name := &feedName{
			parent: strings.Join(tokens[0:3], "/"),
			feedID: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}



