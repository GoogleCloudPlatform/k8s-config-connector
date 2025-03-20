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

// +tool:mockgcp-support
// proto.service: google.cloud.metastore.v1.DataprocMetastore
// proto.message: google.cloud.metastore.v1.Service

package mockmetastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/metastore/v1"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DataprocMetastoreV1) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Service %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataprocMetastoreV1) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*longrunningpb.Operation, error) {
	// This is a test comment.
	reqName := req.Parent + "/services/" + req.ServiceId
	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Service).(*pb.Service)
	// deletionProtection is not a supported field.
	obj.DeletionProtection = false
	obj.Name = fqn

	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Service_CREATING
	obj.ArtifactGcsUri = "gs://gcs-bucket-" + name.Name + "/hive-warehouse"
	obj.EndpointUri = "thrift://mock-endpoint:9083"
	obj.DatabaseType = pb.Service_MYSQL
	obj.Port = 9083
	obj.StateMessage = "The service is being created"
	obj.ReleaseChannel = pb.Service_STABLE
	obj.Network = "projects/" + name.Project.ID + "/global/networks/default"
	obj.TelemetryConfig = &pb.TelemetryConfig{
		LogFormat: pb.TelemetryConfig_JSON,
	}

	// Remove unnecessary fields

	// Add HiveMetastoreConfig with endpointProtocol
	obj.MetastoreConfig = &pb.Service_HiveMetastoreConfig{
		HiveMetastoreConfig: &pb.HiveMetastoreConfig{
			EndpointProtocol: pb.HiveMetastoreConfig_THRIFT,
			Version:          "3.1.2",
			ConfigOverrides: map[string]string{
				"hive.metastore.warehouse.dir": "gs://gcs-bucket-" + name.Name + "/hive-warehouse",
			},
		},
	}

	// Generate a UID if not present
	if obj.Uid == "" {
		obj.Uid = uuid.New().String()
	}

	// Add tier if not set
	if obj.Tier == pb.Service_TIER_UNSPECIFIED {
		obj.Tier = pb.Service_DEVELOPER
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}
	lro, err := s.operations.NewLRO(ctx)
	lro.Done = false
	if err != nil {
		return nil, err
	}
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	// Use the fully qualified type name to ensure compatibility with the expected output.
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.metastore.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		updated, err := s.updateService(ctx, fqn, func(obj *pb.Service) {
			obj.State = pb.Service_ACTIVE
			obj.StateMessage = "The service is ready to use"
		})
		if err != nil {
			return nil, err
		}
		return updated, err
	})
}

func (s *DataprocMetastoreV1) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceName(req.GetService().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	update := func(obj *pb.Service) {
		if req.UpdateMask == nil || len(req.UpdateMask.Paths) == 0 {
			return
		}
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "tier":
				obj.Tier = req.Service.Tier
			case "hive_metastore_config":
				obj.MetastoreConfig = req.Service.MetastoreConfig
			case "maintenance_window":
				obj.MaintenanceWindow = req.Service.MaintenanceWindow
			case "labels":
				obj.Labels = req.Service.Labels
			case "network_config":
				obj.NetworkConfig = req.Service.NetworkConfig
			case "scaling_config":
				obj.ScalingConfig = req.Service.ScalingConfig
			case "deletionProtection":
				// Ignore deletionProtection in update mask
				continue

			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	updated, err := s.updateService(ctx, fqn, update)
	if err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	// // Returns with no createTime
	// lroRet := proto.Clone(obj).(*pb.Workflow)
	// lroRet.CreateTime = nil
	// lroRet.UpdateTime = nil
	// lroRet.RevisionCreateTime = nil

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *DataprocMetastoreV1) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Service{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

// updateService will read-modify-write the object with optimistic locking
func (s *DataprocMetastoreV1) updateService(ctx context.Context, fqn string, update func(obj *pb.Service)) (*pb.Service, error) {
	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	update(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type serviceName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *serviceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", n.Project.ID, n.Location, n.Name)
}

// parseServiceName parses a string into an serviceName.
// The expected form is `projects/*/locations/*/services/*`.
func (s *DataprocMetastoreV1) parseServiceName(name string) (*serviceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "services" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &serviceName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// buildServiceName builds a serviceName from the components.
func (s *MockService) buildServiceName(projectName, region, service string) (*serviceName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &serviceName{
		Project:  project,
		Location: region,
		Name:     service,
	}, nil
}
