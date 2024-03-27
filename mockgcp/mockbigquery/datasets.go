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
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

type datasetsServer struct {
	*MockService
	pb.UnimplementedDatasetsServer
}

func (s *datasetsServer) GetDataset(ctx context.Context, req *pb.GetDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.ProjectId, req.DatasetId)
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

func (s *datasetsServer) InsertDataset(ctx context.Context, req *pb.InsertDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.ProjectId, req.GetDataset().GetDatasetReference().GetDatasetId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetDataset()).(*pb.Dataset)

	if obj.DatasetReference == nil {
		obj.DatasetReference = &pb.DatasetReference{}
	}
	if obj.GetDatasetReference().GetProjectId() == "" {
		obj.DatasetReference.ProjectId = req.ProjectId
	}
	obj.CreationTime = now.UnixMilli()
	obj.LastModifiedTime = now.UnixMilli()
	obj.Id = obj.GetDatasetReference().GetProjectId() + ":" + obj.GetDatasetReference().GetDatasetId()
	obj.Kind = "bigquery#dataset"
	if obj.GetLocation() == "" {
		obj.Location = "US"
	}
	if obj.GetType() == "" {
		obj.Type = "DEFAULT"
	}
	if len(obj.Access) == 0 {
		obj.Access = []*pb.DatasetAccess{
			{
				Role:         "WRITER",
				SpecialGroup: "projectWriters",
			},
			{
				Role:         "OWNER",
				SpecialGroup: "projectOwners",
			},
			{
				Role:        "OWNER",
				UserByEmail: "me@example.com",
			},
			{
				Role:         "READER",
				SpecialGroup: "projectReaders",
			},
		}
	}

	obj.SelfLink = "https://bigquery.googleapis.com/bigquery/v2/" + name.String()

	sortAccess(obj)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating dataset: %v", err)
	}

	return obj, nil
}

func sortAccess(obj *pb.Dataset) {
	// I haven't found any docs on the actual sort order,
	// and it shouldn't actually matter, but it helps our golden testing.
	key := func(a *pb.DatasetAccess) string {
		return a.Role + ":" + a.IamMember + a.UserByEmail + a.Domain + a.GroupByEmail + a.SpecialGroup
	}

	sort.SliceStable(obj.Access, func(i, j int) bool {
		return key(obj.Access[i]) < key(obj.Access[j])
	})
}

func (s *datasetsServer) UpdateDataset(ctx context.Context, req *pb.UpdateDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.ProjectId, req.DatasetId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := req.GetDataset()
	updated.DatasetReference = existing.DatasetReference

	updated.CreationTime = existing.CreationTime
	updated.LastModifiedTime = now.UnixMilli()
	updated.Id = existing.GetDatasetReference().GetProjectId() + ":" + existing.GetDatasetReference().GetDatasetId()
	updated.Kind = "bigquery#dataset"
	updated.Location = existing.Location
	updated.Type = existing.Type
	updated.SelfLink = "https://bigquery.googleapis.com/bigquery/v2/" + name.String()

	sortAccess(updated)

	updated.Etag = computeEtag(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *datasetsServer) DeleteDataset(ctx context.Context, req *pb.DeleteDatasetRequest) (*empty.Empty, error) {
	name, err := s.buildDatasetName(req.ProjectId, req.DatasetId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Dataset{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type datasetName struct {
	Project   *projects.ProjectData
	DatasetID string
}

func (n *datasetName) String() string {
	return "projects/" + n.Project.ID + "/datasets/" + n.DatasetID
}

// parseDatasetName parses a string into a datasetName.
// The expected form is projects/<projectID>/datasets/<DatasetID>
func (s *MockService) parseDatasetName(name string) (*datasetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "datasets" {
		return s.buildDatasetName(tokens[1], tokens[3])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildDatasetName(projectName string, datasetID string) (*datasetName, error) {
	project, err := s.projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	name := &datasetName{
		Project:   project,
		DatasetID: datasetID,
	}

	return name, nil
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
