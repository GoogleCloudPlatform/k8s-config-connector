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

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type FeatureViewName struct {
	Project              *projects.ProjectData
	Location             string
	FeatureOnlineStoreID string
	FeatureViewID        string
}

func (n *FeatureViewName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/featureOnlineStores/%s/featureViews/%s", n.Project.Number, n.Location, n.FeatureOnlineStoreID, n.FeatureViewID)
}

func (s *MockService) parseFeatureViewName(name string) (*FeatureViewName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "featureOnlineStores" && tokens[6] == "featureViews" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &FeatureViewName{
			Project:              project,
			Location:             tokens[3],
			FeatureOnlineStoreID: tokens[5],
			FeatureViewID:        tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *featureOnlineStoreAdminService) GetFeatureView(ctx context.Context, req *pb.GetFeatureViewRequest) (*pb.FeatureView, error) {
	name, err := s.parseFeatureViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FeatureView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The FeatureView %s does not exist.", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *featureOnlineStoreAdminService) CreateFeatureView(ctx context.Context, req *pb.CreateFeatureViewRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/featureViews/" + req.FeatureViewId
	name, err := s.parseFeatureViewName(reqName)
	if err != nil {
		return nil, err
	}

	// Verify parent FeatureOnlineStore exists
	parentName, err := s.parseFeatureOnlineStoreName(req.Parent)
	if err != nil {
		return nil, err
	}
	parentFqn := parentName.String()
	parentObj := &pb.FeatureOnlineStore{}
	if err := s.storage.Get(ctx, parentFqn, parentObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The FeatureOnlineStore %q does not exist.", req.Parent)
		}
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.FeatureView)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateFeatureViewOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return &pb.FeatureView{Name: fqn}, nil
	})
}

func (s *featureOnlineStoreAdminService) UpdateFeatureView(ctx context.Context, req *pb.UpdateFeatureViewRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeatureViewName(req.GetFeatureView().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.FeatureView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	oldSyncConfig := obj.SyncConfig

	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "labels":
			obj.Labels = req.GetFeatureView().GetLabels()
		case "sync_config", "syncConfig":
			obj.SyncConfig = req.GetFeatureView().GetSyncConfig()
		case "big_query_source", "bigQuerySource":
			obj.Source = req.GetFeatureView().GetSource()
		case "feature_registry_source", "featureRegistrySource":
			obj.Source = req.GetFeatureView().GetSource()
		case "vertex_rag_source", "vertexRagSource":
			obj.Source = req.GetFeatureView().GetSource()
		case "index_config", "indexConfig":
			obj.IndexConfig = req.GetFeatureView().GetIndexConfig()
		case "optimized_config", "optimizedConfig":
			obj.OptimizedConfig = req.GetFeatureView().GetOptimizedConfig()
		case "service_agent_type", "serviceAgentType":
			obj.ServiceAgentType = req.GetFeatureView().GetServiceAgentType()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateFeatureViewOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	if obj.SyncConfig != nil && oldSyncConfig != nil {
		return s.operations.DoneLROWithMetadata(ctx, opPrefix, op, &pb.FeatureView{Name: fqn})
	}
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return &pb.FeatureView{Name: fqn}, nil
	})
}

func (s *featureOnlineStoreAdminService) DeleteFeatureView(ctx context.Context, req *pb.DeleteFeatureViewRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeatureViewName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.FeatureView{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s/featureOnlineStores/%s", name.Project.Number, name.Location, name.FeatureOnlineStoreID)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *featureOnlineStoreAdminService) ListFeatureViews(ctx context.Context, req *pb.ListFeatureViewsRequest) (*pb.ListFeatureViewsResponse, error) {
	name, err := s.parseFeatureOnlineStoreName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListFeatureViewsResponse{}
	prefix := name.String() + "/featureViews/"
	kind := (&pb.FeatureView{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		view := obj.(*pb.FeatureView)
		response.FeatureViews = append(response.FeatureViews, view)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
