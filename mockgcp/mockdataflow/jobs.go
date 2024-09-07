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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"
)

type jobsServer struct {
	*MockService
	pb.UnimplementedJobsV1Beta3Server
}

func findJobByJobName(ctx context.Context, jobStore storage.Storage, projectID string, location string, jobName string) (*pb.Job, error) {
	prefix := fmt.Sprintf("projects/%s/locations/%s/jobs/", projectID, location)
	var matches []*pb.Job

	findKind := (&pb.Job{}).ProtoReflect().Descriptor()
	if err := jobStore.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		job := obj.(*pb.Job)
		if job.Name == jobName {
			matches = append(matches, job)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, nil
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	// Probably need to filter by state
	return nil, fmt.Errorf("multiple matches for jobName %q", jobName)
}

func (r *jobsServer) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.Job, error) {
	reqName := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", req.GetProjectId(), req.GetLocation(), req.GetJobId())
	name, err := r.parseJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Job '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *jobsServer) UpdateJob(ctx context.Context, req *pb.UpdateJobRequest) (*pb.Job, error) {
	reqName := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", req.GetProjectId(), req.GetLocation(), req.GetJobId())
	name, err := r.parseJobName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Job{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Only the job state is updateable
	switch req.GetJob().GetRequestedState() {
	case pb.JobState_JOB_STATE_CANCELLED:
		obj.CurrentState = pb.JobState_JOB_STATE_CANCELLING
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_CANCELLED
	case pb.JobState_JOB_STATE_DRAINING:
		obj.CurrentState = pb.JobState_JOB_STATE_DRAINING
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_DRAINING
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unhandled requestedState %v in mock", req.GetJob())
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(10 * time.Second)
		if err := r.StopJob(fqn); err != nil {
			klog.Fatalf("error stopping job: %v", err)
		}
	}()

	// This method returns only a few fields
	ret := &pb.Job{
		// Doesn't seem to return the actual job type, seems to always return JOB_TYPE_BATCH
		// Type: obj.Type,
		Type: pb.JobType_JOB_TYPE_BATCH,
	}
	return ret, nil
}

type jobName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *jobName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/jobs/" + n.Name
}

// parseJobName parses a string into an jobName.
// The expected form is `projects/*/locations/*/jobs/*`.
func (r *jobsServer) parseJobName(name string) (*jobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "jobs" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &jobName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
