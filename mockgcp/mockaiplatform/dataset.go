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
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
	"github.com/google/uuid"
)

type datasetService struct {
	*MockService
	pb.UnimplementedDatasetServiceServer
}

func (s *datasetService) GetDataset(ctx context.Context, req *pb.GetDatasetRequest) (*pb.Dataset, error) {
	name, err := s.parseDatasetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *datasetService) CreateDataset(ctx context.Context, req *pb.CreateDatasetRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/datasets/" + id
	name, err := s.parseDatasetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Dataset).(*pb.Dataset)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	bucketID := uuid.NewString()
	obj.Metadata = structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"dataItemSchemaUri": structpb.NewStringValue("gs://google-cloud-aiplatform/schema/dataset/dataitem/image_1.0.0.yaml"),
			"gcsBucket":         structpb.NewStringValue("cloud-ai-platform-" + bucketID),
		},
	})
	artifactID := uuid.NewString()
	obj.MetadataArtifact = fmt.Sprintf("projects/%d/locations/%s/metadataStores/default/artifacts/%s", name.Project.Number, name.Location, artifactID)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateDatasetOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.Dataset)
		result.CreateTime = nil
		result.UpdateTime = nil
		result.Etag = ""

		result.Labels = map[string]string{
			"aiplatform.googleapis.com/dataset_metadata_schema": "IMAGE",
		}
		result.MetadataArtifact = ""
		metadataStruct := result.GetMetadata().GetStructValue()
		if metadataStruct != nil {
			delete(metadataStruct.Fields, "gcsBucket")
		}
		return result, nil
	})
}

func (s *datasetService) UpdateDataset(ctx context.Context, req *pb.UpdateDatasetRequest) (*pb.Dataset, error) {
	name, err := s.parseDatasetName(req.GetDataset().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetDataset().GetDisplayName()

		case "description":
			obj.Description = req.GetDataset().GetDescription()

		case "labels":
			obj.Labels = req.GetDataset().GetLabels()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *datasetService) DeleteDataset(ctx context.Context, req *pb.DeleteDatasetRequest) (*longrunning.Operation, error) {
	name, err := s.parseDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Dataset{}
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

type DatasetName struct {
	Project   *projects.ProjectData
	Location  string
	DatasetID string
}

func (n *DatasetName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/datasets/%s", n.Project.Number, n.Location, n.DatasetID)
}

// parseDatasetName parses a string into a DatasetName.
// The expected form of input string is projects/<projectID>/locations/<location>/datasets/<DatasetID>
func (s *MockService) parseDatasetName(name string) (*DatasetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "datasets" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &DatasetName{
			Project:   project,
			Location:  tokens[3],
			DatasetID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
