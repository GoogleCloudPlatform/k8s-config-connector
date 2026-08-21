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

package mockrapidmigrationassessment

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/rapidmigrationassessment/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type collectorName struct {
	Project     *projects.ProjectData
	Location    string
	CollectorID string
}

func (n *collectorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collectors/%s", n.Project.ID, n.Location, n.CollectorID)
}

func (s *MockService) parseCollectorName(name string) (*collectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collectors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &collectorName{
			Project:     project,
			Location:    tokens[3],
			CollectorID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *RapidMigrationAssessmentV1) GetCollector(ctx context.Context, req *pb.GetCollectorRequest) (*pb.Collector, error) {
	name, err := s.parseCollectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Collector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RapidMigrationAssessmentV1) CreateCollector(ctx context.Context, req *pb.CreateCollectorRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/collectors/%s", req.GetParent(), req.GetCollectorId())
	name, err := s.parseCollectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetCollector()).(*pb.Collector)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// Populate output-only/read-only fields
	obj.State = pb.Collector_STATE_READY_TO_USE
	obj.Bucket = "normalized-bucket"
	obj.GuestOsScan = &pb.GuestOsScan{
		CoreSource: fmt.Sprintf("projects/%d/locations/%s/sources/normalized-guest-os-scan-source", name.Project.Number, name.Location),
	}
	obj.VsphereScan = &pb.VSphereScan{
		CoreSource: fmt.Sprintf("projects/%d/locations/%s/sources/normalized-vsphere-scan-source", name.Project.Number, name.Location),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		responseObj := proto.Clone(obj).(*pb.Collector)
		responseObj.Labels = nil
		return responseObj, nil
	})
}

func (s *RapidMigrationAssessmentV1) UpdateCollector(ctx context.Context, req *pb.UpdateCollectorRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCollectorName(req.GetCollector().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Collector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetCollector().GetLabels()
		case "description":
			obj.Description = req.GetCollector().GetDescription()
		case "display_name", "displayName":
			obj.DisplayName = req.GetCollector().GetDisplayName()
		case "expected_asset_count", "expectedAssetCount":
			obj.ExpectedAssetCount = req.GetCollector().GetExpectedAssetCount()
		case "collection_days", "collectionDays":
			obj.CollectionDays = req.GetCollector().GetCollectionDays()
		case "eula_uri", "eulaUri":
			obj.EulaUri = req.GetCollector().GetEulaUri()
		case "service_account", "serviceAccount":
			obj.ServiceAccount = req.GetCollector().GetServiceAccount()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported in mock", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		responseObj := proto.Clone(obj).(*pb.Collector)
		responseObj.Labels = nil
		return responseObj, nil
	})
}

func (s *RapidMigrationAssessmentV1) DeleteCollector(ctx context.Context, req *pb.DeleteCollectorRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCollectorName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Collector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	now := time.Now()

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update", // Real GCP has verb "update" for delete operations
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		responseObj := proto.Clone(obj).(*pb.Collector)
		responseObj.Labels = nil
		return responseObj, nil
	})
}
