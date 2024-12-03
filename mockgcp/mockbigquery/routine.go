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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var defaultLanguage = "SQL"
var DeterminismLevelUnspecified = "DETERMINISM_LEVEL_UNSPECIFIED"

type routinesServer struct {
	*MockService
	pb.UnimplementedRoutinesServerServer
}

func (s *routinesServer) GetRoutine(ctx context.Context, req *pb.GetRoutineRequest) (*pb.Routine, error) {
	name, err := s.buildRoutineName(req.GetProjectId(), req.GetDatasetId(), req.GetRoutineId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Routine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Routine %s:%s.%s", name.Project.ID, name.DatasetID, name.RoutineID)
		}
		return nil, err
	}
	if obj.Language == nil {
		obj.Language = &defaultLanguage
	}
	if obj.DeterminismLevel != nil && *obj.DeterminismLevel == DeterminismLevelUnspecified {
		obj.DeterminismLevel = nil
	}

	return obj, nil
}

func (s *routinesServer) InsertRoutine(ctx context.Context, req *pb.InsertRoutineRequest) (*pb.Routine, error) {
	name, err := s.buildRoutineName(req.GetProjectId(), req.GetDatasetId(), req.GetRoutine().GetRoutineReference().GetRoutineId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetRoutine()).(*pb.Routine)

	if obj.RoutineReference == nil {
		obj.RoutineReference = &pb.RoutineReference{}
	}
	if obj.GetRoutineReference().ProjectId == nil {
		obj.RoutineReference.ProjectId = req.ProjectId
	}
	obj.CreationTime = PtrTo(now.UnixMilli())
	obj.LastModifiedTime = PtrTo(now.UnixMilli())
	obj.Etag = PtrTo(computeEtag(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating routine: %v", err)
	}
	if obj.Language == nil {
		obj.Language = &defaultLanguage
	}
	if obj.DeterminismLevel != nil && *obj.DeterminismLevel == DeterminismLevelUnspecified {
		obj.DeterminismLevel = nil
	}
	return obj, nil
}

func (s *routinesServer) UpdateRoutine(ctx context.Context, req *pb.UpdateRoutineRequest) (*pb.Routine, error) {
	name, err := s.buildRoutineName(req.GetProjectId(), req.GetDatasetId(), req.GetRoutineId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Routine{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := req.GetRoutine()
	updated.RoutineReference = existing.RoutineReference

	updated.CreationTime = existing.CreationTime
	updated.LastModifiedTime = PtrTo(now.UnixMilli())
	updated.RoutineType = existing.RoutineType
	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *routinesServer) DeleteRoutine(ctx context.Context, req *pb.DeleteRoutineRequest) (*empty.Empty, error) {
	name, err := s.buildRoutineName(req.GetProjectId(), req.GetDatasetId(), req.GetRoutineId())
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

type routineName struct {
	Project   *projects.ProjectData
	DatasetID string
	RoutineID string
}

func (n *routineName) String() string {
	return "projects/" + n.Project.ID + "/datasets/" + n.DatasetID + "/routines/" + n.RoutineID
}

func (s *MockService) buildRoutineName(projectName string, datasetID string, routineID string) (*routineName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	name := &routineName{
		Project:   project,
		DatasetID: datasetID,
		RoutineID: routineID,
	}

	return name, nil
}
