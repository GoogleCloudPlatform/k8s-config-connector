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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type featurestoreService struct {
	*MockService
	pb.UnimplementedFeaturestoreServiceServer
}

func (s *featurestoreService) GetFeaturestore(ctx context.Context, req *pb.GetFeaturestoreRequest) (*pb.Featurestore, error) {
	name, err := s.parseFeaturestoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Featurestore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The Featurestore does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *featurestoreService) CreateFeaturestore(ctx context.Context, req *pb.CreateFeaturestoreRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/featurestores/" + req.FeaturestoreId
	name, err := s.parseFeaturestoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.Featurestore)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)
	obj.State = pb.Featurestore_STABLE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateFeaturestoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// LRO response only contains name field
		result := &pb.Featurestore{}
		result.Name = fqn
		return result, nil
	})
}

func (s *featurestoreService) UpdateFeaturestore(ctx context.Context, req *pb.UpdateFeaturestoreRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeaturestoreName(req.GetFeaturestore().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Featurestore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "etag":
			obj.Etag = req.GetFeaturestore().GetEtag()
		case "labels":
			obj.Labels = req.GetFeaturestore().GetLabels()
		case "onlineServingConfig.fixedNodeCount":
			obj.OnlineServingConfig.FixedNodeCount = req.GetFeaturestore().GetOnlineServingConfig().GetFixedNodeCount()
		case "onlineStorageTtlDays":
			obj.OnlineStorageTtlDays = req.GetFeaturestore().GetOnlineStorageTtlDays()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateFeaturestoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// LRO response only contains name field
		result := &pb.Featurestore{}
		result.Name = fqn
		return result, nil
	})
}

func (s *featurestoreService) DeleteFeaturestore(ctx context.Context, req *pb.DeleteFeaturestoreRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeaturestoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Featurestore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type FeaturestoreName struct {
	Project        *projects.ProjectData
	Location       string
	FeaturestoreID string
}

func (n *FeaturestoreName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/featurestores/%s", n.Project.Number, n.Location, n.FeaturestoreID)
}

// parseFeaturestoreName parses a string into a FeaturestoreName.
// The expected form of input string is projects/<projectID>/locations/<location>/featurestores/<featurestoreID>
func (s *MockService) parseFeaturestoreName(name string) (*FeaturestoreName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "featurestores" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &FeaturestoreName{
			Project:        project,
			Location:       tokens[3],
			FeaturestoreID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *featurestoreService) GetEntityType(ctx context.Context, req *pb.GetEntityTypeRequest) (*pb.EntityType, error) {
	name, err := s.parseEntityTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntityType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType %s is not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *featurestoreService) CreateEntityType(ctx context.Context, req *pb.CreateEntityTypeRequest) (*longrunning.Operation, error) {
	parent, err := s.parseFeaturestoreName(req.Parent)
	if err != nil {
		return nil, err
	}

	entityTypeID := req.EntityTypeId
	if entityTypeID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "entityTypeId is required")
	}

	name := &EntityTypeName{
		Project:        parent.Project,
		Location:       parent.Location,
		FeaturestoreID: parent.FeaturestoreID,
		EntityTypeID:   entityTypeID,
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.EntityType).(*pb.EntityType)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	if obj.MonitoringConfig == nil {
		obj.MonitoringConfig = &pb.FeaturestoreMonitoringConfig{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateEntityTypeOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := &pb.EntityType{}
		result.Name = fqn
		return result, nil
	})
}

func (s *featurestoreService) UpdateEntityType(ctx context.Context, req *pb.UpdateEntityTypeRequest) (*pb.EntityType, error) {
	reqObj := req.GetEntityType()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "entity_type is required")
	}

	name, err := s.parseEntityTypeName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntityType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType %s is not found.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, reqObj, paths); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update EntityType fields: %v", err)
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *featurestoreService) DeleteEntityType(ctx context.Context, req *pb.DeleteEntityTypeRequest) (*longrunning.Operation, error) {
	name, err := s.parseEntityTypeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.EntityType{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type EntityTypeName struct {
	Project        *projects.ProjectData
	Location       string
	FeaturestoreID string
	EntityTypeID   string
}

func (n *EntityTypeName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/featurestores/%s/entityTypes/%s", n.Project.Number, n.Location, n.FeaturestoreID, n.EntityTypeID)
}

func (s *MockService) parseEntityTypeName(name string) (*EntityTypeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "featurestores" && tokens[6] == "entityTypes" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &EntityTypeName{
			Project:        project,
			Location:       tokens[3],
			FeaturestoreID: tokens[5],
			EntityTypeID:   tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *featurestoreService) GetFeature(ctx context.Context, req *pb.GetFeatureRequest) (*pb.Feature, error) {
	name, err := s.parseFeatureName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Feature %s is not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *featurestoreService) CreateFeature(ctx context.Context, req *pb.CreateFeatureRequest) (*longrunning.Operation, error) {
	parent, err := s.parseEntityTypeName(req.Parent)
	if err != nil {
		return nil, err
	}

	featureID := req.FeatureId
	if featureID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "featureId is required")
	}

	name := &FeatureName{
		Project:        parent.Project,
		Location:       parent.Location,
		FeaturestoreID: parent.FeaturestoreID,
		EntityTypeID:   parent.EntityTypeID,
		FeatureID:      featureID,
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Feature).(*pb.Feature)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	if obj.MonitoringConfig == nil {
		obj.MonitoringConfig = &pb.FeaturestoreMonitoringConfig{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateFeatureOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := &pb.Feature{}
		result.Name = fqn
		return result, nil
	})
}

func (s *featurestoreService) UpdateFeature(ctx context.Context, req *pb.UpdateFeatureRequest) (*pb.Feature, error) {
	reqObj := req.GetFeature()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "feature is required")
	}

	name, err := s.parseFeatureName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Feature %s is not found.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, reqObj, paths); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update Feature fields: %v", err)
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *featurestoreService) DeleteFeature(ctx context.Context, req *pb.DeleteFeatureRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeatureName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Feature{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	// Note: delete feature LRO name is projects/{project}/locations/{location}/featurestores/{fs}/entityTypes/{et}/operations/{op}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s/featurestores/%s/entityTypes/%s", name.Project.Number, name.Location, name.FeaturestoreID, name.EntityTypeID)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type FeatureName struct {
	Project        *projects.ProjectData
	Location       string
	FeaturestoreID string
	EntityTypeID   string
	FeatureID      string
}

func (n *FeatureName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/featurestores/%s/entityTypes/%s/features/%s", n.Project.Number, n.Location, n.FeaturestoreID, n.EntityTypeID, n.FeatureID)
}

func (s *MockService) parseFeatureName(name string) (*FeatureName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "featurestores" && tokens[6] == "entityTypes" && tokens[8] == "features" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &FeatureName{
			Project:        project,
			Location:       tokens[3],
			FeaturestoreID: tokens[5],
			EntityTypeID:   tokens[7],
			FeatureID:      tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
