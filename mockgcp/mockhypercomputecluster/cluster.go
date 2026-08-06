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
// proto.service: google.cloud.hypercomputecluster.v1.HypercomputeCluster
// proto.message: google.cloud.hypercomputecluster.v1.Cluster

package mockhypercomputecluster

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type hypercomputecluster struct {
	*MockService
	pb.UnimplementedHypercomputeClusterServer
}

func (s *hypercomputecluster) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/clusters/%s", req.GetParent(), req.GetClusterId())
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	// Validate partitions must have node_set_ids
	if req.GetCluster().GetOrchestrator().GetSlurm() != nil {
		slurm := req.GetCluster().GetOrchestrator().GetSlurm()
		for _, p := range slurm.GetPartitions() {
			if len(p.GetNodeSetIds()) == 0 {
				st := status.New(codes.InvalidArgument, fmt.Sprintf("The request was invalid: The partition %q is invalid because its node_set_ids list is empty. Each partition must be associated with at least one node set. Add one or more valid node set IDs.", p.GetId()))
				br := &errdetails.BadRequest{}
				br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       "cluster.orchestrator.slurm.partitions[0].node_set_ids",
					Description: "",
				})
				ri := &errdetails.RequestInfo{
					RequestId: "8fb3ac4b90874bb8",
				}
				st, _ = st.WithDetails(br, ri)
				return nil, st.Err()
			}
		}
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Reconciling = true

	// Populate networkResources sub-field "network" matching "config"
	if obj.NetworkResources != nil {
		for _, netRes := range obj.NetworkResources {
			if netRes.Config != nil && netRes.Config.GetExistingNetwork() != nil {
				netRes.Reference = &pb.NetworkResource_Network{
					Network: &pb.NetworkReference{
						Network:    netRes.Config.GetExistingNetwork().Network,
						Subnetwork: netRes.Config.GetExistingNetwork().Subnetwork,
					},
				}
			}
		}
	}

	// Populate login nodes instances suffix
	if obj.GetOrchestrator().GetSlurm().GetLoginNodes() != nil {
		loginNodes := obj.Orchestrator.GetSlurm().LoginNodes
		loginNodes.Instances = []*pb.ComputeInstance{
			{
				Instance: fmt.Sprintf("projects/%s/zones/%s/instances/%s-login-001", name.Project.ID, loginNodes.GetZone(), name.Cluster),
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *hypercomputecluster) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *hypercomputecluster) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.GetCluster().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// Apply updates
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetCluster().GetLabels()
		case "description":
			obj.Description = req.GetCluster().GetDescription()
		case "orchestrator", "orchestrator.slurm.node_sets":
			obj.Orchestrator = req.GetCluster().GetOrchestrator()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *hypercomputecluster) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
}

func (n *clusterName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", n.Project.ID, n.Location, n.Cluster)
}

func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &clusterName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid cluster name %q", name)
}
