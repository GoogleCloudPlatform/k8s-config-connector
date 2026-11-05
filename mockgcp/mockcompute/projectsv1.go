// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcompute

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type ProjectsV1 struct {
	*MockService
	pb.UnimplementedProjectsServer
}

func (s *ProjectsV1) Get(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	project, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found: %v", req.GetProject(), err)
	}

	fqn := "projects/" + project.ID
	obj := &pb.Project{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj = s.defaultProject(ctx, project.ID, project.Number)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func (s *ProjectsV1) SetCommonInstanceMetadata(ctx context.Context, req *pb.SetCommonInstanceMetadataProjectRequest) (*pb.Operation, error) {
	project, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found: %v", req.GetProject(), err)
	}

	fqn := "projects/" + project.ID
	obj := &pb.Project{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj = s.defaultProject(ctx, project.ID, project.Number)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	metadataReq := req.GetMetadataResource()
	if metadataReq != nil {
		if metadataReq.Fingerprint != nil && *metadataReq.Fingerprint != "" {
			if obj.CommonInstanceMetadata != nil && obj.CommonInstanceMetadata.Fingerprint != nil {
				if *metadataReq.Fingerprint != *obj.CommonInstanceMetadata.Fingerprint {
					return nil, status.Errorf(codes.FailedPrecondition, "Condition Not Met: fingerprint mismatch")
				}
			}
		}

		newMetadata := proto.CloneOf(metadataReq)
		newMetadata.Kind = PtrTo("compute#metadata")
		newMetadata.Fingerprint = PtrTo(computeFingerprint(newMetadata))
		obj.CommonInstanceMetadata = newMetadata
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	selfLink := fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/%s", project.ID)
	op := &pb.Operation{
		OperationType: PtrTo("compute.projects.setCommonInstanceMetadata"),
		TargetLink:    PtrTo(selfLink),
		TargetId:      obj.Id,
		User:          PtrTo("user@example.com"),
	}

	return s.startGlobalLRO(ctx, project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ProjectsV1) defaultProject(ctx context.Context, projectID string, projectNumber int64) *pb.Project {
	id := uint64(projectNumber)
	if id == 0 {
		id = uint64(s.generateID())
	}
	selfLink := fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/%s", projectID)
	meta := &pb.Metadata{
		Kind: PtrTo("compute#metadata"),
	}
	meta.Fingerprint = PtrTo(computeFingerprint(meta))

	return &pb.Project{
		Id:                     PtrTo(id),
		CreationTimestamp:      PtrTo(s.nowString()),
		Name:                   PtrTo(projectID),
		SelfLink:               PtrTo(selfLink),
		Kind:                   PtrTo("compute#project"),
		CommonInstanceMetadata: meta,
		DefaultNetworkTier:     PtrTo("PREMIUM"),
		DefaultServiceAccount:  PtrTo(strconv.FormatInt(projectNumber, 10) + "-compute@developer.gserviceaccount.com"),
	}
}
