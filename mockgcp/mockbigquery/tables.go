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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "cloud.google.com/go/bigquery/v2/apiv2/bigquerypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type tablesServer struct {
	*MockService
	pb.UnimplementedTableServiceServer
}

func (s *tablesServer) PatchTable(ctx context.Context, req *pb.UpdateOrPatchTableRequest) (*pb.Table, error) {
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
	updated.LastModifiedTime = uint64(now.UnixMilli())

	updated.FriendlyName = req.GetTable().FriendlyName
	updated.Description = req.GetTable().Description
	updated.Schema = req.GetTable().Schema
	if updated.GetExternalDataConfiguration() != nil {
		updated.RequirePartitionFilter = req.GetTable().RequirePartitionFilter
	}

	updated.Etag = computeEtag(updated)

	s.normalizeTable(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
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

	s.normalizeTable(obj)

	return obj, nil
}

func (s *tablesServer) InsertTable(ctx context.Context, req *pb.InsertTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTable().GetTableReference().GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetTable())

	datasetServer := &datasetsServer{MockService: s.MockService}
	datasetName, err := datasetServer.buildDatasetName(req.GetProjectId(), req.GetDatasetId())
	if err != nil {
		return nil, err
	}
	dataset := &pb.Dataset{}
	if err := s.storage.Get(ctx, datasetName.String(), dataset); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Dataset %s:%s", name.Project.ID, name.DatasetID)
		}
		return nil, err
	}

	if dataset.DefaultEncryptionConfiguration != nil && obj.EncryptionConfiguration == nil {
		obj.EncryptionConfiguration = dataset.DefaultEncryptionConfiguration
	}

	if obj.TableReference == nil {
		obj.TableReference = &pb.TableReference{}
	}
	if obj.GetTableReference().GetProjectId() == "" {
		obj.TableReference.ProjectId = req.ProjectId
	}
	obj.CreationTime = now.UnixMilli()
	obj.LastModifiedTime = uint64(now.UnixMilli())
	obj.Id = obj.GetTableReference().GetProjectId() + ":" + obj.GetTableReference().GetTableId()
	obj.Kind = "bigquery#table"

	if obj.TimePartitioning != nil {
		col := pb.PartitionedColumn{}
		if obj.TimePartitioning.Field != nil {
			col.Field = &obj.TimePartitioning.Field.Value
		}
		obj.PartitionDefinition = &pb.PartitioningDefinition{PartitionedColumn: []*pb.PartitionedColumn{&col}}
	}
	if obj.Location == "" {
		obj.Location = "us-central1"
	}

	if obj.GetExternalDataConfiguration() != nil {
		obj.Type = "EXTERNAL"
	} else if obj.GetView() != nil {
		obj.Type = "VIEW"
	} else {
		obj.Type = "TABLE"
	}

	if obj.NumActiveLogicalBytes == nil {
		obj.NumActiveLogicalBytes = wrapperspb.Int64(0)
	}
	if obj.NumBytes == nil {
		obj.NumBytes = wrapperspb.Int64(0)
	}
	if obj.NumLongTermBytes == nil {
		obj.NumLongTermBytes = wrapperspb.Int64(0)
	}
	if obj.NumLongTermLogicalBytes == nil {
		obj.NumLongTermLogicalBytes = wrapperspb.Int64(0)
	}
	if obj.NumRows == nil {
		obj.NumRows = wrapperspb.UInt64(0)
	}
	if obj.NumTotalLogicalBytes == nil {
		obj.NumTotalLogicalBytes = wrapperspb.Int64(0)
	}

	if obj.GetExternalDataConfiguration() != nil {
		if obj.RequirePartitionFilter == nil {
			obj.RequirePartitionFilter = wrapperspb.Bool(false)
		}

		if obj.Schema == nil {
			if obj.GetExternalDataConfiguration().GetAutodetect().GetValue() {
				obj.Schema = &pb.TableSchema{}
				sourceURI := ""
				if len(obj.GetExternalDataConfiguration().SourceUris) == 1 {
					sourceURI = obj.GetExternalDataConfiguration().SourceUris[0]
				}
				switch sourceURI {
				case "gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv":
					obj.Schema.Fields = []*pb.TableFieldSchema{
						{
							Mode: "NULLABLE",
							Name: "name",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "post_abbr",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "date",
							Type: "DATE",
						},
					}
				case "gs://cloud-samples-data/bigquery/us-states/us-states.avro":
					obj.Schema.Fields = []*pb.TableFieldSchema{
						{
							Mode:        "REQUIRED",
							Name:        "name",
							Type:        "STRING",
							Description: wrapperspb.String("The common name of the state."),
						},
						{
							Mode:        "REQUIRED",
							Name:        "post_abbr",
							Type:        "STRING",
							Description: wrapperspb.String("The postal code abbreviation of the state."),
						},
					}
				case "gs://cloud-samples-data/bigquery/us-states/us-states.parquet":
					obj.Schema.Fields = []*pb.TableFieldSchema{
						{
							Mode: "NULLABLE",
							Name: "name",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "post_abbr",
							Type: "STRING",
						},
					}
				default:
					// Schema for "gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_ANG.txt"
					obj.Schema.Fields = []*pb.TableFieldSchema{
						{
							Mode: "NULLABLE",
							Name: "string_field_0",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_1",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_2",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_3",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_4",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_5",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_6",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_7",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_8",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_9",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_10",
							Type: "STRING",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_11",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "int64_field_12",
							Type: "INTEGER",
						},
						{
							Mode: "NULLABLE",
							Name: "string_field_13",
							Type: "STRING",
						},
					}
				}
			}
		}
	}
	if obj.GetView() != nil {
		if strings.HasPrefix(obj.View.Query, "SELECT distinct dt, user_id FROM") {
			obj.Schema = &pb.TableSchema{}
			obj.Schema.Fields = []*pb.TableFieldSchema{
				{
					Mode: "NULLABLE",
					Name: "dt",
					Type: "DATE",
				},
				{
					Mode: "NULLABLE",
					Name: "user_id",
					Type: "STRING",
				},
			}
		}
	}
	if obj.MaterializedView != nil {
		obj.Type = "MATERIALIZED_VIEW"
		obj.MaterializedView.LastRefreshTime = now.UnixMilli()
		obj.MaterializedViewStatus = &pb.MaterializedViewStatus{
			RefreshWatermark: timestamppb.New(now),
		}
		if obj.Schema == nil {
			obj.Schema = &pb.TableSchema{}
			// TODO: Find a way to convert sql string query to actual object instead of hardcoding.
			// schema of the query in bigquerytable-view test case
			obj.Schema.Fields = []*pb.TableFieldSchema{
				{
					Mode: "NULLABLE",
					Name: "dt",
					Type: "DATE",
				},
				{
					Mode: "NULLABLE",
					Name: "user_id",
					Type: "STRING",
				},
			}
		}
	}

	obj.SelfLink = "https://bigquery.googleapis.com/bigquery/v2/" + name.String()

	obj.Etag = computeEtag(obj)

	s.normalizeTable(obj)

	ret := CloneProto(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating Table: %v", err)
	}

	// Return value has empty schema populated, even though other methods do not
	if ret.Schema == nil {
		ret.Schema = &pb.TableSchema{}
	}
	return ret, nil
}

func (s *tablesServer) UpdateTable(ctx context.Context, req *pb.UpdateOrPatchTableRequest) (*pb.Table, error) {
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
	updated.LastModifiedTime = uint64(now.UnixMilli())
	updated.Description = req.GetTable().Description
	updated.FriendlyName = req.GetTable().FriendlyName
	updated.Schema = req.GetTable().GetSchema()
	if updated.GetExternalDataConfiguration() != nil {
		updated.RequirePartitionFilter = req.GetTable().RequirePartitionFilter
		updated.ExternalDataConfiguration = req.GetTable().ExternalDataConfiguration
	}
	updated.Schema = req.GetTable().Schema
	updated.ExpirationTime = req.GetTable().ExpirationTime

	updated.Etag = computeEtag(updated)
	updated.Labels = req.GetTable().Labels

	updated.TableConstraints = req.GetTable().TableConstraints

	updated.View = req.GetTable().View

	if req.GetTable().View != nil {
		if strings.HasPrefix(req.GetTable().View.Query, "SELECT distinct dt, user_id, guid FROM") {
			updated.Schema = &pb.TableSchema{}
			updated.Schema.Fields = []*pb.TableFieldSchema{
				{
					Mode:        "NULLABLE",
					Name:        "dt",
					Type:        "DATE",
					Description: wrapperspb.String("dt"),
				},
				{
					Mode:        "NULLABLE",
					Name:        "user_id",
					Type:        "STRING",
					Description: wrapperspb.String("user_id"),
				},
				{
					Mode:        "NULLABLE",
					Name:        "guid",
					Type:        "STRING",
					Description: wrapperspb.String("guid"),
				},
			}
		}
	}

	s.normalizeTable(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *tablesServer) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*emptypb.Empty, error) {
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

	return &emptypb.Empty{}, nil
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

func (s *tablesServer) normalizeTable(table *pb.Table) {
	if table.Schema != nil {
		for _, field := range table.Schema.Fields {
			s.normalizeTableField(field)
		}
	}
}

func (s *tablesServer) normalizeTableField(field *pb.TableFieldSchema) {
	if field.PolicyTags != nil && len(field.PolicyTags.Names) == 0 {
		field.PolicyTags = nil
	}
	for _, subField := range field.Fields {
		s.normalizeTableField(subField)
	}
}
