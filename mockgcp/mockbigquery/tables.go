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
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

type tablesServer struct {
	*MockService
	pb.UnimplementedTablesServerServer
}

func (s *tablesServer) PatchTable(ctx context.Context, req *pb.PatchTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := CloneProto(existing)
	updated.LastModifiedTime = PtrTo(uint64(now.UnixMilli()))

	updated.FriendlyName = req.GetTable().FriendlyName
	if updated.GetExternalDataConfiguration() != nil {
		updated.RequirePartitionFilter = PtrTo(req.GetTable().GetRequirePartitionFilter())
	}

	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *tablesServer) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Table %s:%s.%s", name.Project.ID, name.DatasetID, name.TableID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tablesServer) InsertTable(ctx context.Context, req *pb.InsertTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTable().GetTableReference().GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetTable()).(*pb.Table)

	if obj.TableReference == nil {
		obj.TableReference = &pb.TableReference{}
	}
	if obj.GetTableReference().GetProjectId() == "" {
		obj.TableReference.ProjectId = req.ProjectId
	}
	obj.CreationTime = PtrTo(now.UnixMilli())
	obj.LastModifiedTime = PtrTo(uint64(now.UnixMilli()))
	obj.Id = PtrTo(obj.GetTableReference().GetProjectId() + ":" + obj.GetTableReference().GetTableId())
	obj.Kind = PtrTo("bigquery#table")

	if obj.TimePartitioning != nil {
		col := pb.PartitionedColumn{Field: obj.TimePartitioning.Field}
		obj.PartitionDefinition = &pb.PartitioningDefinition{PartitionedColumn: []*pb.PartitionedColumn{&col}}
	}
	if obj.Location == nil {
		obj.Location = PtrTo("us-central1")
	}

	if obj.GetExternalDataConfiguration() != nil {
		obj.Type = PtrTo("EXTERNAL")
	} else {
		obj.Type = PtrTo("TABLE")
	}

	if obj.NumActiveLogicalBytes == nil {
		obj.NumActiveLogicalBytes = PtrTo(int64(0))
	}
	if obj.NumBytes == nil {
		obj.NumBytes = PtrTo(int64(0))
	}
	if obj.NumLongTermBytes == nil {
		obj.NumLongTermBytes = PtrTo(int64(0))
	}
	if obj.NumLongTermLogicalBytes == nil {
		obj.NumLongTermLogicalBytes = PtrTo(int64(0))
	}
	if obj.NumRows == nil {
		obj.NumRows = PtrTo(uint64(0))
	}
	if obj.NumTotalLogicalBytes == nil {
		obj.NumTotalLogicalBytes = PtrTo(int64(0))
	}

	if obj.GetExternalDataConfiguration() != nil {
		if obj.RequirePartitionFilter == nil {
			obj.RequirePartitionFilter = PtrTo(false)
		}

		if obj.Schema == nil {
			if obj.GetExternalDataConfiguration().GetAutodetect() {
				obj.Schema = &pb.TableSchema{}

				// Schema for "gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_ANG.txt"
				obj.Schema.Fields = []*pb.TableFieldSchema{
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_0"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_1"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_2"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_3"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_4"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_5"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_6"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_7"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_8"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_9"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_10"),
						Type: PtrTo("STRING"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_11"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("int64_field_12"),
						Type: PtrTo("INTEGER"),
					},
					{
						Mode: PtrTo("NULLABLE"),
						Name: PtrTo("string_field_13"),
						Type: PtrTo("STRING"),
					},
				}
			}
		}
	}

	if obj.MaterializedView != nil {
		obj.Type = PtrTo("MATERIALIZED_VIEW")
		obj.MaterializedView.LastRefreshTime = PtrTo(now.UnixMilli())
		obj.MaterializedViewStatus = &pb.MaterializedViewStatus{
			RefreshWatermark: timestamppb.New(now),
		}
		if obj.Schema == nil {
			obj.Schema = &pb.TableSchema{}
			// TODO: Find a way to convert sql string query to actual object instead of hardcoding.
			// schema of the query in bigquerytable-view test case
			obj.Schema.Fields = []*pb.TableFieldSchema{
				{
					Mode: PtrTo("NULLABLE"),
					Name: PtrTo("dt"),
					Type: PtrTo("DATE"),
				},
				{
					Mode: PtrTo("NULLABLE"),
					Name: PtrTo("user_id"),
					Type: PtrTo("STRING"),
				},
			}
		}
	}

	obj.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	obj.Etag = PtrTo(computeEtag(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating Table: %v", err)
	}

	ret := CloneProto(obj)
	// Return value has empty schema populated, even though other methods do not
	if ret.Schema == nil {
		ret.Schema = &pb.TableSchema{}
	}
	return ret, nil
}

func (s *tablesServer) UpdateTable(ctx context.Context, req *pb.UpdateTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := CloneProto(existing)
	updated.LastModifiedTime = PtrTo(uint64(now.UnixMilli()))
	updated.Description = req.GetTable().Description
	updated.FriendlyName = req.GetTable().FriendlyName
	if updated.GetExternalDataConfiguration() != nil {
		updated.RequirePartitionFilter = PtrTo(req.GetTable().GetRequirePartitionFilter())
	}

	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *tablesServer) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*empty.Empty, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Table{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &empty.Empty{}, nil
}

type tableName struct {
	Project   *projects.ProjectData
	DatasetID string
	TableID   string
}

func (n *tableName) String() string {
	return "projects/" + n.Project.ID + "/datasets/" + n.DatasetID + "/tables/" + n.TableID
}

// parseTableName parses a string into a tableName.
// The expected form is projects/<projectID>/datasets/<DatasetID>/tables/<TableID> --> TODO
func (s *MockService) parseTableName(name string) (*tableName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "datasets" && tokens[4] == "tables" {
		return s.buildTableName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildTableName(projectName string, datasetID string, tableID string) (*tableName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	name := &tableName{
		Project:   project,
		DatasetID: datasetID,
		TableID:   tableID,
	}

	return name, nil
}
