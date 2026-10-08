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

type resourceBundleName struct {
	Project        *projects.ProjectData
	Location       string
	ResourceBundle string
}

func (n *resourceBundleName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/resourceBundles/" + n.ResourceBundle
}

func (s *MockService) parseResourceBundleName(name string) (*resourceBundleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "resourceBundles" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &resourceBundleName{
			Project:        project,
			Location:       tokens[3],
			ResourceBundle: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigDeliveryServer) GetResourceBundle(ctx context.Context, req *pb.GetResourceBundleRequest) (*pb.ResourceBundle, error) {
	name, err := s.parseResourceBundleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ResourceBundle{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigDeliveryServer) ListResourceBundles(ctx context.Context, req *pb.ListResourceBundlesRequest) (*pb.ListResourceBundlesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/", project.ID, tokens[3])

	response := &pb.ListResourceBundlesResponse{}
	kind := (&pb.ResourceBundle{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.ResourceBundle)
		response.ResourceBundles = append(response.ResourceBundles, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ConfigDeliveryServer) CreateResourceBundle(ctx context.Context, req *pb.CreateResourceBundleRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/resourceBundles/" + req.ResourceBundleId
	name, err := s.parseResourceBundleName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.ResourceBundle).(*pb.ResourceBundle)
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
		lroResponse := proto.Clone(obj).(*pb.ResourceBundle)
		lroResponse.Labels = nil
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) UpdateResourceBundle(ctx context.Context, req *pb.UpdateResourceBundleRequest) (*longrunning.Operation, error) {
	reqName := req.GetResourceBundle().GetName()

	name, err := s.parseResourceBundleName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := &pb.ResourceBundle{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Description = req.GetResourceBundle().GetDescription()
		obj.Labels = req.GetResourceBundle().GetLabels()
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetResourceBundle().GetDescription()
			case "labels":
				obj.Labels = req.GetResourceBundle().GetLabels()
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
		lroResponse := proto.Clone(obj).(*pb.ResourceBundle)
		lroResponse.Labels = nil
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) DeleteResourceBundle(ctx context.Context, req *pb.DeleteResourceBundleRequest) (*longrunning.Operation, error) {
	name, err := s.parseResourceBundleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.ResourceBundle{}
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
