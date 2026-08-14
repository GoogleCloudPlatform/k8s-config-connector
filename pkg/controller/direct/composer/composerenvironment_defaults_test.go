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

package composer

import (
	"strings"
	"testing"

	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestPopulateDefaults(t *testing.T) {
	t.Run("populates server defaults for omitted fields without diff", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					NodeConfig: &krm.NodeConfig{
						ServiceAccountRef: &refs.IAMServiceAccountRef{External: "sa@p1.iam.gserviceaccount.com"},
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Name:       "projects/p1/locations/us-central1/environments/env1",
			Uuid:       "test-uuid",
			State:      composerpb.Environment_RUNNING,
			CreateTime: direct.StringTimestamp_ToProto(&direct.MapContext{}, direct.LazyPtr("2024-04-01T12:34:56.123456Z")),
			UpdateTime: direct.StringTimestamp_ToProto(&direct.MapContext{}, direct.LazyPtr("2024-04-01T12:34:56.123456Z")),
			StorageConfig: &composerpb.StorageConfig{
				Bucket: "us-central1-auto-bucket",
			},
			Config: &composerpb.EnvironmentConfig{
				AirflowUri:      "https://test.composer.googleusercontent.com",
				AirflowByoidUri: "https://test.composer.byoid.googleusercontent.com",
				DagGcsPrefix:    "gs://us-central1-auto-bucket/dags",
				EnvironmentSize: composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL,
				NodeConfig: &composerpb.NodeConfig{
					ServiceAccount:                "sa@p1.iam.gserviceaccount.com",
					ComposerInternalIpv4CidrBlock: "100.64.128.0/20",
					Network:                       "projects/p1/global/networks/default",
					MachineType:                   "n1-standard-1",
					DiskSizeGb:                    100,
				},
				SoftwareConfig: &composerpb.SoftwareConfig{
					ImageVersion:         "composer-3-airflow-2.11.1-build.14",
					WebServerPluginsMode: composerpb.SoftwareConfig_PLUGINS_ENABLED,
					PythonVersion:        "3",
					SchedulerCount:       1,
				},
				WebServerNetworkAccessControl: &composerpb.WebServerNetworkAccessControl{
					AllowedIpRanges: []*composerpb.WebServerNetworkAccessControl_AllowedIpRange{
						{
							Value:       "0.0.0.0/0",
							Description: "Allows access from all IPv4 addresses (default value)",
						},
						{
							Value:       "::0/0",
							Description: "Allows access from all IPv6 addresses (default value)",
						},
					},
				},
				DatabaseConfig: &composerpb.DatabaseConfig{
					MachineType: "db-custom-2-7680",
				},
				MaintenanceWindow: &composerpb.MaintenanceWindow{
					Recurrence: "FREQ=WEEKLY;BYDAY=FR,SA,SU",
				},
				PrivateEnvironmentConfig: &composerpb.PrivateEnvironmentConfig{
					NetworkingConfig: &composerpb.NetworkingConfig{
						ConnectionType: composerpb.NetworkingConfig_CONNECTION_TYPE_UNSPECIFIED,
					},
				},
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					Scheduler: &composerpb.WorkloadsConfig_SchedulerResource{
						Count:     1,
						Cpu:       0.5,
						MemoryGb:  2,
						StorageGb: 1,
					},
				},
				DataRetentionConfig: &composerpb.DataRetentionConfig{
					AirflowMetadataRetentionConfig: &composerpb.AirflowMetadataRetentionPolicyConfig{
						RetentionDays: 60,
						RetentionMode: composerpb.AirflowMetadataRetentionPolicyConfig_RETENTION_MODE_ENABLED,
					},
				},
			},
		}

		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		populateDesiredWithDefaults(desired, desiredPb)
		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		err := validateUpdatableFields(desiredPb, actual)
		if err != nil {
			t.Errorf("expected nil error after populating defaults, got %v", err)
		}
	})

	t.Run("preserves explicit desired modifications on immutable fields", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					NodeConfig: &krm.NodeConfig{
						MachineType: direct.LazyPtr("n1-standard-4"),
					},
				},
			},
		}
		actual := defaultEnvironmentPb()
		actual.Name = "projects/p1/locations/us-central1/environments/env1"
		actual.Config.NodeConfig = &composerpb.NodeConfig{
			MachineType: "n1-standard-1",
		}

		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		populateDesiredWithDefaults(desired, desiredPb)
		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for machineType modification, got nil")
		}
		expectedMsg := `updating field(s) [config.node_config.machine_type] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})

	t.Run("defaultEnvironmentPb returns universal static defaults", func(t *testing.T) {
		defPb := defaultEnvironmentPb()
		cfg := defPb.GetConfig()
		if len(cfg.GetWebServerNetworkAccessControl().GetAllowedIpRanges()) != 2 {
			t.Errorf("expected 2 allowed IP ranges, got %v", len(cfg.GetWebServerNetworkAccessControl().GetAllowedIpRanges()))
		}
		if cfg.GetEnvironmentSize() != 0 {
			t.Errorf("expected unspecified environment size, got %v", cfg.GetEnvironmentSize())
		}
		if cfg.GetMaintenanceWindow() != nil {
			t.Errorf("expected nil maintenance window in template, got %v", cfg.GetMaintenanceWindow())
		}
		if cfg.GetWorkloadsConfig() != nil {
			t.Errorf("expected nil workloads config in template, got %v", cfg.GetWorkloadsConfig())
		}
	})

	t.Run("populateDesiredWithDefaults sets universal static defaults only", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{},
		}
		desiredPb := &composerpb.Environment{}
		populateDesiredWithDefaults(desired, desiredPb)

		cfg := desiredPb.GetConfig()
		if len(cfg.GetWebServerNetworkAccessControl().GetAllowedIpRanges()) != 2 {
			t.Errorf("expected 2 allowed IP ranges, got %v", len(cfg.GetWebServerNetworkAccessControl().GetAllowedIpRanges()))
		}
		// Dynamic fields should NOT be set by populateDesiredWithDefaults
		if cfg.GetEnvironmentSize() != 0 {
			t.Errorf("expected 0 environmentSize in desiredPb from populateDesiredWithDefaults, got %v", cfg.GetEnvironmentSize())
		}
		if cfg.GetWorkloadsConfig() != nil {
			t.Errorf("expected nil workloadsConfig in desiredPb from populateDesiredWithDefaults, got %v", cfg.GetWorkloadsConfig())
		}
		if cfg.GetMaintenanceWindow() != nil {
			t.Errorf("expected nil maintenanceWindow in desiredPb from populateDesiredWithDefaults, got %v", cfg.GetMaintenanceWindow())
		}
		if cfg.GetNodeConfig().GetNetwork() != "" {
			t.Errorf("expected empty network in desiredPb, got %v", cfg.GetNodeConfig().GetNetwork())
		}
		if cfg.GetNodeConfig().GetMachineType() != "" {
			t.Errorf("expected empty machineType in desiredPb, got %v", cfg.GetNodeConfig().GetMachineType())
		}
		if cfg.GetNodeConfig().GetDiskSizeGb() != 0 {
			t.Errorf("expected 0 diskSizeGb in desiredPb, got %v", cfg.GetNodeConfig().GetDiskSizeGb())
		}
		if cfg.GetSoftwareConfig().GetPythonVersion() != "" {
			t.Errorf("expected empty pythonVersion in desiredPb, got %v", cfg.GetSoftwareConfig().GetPythonVersion())
		}
		if cfg.GetNodeConfig().GetComposerInternalIpv4CidrBlock() != "" {
			t.Errorf("expected empty CIDR block in desired, got %v", cfg.GetNodeConfig().GetComposerInternalIpv4CidrBlock())
		}
		if desiredPb.GetStorageConfig().GetBucket() != "" {
			t.Errorf("expected empty bucket in desired, got %v", desiredPb.GetStorageConfig().GetBucket())
		}
	})

	t.Run("populateDesiredWithActualIfComputed copies dynamic defaults", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{},
		}
		desiredPb := &composerpb.Environment{}
		actual := &composerpb.Environment{
			StorageConfig: &composerpb.StorageConfig{
				Bucket: "auto-bucket-123",
			},
			Config: &composerpb.EnvironmentConfig{
				NodeConfig: &composerpb.NodeConfig{
					ComposerInternalIpv4CidrBlock: "100.64.128.0/20",
					Network:                       "projects/p1/global/networks/default",
					MachineType:                   "n1-standard-1",
					DiskSizeGb:                    100,
				},
				SoftwareConfig: &composerpb.SoftwareConfig{
					ImageVersion:         "composer-2.11.3-airflow-2.10.2",
					PythonVersion:        "3",
					SchedulerCount:       1,
					WebServerPluginsMode: composerpb.SoftwareConfig_PLUGINS_ENABLED,
				},
				DatabaseConfig: &composerpb.DatabaseConfig{
					MachineType: "db-custom-2-7680",
					Zone:        "us-central1-a",
				},
				DataRetentionConfig: &composerpb.DataRetentionConfig{
					AirflowMetadataRetentionConfig: &composerpb.AirflowMetadataRetentionPolicyConfig{
						RetentionDays: 60,
						RetentionMode: composerpb.AirflowMetadataRetentionPolicyConfig_RETENTION_MODE_ENABLED,
					},
				},
			},
		}

		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		if desiredPb.GetStorageConfig().GetBucket() != "auto-bucket-123" {
			t.Errorf("expected bucket auto-bucket-123, got %v", desiredPb.GetStorageConfig().GetBucket())
		}
		if desiredPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock() != "100.64.128.0/20" {
			t.Errorf("expected CIDR 100.64.128.0/20, got %v", desiredPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock())
		}
		if desiredPb.GetConfig().GetNodeConfig().GetNetwork() != "projects/p1/global/networks/default" {
			t.Errorf("expected default network, got %v", desiredPb.GetConfig().GetNodeConfig().GetNetwork())
		}
		if desiredPb.GetConfig().GetNodeConfig().GetMachineType() != "n1-standard-1" {
			t.Errorf("expected machineType n1-standard-1, got %v", desiredPb.GetConfig().GetNodeConfig().GetMachineType())
		}
		if desiredPb.GetConfig().GetNodeConfig().GetDiskSizeGb() != 100 {
			t.Errorf("expected diskSizeGb 100, got %v", desiredPb.GetConfig().GetNodeConfig().GetDiskSizeGb())
		}
		if desiredPb.GetConfig().GetSoftwareConfig().GetImageVersion() != "composer-2.11.3-airflow-2.10.2" {
			t.Errorf("expected imageVersion composer-2.11.3-airflow-2.10.2, got %v", desiredPb.GetConfig().GetSoftwareConfig().GetImageVersion())
		}
		if desiredPb.GetConfig().GetSoftwareConfig().GetPythonVersion() != "3" {
			t.Errorf("expected pythonVersion 3, got %v", desiredPb.GetConfig().GetSoftwareConfig().GetPythonVersion())
		}
		if desiredPb.GetConfig().GetSoftwareConfig().GetSchedulerCount() != 1 {
			t.Errorf("expected schedulerCount 1, got %v", desiredPb.GetConfig().GetSoftwareConfig().GetSchedulerCount())
		}
		if desiredPb.GetConfig().GetDatabaseConfig().GetMachineType() != "db-custom-2-7680" {
			t.Errorf("expected machineType db-custom-2-7680, got %v", desiredPb.GetConfig().GetDatabaseConfig().GetMachineType())
		}
		if desiredPb.GetConfig().GetDataRetentionConfig().GetAirflowMetadataRetentionConfig().GetRetentionDays() != 60 {
			t.Errorf("expected retention days 60, got %v", desiredPb.GetConfig().GetDataRetentionConfig().GetAirflowMetadataRetentionConfig().GetRetentionDays())
		}
	})

	t.Run("populateDesiredWithActualIfComputed preserves explicit user values for dynamic fields", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				StorageConfig: &krm.StorageConfig{
					BucketRef: &storagev1beta1.StorageBucketRef{
						External: "gs://my-custom-bucket",
					},
				},
				Config: &krm.EnvironmentConfig{
					DatabaseConfig: &krm.DatabaseConfig{
						Zone: direct.LazyPtr("us-central1-c"),
					},
					WorkloadsConfig: &krm.WorkloadsConfig{
						DagProcessor: &krm.WorkloadsConfig_DagProcessorResource{
							CPU: direct.LazyPtr("1.5"),
						},
					},
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		actual := &composerpb.Environment{
			StorageConfig: &composerpb.StorageConfig{
				Bucket: "auto-bucket",
			},
			Config: &composerpb.EnvironmentConfig{
				DatabaseConfig: &composerpb.DatabaseConfig{
					Zone: "us-central1-a",
				},
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					DagProcessor: &composerpb.WorkloadsConfig_DagProcessorResource{
						Cpu: 0.5,
					},
					Worker: &composerpb.WorkloadsConfig_WorkerResource{
						Cpu: 1.0,
					},
				},
			},
		}

		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		if desiredPb.GetStorageConfig().GetBucket() != "my-custom-bucket" {
			t.Errorf("expected bucket my-custom-bucket, got %v", desiredPb.GetStorageConfig().GetBucket())
		}
		if desiredPb.GetConfig().GetDatabaseConfig().GetZone() != "us-central1-c" {
			t.Errorf("expected database zone us-central1-c, got %v", desiredPb.GetConfig().GetDatabaseConfig().GetZone())
		}
		if desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCpu() != 1.5 {
			t.Errorf("expected DagProcessor CPU 1.5, got %v", desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCpu())
		}
		if desiredPb.GetConfig().GetWorkloadsConfig().GetWorker().GetCpu() != 1.0 {
			t.Errorf("expected Worker CPU 1.0 from actual, got %v", desiredPb.GetConfig().GetWorkloadsConfig().GetWorker().GetCpu())
		}
	})

	t.Run("single-pass path-based dynamic inheritance handles Composer 1, 2, and 3 fields", func(t *testing.T) {
		// Desired spec with Composer 1 nodeCount explicitly set, but omitting environmentSize and workloadsConfig
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					NodeCount: direct.LazyPtr(int32(5)),
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				NodeCount: 3,
				WebServerConfig: &composerpb.WebServerConfig{
					MachineType: "composer-n1-webserver-2",
				},
				MaintenanceWindow: &composerpb.MaintenanceWindow{
					Recurrence: "FREQ=WEEKLY;BYDAY=FR,SA,SU",
				},
				PrivateEnvironmentConfig: &composerpb.PrivateEnvironmentConfig{
					CloudComposerNetworkIpv4CidrBlock: "172.31.245.0/24",
					CloudSqlIpv4CidrBlock:             "10.0.0.0/12",
					PrivateClusterConfig:              &composerpb.PrivateClusterConfig{},
				},
			},
		}

		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		// Explicit user value (NodeCount = 5) must be preserved
		if desiredPb.GetConfig().GetNodeCount() != 5 {
			t.Errorf("expected NodeCount 5, got %v", desiredPb.GetConfig().GetNodeCount())
		}
		// Unset fields must be copied from actualPb
		if desiredPb.GetConfig().GetWebServerConfig().GetMachineType() != "composer-n1-webserver-2" {
			t.Errorf("expected WebServerConfig machineType composer-n1-webserver-2, got %v", desiredPb.GetConfig().GetWebServerConfig().GetMachineType())
		}
		if desiredPb.GetConfig().GetMaintenanceWindow().GetRecurrence() != "FREQ=WEEKLY;BYDAY=FR,SA,SU" {
			t.Errorf("expected weekly recurrence, got %v", desiredPb.GetConfig().GetMaintenanceWindow().GetRecurrence())
		}
		if desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudComposerNetworkIpv4CidrBlock() != "172.31.245.0/24" {
			t.Errorf("expected CIDR 172.31.245.0/24, got %v", desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudComposerNetworkIpv4CidrBlock())
		}
		if desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudSqlIpv4CidrBlock() != "10.0.0.0/12" {
			t.Errorf("expected CloudSQL CIDR 10.0.0.0/12, got %v", desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudSqlIpv4CidrBlock())
		}
	})

	t.Run("all acronym fields are populated correctly from actual when omitted", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{},
		}
		desiredPb := &composerpb.Environment{}

		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				NodeConfig: &composerpb.NodeConfig{
					ComposerInternalIpv4CidrBlock: "100.64.128.0/20",
					DiskSizeGb:                    100,
					IpAllocationPolicy: &composerpb.IPAllocationPolicy{
						UseIpAliases: true,
					},
				},
				PrivateEnvironmentConfig: &composerpb.PrivateEnvironmentConfig{
					CloudComposerNetworkIpv4CidrBlock: "172.31.245.0/24",
					CloudSqlIpv4CidrBlock:             "10.0.0.0/12",
					WebServerIpv4CidrBlock:            "172.31.250.0/24",
				},
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					DagProcessor: &composerpb.WorkloadsConfig_DagProcessorResource{
						Cpu:       1.5,
						MemoryGb:  2.0,
						StorageGb: 2.0,
						Count:     2,
					},
				},
			},
		}

		populateDesiredWithActualIfComputed(desired, desiredPb, actual)

		// Verify NodeConfig acronym fields (IPv4, CIDR, GB, IP)
		if desiredPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock() != "100.64.128.0/20" {
			t.Errorf("expected NodeConfig ComposerInternalIPv4CIDRBlock 100.64.128.0/20, got %v", desiredPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock())
		}
		if desiredPb.GetConfig().GetNodeConfig().GetDiskSizeGb() != 100 {
			t.Errorf("expected NodeConfig DiskSizeGB 100, got %v", desiredPb.GetConfig().GetNodeConfig().GetDiskSizeGb())
		}
		if !desiredPb.GetConfig().GetNodeConfig().GetIpAllocationPolicy().GetUseIpAliases() {
			t.Errorf("expected NodeConfig IPAllocationPolicy.useIpAliases true, got false")
		}

		// Verify PrivateEnvironmentConfig acronym fields (IPv4, CIDR, SQL)
		if desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudComposerNetworkIpv4CidrBlock() != "172.31.245.0/24" {
			t.Errorf("expected PrivateEnvironmentConfig CloudComposerNetworkIPv4CIDRBlock 172.31.245.0/24, got %v", desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudComposerNetworkIpv4CidrBlock())
		}
		if desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudSqlIpv4CidrBlock() != "10.0.0.0/12" {
			t.Errorf("expected PrivateEnvironmentConfig CloudSQLIPv4CIDRBlock 10.0.0.0/12, got %v", desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetCloudSqlIpv4CidrBlock())
		}
		if desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetWebServerIpv4CidrBlock() != "172.31.250.0/24" {
			t.Errorf("expected PrivateEnvironmentConfig WebServerIPv4CIDRBlock 172.31.250.0/24, got %v", desiredPb.GetConfig().GetPrivateEnvironmentConfig().GetWebServerIpv4CidrBlock())
		}

		// Verify WorkloadsConfig acronym fields (DagProcessor)
		if desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCpu() != 1.5 {
			t.Errorf("expected WorkloadsConfig DagProcessor CPU 1.5, got %v", desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCpu())
		}
		if desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCount() != 2 {
			t.Errorf("expected WorkloadsConfig DagProcessor Count 2, got %v", desiredPb.GetConfig().GetWorkloadsConfig().GetDagProcessor().GetCount())
		}
	})
}

func TestComputedFieldPathsValidity(t *testing.T) {
	env := &composerpb.Environment{
		StorageConfig: &composerpb.StorageConfig{},
		Config: &composerpb.EnvironmentConfig{
			NodeConfig:               &composerpb.NodeConfig{},
			SoftwareConfig:           &composerpb.SoftwareConfig{},
			DatabaseConfig:           &composerpb.DatabaseConfig{},
			WebServerConfig:          &composerpb.WebServerConfig{},
			PrivateEnvironmentConfig: &composerpb.PrivateEnvironmentConfig{},
			WorkloadsConfig:          &composerpb.WorkloadsConfig{},
		},
	}
	parentMap := buildParentMap(env, env)

	for _, path := range computedFieldPaths {
		lastDot := strings.LastIndex(path, ".")
		var parentPath, leafName string
		if lastDot == -1 {
			parentPath = ""
			leafName = path
		} else {
			parentPath = path[:lastDot]
			leafName = path[lastDot+1:]
		}

		pair, ok := parentMap[parentPath]
		if !ok {
			t.Fatalf("path %q: parent path %q is not registered in buildParentMap", path, parentPath)
		}
		fd := findProtoField(pair.actual.Descriptor(), leafName)
		if fd == nil {
			t.Fatalf("path %q: field %q does not exist on proto message %s", path, leafName, pair.actual.Descriptor().FullName())
		}
	}
}

func TestFindProtoField(t *testing.T) {
	nodeDesc := (&composerpb.NodeConfig{}).ProtoReflect().Descriptor()
	privDesc := (&composerpb.PrivateEnvironmentConfig{}).ProtoReflect().Descriptor()
	workloadDesc := (&composerpb.WorkloadsConfig_DagProcessorResource{}).ProtoReflect().Descriptor()

	tests := []struct {
		desc      protoreflect.MessageDescriptor
		krmLeaf   string
		wantField protoreflect.Name
	}{
		{nodeDesc, "serviceAccountRef", "service_account"},
		{nodeDesc, "serviceAccount", "service_account"},
		{nodeDesc, "ComposerInternalIPv4CIDRBlock", "composer_internal_ipv4_cidr_block"},
		{nodeDesc, "composerInternalIpv4CidrBlock", "composer_internal_ipv4_cidr_block"},
		{nodeDesc, "diskSizeGb", "disk_size_gb"},
		{nodeDesc, "DiskSizeGB", "disk_size_gb"},
		{privDesc, "CloudComposerNetworkIPv4CIDRBlock", "cloud_composer_network_ipv4_cidr_block"},
		{privDesc, "CloudSQLIPv4CIDRBlock", "cloud_sql_ipv4_cidr_block"},
		{privDesc, "WebServerIPv4CIDRBlock", "web_server_ipv4_cidr_block"},
		{privDesc, "CloudComposerConnectionSubnetworkRef", "cloud_composer_connection_subnetwork"},
		{workloadDesc, "cpu", "cpu"},
		{workloadDesc, "CPU", "cpu"},
		{workloadDesc, "memoryGb", "memory_gb"},
		{workloadDesc, "storageGb", "storage_gb"},
	}

	for _, tc := range tests {
		fd := findProtoField(tc.desc, tc.krmLeaf)
		if fd == nil {
			t.Errorf("findProtoField(%s, %q) = nil, want %q", tc.desc.FullName(), tc.krmLeaf, tc.wantField)
			continue
		}
		if fd.Name() != tc.wantField {
			t.Errorf("findProtoField(%s, %q) = %q, want %q", tc.desc.FullName(), tc.krmLeaf, fd.Name(), tc.wantField)
		}
	}

	// Non-existent field should return nil
	if fd := findProtoField(nodeDesc, "nonExistentField"); fd != nil {
		t.Errorf("expected nil for nonExistentField, got %q", fd.Name())
	}
}
