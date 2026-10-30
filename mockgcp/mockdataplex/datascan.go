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
// proto.service: google.cloud.dataplex.v1.DataScanService
// proto.message: google.cloud.dataplex.v1.DataScan

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

type DataScanService struct {
	*MockService
	pb.UnimplementedDataScanServiceServer
}

func (s *DataScanService) GetDataScan(ctx context.Context, req *pb.GetDataScanRequest) (*pb.DataScan, error) {
	name, err := s.parseDataScanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataScan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataScanService) ListDataScans(ctx context.Context, req *pb.ListDataScansRequest) (*pb.ListDataScansResponse, error) {
	response := &pb.ListDataScansResponse{}

	kind := (&pb.DataScan{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		ds := obj.(*pb.DataScan)
		if strings.HasPrefix(ds.GetName(), req.Parent) {
			response.DataScans = append(response.DataScans, ds)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DataScanService) CreateDataScan(ctx context.Context, req *pb.CreateDataScanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/dataScans/%s", req.GetParent(), req.GetDataScanId())
	name, err := s.parseDataScanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetDataScan()).(*pb.DataScan)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.State_ACTIVE

	s.populateDefaultsForDataScan(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataScanService) UpdateDataScan(ctx context.Context, req *pb.UpdateDataScanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataScanName(req.GetDataScan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.DataScan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, req.GetDataScan())
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetDataScan().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetDataScan().GetDisplayName()
			case "labels":
				obj.Labels = req.GetDataScan().GetLabels()
			case "data":
				obj.Data = req.GetDataScan().GetData()
			case "execution_spec", "executionSpec":
				obj.ExecutionSpec = req.GetDataScan().GetExecutionSpec()
			case "data_profile_spec", "dataProfileSpec":
				obj.Spec = req.GetDataScan().GetSpec()
			case "data_quality_spec", "dataQualitySpec":
				obj.Spec = req.GetDataScan().GetSpec()
			case "data_discovery_spec", "dataDiscoverySpec":
				obj.Spec = req.GetDataScan().GetSpec()
			case "data_documentation_spec", "dataDocumentationSpec":
				obj.Spec = req.GetDataScan().GetSpec()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	s.populateDefaultsForDataScan(obj)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataScanService) DeleteDataScan(ctx context.Context, req *pb.DeleteDataScanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataScanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.DataScan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MockService) populateDefaultsForDataScan(obj *pb.DataScan) {
	if obj.GetDataProfileSpec() != nil {
		obj.Type = pb.DataScanType_DATA_PROFILE
	} else if obj.GetDataQualitySpec() != nil {
		obj.Type = pb.DataScanType_DATA_QUALITY
	} else if obj.GetDataDiscoverySpec() != nil {
		obj.Type = pb.DataScanType_DATA_DISCOVERY
	} else if obj.GetDataDocumentationSpec() != nil {
		obj.Type = pb.DataScanType_DATA_DOCUMENTATION
	}

	if obj.ExecutionSpec == nil {
		obj.ExecutionSpec = &pb.DataScan_ExecutionSpec{
			Trigger: &pb.Trigger{
				Mode: &pb.Trigger_OnDemand_{
					OnDemand: &pb.Trigger_OnDemand{},
				},
			},
		}
	} else if obj.ExecutionSpec.Trigger == nil {
		obj.ExecutionSpec.Trigger = &pb.Trigger{
			Mode: &pb.Trigger_OnDemand_{
				OnDemand: &pb.Trigger_OnDemand{},
			},
		}
	}

	if obj.ExecutionStatus == nil {
		obj.ExecutionStatus = &pb.DataScan_ExecutionStatus{}
	}
}

type dataScanName struct {
	Project    *projects.ProjectData
	Location   string
	DataScanID string
}

func (n *dataScanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", n.Project.ID, n.Location, n.DataScanID)
}

func (s *MockService) parseDataScanName(name string) (*dataScanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataScans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &dataScanName{
			Project:    project,
			Location:   tokens[3],
			DataScanID: tokens[5],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
