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
	"sort"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

var defaultMaxTimeTravelHours = int64(168)

type datasetsServer struct {
	*MockService
	pb.UnimplementedDatasetsServerServer
}

func (s *datasetsServer) GetDataset(ctx context.Context, req *pb.GetDatasetRequest) (*pb.Dataset, error) {
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
	if obj.MaxTimeTravelHours == nil {
		obj.MaxTimeTravelHours = &defaultMaxTimeTravelHours
	}

	return obj, nil
}

func (s *datasetsServer) InsertDataset(ctx context.Context, req *pb.InsertDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.GetProjectId(), req.GetDataset().GetDatasetReference().GetDatasetId())
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
	obj.CreationTime = PtrTo(now.UnixMilli())
	obj.LastModifiedTime = PtrTo(now.UnixMilli())
	obj.Id = PtrTo(obj.GetDatasetReference().GetProjectId() + ":" + obj.GetDatasetReference().GetDatasetId())
	obj.Kind = PtrTo("bigquery#dataset")
	if obj.Location == nil {
		obj.Location = PtrTo("US")
	}
	if obj.Type == nil {
		obj.Type = PtrTo("DEFAULT")
	}
	if len(obj.Access) == 0 {
		obj.Access = []*pb.DatasetAccess{
			{
				Role:         PtrTo("WRITER"),
				SpecialGroup: PtrTo("projectWriters"),
			},
			{
				Role:         PtrTo("OWNER"),
				SpecialGroup: PtrTo("projectOwners"),
			},
			{
				Role:        PtrTo("OWNER"),
				UserByEmail: PtrTo("me@example.com"),
			},
			{
				Role:         PtrTo("READER"),
				SpecialGroup: PtrTo("projectReaders"),
			},
		}
	}

	obj.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	sortAccess(obj)

	obj.Etag = PtrTo(computeEtag(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating dataset: %v", err)
	}

	return obj, nil
}

func sortAccess(obj *pb.Dataset) {
	// I haven't found any docs on the actual sort order,
	// and it shouldn't actually matter, but it helps our golden testing.
	// This order seems to keep our test data happy

	getRoleKey := func(a *pb.DatasetAccess) string {
		roleKey := a.GetRole()
		switch roleKey {
		case "WRITER":
			roleKey = "roles/bigquery.dataEditor"
		case "OWNER":
			roleKey = "roles/bigquery.dataOwner"
		case "READER":
			roleKey = "roles/bigquery.dataViewer"
		}
		return roleKey
	}

	sort.SliceStable(obj.Access, func(i, j int) bool {
		l := obj.Access[i]
		r := obj.Access[j]

		if lv, rv := getRoleKey(l), getRoleKey(r); lv != rv {
			return lv < rv
		}

		if lv, rv := l.GetGroupByEmail(), r.GetGroupByEmail(); lv != rv {
			return lv < rv
		}

		if lv, rv := l.GetUserByEmail(), r.GetUserByEmail(); lv != rv {
			return lv < rv
		}

		if lv, rv := l.GetDomain(), r.GetDomain(); lv != rv {
			return lv < rv
		}

		if lv, rv := l.GetSpecialGroup(), r.GetSpecialGroup(); lv != rv {
			return lv < rv
		}

		if lv, rv := l.GetIamMember(), r.GetIamMember(); lv != rv {
			return lv < rv
		}

		return false
	})
}

func (s *datasetsServer) UpdateDataset(ctx context.Context, req *pb.UpdateDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.GetProjectId(), req.GetDatasetId())
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
	updated.LastModifiedTime = PtrTo(now.UnixMilli())
	updated.Id = PtrTo(existing.GetDatasetReference().GetProjectId() + ":" + existing.GetDatasetReference().GetDatasetId())
	updated.Kind = PtrTo("bigquery#dataset")
	updated.Location = existing.Location
	updated.Type = existing.Type
	updated.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	sortAccess(updated)

	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *datasetsServer) PatchDataset(ctx context.Context, req *pb.PatchDatasetRequest) (*pb.Dataset, error) {
	name, err := s.buildDatasetName(req.GetProjectId(), req.GetDatasetId())
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
	updated.LastModifiedTime = PtrTo(now.UnixMilli())
	updated.Id = PtrTo(existing.GetDatasetReference().GetProjectId() + ":" + existing.GetDatasetReference().GetDatasetId())
	updated.Kind = PtrTo("bigquery#dataset")
	updated.Location = existing.Location
	updated.Type = existing.Type
	updated.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	sortAccess(updated)

	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *datasetsServer) DeleteDataset(ctx context.Context, req *pb.DeleteDatasetRequest) (*empty.Empty, error) {
	name, err := s.buildDatasetName(req.GetProjectId(), req.GetDatasetId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Dataset{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

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
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	name := &datasetName{
		Project:   project,
		DatasetID: datasetID,
	}

	return name, nil
}
