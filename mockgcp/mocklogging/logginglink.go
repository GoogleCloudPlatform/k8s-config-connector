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
	// "fmt"
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

type linkService struct {
	*MockService
	pb.UnimplementedConfigServiceV2Server
}

/*

// createLinkDefaultObjects will ensure that the default log bucket is created for the folder/project/org
// The input to this is probably linkName not bucketNmae
func (s *linkService) createLinkDefaultObjects(ctx context.Context, name *loggingLinkName) error {
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

*/

func (s *linkService) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.Link, error) {
	name, err := s.parseLoggingLinkName(req.Name)
	fmt.Printf("MOCK LOGGING GET LINK")

	if err != nil {
		return nil, err
	}
	/*
		if err := s.createLinkDefaultObjects(ctx, name); err != nil {
			return nil, err
		}
	*/
	fqn := name.String()
	obj := &pb.Link{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Bucket `%s` does not exist", name.BucketName)
		}
		return nil, err
	}
	return obj, nil
}

func (s *linkService) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.Link, error) {
	fmt.Printf("MOCK LOGGING CREATE LINK")
	reqName := req.Parent + "/links/" + req.GetLinkId()
	name, err := s.parseLoggingLinkName(reqName)
	if err != nil {
		return nil, err
	}
	/*
		if err := s.createLinkDefaultObjects(ctx, name); err != nil {
			return nil, err
		}
	*/
	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetLink()).(*pb.Link)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	// s.populateDefaultsForLogBucket(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *linkService) populateDefaultsForLoggingLink(obj *pb.Link) {
	if obj.LifecycleState == pb.LifecycleState_LIFECYCLE_STATE_UNSPECIFIED {
		obj.LifecycleState = pb.LifecycleState_ACTIVE
	}
}

func (s *linkService) DeleteLink(ctx context.Context, req *pb.DeleteLinkRequest) (*empty.Empty, error) {
	name, err := s.parseLoggingLinkName(req.Name)
	if err != nil {
		return nil, err
	}
	/*
		if err := s.createLinkDefaultObjects(ctx, name); err != nil {
			return nil, err
		}
	*/
	fqn := name.String()
	deletedObj := &pb.Link{}
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
