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
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestPopulateDesiredWithDefaults(t *testing.T) {
	desired := &krm.ComposerEnvironment{
		Spec: krm.ComposerEnvironmentSpec{},
	}
	desiredPb := &composerpb.Environment{}
	populateDesiredWithDefaults(desired, desiredPb)

	wantPb := &composerpb.Environment{
		Config: &composerpb.EnvironmentConfig{
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
		},
	}
	if diff := cmp.Diff(wantPb, desiredPb, protocmp.Transform()); diff != "" {
		t.Errorf("populateDesiredWithDefaults() diff (-want +got):\n%s", diff)
	}
}

func TestPopulateDesiredWithActualIfComputed_CopiesDynamicDefaults(t *testing.T) {
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

	if diff := cmp.Diff(actual, desiredPb, protocmp.Transform()); diff != "" {
		t.Errorf("populateDesiredWithActualIfComputed diff (-want +got):\n%s", diff)
	}
}

func TestPopulateDesiredWithActualIfComputed_PreservesExplicitUserValues(t *testing.T) {
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

	wantPb := &composerpb.Environment{
		StorageConfig: &composerpb.StorageConfig{
			Bucket: "my-custom-bucket",
		},
		Config: &composerpb.EnvironmentConfig{
			DatabaseConfig: &composerpb.DatabaseConfig{
				Zone: "us-central1-c",
			},
			WorkloadsConfig: &composerpb.WorkloadsConfig{
				DagProcessor: &composerpb.WorkloadsConfig_DagProcessorResource{
					Cpu: 1.5,
				},
				Worker: &composerpb.WorkloadsConfig_WorkerResource{
					Cpu: 1.0,
				},
			},
		},
	}
	if diff := cmp.Diff(wantPb, desiredPb, protocmp.Transform()); diff != "" {
		t.Errorf("populateDesiredWithActualIfComputed diff (-want +got):\n%s", diff)
	}
}

func TestPopulateDesiredWithActualIfComputed_Generations(t *testing.T) {
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

	wantPb := &composerpb.Environment{
		Config: &composerpb.EnvironmentConfig{
			NodeCount: 5,
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
	if diff := cmp.Diff(wantPb, desiredPb, protocmp.Transform()); diff != "" {
		t.Errorf("populateDesiredWithActualIfComputed diff (-want +got):\n%s", diff)
	}
}

func TestPopulateDesiredWithActualIfComputed_Acronyms(t *testing.T) {
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

	if diff := cmp.Diff(actual, desiredPb, protocmp.Transform()); diff != "" {
		t.Errorf("populateDesiredWithActualIfComputed diff (-want +got):\n%s", diff)
	}
}

func TestPopulateDefaults_EndToEnd_NoDiff(t *testing.T) {
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
}

func TestPopulateDefaults_EndToEnd_PreservesExplicitModification(t *testing.T) {
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
