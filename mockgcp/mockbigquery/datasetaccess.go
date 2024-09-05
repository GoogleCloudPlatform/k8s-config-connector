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

package mockbigquery

import (
	"context"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *datasetsServer) GetAllDatasetAccess(ctx context.Context, req *pb.GetDatasetRequest) (*pb.DatasetAccess[], error) {
	name, err := s.buildDatasetName(req.GetProjectId(), req.GetDatasetId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Dataset %s:%s", name.Project.ID, name.DatasetID)
		}
		return nil, err
	}

	return obj.GetAccess(), nil
}

func (s *datasetsServer) InsertDatasetAccess(ctx context.Context, req *pb.InsertDatasetAccessRequest) (*pb.Dataset, error) {
	datasetObj, err := s.GetDataset(ctx, &pb.GetDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Dataset %s:%s", name.Project.ID, name.DatasetID)
		}
		return nil, err
	}
	datasetObj.Access = append(datasetObj.Access, req.GetDatasetAccess())
	if updatedDatasetObj, err := s.UpdateDataset(ctx, &pb.UpdateDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
		Dataset: datasetObj,
	}); err != nil {
		return nil, err
	}
	return updatedDatasetObj, nil
}

func (s *datasetsServer) UpdateDataset(ctx context.Context, req *pb.UpdateDatasetAccessRequest) (*pb.Dataset, error) {
	datasetObj, err := s.GetDataset(ctx, &pb.GetDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Dataset %s:%s", name.Project.ID, name.DatasetID)
		}
		return nil, err
	}
	//TODO: Find the unique identifier for BigQueryDatasetAccess
	for index, access := range datasetObj.Access {
		if foundDatasetAccess(access, req.GetIdentifier()) {
			datasetObj.Access[index] = req.GetDatasetAccess()
			break
		}
	}
	datasetObj.Access = append(datasetObj.Access, req.GetDatasetAccess())
	if updatedDatasetObj, err := s.UpdateDataset(ctx, &pb.UpdateDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
		Dataset: datasetObj,
	}); err != nil {
		return nil, err
	}
	return updatedDatasetObj, nil
}

func (s *datasetsServer) DeleteDatasetAccess(ctx context.Context, req *pb.DeleteDatasetAccessRequest) (*empty.Empty, error) {
	datasetObj, err := s.GetDataset(ctx, &pb.GetDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Dataset %s:%s", name.Project.ID, name.DatasetID)
		}
		return nil, err
	}
	newDatasetAccessList := []*pb.DatasetAccess
	for _, access := range datasetObj.GetAccess() {
		if foundDatasetAccess(access, req.GetIdentifier()) {
			continue
		}
		newDatasetAccessList = append(newDatasetAccessList, access)
	}
	datasetObj.Access = newDatasetAccessList
	if _, err := s.UpdateDataset(ctx, &pb.UpdateDatasetRequest{
		DatasetId: req.GetDatasetId(),
		ProjectId: req.GetProjectId(),
		Dataset: datasetObj,
	}); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func foundDatasetAccess(actual *pb.DatasetAccess, datasetAccessIdentifier string) (bool, error) {
	return true, nil
}
