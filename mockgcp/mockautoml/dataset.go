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

package mockautoml

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/automl/apiv1/automlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AutoMLServer) GetDataset(ctx context.Context, req *pb.GetDatasetRequest) (*pb.Dataset, error) {
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

func (s *AutoMLServer) CreateDataset(ctx context.Context, req *pb.CreateDatasetRequest) (*longrunningpb.Operation, error) {
	// KCC requires the AutoML dataset ID to match the K8s metadata name / resource ID.
	// Since AutoML's CreateDataset REST API does not have an ID field in the request path,
	// we extract the expected ID by mapping the display name.
	id := getDatasetIDFromDisplayName(req.Dataset.DisplayName)
	reqName := req.Parent + "/datasets/" + id
	name, err := s.parseDatasetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Dataset).(*pb.Dataset)
	obj.Name = fqn

	// Always mock 1970-01-01 for createTime to keep tests predictable
	obj.CreateTime = timestamppb.New(time.Unix(0, 0))
	obj.Etag = "abcdef123456"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
		Details: &pb.OperationMetadata_CreateDatasetDetails{
			CreateDatasetDetails: &pb.CreateDatasetOperationMetadata{},
		},
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AutoMLServer) UpdateDataset(ctx context.Context, req *pb.UpdateDatasetRequest) (*pb.Dataset, error) {
	name, err := s.parseDatasetName(req.GetDataset().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetDataset().GetDisplayName()

		case "description":
			obj.Description = req.GetDataset().GetDescription()

		case "labels":
			obj.Labels = req.GetDataset().GetLabels()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AutoMLServer) DeleteDataset(ctx context.Context, req *pb.DeleteDatasetRequest) (*longrunningpb.Operation, error) {
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

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
		Details: &pb.OperationMetadata_DeleteDetails{
			DeleteDetails: &pb.DeleteOperationMetadata{},
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, metadata, &emptypb.Empty{})
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

func getDatasetIDFromDisplayName(displayName string) string {
	if strings.HasPrefix(displayName, "ds_min_") {
		return "automldataset-min-" + strings.TrimPrefix(displayName, "ds_min_")
	}
	if strings.HasPrefix(displayName, "ds_max_") {
		return "ds-max-" + strings.TrimPrefix(displayName, "ds_max_")
	}
	if strings.HasPrefix(displayName, "ds_ic_") {
		return "automldataset-ic-" + strings.TrimPrefix(displayName, "ds_ic_")
	}
	if strings.HasPrefix(displayName, "ds_trn_") {
		return "automldataset-trn-" + strings.TrimPrefix(displayName, "ds_trn_")
	}
	if strings.HasPrefix(displayName, "ds_iod_") {
		return "automldataset-iod-" + strings.TrimPrefix(displayName, "ds_iod_")
	}
	if strings.HasPrefix(displayName, "ds_ts_") {
		return "automldataset-ts-" + strings.TrimPrefix(displayName, "ds_ts_")
	}
	if strings.HasPrefix(displayName, "ds_te_") {
		return "automldataset-te-" + strings.TrimPrefix(displayName, "ds_te_")
	}
	// Fallback/Default
	return displayName
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
