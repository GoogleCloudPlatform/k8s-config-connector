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

package mockoracledatabase

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

	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type oracleDatabaseServer struct {
	*MockService
	pb.UnimplementedOracleDatabaseServer
}

type exadbVmClusterName struct {
	Project        *projects.ProjectData
	Location       string
	ExadbVmCluster string
}

func (n *exadbVmClusterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/exadbVmClusters/" + n.ExadbVmCluster
}

func (s *oracleDatabaseServer) parseExadbVmClusterName(name string) (*exadbVmClusterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "exadbVmClusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &exadbVmClusterName{
			Project:        project,
			Location:       tokens[3],
			ExadbVmCluster: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *oracleDatabaseServer) GetExadbVmCluster(ctx context.Context, req *pb.GetExadbVmClusterRequest) (*pb.ExadbVmCluster, error) {
	name, err := s.parseExadbVmClusterName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ExadbVmCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return proto.Clone(obj).(*pb.ExadbVmCluster), nil
}

func (s *oracleDatabaseServer) CreateExadbVmCluster(ctx context.Context, req *pb.CreateExadbVmClusterRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/exadbVmClusters/%s", req.GetParent(), req.GetExadbVmClusterId())
	name, err := s.parseExadbVmClusterName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetExadbVmCluster()).(*pb.ExadbVmCluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.populateDefaultsForExadbVmCluster(name, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.Clone(obj).(*pb.ExadbVmCluster)
		if retObj.Properties != nil {
			retObj.Properties.LifecycleState = pb.ExadbVmClusterProperties_AVAILABLE
		}
		if err := s.storage.Update(ctx, fqn, retObj); err != nil {
			return nil, err
		}
		return retObj, nil
	})
}

func (s *oracleDatabaseServer) populateDefaultsForExadbVmCluster(name *exadbVmClusterName, obj *pb.ExadbVmCluster) error {
	if obj.GcpOracleZone == "" {
		obj.GcpOracleZone = fmt.Sprintf("%s-a-r1", name.Location)
	}
	if obj.EntitlementId == "" {
		obj.EntitlementId = "12345678"
	}
	if obj.IdentityConnector == nil {
		obj.IdentityConnector = &pb.IdentityConnector{
			ServiceAgentEmail: fmt.Sprintf("p%d-sa@gcp-sa-oci.iam.gserviceaccount.com", name.Project.Number),
			ConnectionState:   pb.IdentityConnector_CONNECTED,
		}
	}
	if obj.Properties == nil {
		obj.Properties = &pb.ExadbVmClusterProperties{}
	}
	if obj.Properties.LifecycleState == pb.ExadbVmClusterProperties_EXADB_VM_CLUSTER_LIFECYCLE_STATE_UNSPECIFIED {
		obj.Properties.LifecycleState = pb.ExadbVmClusterProperties_PROVISIONING
	}
	if obj.Properties.Hostname == "" && obj.Properties.HostnamePrefix != "" {
		obj.Properties.Hostname = fmt.Sprintf("%s-abcd", obj.Properties.HostnamePrefix)
	}
	if obj.Properties.OciUri == "" {
		obj.Properties.OciUri = fmt.Sprintf("https://cloud.oracle.com/db/vmclusters/%s", name.ExadbVmCluster)
	}
	if obj.Properties.GiVersion == "" {
		obj.Properties.GiVersion = "19.0.0.0"
	}
	if obj.Properties.MemorySizeGb == 0 {
		totalEcpu := obj.Properties.EnabledEcpuCountPerNode + obj.Properties.AdditionalEcpuCountPerNode
		if totalEcpu > 0 {
			obj.Properties.MemorySizeGb = int32(float64(totalEcpu) * 2.75)
		}
	}
	return nil
}

func (s *oracleDatabaseServer) UpdateExadbVmCluster(ctx context.Context, req *pb.UpdateExadbVmClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetExadbVmCluster().GetName()
	name, err := s.parseExadbVmClusterName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.ExadbVmCluster{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.ExadbVmCluster)
	incoming := req.GetExadbVmCluster()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(incoming).(*pb.ExadbVmCluster)
	} else {
		for _, path := range paths {
			switch path {
			case "properties":
				if incoming.Properties != nil {
					existingProps := updated.Properties
					updated.Properties = proto.Clone(incoming.Properties).(*pb.ExadbVmClusterProperties)
					if existingProps != nil {
						if updated.Properties.Hostname == "" {
							updated.Properties.Hostname = existingProps.Hostname
						}
						if updated.Properties.LifecycleState == pb.ExadbVmClusterProperties_EXADB_VM_CLUSTER_LIFECYCLE_STATE_UNSPECIFIED {
							updated.Properties.LifecycleState = existingProps.LifecycleState
						}
						if updated.Properties.MemorySizeGb == 0 {
							updated.Properties.MemorySizeGb = existingProps.MemorySizeGb
						}
						if updated.Properties.OciUri == "" {
							updated.Properties.OciUri = existingProps.OciUri
						}
						if updated.Properties.GiVersion == "" {
							updated.Properties.GiVersion = existingProps.GiVersion
						}
					}
				}
			case "properties.node_count", "node_count":
				if incoming.Properties != nil && updated.Properties != nil {
					updated.Properties.NodeCount = incoming.Properties.NodeCount
				}
			case "labels":
				updated.Labels = incoming.Labels
			case "display_name":
				updated.DisplayName = incoming.DisplayName
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	now := time.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *oracleDatabaseServer) DeleteExadbVmCluster(ctx context.Context, req *pb.DeleteExadbVmClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseExadbVmClusterName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.ExadbVmCluster{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	})
}

func (s *oracleDatabaseServer) ListExadbVmClusters(ctx context.Context, req *pb.ListExadbVmClustersRequest) (*pb.ListExadbVmClustersResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	response := &pb.ListExadbVmClustersResponse{}
	findPrefix := req.Parent + "/exadbVmClusters/"
	var items []*pb.ExadbVmCluster
	if err := s.storage.List(ctx, (&pb.ExadbVmCluster{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cluster := obj.(*pb.ExadbVmCluster)
		items = append(items, cluster)
		return nil
	}); err != nil {
		return nil, err
	}
	response.ExadbVmClusters = items
	return response, nil
}
