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

package mockaistreams

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/partner/aistreams/v1alpha1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
}

func (n *clusterName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", n.Project.ID, n.Location, n.Cluster)
}

func (n *clusterName) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseClusterName parses projects/*/locations/*/clusters/*.
func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *AIStreamsV1) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Cluster %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AIStreamsV1) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/clusters/%s", req.GetParent(), req.GetClusterId())
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Certificate = "dummy-cert"
	obj.ServiceEndpoint = fmt.Sprintf("%s.aistreams.googleapis.com", name.Cluster)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1alpha1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *AIStreamsV1) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.GetCluster().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	existing := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// For simple mock implementation, overwrite fields or apply masks.
	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(now)
	obj.Certificate = existing.Certificate
	obj.ServiceEndpoint = existing.ServiceEndpoint

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1alpha1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *AIStreamsV1) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := name.Parent()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1alpha1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}
