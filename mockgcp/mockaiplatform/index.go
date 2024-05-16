// Copyright 2024 Google LLC
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

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
)

type indexService struct {
	*MockService
	pb.UnimplementedIndexServiceServer
}

func (s *indexService) GetIndex(ctx context.Context, req *pb.GetIndexRequest) (*pb.Index, error) {
	name, err := s.parseIndexName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Index{}
	fmt.Printf("[debug] getting index %s\n", fqn) // TODO(jingyih): remove
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *indexService) CreateIndex(ctx context.Context, req *pb.CreateIndexRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/indexes/" + id
	name, err := s.parseIndexName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Index).(*pb.Index)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	fmt.Printf("[debug] creating index %s\n", fqn) // TODO(jingyih): remove
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateIndexOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.Index)
		result.CreateTime = nil
		result.UpdateTime = nil
		result.Etag = ""

		return result, nil
	})
}

func (s *indexService) UpdateIndex(ctx context.Context, req *pb.UpdateIndexRequest) (*longrunning.Operation, error) {
	name, err := s.parseIndexName(req.GetIndex().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Index{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetIndex().GetDisplayName()

		case "description":
			obj.Description = req.GetIndex().GetDescription()

		case "labels":
			obj.Labels = req.GetIndex().GetLabels()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	fmt.Printf("[debug] updating index %s\n", fqn) // TODO(jingyih): remove
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateIndexOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, obj)
}

func (s *indexService) DeleteIndex(ctx context.Context, req *pb.DeleteIndexRequest) (*longrunning.Operation, error) {
	name, err := s.parseIndexName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Index{}
	fmt.Printf("[debug] deleting index %s\n", fqn) // TODO(jingyih): remove
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type IndexName struct {
	Project  *projects.ProjectData
	Location string
	IndexID  string
}

func (n *IndexName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/indexes/%s", n.Project.Number, n.Location, n.IndexID)
}

// parseIndexName parses a string into a IndexName.
// The expected form of input string is projects/<projectID>/locations/<location>/indexes/<indexID>
func (s *MockService) parseIndexName(name string) (*IndexName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "indexes" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &IndexName{
			Project:  project,
			Location: tokens[3],
			IndexID:  tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
