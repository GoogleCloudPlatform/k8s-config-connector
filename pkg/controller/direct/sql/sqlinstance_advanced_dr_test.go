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

package sql

import (
	"testing"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	api "google.golang.org/api/sqladmin/v1beta4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsAdvancedDREnabled(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        bool
	}{
		{
			name:        "nil object",
			annotations: nil,
			want:        false,
		},
		{
			name:        "no annotations",
			annotations: map[string]string{},
			want:        false,
		},
		{
			name: "disabled",
			annotations: map[string]string{
				EnableAdvancedDRAnnotation: "disabled",
			},
			want: false,
		},
		{
			name: "enabled",
			annotations: map[string]string{
				EnableAdvancedDRAnnotation: "enabled",
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var obj *krm.SQLInstance
			if tc.annotations != nil {
				obj = &krm.SQLInstance{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: tc.annotations,
					},
				}
			}
			got := IsAdvancedDREnabled(obj)
			if got != tc.want {
				t.Errorf("IsAdvancedDREnabled() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsDesiredPrimaryWithDR(t *testing.T) {
	tests := []struct {
		name string
		inst *api.DatabaseInstance
		want bool
	}{
		{
			name: "nil instance",
			inst: nil,
			want: false,
		},
		{
			name: "standard instance",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
			},
			want: false,
		},
		{
			name: "read replica instance",
			inst: &api.DatabaseInstance{
				InstanceType: "READ_REPLICA_INSTANCE",
			},
			want: false,
		},
		{
			name: "primary with replica ref but master set",
			inst: &api.DatabaseInstance{
				InstanceType:       "CLOUD_SQL_INSTANCE",
				MasterInstanceName: "some-master",
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "replica-1",
				},
			},
			want: false,
		},
		{
			name: "primary with empty replication cluster",
			inst: &api.DatabaseInstance{
				InstanceType:       "CLOUD_SQL_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{},
			},
			want: false,
		},
		{
			name: "valid primary with DR",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "replica-1",
					DrReplica:             false,
				},
			},
			want: true,
		},
		{
			name: "primary but marked as drReplica",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "replica-1",
					DrReplica:             true,
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsDesiredPrimaryWithDR(tc.inst)
			if got != tc.want {
				t.Errorf("IsDesiredPrimaryWithDR() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsDesiredReplicaWithDR(t *testing.T) {
	tests := []struct {
		name string
		inst *api.DatabaseInstance
		want bool
	}{
		{
			name: "nil instance",
			inst: nil,
			want: false,
		},
		{
			name: "primary instance",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
			},
			want: false,
		},
		{
			name: "read replica but no master",
			inst: &api.DatabaseInstance{
				InstanceType: "READ_REPLICA_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: true,
				},
			},
			want: false,
		},
		{
			name: "read replica but no replication cluster",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
			},
			want: true,
		},
		{
			name: "read replica but not drReplica",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: false,
				},
			},
			want: true,
		},
		{
			name: "dr replica but designates another failover replica",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica:             true,
					FailoverDrReplicaName: "replica-2",
				},
			},
			want: false,
		},
		{
			name: "valid DR replica",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: true,
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsDesiredReplicaWithDR(tc.inst)
			if got != tc.want {
				t.Errorf("IsDesiredReplicaWithDR() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsActualPrimaryWithDR(t *testing.T) {
	tests := []struct {
		name string
		inst *api.DatabaseInstance
		want bool
	}{
		{
			name: "not desired primary",
			inst: &api.DatabaseInstance{
				InstanceType: "READ_REPLICA_INSTANCE",
			},
			want: false,
		},
		{
			name: "desired primary but drReplica is true",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "replica-1",
					DrReplica:             true,
				},
			},
			want: false,
		},
		{
			name: "desired primary and drReplica is false",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "replica-1",
					DrReplica:             false,
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsActualPrimaryWithDR(tc.inst)
			if got != tc.want {
				t.Errorf("IsActualPrimaryWithDR() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsActualReplicaWithDR(t *testing.T) {
	tests := []struct {
		name string
		inst *api.DatabaseInstance
		want bool
	}{
		{
			name: "not desired replica",
			inst: &api.DatabaseInstance{
				InstanceType: "CLOUD_SQL_INSTANCE",
			},
			want: false,
		},
		{
			name: "desired replica but replication cluster is nil",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: nil,
			},
			want: false,
		},
		{
			name: "desired replica but drReplica is false",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: false,
				},
			},
			want: false,
		},
		{
			name: "desired replica and drReplica is true",
			inst: &api.DatabaseInstance{
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "master-1",
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: true,
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsActualReplicaWithDR(tc.inst)
			if got != tc.want {
				t.Errorf("IsActualReplicaWithDR() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPreprocessDRForCreateIfEnabled(t *testing.T) {
	desired := &krm.SQLInstance{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				EnableAdvancedDRAnnotation: "enabled",
			},
		},
		Spec: krm.SQLInstanceSpec{
			ReplicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{External: "some-replica"},
			},
		},
	}

	PreprocessDRForCreateIfEnabled(desired)

	if desired.Spec.ReplicationCluster.FailoverDRReplicaRef != nil {
		t.Errorf("expected FailoverDRReplicaRef to be nil, got %v", desired.Spec.ReplicationCluster.FailoverDRReplicaRef)
	}

	// Disabled case
	desiredDisabled := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			ReplicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{External: "some-replica"},
			},
		},
	}

	PreprocessDRForCreateIfEnabled(desiredDisabled)

	if desiredDisabled.Spec.ReplicationCluster.FailoverDRReplicaRef == nil {
		t.Error("expected FailoverDRReplicaRef to remain non-nil when Advanced DR is disabled")
	}
}

func TestShouldSkipUpdateForAdvancedDR(t *testing.T) {
	desiredKRM := &krm.SQLInstance{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				EnableAdvancedDRAnnotation: "enabled",
			},
		},
	}

	desiredKRMDisabled := &krm.SQLInstance{}

	primaryDRProto := &api.DatabaseInstance{
		InstanceType: "CLOUD_SQL_INSTANCE",
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "replica-1",
			DrReplica:             false,
		},
	}

	replicaDRProto := &api.DatabaseInstance{
		InstanceType:       "READ_REPLICA_INSTANCE",
		MasterInstanceName: "primary-1",
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica: true,
		},
	}

	replicaDRProtoNoCluster := &api.DatabaseInstance{
		InstanceType:       "READ_REPLICA_INSTANCE",
		MasterInstanceName: "primary-1",
		ReplicationCluster: nil,
	}

	tests := []struct {
		name       string
		desiredKRM *krm.SQLInstance
		desired    *api.DatabaseInstance
		actual     *api.DatabaseInstance
		want       bool
	}{
		{
			name:       "disabled",
			desiredKRM: desiredKRMDisabled,
			desired:    primaryDRProto,
			actual:     replicaDRProto,
			want:       false,
		},
		{
			name:       "nil protos",
			desiredKRM: desiredKRM,
			desired:    nil,
			actual:     nil,
			want:       false,
		},
		{
			name:       "no role switch (primary matches primary)",
			desiredKRM: desiredKRM,
			desired:    primaryDRProto,
			actual:     primaryDRProto,
			want:       false,
		},
		{
			name:       "no role switch (replica matches replica)",
			desiredKRM: desiredKRM,
			desired:    replicaDRProto,
			actual:     replicaDRProto,
			want:       false,
		},
		{
			name:       "demoted primary to replica",
			desiredKRM: desiredKRM,
			desired:    primaryDRProto,
			actual:     replicaDRProto,
			want:       true,
		},
		{
			name:       "promoted replica to primary",
			desiredKRM: desiredKRM,
			desired:    replicaDRProto,
			actual:     primaryDRProto,
			want:       true,
		},
		{
			name:       "promoted replica (no replication cluster in desired) to primary",
			desiredKRM: desiredKRM,
			desired:    replicaDRProtoNoCluster,
			actual:     primaryDRProto,
			want:       true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ShouldSkipUpdateForAdvancedDR(tc.desiredKRM, tc.desired, tc.actual)
			if got != tc.want {
				t.Errorf("ShouldSkipUpdateForAdvancedDR() = %v, want %v", got, tc.want)
			}
		})
	}
}
