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

package mockdataflow

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/dataflow/v1beta3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"
)

type flexTemplatesServer struct {
	*MockService
	pb.UnimplementedFlexTemplatesServiceServer
}

func (r *flexTemplatesServer) LaunchFlexTemplate(ctx context.Context, req *pb.LaunchFlexTemplateRequest) (*pb.LaunchFlexTemplateResponse, error) {
	now := time.Now()
	jobID := now.Format("2006-01-02-15_04_05") + fmt.Sprintf("-%d", now.UnixNano())
	reqName := fmt.Sprintf("projects/%s/locations/%s/flexTemplates/%s", req.GetProjectId(), req.GetLocation(), jobID)
	name, err := r.parseFlexTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	jobName := &jobName{
		Project:  name.Project,
		Location: name.Location,
		Name:     jobID,
	}

	fqn := jobName.String()

	job := &pb.Job{}
	job.Name = name.Name
	job.CreateTime = timestamppb.New(now)
	job.CurrentStateTime = timestamppb.New(time.Unix(0, 0))
	job.Id = jobID
	job.Location = name.Location
	job.ProjectId = name.Project.ID
	job.StartTime = timestamppb.New(now)

	if launchParameter := req.GetLaunchParameter(); launchParameter != nil {
		job.Name = launchParameter.GetJobName()
	}

	if req.GetLaunchParameter().GetUpdate() {
		existingJob, err := findJobByJobName(ctx, r.storage, name.Project.ID, name.Location, job.Name)
		if err != nil {
			return nil, err
		}
		if existingJob == nil {
			return nil, fmt.Errorf("existing job not found")
		}
		job.ReplaceJobId = existingJob.GetId()
		// job.CurrentState = pb.JobState_JOB_STATE_QUEUED
	}

	if err := r.storage.Create(ctx, fqn, job); err != nil {
		return nil, err
	}

	go func() {
		if err := r.StartJob(fqn, name.Project, req); err != nil {
			klog.Fatalf("error starting job: %v", err)
		}
	}()

	retVal := &pb.LaunchFlexTemplateResponse{
		Job: job,
	}
	retVal = proto.Clone(retVal).(*pb.LaunchFlexTemplateResponse)
	retVal.Job.ReplaceJobId = ""
	return retVal, nil
}

type flexTemplateName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *flexTemplateName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/flexTemplates/" + n.Name
}

// parseTemplateName parses a string into an clusterName.
// The expected form is `projects/*/locations/*/flexTemplates/*`.
func (r *flexTemplatesServer) parseFlexTemplateName(name string) (*flexTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "flexTemplates" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &flexTemplateName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
