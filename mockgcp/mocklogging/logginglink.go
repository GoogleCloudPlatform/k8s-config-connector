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

// +tool:mockgcp-support
// krm.apiVersion: logging.cnrm.cloud.google.com/v1beta1
// krm.kind: LoggingLoggingLink
// proto.service: google.logging.v2.ConfigServiceV2
// proto.resource: LoggingLink

package mocklogging

import (
	"context"
	"fmt"
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

/*

//I dont think I need to declare this because it is in LogBucket

type configService struct {
	*MockService
	pb.UnimplementedConfigServiceV2Server
}
*/

// createLinkDefaultObjects will ensure that the default log bucket is created for the folder/project/org
// The input to this is probably linkName not bucketNmae
func (s *configService) createLinkDefaultObjects(ctx context.Context, name *loggingLinkName) error {
	// Create the default bucket
	{
		bucket := &pb.LogBucket{
			Description:    "Default bucket",
			LifecycleState: pb.LifecycleState_ACTIVE,
			RetentionDays:  30,
		}
		if name.folder != "" {
			bucket.Name = fmt.Sprintf("folders/%s/locations/global/buckets/_Default", name.folder)
		}

		// This function exists in LogBucket, assuming I can call it here
		if err := s.createBucketIfNotExists(ctx, bucket); err != nil {
			return err
		}
	}

	return nil
}

func (s *configService) GetLoggingLink(ctx context.Context, req *pb.GetLoggingLinkRequest) (*pb.LoggingLink, error) {
	name, err := s.parseLoggingLinkName(req.Name)
	if err != nil {
		return nil, err
	}
	if err := s.createLinkDefaultObjects(ctx, name); err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.LoggingLink{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Bucket `%s` does not exist", name.BucketName)
		}
		return nil, err
	}
	return obj, nil
}

func (s *configService) CreateLoggingLink(ctx context.Context, req *pb.CreateLoggingLinkRequest) (*pb.LoggingLink, error) {
	reqName := req.Parent + "/buckets/" + req.GetBucketId() + "/links/" + req.GetLinkId()
	name, err := s.parseLoggingLinkName(reqName)
	if err != nil {
		return nil, err
	}
	if err := s.createLinkDefaultObjects(ctx, name); err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetLoggingLink()).(*pb.LoggingLink)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForLogBucket(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *configService) populateDefaultsForLoggingLink(obj *pb.LoggingLink) {
	if obj.LifecycleState == pb.LifecycleState_LIFECYCLE_STATE_UNSPECIFIED {
		obj.LifecycleState = pb.LifecycleState_ACTIVE
	}
}

func (s *configService) UpdateLoggingLink(ctx context.Context, req *pb.UpdateLoggingLinkRequest) (*pb.LoggingLink, error) {
	reqName := req.Name
	name, err := s.parseLoggingLinkName(reqName)
	if err != nil {
		return nil, err
	}
	if err := s.createLinkDefaultObjects(ctx, name); err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.LoggingLink{}
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
			updated.Description = req.GetLoggingLink().GetDescription()
		// case "labels":
		// 	updated.Labels = req.GetDnsAuthorization().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	s.populateDefaultsForLoggingLink(updated)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configService) DeleteLoggingLink(ctx context.Context, req *pb.DeleteLoggingLinkRequest) (*empty.Empty, error) {
	name, err := s.parseLoggingLinkName(req.Name)
	if err != nil {
		return nil, err
	}
	if err := s.createLinkDefaultObjects(ctx, name); err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.LoggingLink{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

type loggingLinkName struct {
	// Only one of project/folder/organization/billingAccount should be set
	project        *projects.ProjectData
	folder         string
	organization   string
	billingAccount string

	location   string
	BucketName string
	LinkName   string
}

func (n *loggingLinkName) String() string {
	if n.organization != "" {
		return "organizations/" + n.organization + "/locations/" + n.location + "/buckets/" + n.BucketName + "/links/" + n.LinkName
	}
	if n.folder != "" {
		return "folders/" + n.folder + "/locations/" + n.location + "/buckets/" + n.BucketName + "/links/" + n.LinkName
	}
	if n.billingAccount != "" {
		return "billingAccounts/" + n.billingAccount + "/locations/" + n.location + "/buckets/" + n.BucketName + "/links/" + n.LinkName
	}
	return "projects/" + n.project.ID + "/locations/" + n.location + "/buckets/" + n.BucketName + "/links/" + n.LinkName
}

// parseLoggingLinkName parses a string into a loggingLinkName.
// The expected form is `projects/*/locations/*/buckets/*/links/*`
func (s *MockService) parseLoggingLinkName(name string) (*loggingLinkName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &loggingLinkName{
			project:    project,
			location:   tokens[3],
			BucketName: tokens[5],
			LinkName:   tokens[7],
		}
		return name, nil
	} else if len(tokens) == 8 && tokens[0] == "folders" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &loggingLinkName{
			folder:     tokens[1],
			location:   tokens[3],
			BucketName: tokens[5],
			LinkName:   tokens[7],
		}
		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &loggingLinkName{
			organization: tokens[1],
			location:     tokens[3],
			BucketName:   tokens[5],
			LinkName:     tokens[7],
		}
		return name, nil
	} else if len(tokens) == 6 && tokens[0] == "billingAccounts" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &loggingLinkName{
			billingAccount: tokens[1],
			location:       tokens[3],
			BucketName:     tokens[5],
			LinkName:       tokens[7],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
