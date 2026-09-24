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
// proto.message: google.cloud.dataplex.v1.DataTaxonomy

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

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

type DataTaxonomyService struct {
	*MockService
	pb.UnimplementedDataTaxonomyServiceServer
}

func (s *DataTaxonomyService) GetDataTaxonomy(ctx context.Context, req *pb.GetDataTaxonomyRequest) (*pb.DataTaxonomy, error) {
	name, err := s.parseDataTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataTaxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataTaxonomyService) CreateDataTaxonomy(ctx context.Context, req *pb.CreateDataTaxonomyRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/dataTaxonomies/" + req.DataTaxonomyId
	name, err := s.parseDataTaxonomyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.DataTaxonomy).(*pb.DataTaxonomy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Uid = "dataTaxonomy-" + name.DataTaxonomyID

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

func (s *DataTaxonomyService) UpdateDataTaxonomy(ctx context.Context, req *pb.UpdateDataTaxonomyRequest) (*longrunning.Operation, error) {
	name, err := s.parseDataTaxonomyName(req.GetDataTaxonomy().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.DataTaxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	updateMask := req.GetUpdateMask()

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			obj.Description = req.GetDataTaxonomy().GetDescription()
		case "display_name", "displayName":
			obj.DisplayName = req.GetDataTaxonomy().GetDisplayName()
		case "labels":
			obj.Labels = req.GetDataTaxonomy().GetLabels()
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

func (s *DataTaxonomyService) DeleteDataTaxonomy(ctx context.Context, req *pb.DeleteDataTaxonomyRequest) (*longrunning.Operation, error) {
	name, err := s.parseDataTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DataTaxonomy{}
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

type dataTaxonomyName struct {
	Project        *projects.ProjectData
	Location       string
	DataTaxonomyID string
}

func (n *dataTaxonomyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", n.Project.ID, n.Location, n.DataTaxonomyID)
}

func (s *MockService) parseDataTaxonomyName(name string) (*dataTaxonomyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataTaxonomies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &dataTaxonomyName{
			Project:        project,
			Location:       tokens[3],
			DataTaxonomyID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
