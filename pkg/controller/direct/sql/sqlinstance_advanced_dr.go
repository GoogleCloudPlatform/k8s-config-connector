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
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	api "google.golang.org/api/sqladmin/v1beta4"
)

// EnableAdvancedDRAnnotation is the opt-in annotation for Cloud SQL Advanced Disaster Recovery (DR) support.
// When set to "enabled", Config Connector accommodates cross-region DR setups by:
// 1. Allowing circular-dependency-free initial creation of Primary and DR Replica instances.
// 2. Preventing destructive role reversions when an out-of-band switchover or failover occurs in GCP.
const EnableAdvancedDRAnnotation = "cnrm.cloud.google.com/sqlinstance-advanced-dr"

// IsAdvancedDREnabled checks whether the Advanced DR feature is enabled on the KRM SQLInstance resource.
// Users must explicitly set `cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"`.
func IsAdvancedDREnabled(obj *krmv1beta1.SQLInstance) bool {
	if obj == nil {
		return false
	}
	return obj.GetAnnotations()[EnableAdvancedDRAnnotation] == "enabled"
}

// IsPrimaryWithDR returns true if the instance is a valid Primary instance configured for Advanced DR.
//
// In Cloud SQL:
// 1. A primary instance has instanceType "CLOUD_SQL_INSTANCE".
// 2. A primary instance does NOT have a master (masterInstanceName must be empty).
// 3. A primary instance configures replicationCluster.failoverDrReplicaName pointing to its DR replica.
// 4. A primary instance is NOT itself a DR replica (replicationCluster.drReplica must be false).
func IsPrimaryWithDR(inst *api.DatabaseInstance) bool {
	if inst == nil {
		return false
	}
	if inst.InstanceType != "CLOUD_SQL_INSTANCE" {
		return false
	}
	if inst.MasterInstanceName != "" {
		return false
	}
	if inst.ReplicationCluster == nil {
		return false
	}
	if inst.ReplicationCluster.FailoverDrReplicaName == "" {
		return false
	}
	if inst.ReplicationCluster.DrReplica {
		return false
	}
	return true
}

// IsReplicaWithDR returns true if the instance is a valid Replica designated for Advanced DR.
//
// In Cloud SQL:
// 1. A replica instance has instanceType "READ_REPLICA_INSTANCE".
// 2. A replica instance MUST have a master configured (masterInstanceName must be non-empty).
// 3. A replica instance has replicationCluster.drReplica set to true to indicate it is the DR replica.
// 4. A replica instance does NOT designate another failover DR replica (replicationCluster.failoverDrReplicaName must be empty).
func IsReplicaWithDR(inst *api.DatabaseInstance) bool {
	if inst == nil {
		return false
	}
	if inst.InstanceType != "READ_REPLICA_INSTANCE" {
		return false
	}
	if inst.MasterInstanceName == "" {
		return false
	}
	if inst.ReplicationCluster == nil {
		return false
	}
	if !inst.ReplicationCluster.DrReplica {
		return false
	}
	if inst.ReplicationCluster.FailoverDrReplicaName != "" {
		return false
	}
	return true
}

// PreprocessDRForCreateIfEnabled unsets failoverDRReplicaRef during initial creation to break the
// circular dependency between Primary and DR Replica instances.
//
// In Advanced DR, the Primary instance references the DR Replica via spec.replicationCluster.failoverDRReplicaRef,
// while the DR Replica references the Primary via spec.masterInstanceRef. If both references are strictly enforced
// at creation time, neither instance can be created first.
//
// When Advanced DR is enabled, clearing failoverDRReplicaRef on initial creation allows the Primary instance to be
// provisioned in GCP as a standalone Enterprise Plus instance first. Once the Primary is ready, the DR Replica can
// be created with masterInstanceRef pointing to the Primary, completing the DR pairing.
func PreprocessDRForCreateIfEnabled(desired *krmv1beta1.SQLInstance) {
	if !IsAdvancedDREnabled(desired) {
		return
	}
	if desired.Spec.ReplicationCluster == nil {
		return
	}
	desired.Spec.ReplicationCluster.FailoverDRReplicaRef = nil
}

// ShouldSkipUpdateForAdvancedDR determines if the controller should skip mutating the GCP resource
// during Update() because an out-of-band DR operation (switchover or failover) has inverted the instance roles.
//
// When a switchover or failover is performed out-of-band in GCP:
// - The former Primary becomes a DR Replica (instanceType: "READ_REPLICA_INSTANCE", masterInstanceName != "", drReplica: true).
// - The former DR Replica is promoted to Primary (instanceType: "CLOUD_SQL_INSTANCE", masterInstanceName == "", failoverDrReplicaName != "").
//
// Without this check, KCC would treat the swapped roles as drift against the static KRM specs and attempt
// to send GCP update requests to revert them.
//
// When this function returns true:
//  1. The controller skips sending role mutation requests to GCP.
//  2. The controller skips SetLastModifiedCookie to avoid storing an invalid state hash.
//  3. The controller syncs status.observedState and status.instanceType so KRM status accurately reflects
//     the new live role.
func ShouldSkipUpdateForAdvancedDR(desiredKRM *krmv1beta1.SQLInstance, desiredProto *api.DatabaseInstance, actualProto *api.DatabaseInstance) bool {
	if !IsAdvancedDREnabled(desiredKRM) {
		return false
	}
	if desiredProto == nil || actualProto == nil {
		return false
	}

	// Case 1: KRM spec defines Primary with DR, but GCP state is now a DR Replica (instance demoted via switchover).
	if IsPrimaryWithDR(desiredProto) && IsReplicaWithDR(actualProto) {
		return true
	}

	// Case 2: KRM spec defines DR Replica, but GCP state is now Primary (instance promoted via switchover/failover).
	if IsReplicaWithDR(desiredProto) && IsPrimaryWithDR(actualProto) {
		return true
	}

	return false
}
