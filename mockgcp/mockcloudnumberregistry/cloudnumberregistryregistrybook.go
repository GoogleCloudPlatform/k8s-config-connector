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

package mockcloudnumberregistry

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudnumberregistry/pb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CloudNumberRegistryServer struct {
	*MockService
	pb.UnimplementedCloudNumberRegistryServer
}

func (s *CloudNumberRegistryServer) CreateRegistryBook(ctx context.Context, req *pb.CreateRegistryBookRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/registryBooks/" + req.RegistryBookId
	name, err := s.parseRegistryBookName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.RegistryBook).(*pb.RegistryBook)
	obj.Name = fqn
	obj.ClaimedScopes = s.normalizeClaimedScopes(obj.ClaimedScopes)
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1alpha",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj).(*pb.RegistryBook)
		result.Labels = nil
		return result, nil
	})
}

func (s *CloudNumberRegistryServer) GetRegistryBook(ctx context.Context, req *pb.GetRegistryBookRequest) (*pb.RegistryBook, error) {
	name, err := s.parseRegistryBookName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RegistryBook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *CloudNumberRegistryServer) UpdateRegistryBook(ctx context.Context, req *pb.UpdateRegistryBookRequest) (*longrunning.Operation, error) {
	name, err := s.parseRegistryBookName(req.GetRegistryBook().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.RegistryBook{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pb.RegistryBook)
	updated.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If update_mask is not provided, all fields should be overwritten.
		updated = proto.Clone(req.GetRegistryBook()).(*pb.RegistryBook)
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.New(time.Now())
		updated.Name = obj.Name
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				updated.Labels = req.GetRegistryBook().GetLabels()
			case "claimed_scopes", "claimedScopes":
				updated.ClaimedScopes = req.GetRegistryBook().GetClaimedScopes()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	updated.ClaimedScopes = s.normalizeClaimedScopes(updated.ClaimedScopes)

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1alpha",
	}
	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(updated).(*pb.RegistryBook)
		result.Labels = nil
		return result, nil
	})
}

func (s *CloudNumberRegistryServer) DeleteRegistryBook(ctx context.Context, req *pb.DeleteRegistryBookRequest) (*longrunning.Operation, error) {
	name, err := s.parseRegistryBookName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.RegistryBook{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1alpha",
	}
	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *CloudNumberRegistryServer) normalizeClaimedScopes(claimedScopes []string) []string {
	var normalized []string
	for _, scope := range claimedScopes {
		// Expecting format "projects/{project}"
		if strings.HasPrefix(scope, "projects/") {
			projectIDOrNumber := strings.TrimPrefix(scope, "projects/")
			project, err := s.Projects.GetProjectByIDOrNumber(projectIDOrNumber)
			if err == nil {
				normalized = append(normalized, fmt.Sprintf("projects/%d", project.Number))
				continue
			}
		}
		normalized = append(normalized, scope)
	}
	return normalized
}

type registryBookName struct {
	Project      *projects.ProjectData
	Location     string
	RegistryBook string
}

func (n *registryBookName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/registryBooks/" + n.RegistryBook
}

func (s *CloudNumberRegistryServer) parseRegistryBookName(name string) (*registryBookName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "registryBooks" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &registryBookName{
			Project:      project,
			Location:     tokens[3],
			RegistryBook: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
