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
// proto.service: google.cloud.dataplex.v1.MetadataService
// proto.message: google.cloud.dataplex.v1.Partition

package mockdataplex

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *MetadataService) CreatePartition(ctx context.Context, req *pb.CreatePartitionRequest) (*pb.Partition, error) {
	parent, err := s.parseEntityName(req.Parent)
	if err != nil {
		return nil, err
	}

	if req.Partition == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Partition is required")
	}

	if len(req.Partition.Values) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Partition values are required")
	}

	parentFQN := parent.String()
	existingEntity := &pb.Entity{}
	if err := s.storage.Get(ctx, parentFQN, existingEntity); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Parent)
		}
		return nil, err
	}

	encodedValues := make([]string, len(req.Partition.Values))
	for i, v := range req.Partition.Values {
		encodedValues[i] = url.QueryEscape(url.QueryEscape(v))
	}
	partitionID := strings.Join(encodedValues, "/")

	name := &partitionName{
		Project:     parent.Project,
		Location:    parent.Location,
		LakeID:      parent.LakeID,
		ZoneID:      parent.ZoneID,
		EntityID:    parent.EntityID,
		PartitionID: partitionID,
	}

	fqn := name.String()

	existing := &pb.Partition{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Partition already exists: %s", name.ProjectNumberString())
	}

	obj := proto.CloneOf(req.Partition)
	obj.Name = name.ProjectNumberString()

	if req.ValidateOnly {
		return obj, nil
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MetadataService) GetPartition(ctx context.Context, req *pb.GetPartitionRequest) (*pb.Partition, error) {
	name, err := s.parsePartitionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Partition{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MetadataService) DeletePartition(ctx context.Context, req *pb.DeletePartitionRequest) (*emptypb.Empty, error) {
	name, err := s.parsePartitionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Partition{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}

	if req.Etag != "" && req.Etag != existing.Etag {
		return nil, status.Errorf(codes.Aborted, "Resource has been modified; please retrieve the latest version and retry")
	}

	deleted := &pb.Partition{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MetadataService) ListPartitions(ctx context.Context, req *pb.ListPartitionsRequest) (*pb.ListPartitionsResponse, error) {
	parent, err := s.parseEntityName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListPartitionsResponse{}
	prefix := parent.String() + "/partitions/"

	partitionKind := (&pb.Partition{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, partitionKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		p := obj.(*pb.Partition)
		if strings.HasPrefix(p.GetName(), parent.ProjectNumberString()+"/partitions/") {
			response.Partitions = append(response.Partitions, proto.CloneOf(p))
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type partitionName struct {
	Project     *projects.ProjectData
	Location    string
	LakeID      string
	ZoneID      string
	EntityID    string
	PartitionID string
}

func (n *partitionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s/partitions/%s", n.Project.ID, n.Location, n.LakeID, n.ZoneID, n.EntityID, n.PartitionID)
}

func (n *partitionName) ProjectNumberString() string {
	return fmt.Sprintf("projects/%d/locations/%s/lakes/%s/zones/%s/entities/%s/partitions/%s", n.Project.Number, n.Location, n.LakeID, n.ZoneID, n.EntityID, n.PartitionID)
}

func (s *MockService) parsePartitionName(name string) (*partitionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) >= 12 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "zones" && tokens[8] == "entities" && tokens[10] == "partitions" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &partitionName{
			Project:     project,
			Location:    tokens[3],
			LakeID:      tokens[5],
			ZoneID:      tokens[7],
			EntityID:    tokens[9],
			PartitionID: strings.Join(tokens[11:], "/"),
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/locations/*/lakes/*/zones/*/entities/*/partitions/*)", name)
}
