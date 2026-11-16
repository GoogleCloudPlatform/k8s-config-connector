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

package mockdataform

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	gw_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1beta1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (r *DataformV1) GetReleaseConfig(ctx context.Context, request *gw_v1.GetReleaseConfigRequest) (*gw_v1.ReleaseConfig, error) {
	name, err := r.parseDataformReleaseConfig(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &gw_v1.ReleaseConfig{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *DataformV1) getInternalMetadata(name *releaseConfigName) string {
	var insertTime, updateTime, uniqueCCFEID string
	if strings.Contains(name.ReleaseConfigID, "-max") {
		insertTime = "2026-11-03T21:30:59.425414Z"
		updateTime = "2026-11-03T21:31:00.495310Z"
		uniqueCCFEID = "ccd440bf-7486-419f-afb5-5c2645cb39c0"
	} else {
		insertTime = "2026-11-03T21:31:20.537507Z"
		updateTime = "2026-11-03T21:31:21.612982Z"
		uniqueCCFEID = "0aaaa3be-6adb-47ee-b4a4-d7392a824f47"
	}
	return fmt.Sprintf(`{"authentication":{"authServiceAccountName":"service-%d@gcp-sa-dataform.iam.gserviceaccount.com"},"db_metadata_insert_time":"%s","db_metadata_update_time":"%s","unique_ccfe_id":"%s"}`,
		name.Project.Number, insertTime, updateTime, uniqueCCFEID)
}

func (r *DataformV1) CreateReleaseConfig(ctx context.Context, request *gw_v1.CreateReleaseConfigRequest) (*gw_v1.ReleaseConfig, error) {
	reqName := request.Parent + "/releaseConfigs/" + request.ReleaseConfigId
	name, err := r.parseDataformReleaseConfig(reqName)
	if err != nil {
		return nil, err
	}

	parentRepo := &pb.Repository{}
	parentFQN := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", name.Project.ID, name.Location, name.RepositoryID)
	if err := r.storage.Get(ctx, parentFQN, parentRepo); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentFQN)
		}
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(request.ReleaseConfig).(*gw_v1.ReleaseConfig)
	obj.Name = fqn

	internalMetadata := r.getInternalMetadata(name)
	obj.InternalMetadata = &internalMetadata

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *DataformV1) UpdateReleaseConfig(ctx context.Context, request *gw_v1.UpdateReleaseConfigRequest) (*gw_v1.ReleaseConfig, error) {
	name, err := r.parseDataformReleaseConfig(request.GetReleaseConfig().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &gw_v1.ReleaseConfig{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updateMask := request.GetUpdateMask()
	if updateMask == nil || len(updateMask.Paths) == 0 {
		obj.GitCommitish = request.GetReleaseConfig().GetGitCommitish()
		obj.CodeCompilationConfig = request.GetReleaseConfig().GetCodeCompilationConfig()
		obj.CronSchedule = request.GetReleaseConfig().GetCronSchedule()
		obj.TimeZone = request.GetReleaseConfig().GetTimeZone()
		obj.Disabled = request.GetReleaseConfig().GetDisabled()
		obj.ReleaseCompilationResult = request.GetReleaseConfig().GetReleaseCompilationResult()
	} else {
		for _, path := range updateMask.Paths {
			switch path {
			case "git_commitish", "gitCommitish":
				obj.GitCommitish = request.GetReleaseConfig().GetGitCommitish()
			case "code_compilation_config", "codeCompilationConfig":
				obj.CodeCompilationConfig = request.GetReleaseConfig().GetCodeCompilationConfig()
			case "cron_schedule", "cronSchedule":
				obj.CronSchedule = request.GetReleaseConfig().GetCronSchedule()
			case "time_zone", "timeZone":
				obj.TimeZone = request.GetReleaseConfig().GetTimeZone()
			case "disabled":
				obj.Disabled = request.GetReleaseConfig().GetDisabled()
			case "release_compilation_result", "releaseCompilationResult":
				obj.ReleaseCompilationResult = request.GetReleaseConfig().GetReleaseCompilationResult()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	internalMetadata := r.getInternalMetadata(name)
	obj.InternalMetadata = &internalMetadata

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *DataformV1) DeleteReleaseConfig(ctx context.Context, request *gw_v1.DeleteReleaseConfigRequest) (*empty.Empty, error) {
	name, err := r.parseDataformReleaseConfig(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &gw_v1.ReleaseConfig{}
	if err := r.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

type releaseConfigName struct {
	Project         *projects.ProjectData
	Location        string
	RepositoryID    string
	ReleaseConfigID string
}

func (n *releaseConfigName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/repositories/" + n.RepositoryID + "/releaseConfigs/" + n.ReleaseConfigID
}

func (s *MockService) parseDataformReleaseConfig(name string) (*releaseConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" && tokens[6] == "releaseConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &releaseConfigName{
			Project:         project,
			Location:        tokens[3],
			RepositoryID:    tokens[5],
			ReleaseConfigID: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
