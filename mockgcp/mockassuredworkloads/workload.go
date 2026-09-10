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

package mockassuredworkloads

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	v3 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
)

type AssuredWorkloadsServer struct {
	*MockService
	pb.UnimplementedAssuredWorkloadsServiceServer
}

func (s *AssuredWorkloadsServer) generateResourceID() int64 {
	// 12-digit folder or project ID
	return 100000000000 + rand.Int63n(900000000000)
}

func (s *AssuredWorkloadsServer) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AssuredWorkloadsServer) CreateWorkload(ctx context.Context, req *pb.CreateWorkloadRequest) (*longrunning.Operation, error) {
	if req.Workload == nil {
		return nil, status.Errorf(codes.InvalidArgument, "workload is required")
	}

	// Generate a random workload ID
	workloadID := fmt.Sprintf("workload-%06d", rand.Intn(1000000))
	reqName := req.Parent + "/workloads/" + workloadID
	name, err := s.parseWorkloadName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Workload).(*pb.Workload)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	// Populate simulated resource info
	folderParent := req.Parent
	if idx := strings.Index(folderParent, "/locations/"); idx != -1 {
		folderParent = folderParent[:idx]
	}

	if len(obj.ResourceSettings) == 0 {
		resourceID := s.generateResourceID()
		obj.Resources = []*pb.Workload_ResourceInfo{
			{
				ResourceId:   resourceID,
				ResourceType: pb.Workload_ResourceInfo_CONSUMER_FOLDER,
			},
		}
		// Create mock folder for CONSUMER_FOLDER
		folderFqn := fmt.Sprintf("folders/%d", resourceID)
		folderObj := &v3.Folder{
			Name:        folderFqn,
			Parent:      folderParent,
			DisplayName: obj.DisplayName,
			State:       v3.Folder_ACTIVE,
			CreateTime:  timestamppb.New(now),
			UpdateTime:  timestamppb.New(now),
			Etag:        "abcdef0123A=",
		}
		if err := s.storage.Create(ctx, folderFqn, folderObj); err != nil {
			return nil, err
		}
	} else {
		for _, setting := range obj.ResourceSettings {
			resourceID := s.generateResourceID()
			obj.Resources = append(obj.Resources, &pb.Workload_ResourceInfo{
				ResourceId:   resourceID,
				ResourceType: setting.ResourceType,
			})
			if setting.ResourceType == pb.Workload_ResourceInfo_CONSUMER_FOLDER {
				folderFqn := fmt.Sprintf("folders/%d", resourceID)
				folderObj := &v3.Folder{
					Name:        folderFqn,
					Parent:      folderParent,
					DisplayName: setting.DisplayName,
					State:       v3.Folder_ACTIVE,
					CreateTime:  timestamppb.New(now),
					UpdateTime:  timestamppb.New(now),
					Etag:        "abcdef0123A=",
				}
				if err := s.storage.Create(ctx, folderFqn, folderObj); err != nil {
					return nil, err
				}
			} else if setting.ResourceType == pb.Workload_ResourceInfo_CONSUMER_PROJECT || setting.ResourceType == pb.Workload_ResourceInfo_ENCRYPTION_KEYS_PROJECT {
				projectFqn := fmt.Sprintf("projects/mock-project-%d", resourceID)
				projectObj := &v3.Project{
					Name:        fmt.Sprintf("projects/%d", resourceID),
					ProjectId:   fmt.Sprintf("mock-project-%d", resourceID),
					DisplayName: setting.DisplayName,
					State:       v3.Project_ACTIVE,
					CreateTime:  timestamppb.New(now),
					Etag:        "abcdef0123A=",
				}
				if err := s.storage.Create(ctx, projectFqn, projectObj); err != nil {
					return nil, err
				}
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.CreateWorkloadOperationMetadata{
		CreateTime:       timestamppb.New(now),
		DisplayName:      obj.DisplayName,
		Parent:           req.Parent,
		ComplianceRegime: obj.ComplianceRegime,
	}

	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AssuredWorkloadsServer) UpdateWorkload(ctx context.Context, req *pb.UpdateWorkloadRequest) (*pb.Workload, error) {
	if req.Workload == nil {
		return nil, status.Errorf(codes.InvalidArgument, "workload is required")
	}

	name, err := s.parseWorkloadName(req.Workload.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Paths in field_mask are typically like "workload.display_name", "workload.labels"
	paths := req.UpdateMask.GetPaths()
	if len(paths) == 0 {
		// If update_mask is empty, update display name and labels by default
		paths = []string{"workload.display_name", "workload.labels"}
	}

	for _, path := range paths {
		switch path {
		case "workload.display_name", "display_name", "workload.displayName", "displayName":
			existing.DisplayName = req.Workload.DisplayName
		case "workload.labels", "labels":
			existing.Labels = req.Workload.Labels
		}
	}

	existing.Etag = fields.ComputeWeakEtag(existing)

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *AssuredWorkloadsServer) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
