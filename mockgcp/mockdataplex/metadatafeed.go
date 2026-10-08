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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.MetadataFeed

package mockdataplex

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *CatalogService) GetMetadataFeed(ctx context.Context, req *pb.GetMetadataFeedRequest) (*pb.MetadataFeed, error) {
	name, err := s.parseMetadataFeedName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MetadataFeed{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) ListMetadataFeeds(ctx context.Context, req *pb.ListMetadataFeedsRequest) (*pb.ListMetadataFeedsResponse, error) {
	response := &pb.ListMetadataFeedsResponse{}

	metadataFeedKind := (&pb.MetadataFeed{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, metadataFeedKind, storage.ListOptions{}, func(obj proto.Message) error {
		metadataFeed := obj.(*pb.MetadataFeed)
		if strings.HasPrefix(metadataFeed.GetName(), req.Parent) {
			response.MetadataFeeds = append(response.MetadataFeeds, metadataFeed)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *CatalogService) CreateMetadataFeed(ctx context.Context, req *pb.CreateMetadataFeedRequest) (*longrunningpb.Operation, error) {
	feedID := req.GetMetadataFeedId()
	if feedID == "" {
		feedID = fmt.Sprintf("metadata-feed-%d", time.Now().UnixNano())
	}
	reqName := fmt.Sprintf("%s/metadataFeeds/%s", req.GetParent(), feedID)
	name, err := s.parseMetadataFeedName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetMetadataFeed())
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateDefaultsForMetadataFeed(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *MockService) populateDefaultsForMetadataFeed(obj *pb.MetadataFeed) {
	if obj.Filters == nil {
		obj.Filters = &pb.MetadataFeed_Filters{}
	}
	if obj.Scope != nil {
		for i, p := range obj.Scope.Projects {
			obj.Scope.Projects[i] = s.replaceProjectIDWithNumber(p)
		}
		for i, eg := range obj.Scope.EntryGroups {
			obj.Scope.EntryGroups[i] = s.replaceProjectIDWithNumber(eg)
		}
	}
}

func (s *MockService) replaceProjectIDWithNumber(link string) string {
	tokens := strings.Split(link, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err == nil {
			tokens[1] = strconv.FormatInt(project.Number, 10)
			return strings.Join(tokens, "/")
		}
	}
	return link
}

func (s *CatalogService) UpdateMetadataFeed(ctx context.Context, req *pb.UpdateMetadataFeedRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMetadataFeedName(req.GetMetadataFeed().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.MetadataFeed{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, req.GetMetadataFeed())
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetMetadataFeed().GetLabels()
			case "filters":
				obj.Filters = req.GetMetadataFeed().GetFilters()
			case "scope":
				obj.Scope = req.GetMetadataFeed().GetScope()
			case "pubsub_topic", "pubsubTopic", "endpoint":
				obj.Endpoint = req.GetMetadataFeed().GetEndpoint()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	s.populateDefaultsForMetadataFeed(obj)

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) DeleteMetadataFeed(ctx context.Context, req *pb.DeleteMetadataFeedRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMetadataFeedName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.MetadataFeed{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type metadataFeedName struct {
	Project        *projects.ProjectData
	Location       string
	MetadataFeedID string
}

func (n *metadataFeedName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/metadataFeeds/%s", n.Project.ID, n.Location, n.MetadataFeedID)
}

// parseMetadataFeedName parses a string into a metadataFeedName.
// The expected form is `projects/*/locations/*/metadataFeeds/*`.
func (s *MockService) parseMetadataFeedName(name string) (*metadataFeedName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataFeeds" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &metadataFeedName{
			Project:        project,
			Location:       tokens[3],
			MetadataFeedID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
