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

const (
	// DefaultComposerInternalIpv4CidrBlock is the mock-simulated GCP default CIDR block.
	DefaultComposerInternalIpv4CidrBlock = "172.31.251.0/24"
	// DefaultComposerNetworkAttachment is the mock-simulated GCP default network attachment.
	DefaultComposerNetworkAttachment = "projects/${projectId}/regions/us-central1/networkAttachments/test-attachment"
)

func (s *ComposerV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No such environment found: %s", fqn)
		}
		return nil, err
	}

	return proto.CloneOf(obj), nil
}

func (s *ComposerV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {

	name, err := s.parseEnvironmentName(req.Environment.Name)
	if err != nil {
		return nil, err
	}

	if bucket := req.GetEnvironment().GetStorageConfig().GetBucket(); strings.Contains(bucket, "/") {
		return nil, status.Errorf(codes.InvalidArgument, "Composer doesn't have permission to read the requested Cloud Storage bucket: gs://%s", bucket)
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.Environment)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Environment_RUNNING
	obj.Uuid = "7eca3d3d-2b50-473a-b91f-bf106a0deb91"
	s.populateDefaultsForEnvironment(obj, name)

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
		ResourceUuid:  obj.Uuid,
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		resp := proto.CloneOf(obj)
		resp.State = pb.Environment_STATE_UNSPECIFIED
		return resp, nil
	})
}

func (s *ComposerV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.GetName())
	if err != nil {
		return nil, err
	}
	if bucket := req.GetEnvironment().GetStorageConfig().GetBucket(); strings.Contains(bucket, "/") {
		return nil, status.Errorf(codes.InvalidArgument, "Composer doesn't have permission to read the requested Cloud Storage bucket: gs://%s", bucket)
	}
	fqn := name.String()

	existing := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.CloneOf(existing)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		tokens := strings.Split(path, ".")
		switch normalizeField(tokens[0]) {
		case "labels":
			updated.Labels = req.GetEnvironment().GetLabels()
		case "config":
			if len(tokens) > 1 {
				switch normalizeField(tokens[1]) {
				case "nodecount":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.NodeCount = req.GetEnvironment().GetConfig().GetNodeCount()
				case "environmentsize":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.EnvironmentSize = req.GetEnvironment().GetConfig().GetEnvironmentSize()
				case "workloadsconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.WorkloadsConfig = req.GetEnvironment().GetConfig().GetWorkloadsConfig()
				case "maintenancewindow":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.MaintenanceWindow = req.GetEnvironment().GetConfig().GetMaintenanceWindow()
				case "softwareconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					if len(tokens) > 2 {
						switch normalizeField(tokens[2]) {
						case "imageversion":
							if updated.Config.SoftwareConfig == nil {
								updated.Config.SoftwareConfig = &pb.SoftwareConfig{}
							}
							updated.Config.SoftwareConfig.ImageVersion = req.GetEnvironment().GetConfig().GetSoftwareConfig().GetImageVersion()
						case "pypipackages":
							if updated.Config.SoftwareConfig == nil {
								updated.Config.SoftwareConfig = &pb.SoftwareConfig{}
							}
							updated.Config.SoftwareConfig.PypiPackages = req.GetEnvironment().GetConfig().GetSoftwareConfig().GetPypiPackages()
						case "airflowconfigoverrides":
							if updated.Config.SoftwareConfig == nil {
								updated.Config.SoftwareConfig = &pb.SoftwareConfig{}
							}
							updated.Config.SoftwareConfig.AirflowConfigOverrides = req.GetEnvironment().GetConfig().GetSoftwareConfig().GetAirflowConfigOverrides()
						case "envvariables":
							if updated.Config.SoftwareConfig == nil {
								updated.Config.SoftwareConfig = &pb.SoftwareConfig{}
							}
							updated.Config.SoftwareConfig.EnvVariables = req.GetEnvironment().GetConfig().GetSoftwareConfig().GetEnvVariables()
						default:
							updated.Config.SoftwareConfig = req.GetEnvironment().GetConfig().GetSoftwareConfig()
						}
					} else {
						updated.Config.SoftwareConfig = req.GetEnvironment().GetConfig().GetSoftwareConfig()
					}
				case "webservernetworkaccesscontrol":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.WebServerNetworkAccessControl = req.GetEnvironment().GetConfig().GetWebServerNetworkAccessControl()
				case "databaseconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.DatabaseConfig = req.GetEnvironment().GetConfig().GetDatabaseConfig()
				case "webserverconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.WebServerConfig = req.GetEnvironment().GetConfig().GetWebServerConfig()
				case "recoveryconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					if updated.Config.RecoveryConfig == nil {
						updated.Config.RecoveryConfig = &pb.RecoveryConfig{}
					}
					if len(tokens) > 2 {
						switch normalizeField(tokens[2]) {
						case "scheduledsnapshotsconfig":
							updated.Config.RecoveryConfig.ScheduledSnapshotsConfig = req.GetEnvironment().GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig()
						default:
							updated.Config.RecoveryConfig = req.GetEnvironment().GetConfig().GetRecoveryConfig()
						}
					} else {
						updated.Config.RecoveryConfig = req.GetEnvironment().GetConfig().GetRecoveryConfig()
					}
				case "resiliencemode":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.ResilienceMode = req.GetEnvironment().GetConfig().GetResilienceMode()
				case "masterauthorizednetworksconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.MasterAuthorizedNetworksConfig = req.GetEnvironment().GetConfig().GetMasterAuthorizedNetworksConfig()
				case "dataretentionconfig":
					if updated.Config == nil {
						updated.Config = &pb.EnvironmentConfig{}
					}
					updated.Config.DataRetentionConfig = req.GetEnvironment().GetConfig().GetDataRetentionConfig()
				default:
					return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
				}
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	updated.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		OperationType: pb.OperationMetadata_UPDATE,
		CreateTime:    timestamppb.New(now),
		Resource:      name.String(),
		State:         pb.OperationMetadata_PENDING,
		ResourceUuid:  updated.Uuid,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		resp := proto.CloneOf(updated)
		resp.State = pb.Environment_STATE_UNSPECIFIED
		return resp, nil
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
		ResourceUuid:  deleted.Uuid,
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationMetadata_SUCCEEDED
		return &emptypb.Empty{}, nil
	})
}

func (s *ComposerV1) populateDefaultsForEnvironment(obj *pb.Environment, name *environmentName) {
	if obj.StorageConfig == nil {
		obj.StorageConfig = &pb.StorageConfig{}
	}
	if obj.StorageConfig.Bucket == "" {
		obj.StorageConfig.Bucket = "us-central1-composerenviron-cd7334a9-bucket"
	}
	if obj.Config == nil {
		obj.Config = &pb.EnvironmentConfig{}
	}

	s.populateDefaultsForEnvironmentConfig(obj.Config, obj.StorageConfig.Bucket, name)
}

func (s *ComposerV1) populateDefaultsForEnvironmentConfig(config *pb.EnvironmentConfig, bucket string, name *environmentName) {
	config.AirflowByoidUri = "https://7eca3d3d2b50473ab91fbf106a0deb91-dot-us-central1.composer.byoid.googleusercontent.com"
	config.AirflowUri = "https://7eca3d3d2b50473ab91fbf106a0deb91-dot-us-central1.composer.googleusercontent.com"
	config.DagGcsPrefix = fmt.Sprintf("gs://%s/dags", bucket)

	if config.DataRetentionConfig == nil {
		config.DataRetentionConfig = &pb.DataRetentionConfig{}
	}
	if config.DataRetentionConfig.AirflowMetadataRetentionConfig == nil {
		config.DataRetentionConfig.AirflowMetadataRetentionConfig = &pb.AirflowMetadataRetentionPolicyConfig{
			RetentionMode: pb.AirflowMetadataRetentionPolicyConfig_RETENTION_MODE_ENABLED,
			RetentionDays: 60,
		}
	}

	if config.EncryptionConfig == nil {
		config.EncryptionConfig = &pb.EncryptionConfig{}
	}
	if config.EnvironmentSize == pb.EnvironmentConfig_ENVIRONMENT_SIZE_UNSPECIFIED {
		config.EnvironmentSize = pb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL
	}

	if config.NodeConfig == nil {
		config.NodeConfig = &pb.NodeConfig{}
	}
	if config.NodeConfig.ComposerInternalIpv4CidrBlock == "" {
		config.NodeConfig.ComposerInternalIpv4CidrBlock = "100.64.128.0/20"
	}
	if config.NodeConfig.IpAllocationPolicy == nil {
		config.NodeConfig.IpAllocationPolicy = &pb.IPAllocationPolicy{}
	}
	if config.NodeConfig.ServiceAccount == "" && name != nil {
		config.NodeConfig.ServiceAccount = fmt.Sprintf("sa-${uniqueId}@%s.iam.gserviceaccount.com", name.Project.ID)
	}

	if config.PrivateEnvironmentConfig == nil {
		config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
	if config.PrivateEnvironmentConfig.NetworkingConfig == nil {
		config.PrivateEnvironmentConfig.NetworkingConfig = &pb.NetworkingConfig{}
	}
	if config.PrivateEnvironmentConfig.NetworkingType == pb.PrivateEnvironmentConfig_NETWORKING_TYPE_UNSPECIFIED {
		config.PrivateEnvironmentConfig.NetworkingType = pb.PrivateEnvironmentConfig_PUBLIC
	}

	if config.SoftwareConfig == nil {
		config.SoftwareConfig = &pb.SoftwareConfig{}
	}
	if config.SoftwareConfig.AuditLogsReplicationMode == pb.SoftwareConfig_AUDIT_LOGS_REPLICATION_MODE_UNSPECIFIED {
		config.SoftwareConfig.AuditLogsReplicationMode = pb.SoftwareConfig_AUDIT_LOGS_REPLICATION_DISABLED
	}
	if config.SoftwareConfig.CloudDataLineageIntegration == nil {
		config.SoftwareConfig.CloudDataLineageIntegration = &pb.CloudDataLineageIntegration{}
	}
	if config.SoftwareConfig.ImageVersion == "" {
		config.SoftwareConfig.ImageVersion = "composer-3-airflow-2.11.1-build.14"
	}
	if config.SoftwareConfig.WebServerPluginsMode == pb.SoftwareConfig_WEB_SERVER_PLUGINS_MODE_UNSPECIFIED {
		config.SoftwareConfig.WebServerPluginsMode = pb.SoftwareConfig_PLUGINS_ENABLED
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
	if config.WorkloadsConfig.DagProcessor == nil {
		config.WorkloadsConfig.DagProcessor = &pb.WorkloadsConfig_DagProcessorResource{
			Count:     1,
			Cpu:       1,
			MemoryGb:  4,
			StorageGb: 1,
		}
	}
	if config.WorkloadsConfig.Scheduler == nil {
		config.WorkloadsConfig.Scheduler = &pb.WorkloadsConfig_SchedulerResource{
			Count:     1,
			Cpu:       0.5,
			MemoryGb:  2,
			StorageGb: 1,
		}
	}
	if config.WorkloadsConfig.Triggerer == nil {
		config.WorkloadsConfig.Triggerer = &pb.WorkloadsConfig_TriggererResource{
			Count:    1,
			Cpu:      1,
			MemoryGb: 2,
		}
	}
	if config.WorkloadsConfig.WebServer == nil {
		config.WorkloadsConfig.WebServer = &pb.WorkloadsConfig_WebServerResource{
			Cpu:       1,
			MemoryGb:  4,
			StorageGb: 1,
		}
	}
	if config.WorkloadsConfig.Worker == nil {
		config.WorkloadsConfig.Worker = &pb.WorkloadsConfig_WorkerResource{
			Cpu:       0.5,
			MemoryGb:  2,
			MinCount:  1,
			MaxCount:  3,
			StorageGb: 10,
		}
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

func normalizeField(s string) string {
	s = strings.ReplaceAll(s, "_", "")
	return strings.ToLower(s)
}

