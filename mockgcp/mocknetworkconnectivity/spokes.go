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

package mocknetworkconnectivity

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
	"k8s.io/apimachinery/pkg/util/uuid"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type spokesServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsSpokesServerServer
}

func (s *spokesServer) GetProjectsLocationsSpoke(ctx context.Context, req *pb.GetProjectsLocationsSpokeRequest) (*pb.Spoke, error) {
	name, err := s.parseSpokeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Spoke{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *spokesServer) CreateProjectsLocationsSpoke(ctx context.Context, req *pb.CreateProjectsLocationsSpokeRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/spokes/%s", req.GetParent(), req.GetSpokeId())
	name, err := s.parseSpokeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsSpoke()).(*pb.Spoke)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = "ACTIVE"
	obj.UniqueId = string(uuid.NewUUID())

	// Set resolved defaults using project IDs
	obj.SpokeType = "VPC_NETWORK"
	if obj.Hub != "" {
		obj.Group = fmt.Sprintf("%s/groups/default", obj.Hub)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		lroObj := proto.Clone(obj).(*pb.Spoke)

		// Normalize Hub URI to use project number in LRO response
		if lroObj.Hub != "" {
			hubName, err := s.parseHubName(lroObj.Hub)
			if err == nil {
				lroObj.Hub = hubName.StringWithProjectNumber()
			}
		}

		// Normalize Linked VPC Network URI to use project number in LRO response
		if lroObj.LinkedVpcNetwork != nil && lroObj.LinkedVpcNetwork.Uri != "" {
			uri := lroObj.LinkedVpcNetwork.Uri
			parts := strings.Split(uri, "/projects/")
			if len(parts) == 2 {
				subParts := strings.Split(parts[1], "/")
				if len(subParts) > 0 {
					projectID := subParts[0]
					project, err := s.Projects.GetProjectByIDOrNumber(projectID)
					if err == nil {
						subParts[0] = fmt.Sprintf("%d", project.Number)
						lroObj.LinkedVpcNetwork.Uri = parts[0] + "/projects/" + strings.Join(subParts, "/")
					}
				}
			}
		}

		// Update resolved defaults using normalized Hub
		if lroObj.Hub != "" {
			lroObj.Group = fmt.Sprintf("%s/groups/default", lroObj.Hub)
		}

		return lroObj, nil
	})
}

func (s *spokesServer) DeleteProjectsLocationsSpoke(ctx context.Context, req *pb.DeleteProjectsLocationsSpokeRequest) (*longrunning.Operation, error) {
	name, err := s.parseSpokeName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Spoke{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *spokesServer) PatchProjectsLocationsSpoke(ctx context.Context, req *pb.PatchProjectsLocationsSpokeRequest) (*longrunning.Operation, error) {
	name, err := s.parseSpokeName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Spoke{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqResource := req.GetProjectsLocationsSpoke()

	// Apply field mask updates
	paths := strings.Split(req.GetUpdateMask(), ",")
	for _, path := range paths {
		if path == "" {
			continue
		}
		switch path {
		case "description":
			obj.Description = reqResource.Description
		case "labels":
			obj.Labels = reqResource.Labels
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}
