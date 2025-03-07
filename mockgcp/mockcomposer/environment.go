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
// proto.service: google.cloud.orchestration.airflow.service.v1.Environments
// proto.message: google.cloud.orchestration.airflow.service.v1.Environment

package mockcomposer

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
)

func (s *ComposerV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.Config.DatabaseConfig.MachineType == "db-custom-2-7680" {
		obj.Config.DatabaseConfig.MachineType = ""
	}
	if obj.Config.NodeConfig.IpAllocationPolicy.UseIpAliases {
		obj.Config.NodeConfig.IpAllocationPolicy.UseIpAliases = false
	}
	if obj.Config.SoftwareConfig.PythonVersion == "3" {
		obj.Config.SoftwareConfig.PythonVersion = ""
	}
	return obj, nil
}

func (s *ComposerV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {

	name, err := s.parseEnvironmentName(req.Environment.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Environment).(*pb.Environment)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Environment_RUNNING
	obj.Uuid = "test-uuid" // TODO: real value
	s.populateDefaultsForEnvironment(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		OperationType: pb.OperationMetadata_CREATE,
		CreateTime:    timestamppb.New(now),
		Resource:      name.String(),
		State:         pb.OperationMetadata_PENDING,
		ResourceUuid:  "test-uuid",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		return obj, nil
	})
}

func (s *ComposerV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.Environment)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		tokens := strings.Split(path, ".")
		switch tokens[0] {
		case "labels":
			updated.Labels = req.GetEnvironment().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	updated.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	// // Returns with no createTime
	// lroRet := proto.Clone(obj).(*pb.Workflow)
	// lroRet.CreateTime = nil
	// lroRet.UpdateTime = nil
	// lroRet.RevisionCreateTime = nil
	lroMetadata := &pb.OperationMetadata{
		OperationType: pb.OperationMetadata_UPDATE,
		CreateTime:    timestamppb.New(now),
		Resource:      name.String(),
		State:         pb.OperationMetadata_PENDING,
		ResourceUuid:  "test-uuid", // TODO: Real values
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		return updated, nil
	})
}

func (s *ComposerV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Environment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		OperationType: pb.OperationMetadata_DELETE,
		CreateTime:    timestamppb.Now(),
		Resource:      name.String(),
		State:         pb.OperationMetadata_PENDING,
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		return &emptypb.Empty{}, nil
	})
}

func (s *ComposerV1) populateDefaultsForEnvironment(obj *pb.Environment) {
	if obj.StorageConfig == nil {
		obj.StorageConfig = &pb.StorageConfig{}
	}
	if obj.StorageConfig.Bucket == "" {
		obj.StorageConfig.Bucket = "us-central1-test-123456-asdfg-bucket"
	}
	if obj.Config == nil {
		obj.Config = &pb.EnvironmentConfig{}
	}

	s.populateDefaultsForEnvironmentConfig(obj.Config)
}

func (s *ComposerV1) populateDefaultsForEnvironmentConfig(config *pb.EnvironmentConfig) {
	config.AirflowByoidUri = "https://123456qwert-dot-us-central1.composer.byoid.googleusercontent.com"
	config.AirflowUri = "https://123456qwert-dot-us-central1.composer.googleusercontent.com"
	config.DagGcsPrefix = "gs://us-central1-test-123456-asdfg-bucket/dags"
	config.GkeCluster = "projects/${projectId}/locations/us-central1/clusters/us-central1-test-123456-asdfg-gke"
	if config.DataRetentionConfig == nil {
		config.DataRetentionConfig = &pb.DataRetentionConfig{}
	}
	if config.DataRetentionConfig.AirflowMetadataRetentionConfig == nil {
		config.DataRetentionConfig.AirflowMetadataRetentionConfig = &pb.AirflowMetadataRetentionPolicyConfig{}
	}
	if config.DataRetentionConfig.AirflowMetadataRetentionConfig.RetentionMode == pb.AirflowMetadataRetentionPolicyConfig_RETENTION_MODE_UNSPECIFIED {
		config.DataRetentionConfig.AirflowMetadataRetentionConfig.RetentionMode = pb.AirflowMetadataRetentionPolicyConfig_RETENTION_MODE_DISABLED
	}
	if config.DataRetentionConfig.TaskLogsRetentionConfig == nil {
		config.DataRetentionConfig.TaskLogsRetentionConfig = &pb.TaskLogsRetentionConfig{}
	}
	if config.DataRetentionConfig.TaskLogsRetentionConfig.StorageMode == pb.TaskLogsRetentionConfig_TASK_LOGS_STORAGE_MODE_UNSPECIFIED {
		config.DataRetentionConfig.TaskLogsRetentionConfig.StorageMode = pb.TaskLogsRetentionConfig_CLOUD_LOGGING_ONLY
	}
	if config.DatabaseConfig == nil {
		config.DatabaseConfig = &pb.DatabaseConfig{}
	}
	if config.DatabaseConfig.MachineType == "" {
		config.DatabaseConfig.MachineType = "db-custom-2-7680"
	}
	if config.EncryptionConfig == nil {
		config.EncryptionConfig = &pb.EncryptionConfig{}
	}
	if config.EnvironmentSize == pb.EnvironmentConfig_ENVIRONMENT_SIZE_UNSPECIFIED {
		config.EnvironmentSize = pb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL
	}
	if config.MaintenanceWindow == nil {
		config.MaintenanceWindow = &pb.MaintenanceWindow{}
	}
	if config.MaintenanceWindow.StartTime == nil {
		config.MaintenanceWindow.StartTime = timestamppb.New(time.Unix(0, 0)) // "1970-01-01T00:00:00Z"
	}
	if config.MaintenanceWindow.EndTime == nil {
		config.MaintenanceWindow.EndTime = timestamppb.New(time.Unix(14400, 0)) // "1970-01-01T04:00:00Z"
	}
	if config.MaintenanceWindow.Recurrence == "" {
		config.MaintenanceWindow.Recurrence = "FREQ=WEEKLY;BYDAY=FR,SA,SU"
	}

	if config.NodeConfig == nil {
		config.NodeConfig = &pb.NodeConfig{}
	}
	if config.NodeConfig.IpAllocationPolicy == nil {
		config.NodeConfig.IpAllocationPolicy = &pb.IPAllocationPolicy{}
	}
	if !config.NodeConfig.IpAllocationPolicy.UseIpAliases {
		config.NodeConfig.IpAllocationPolicy.UseIpAliases = true
	}
	if config.NodeConfig.Network == "" {
		config.NodeConfig.Network = "projects/${projectId}/global/networks/default"
	}
	if config.NodeConfig.ServiceAccount == "" {
		config.NodeConfig.ServiceAccount = "${projectNumber}-compute@developer.gserviceaccount.com"
	}

	if config.PrivateEnvironmentConfig == nil {
		config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
	if config.PrivateEnvironmentConfig.CloudComposerNetworkIpv4CidrBlock == "" {
		config.PrivateEnvironmentConfig.CloudComposerNetworkIpv4CidrBlock = "172.31.245.0/24"
	}
	if config.PrivateEnvironmentConfig.CloudSqlIpv4CidrBlock == "" {
		config.PrivateEnvironmentConfig.CloudSqlIpv4CidrBlock = "10.0.0.0/12"
	}
	if config.PrivateEnvironmentConfig.PrivateClusterConfig == nil {
		config.PrivateEnvironmentConfig.PrivateClusterConfig = &pb.PrivateClusterConfig{}
	}

	if config.SoftwareConfig == nil {
		config.SoftwareConfig = &pb.SoftwareConfig{}
	}
	if config.SoftwareConfig.CloudDataLineageIntegration == nil {
		config.SoftwareConfig.CloudDataLineageIntegration = &pb.CloudDataLineageIntegration{}
	}
	if config.SoftwareConfig.PythonVersion == "" {
		config.SoftwareConfig.PythonVersion = "3"
	}

	// While 'imageVersion' is unlikely to be unset, the following handles it anyway for completeness.
	// Note:  Consider using a proper default if unset, instead of the value below.
	if config.SoftwareConfig.ImageVersion == "" {
		config.SoftwareConfig.ImageVersion = "composer-2.11.3-airflow-2.10.2"
	}

	if config.WebServerNetworkAccessControl == nil {
		config.WebServerNetworkAccessControl = &pb.WebServerNetworkAccessControl{}
	}
	if len(config.WebServerNetworkAccessControl.AllowedIpRanges) == 0 {
		config.WebServerNetworkAccessControl.AllowedIpRanges = []*pb.WebServerNetworkAccessControl_AllowedIpRange{
			{
				Description: "Allows access from all IPv4 addresses (default value)",
				Value:       "0.0.0.0/0",
			},
			{
				Description: "Allows access from all IPv6 addresses (default value)",
				Value:       "::0/0",
			},
		}
	}

	if config.WorkloadsConfig == nil {
		config.WorkloadsConfig = &pb.WorkloadsConfig{}
	}

	if config.WorkloadsConfig.Scheduler == nil {
		config.WorkloadsConfig.Scheduler = &pb.WorkloadsConfig_SchedulerResource{}
	}
	if config.WorkloadsConfig.Scheduler.Count == 0 {
		config.WorkloadsConfig.Scheduler.Count = 1
	}
	if config.WorkloadsConfig.Scheduler.Cpu == 0 {
		config.WorkloadsConfig.Scheduler.Cpu = 0.5
	}
	if config.WorkloadsConfig.Scheduler.MemoryGb == 0 {
		config.WorkloadsConfig.Scheduler.MemoryGb = 2
	}
	if config.WorkloadsConfig.Scheduler.StorageGb == 0 {
		config.WorkloadsConfig.Scheduler.StorageGb = 1
	}

	if config.WorkloadsConfig.WebServer == nil {
		config.WorkloadsConfig.WebServer = &pb.WorkloadsConfig_WebServerResource{}
	}
	if config.WorkloadsConfig.WebServer.Cpu == 0 {
		config.WorkloadsConfig.WebServer.Cpu = 0.5
	}
	if config.WorkloadsConfig.WebServer.MemoryGb == 0 {
		config.WorkloadsConfig.WebServer.MemoryGb = 2
	}
	if config.WorkloadsConfig.WebServer.StorageGb == 0 {
		config.WorkloadsConfig.WebServer.StorageGb = 1
	}

	if config.WorkloadsConfig.Worker == nil {
		config.WorkloadsConfig.Worker = &pb.WorkloadsConfig_WorkerResource{}
	}
	if config.WorkloadsConfig.Worker.Cpu == 0 {
		config.WorkloadsConfig.Worker.Cpu = 0.5
	}
	if config.WorkloadsConfig.Worker.MaxCount == 0 {
		config.WorkloadsConfig.Worker.MaxCount = 3
	}
	if config.WorkloadsConfig.Worker.MemoryGb == 0 {
		config.WorkloadsConfig.Worker.MemoryGb = 2
	}
	if config.WorkloadsConfig.Worker.MinCount == 0 {
		config.WorkloadsConfig.Worker.MinCount = 1
	}
	if config.WorkloadsConfig.Worker.StorageGb == 0 {
		config.WorkloadsConfig.Worker.StorageGb = 1
	}

	if config.PrivateEnvironmentConfig == nil {
		config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
}

type environmentName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *environmentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/environments/%s", n.Project.ID, n.Location, n.Name)
}

func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "environments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &environmentName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
