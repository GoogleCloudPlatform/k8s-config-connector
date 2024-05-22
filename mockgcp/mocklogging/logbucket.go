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

package mocklogging

import (
	"context"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/logging/v2"
)

type configService struct {
	*MockService
	pb.UnimplementedConfigServiceV2Server
}

func (s *configService) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.LogBucket, error) {
	name, err := s.parseLogBucketName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.LogBucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Bucket `%s` does not exist", name.BucketName)
		}
		return nil, err
	}
	return obj, nil
}

func (s *configService) CreateBucket(ctx context.Context, req *pb.CreateBucketRequest) (*pb.LogBucket, error) {
	reqName := req.Parent + "/buckets/" + req.GetBucketId()
	name, err := s.parseLogBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetBucket()).(*pb.LogBucket)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForLogBucket(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *configService) populateDefaultsForLogBucket(obj *pb.LogBucket) {
	if obj.LifecycleState == pb.LifecycleState_LIFECYCLE_STATE_UNSPECIFIED {
		obj.LifecycleState = pb.LifecycleState_ACTIVE
	}
}

func (s *configService) UpdateBucket(ctx context.Context, req *pb.UpdateBucketRequest) (*pb.LogBucket, error) {
	reqName := req.Name
	name, err := s.parseLogBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.LogBucket{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.LogBucket)
	updated.CreateTime = existing.CreateTime
	updated.UpdateTime = timestamppb.New(now)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetBucket().GetDescription()
		case "retentionDays":
			updated.RetentionDays = req.GetBucket().GetRetentionDays()
		// case "labels":
		// 	updated.Labels = req.GetDnsAuthorization().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	s.populateDefaultsForLogBucket(updated)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configService) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*empty.Empty, error) {
	name, err := s.parseLogBucketName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.LogBucket{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

type logBucketName struct {
	// Only one of project/folder/organization/billingAccount should be set
	project        *projects.ProjectData
	folder         string
	organization   string
	billingAccount string

	location   string
	BucketName string
}

func (n *logBucketName) String() string {
	if n.organization != "" {
		return "organizations/" + n.organization + "/locations/" + n.location + "/buckets/" + n.BucketName
	}
	if n.folder != "" {
		return "folders/" + n.folder + "/locations/" + n.location + "/buckets/" + n.BucketName
	}
	if n.billingAccount != "" {
		return "billingAccounts/" + n.billingAccount + "/locations/" + n.location + "/buckets/" + n.BucketName
	}
	return "projects/" + n.project.ID + "/locations/" + n.location + "/buckets/" + n.BucketName
}

// parseLogBucketName parses a string into a logBucketName.
// The expected form is `projects/*/locations/*/buckets/*`.
func (s *MockService) parseLogBucketName(name string) (*logBucketName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &logBucketName{
			project:    project,
			location:   tokens[3],
			BucketName: tokens[5],
		}
		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "folders" && tokens[2] == "locations" && tokens[4] == "buckets" {
		name := &logBucketName{
			folder:     tokens[1],
			location:   tokens[3],
			BucketName: tokens[5],
		}
		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "buckets" {
		name := &logBucketName{
			organization: tokens[1],
			location:     tokens[3],
			BucketName:   tokens[5],
		}
		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "billingAccounts" && tokens[2] == "locations" && tokens[4] == "buckets" {
		name := &logBucketName{
			billingAccount: tokens[1],
			location:       tokens[3],
			BucketName:     tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
