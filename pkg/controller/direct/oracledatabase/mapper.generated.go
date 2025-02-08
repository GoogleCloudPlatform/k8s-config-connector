// Copyright 2025 Google LLC
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

package oracledatabase

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutonomousDatabaseBackup_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackup) *krm.AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseBackup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AutonomousDatabase = direct.LazyPtr(in.GetAutonomousDatabase())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Properties = AutonomousDatabaseBackupProperties_FromProto(mapCtx, in.GetProperties())
	out.Labels = in.Labels
	return out
}
func AutonomousDatabaseBackup_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseBackup) *pb.AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackup{}
	out.Name = direct.ValueOf(in.Name)
	out.AutonomousDatabase = direct.ValueOf(in.AutonomousDatabase)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Properties = AutonomousDatabaseBackupProperties_ToProto(mapCtx, in.Properties)
	out.Labels = in.Labels
	return out
}
func AutonomousDatabaseBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackup) *krm.AutonomousDatabaseBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseBackupObservedState{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	out.Properties = AutonomousDatabaseBackupPropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	// MISSING: Labels
	return out
}
func AutonomousDatabaseBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseBackupObservedState) *pb.AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackup{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	out.Properties = AutonomousDatabaseBackupPropertiesObservedState_ToProto(mapCtx, in.Properties)
	// MISSING: Labels
	return out
}
func AutonomousDatabaseBackupProperties_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackupProperties) *krm.AutonomousDatabaseBackupProperties {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseBackupProperties{}
	// MISSING: Ocid
	out.RetentionPeriodDays = direct.LazyPtr(in.GetRetentionPeriodDays())
	// MISSING: CompartmentID
	// MISSING: DatabaseSizeTb
	// MISSING: DbVersion
	// MISSING: IsLongTermBackup
	// MISSING: IsAutomaticBackup
	// MISSING: IsRestorable
	out.KeyStoreID = direct.LazyPtr(in.GetKeyStoreId())
	out.KeyStoreWallet = direct.LazyPtr(in.GetKeyStoreWallet())
	out.KMSKeyID = direct.LazyPtr(in.GetKmsKeyId())
	out.KMSKeyVersionID = direct.LazyPtr(in.GetKmsKeyVersionId())
	// MISSING: LifecycleDetails
	// MISSING: LifecycleState
	// MISSING: SizeTb
	// MISSING: AvailableTillTime
	// MISSING: EndTime
	// MISSING: StartTime
	// MISSING: Type
	out.VaultID = direct.LazyPtr(in.GetVaultId())
	return out
}
func AutonomousDatabaseBackupProperties_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseBackupProperties) *pb.AutonomousDatabaseBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackupProperties{}
	// MISSING: Ocid
	out.RetentionPeriodDays = direct.ValueOf(in.RetentionPeriodDays)
	// MISSING: CompartmentID
	// MISSING: DatabaseSizeTb
	// MISSING: DbVersion
	// MISSING: IsLongTermBackup
	// MISSING: IsAutomaticBackup
	// MISSING: IsRestorable
	out.KeyStoreId = direct.ValueOf(in.KeyStoreID)
	out.KeyStoreWallet = direct.ValueOf(in.KeyStoreWallet)
	out.KmsKeyId = direct.ValueOf(in.KMSKeyID)
	out.KmsKeyVersionId = direct.ValueOf(in.KMSKeyVersionID)
	// MISSING: LifecycleDetails
	// MISSING: LifecycleState
	// MISSING: SizeTb
	// MISSING: AvailableTillTime
	// MISSING: EndTime
	// MISSING: StartTime
	// MISSING: Type
	out.VaultId = direct.ValueOf(in.VaultID)
	return out
}
func AutonomousDatabaseBackupPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackupProperties) *krm.AutonomousDatabaseBackupPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseBackupPropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: RetentionPeriodDays
	out.CompartmentID = direct.LazyPtr(in.GetCompartmentId())
	out.DatabaseSizeTb = direct.LazyPtr(in.GetDatabaseSizeTb())
	out.DbVersion = direct.LazyPtr(in.GetDbVersion())
	out.IsLongTermBackup = direct.LazyPtr(in.GetIsLongTermBackup())
	out.IsAutomaticBackup = direct.LazyPtr(in.GetIsAutomaticBackup())
	out.IsRestorable = direct.LazyPtr(in.GetIsRestorable())
	// MISSING: KeyStoreID
	// MISSING: KeyStoreWallet
	// MISSING: KMSKeyID
	// MISSING: KMSKeyVersionID
	out.LifecycleDetails = direct.LazyPtr(in.GetLifecycleDetails())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	out.SizeTb = direct.LazyPtr(in.GetSizeTb())
	out.AvailableTillTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAvailableTillTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: VaultID
	return out
}
func AutonomousDatabaseBackupPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseBackupPropertiesObservedState) *pb.AutonomousDatabaseBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackupProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: RetentionPeriodDays
	out.CompartmentId = direct.ValueOf(in.CompartmentID)
	out.DatabaseSizeTb = direct.ValueOf(in.DatabaseSizeTb)
	out.DbVersion = direct.ValueOf(in.DbVersion)
	out.IsLongTermBackup = direct.ValueOf(in.IsLongTermBackup)
	out.IsAutomaticBackup = direct.ValueOf(in.IsAutomaticBackup)
	out.IsRestorable = direct.ValueOf(in.IsRestorable)
	// MISSING: KeyStoreID
	// MISSING: KeyStoreWallet
	// MISSING: KMSKeyID
	// MISSING: KMSKeyVersionID
	out.LifecycleDetails = direct.ValueOf(in.LifecycleDetails)
	out.LifecycleState = direct.Enum_ToProto[pb.AutonomousDatabaseBackupProperties_State](mapCtx, in.LifecycleState)
	out.SizeTb = direct.ValueOf(in.SizeTb)
	out.AvailableTillTime = direct.StringTimestamp_ToProto(mapCtx, in.AvailableTillTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Type = direct.Enum_ToProto[pb.AutonomousDatabaseBackupProperties_Type](mapCtx, in.Type)
	// MISSING: VaultID
	return out
}
func OracledatabaseAutonomousDatabaseBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackup) *krm.OracledatabaseAutonomousDatabaseBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseBackupObservedState{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	// MISSING: Properties
	// MISSING: Labels
	return out
}
func OracledatabaseAutonomousDatabaseBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseBackupObservedState) *pb.AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackup{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	// MISSING: Properties
	// MISSING: Labels
	return out
}
func OracledatabaseAutonomousDatabaseBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseBackup) *krm.OracledatabaseAutonomousDatabaseBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseBackupSpec{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	// MISSING: Properties
	// MISSING: Labels
	return out
}
func OracledatabaseAutonomousDatabaseBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseBackupSpec) *pb.AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseBackup{}
	// MISSING: Name
	// MISSING: AutonomousDatabase
	// MISSING: DisplayName
	// MISSING: Properties
	// MISSING: Labels
	return out
}
