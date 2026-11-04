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

package mockconfigdelivery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type releaseName struct {
	Project        *projects.ProjectData
	Location       string
	ResourceBundle string
	Release        string
}

func (n *releaseName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/resourceBundles/" + n.ResourceBundle + "/releases/" + n.Release
}

func (s *MockService) parseReleaseName(name string) (*releaseName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "resourceBundles" && tokens[6] == "releases" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &releaseName{
			Project:        project,
			Location:       tokens[3],
			ResourceBundle: tokens[5],
			Release:        tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigDeliveryServer) GetRelease(ctx context.Context, req *pb.GetReleaseRequest) (*pb.Release, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigDeliveryServer) ListReleases(ctx context.Context, req *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "resourceBundles" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/", project.ID, tokens[3], tokens[5])

	response := &pb.ListReleasesResponse{}
	kind := (&pb.Release{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.Release)
		response.Releases = append(response.Releases, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ConfigDeliveryServer) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/releases/" + req.ReleaseId
	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.Release).(*pb.Release)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.Release)
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) UpdateRelease(ctx context.Context, req *pb.UpdateReleaseRequest) (*longrunning.Operation, error) {
	reqName := req.GetRelease().GetName()

	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Version = req.GetRelease().GetVersion()
		obj.Info = req.GetRelease().GetInfo()
		obj.Lifecycle = req.GetRelease().GetLifecycle()
	} else {
		for _, path := range paths {
			switch path {
			case "version":
				obj.Version = req.GetRelease().GetVersion()
			case "info":
				obj.Info = req.GetRelease().GetInfo()
			case "lifecycle":
				obj.Lifecycle = req.GetRelease().GetLifecycle()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.Release)
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) DeleteRelease(ctx context.Context, req *pb.DeleteReleaseRequest) (*longrunning.Operation, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.Release{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
