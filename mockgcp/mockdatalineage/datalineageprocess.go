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

package mockdatalineage

import (
	"context"
	"fmt"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/datacatalog/lineage/apiv1/lineagepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

type lineageServer struct {
	*MockService
	pb.UnimplementedLineageServer
}

func (s *lineageServer) GetProcess(ctx context.Context, req *pb.GetProcessRequest) (*pb.Process, error) {
	name, err := s.parseProcessName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Process{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *lineageServer) CreateProcess(ctx context.Context, req *pb.CreateProcessRequest) (*pb.Process, error) {
	parent, err := s.parseLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	processID := ""
	if req.GetProcess().GetName() != "" {
		pName, err := s.parseProcessName(req.GetProcess().GetName())
		if err == nil {
			processID = pName.Process
		} else {
			processID = req.GetProcess().GetName()
		}
	}
	if processID == "" {
		processID = uuid.New().String()
	}

	name := &processName{
		Project:  parent.Project,
		Location: parent.Location,
		Process:  processID,
	}

	fqn := name.String()

	obj := proto.Clone(req.Process).(*pb.Process)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *lineageServer) UpdateProcess(ctx context.Context, req *pb.UpdateProcessRequest) (*pb.Process, error) {
	reqName := req.GetProcess().GetName()
	name, err := s.parseProcessName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Process{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.Process, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update Process fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.Process).(*pb.Process)
	}

	obj.Name = fqn

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *lineageServer) DeleteProcess(ctx context.Context, req *pb.DeleteProcessRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseProcessName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Process{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			if req.AllowMissing {
				lroPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
				lroMetadata := &pb.OperationMetadata{
					CreateTime:    timestamppb.Now(),
					OperationType: pb.OperationMetadata_DELETE,
					Resource:      fmt.Sprintf("projects/%s/locations/%s/processes/%s", name.Project.ID, name.Location, name.Process),
					ResourceUuid:  name.Process,
					State:         pb.OperationMetadata_SUCCEEDED,
					EndTime:       timestamppb.Now(),
				}
				return s.operations.DoneLROWithMetadata(ctx, lroPrefix, lroMetadata, &emptypb.Empty{})
			}
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:    timestamppb.Now(),
		OperationType: pb.OperationMetadata_DELETE,
		Resource:      fmt.Sprintf("projects/%s/locations/%s/processes/%s", name.Project.ID, name.Location, name.Process),
		ResourceUuid:  name.Process,
		State:         pb.OperationMetadata_SUCCEEDED,
		EndTime:       timestamppb.Now(),
	}
	return s.operations.DoneLROWithMetadata(ctx, lroPrefix, lroMetadata, &emptypb.Empty{})
}
