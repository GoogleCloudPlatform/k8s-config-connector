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
// proto.service: google.cloud.dataplex.v1.DataProductService
// proto.message: google.cloud.dataplex.v1.DataProduct

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

type DataProductService struct {
	*MockService
	pb.UnimplementedDataProductServiceServer
}

func (s *DataProductService) GetDataProduct(ctx context.Context, req *pb.GetDataProductRequest) (*pb.DataProduct, error) {
	name, err := s.parseDataProductName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataProduct{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataProductService) ListDataProducts(ctx context.Context, req *pb.ListDataProductsRequest) (*pb.ListDataProductsResponse, error) {
	response := &pb.ListDataProductsResponse{}

	kind := (&pb.DataProduct{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		dp := obj.(*pb.DataProduct)
		if strings.HasPrefix(dp.GetName(), req.Parent) {
			response.DataProducts = append(response.DataProducts, dp)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DataProductService) CreateDataProduct(ctx context.Context, req *pb.CreateDataProductRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/dataProducts/%s", req.GetParent(), req.GetDataProductId())
	name, err := s.parseDataProductName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetDataProduct()).(*pb.DataProduct)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = uuid.NewString()

	if obj.AccessApprovalConfig == nil {
		obj.AccessApprovalConfig = &pb.DataProduct_AccessApprovalConfig{}
	}
	for k, v := range obj.AccessGroups {
		if v.Id == "" {
			v.Id = k
		}
		if v.DisplayName == "" {
			v.DisplayName = k
		}
		if v.Principal != nil && v.Principal.ServiceAccount == nil {
			v.Principal.ServiceAccount = proto.String("")
		}
	}

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
		lroResponse := proto.Clone(obj).(*pb.DataProduct)
		lroResponse.Labels = nil
		return lroResponse, nil
	})
}

func (s *DataProductService) UpdateDataProduct(ctx context.Context, req *pb.UpdateDataProductRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataProductName(req.GetDataProduct().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.DataProduct{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, req.GetDataProduct())
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetDataProduct().GetDescription()
			case "display_name", "displayName":
				obj.DisplayName = req.GetDataProduct().GetDisplayName()
			case "labels":
				obj.Labels = req.GetDataProduct().GetLabels()
			case "owner_emails", "ownerEmails":
				obj.OwnerEmails = req.GetDataProduct().GetOwnerEmails()
			case "access_approval_config", "accessApprovalConfig":
				obj.AccessApprovalConfig = req.GetDataProduct().GetAccessApprovalConfig()
			case "access_groups", "accessGroups":
				obj.AccessGroups = req.GetDataProduct().GetAccessGroups()
			case "icon":
				obj.Icon = req.GetDataProduct().GetIcon()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	if obj.AccessApprovalConfig == nil {
		obj.AccessApprovalConfig = &pb.DataProduct_AccessApprovalConfig{}
	}
	for k, v := range obj.AccessGroups {
		if v.Id == "" {
			v.Id = k
		}
		if v.DisplayName == "" {
			v.DisplayName = k
		}
		if v.Principal != nil && v.Principal.ServiceAccount == nil {
			v.Principal.ServiceAccount = proto.String("")
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = uuid.NewString()

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
		lroResponse := proto.Clone(obj).(*pb.DataProduct)
		lroResponse.Labels = nil
		return lroResponse, nil
	})
}

func (s *DataProductService) DeleteDataProduct(ctx context.Context, req *pb.DeleteDataProductRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataProductName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.DataProduct{}
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

type dataProductName struct {
	Project       *projects.ProjectData
	Location      string
	DataProductID string
}

func (n *dataProductName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", n.Project.ID, n.Location, n.DataProductID)
}

func (s *MockService) parseDataProductName(name string) (*dataProductName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataProducts" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &dataProductName{
			Project:       project,
			Location:      tokens[3],
			DataProductID: tokens[5],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
