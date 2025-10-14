// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law_or_agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,

// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// +tool:mockgcp-support
// proto.service: google.logging.v2.ConfigServiceV2
// proto.message: google.logging.v2.Link

package mocklogging_config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *configServiceV2) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.Link, error) {
	klog.Infof("tylerreid - GetLink - %+v", &req)
	name, err := s.parseLinkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Link{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *configServiceV2) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*longrunningpb.Operation, error) {
	klog.Infof("tylerreid - CreateLink - %+v", &req)
	reqName := req.Parent + "/links/" + req.GetLinkId()
	name, err := s.parseLinkName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetLink()).(*pb.Link)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.BigqueryDataset = &pb.BigQueryDataset{}
	obj.BigqueryDataset.DatasetId = "bigquery.googleapis.com/projects/${projectId}/datasets/logginglink${uniqueId}"

	s.populateDefaultsForLink(obj)

	metadata := &pb.LinkMetadata{
		StartTime: timestamppb.New(now),
		Request: &pb.LinkMetadata_CreateLinkRequest{
			CreateLinkRequest: &pb.CreateLinkRequest{
				LinkId: req.GetLinkId(),
				Parent: req.GetParent(),
			},
		},
		State: pb.OperationState_OPERATION_STATE_SCHEDULED,
	}

	operationPrefix := fmt.Sprintf("projects/%v/locations/global", name.project.Number)
	return s.MockService.operations.StartLRO(ctx, operationPrefix, metadata, func() (proto.Message, error) {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
		metadata.EndTime = timestamppb.New(now)
		metadata.State = pb.OperationState_OPERATION_STATE_SUCCEEDED
		return obj, nil
	})
}

func (s *configServiceV2) populateDefaultsForLink(obj *pb.Link) {
	if obj.LifecycleState == pb.LifecycleState_LIFECYCLE_STATE_UNSPECIFIED {
		obj.LifecycleState = pb.LifecycleState_ACTIVE
	}
}

func (s *configServiceV2) DeleteLink(ctx context.Context, req *pb.DeleteLinkRequest) (*longrunningpb.Operation, error) {
	klog.Infof("tylerreid - DELETELink - %+v", &req)
	name, err := s.parseLinkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	metadata := &pb.LinkMetadata{
		StartTime: timestamppb.New(now),
		Request: &pb.LinkMetadata_DeleteLinkRequest{
			DeleteLinkRequest: &pb.DeleteLinkRequest{
				Name: req.Name,
			},
		},
		State: pb.OperationState_OPERATION_STATE_SCHEDULED,
	}

	operationPrefix := fmt.Sprintf("projects/%v/locations/global", name.project.Number)
	return s.MockService.operations.StartLRO(ctx, operationPrefix, metadata, func() (proto.Message, error) {
		deletedObj := &pb.Link{}
		if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
			return nil, err
		}
		metadata.EndTime = timestamppb.New(now)
		metadata.State = pb.OperationState_OPERATION_STATE_SUCCEEDED
		return &emptypb.Empty{}, nil
	})
}

type linkName struct {
	// Only one of project/folder/organization/billingAccount should be set
	project        *projects.ProjectData
	folder         string
	organization   string
	billingAccount string

	location   string
	bucketName string
	LinkName   string
}

func (n *linkName) String() string {
	if n.organization != "" {
		return "organizations/" + n.organization + "/locations/" + n.location + "/buckets/" + n.bucketName + "/links/" + n.LinkName
	}
	if n.folder != "" {
		return "folders/" + n.folder + "/locations/" + n.location + "/buckets/" + n.bucketName + "/links/" + n.LinkName
	}
	if n.billingAccount != "" {
		return "billingAccounts/" + n.billingAccount + "/locations/" + n.location + "/buckets/" + n.bucketName + "/links/" + n.LinkName
	}
	return "projects/" + n.project.ID + "/locations/" + n.location + "/buckets/" + n.bucketName + "/links/" + n.LinkName
}

// parseLinkName parses a string into a linkName.
// The expected form is `projects/*/locations/*/buckets/*/links/*`.
func (s *MockService) parseLinkName(name string) (*linkName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &linkName{
			project:    project,
			location:   tokens[3],
			bucketName: tokens[5],
			LinkName:   tokens[7],
		}
		return name, nil
	} else if len(tokens) == 8 && tokens[0] == "folders" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &linkName{
			folder:     tokens[1],
			location:   tokens[3],
			bucketName: tokens[5],
			LinkName:   tokens[7],
		}
		return name, nil
	} else if len(tokens) == 8 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &linkName{
			organization: tokens[1],
			location:     tokens[3],
			bucketName:   tokens[5],
			LinkName:     tokens[7],
		}
		return name, nil
	} else if len(tokens) == 8 && tokens[0] == "billingAccounts" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "links" {
		name := &linkName{
			billingAccount: tokens[1],
			location:       tokens[3],
			bucketName:     tokens[5],
			LinkName:       tokens[7],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
