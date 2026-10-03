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

// +tool:mockgcp-support
// proto.service: google.cloud.dataplex.v1.DataTaxonomyService
// proto.message: google.cloud.dataplex.v1.DataAttributeBinding

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunning "google.golang.org/genproto/googleapis/longrunning"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *DataTaxonomyService) GetDataAttributeBinding(ctx context.Context, req *pb.GetDataAttributeBindingRequest) (*pb.DataAttributeBinding, error) {
	name, err := s.parseDataAttributeBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataAttributeBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataTaxonomyService) CreateDataAttributeBinding(ctx context.Context, req *pb.CreateDataAttributeBindingRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/dataAttributeBindings/" + req.DataAttributeBindingId
	name, err := s.parseDataAttributeBindingName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.DataAttributeBinding).(*pb.DataAttributeBinding)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "dataAttributeBinding-" + name.DataAttributeBindingID

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.RequestedCancellation = false
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataTaxonomyService) UpdateDataAttributeBinding(ctx context.Context, req *pb.UpdateDataAttributeBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseDataAttributeBindingName(req.GetDataAttributeBinding().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataAttributeBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	updateMask := req.GetUpdateMask()

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			obj.Description = req.GetDataAttributeBinding().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetDataAttributeBinding().GetDisplayName()
		case "labels":
			obj.Labels = req.GetDataAttributeBinding().GetLabels()
		case "attributes":
			obj.Attributes = req.GetDataAttributeBinding().GetAttributes()
		case "paths":
			obj.Paths = req.GetDataAttributeBinding().GetPaths()
		default:
			return nil, fmt.Errorf("mock does not implement update of %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataTaxonomyService) DeleteDataAttributeBinding(ctx context.Context, req *pb.DeleteDataAttributeBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseDataAttributeBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DataAttributeBinding{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type dataAttributeBindingName struct {
	Project                *projects.ProjectData
	Location               string
	DataAttributeBindingID string
}

func (n *dataAttributeBindingName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataAttributeBindings/%s", n.Project.ID, n.Location, n.DataAttributeBindingID)
}

func (s *MockService) parseDataAttributeBindingName(name string) (*dataAttributeBindingName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataAttributeBindings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &dataAttributeBindingName{
			Project:                project,
			Location:               tokens[3],
			DataAttributeBindingID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
