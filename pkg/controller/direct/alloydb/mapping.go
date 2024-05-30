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

package alloydb

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/alloydb/v1beta1"
	api "google.golang.org/api/alloydb/v1beta"
)

func ClusterSpecToApi(ctx *MapContext, in *krm.AlloyDBClusterSpec) *api.Cluster {
	if in == nil {
		return nil
	}
	out := &api.Cluster{
		AutomatedBackupPolicy:  AutomatedBackupPolicy_KRMToApi(ctx, in.AutomatedBackupPolicy),
		ClusterType:            ValueOf(in.ClusterType),
		ContinuousBackupConfig: ContinuousBackupConfig_KRMToApi(ctx, in.ContinuousBackupConfig),
		DisplayName:            ValueOf(in.DisplayName),
		EncryptionConfig:       EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		InitialUser:            InitialUser_KRMToApi(ctx, in.InitialUser),
		NetworkConfig:          NetworkConfig_KRMToApi(ctx, in.NetworkConfig),
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}

	return out
}

func NetworkConfig_KRMToApi(ctx *MapContext, in *krm.ClusterNetworkConfig) *api.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &api.NetworkConfig{
		AllocatedIpRange: ValueOf(in.AllocatedIpRange),
		Network:          in.NetworkRef.External,
	}
	return out

}

func AutomatedBackupPolicy_KRMToApi(ctx *MapContext, in *krm.ClusterAutomatedBackupPolicy) *api.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &api.AutomatedBackupPolicy{
		BackupWindow:           ValueOf(in.BackupWindow),
		Enabled:                ValueOf(in.Enabled),
		EncryptionConfig:       EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		Labels:                 in.Labels,
		Location:               ValueOf(in.Location),
		QuantityBasedRetention: QuantityBasedRetention_KRMToApi(ctx, in.QuantityBasedRetention),
		TimeBasedRetention:     TimeBasedRetention_KRMToApi(ctx, in.TimeBasedRetention),
		WeeklySchedule:         WeeklySchedule_KRMToApi(ctx, in.WeeklySchedule),
	}
	return out
}

func WeeklySchedule_KRMToApi(ctx *MapContext, in *krm.ClusterWeeklySchedule) *api.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &api.WeeklySchedule{
		DaysOfWeek: in.DaysOfWeek,
		StartTimes: StartTimes_KRMToApi(ctx, in.StartTimes),
	}

	return out
}

func StartTimes_KRMToApi(ctx *MapContext, in []krm.ClusterStartTimes) []*api.GoogleTypeTimeOfDay {
	out := make([]*api.GoogleTypeTimeOfDay, len(in))
	for i, v := range in {
		out[i] = Time_KRMToApi(ctx, LazyPtr(v))
	}
	return out
}

func Time_KRMToApi(ctx *MapContext, in *krm.ClusterStartTimes) *api.GoogleTypeTimeOfDay {
	if in == nil {
		return nil
	}
	out := &api.GoogleTypeTimeOfDay{
		Hours:   int64(ValueOf(in.Hours)),
		Minutes: int64(ValueOf(in.Minutes)),
		Nanos:   int64(ValueOf(in.Nanos)),
		Seconds: int64(ValueOf(in.Seconds)),
	}
	return out
}

func EncryptionConfig_KRMToApi(ctx *MapContext, in *krm.ClusterEncryptionConfig) *api.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &api.EncryptionConfig{
		KmsKeyName: in.KmsKeyNameRef.External,
	}
	return out
}

func TimeBasedRetention_KRMToApi(ctx *MapContext, in *krm.ClusterTimeBasedRetention) *api.TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &api.TimeBasedRetention{
		RetentionPeriod: ValueOf(in.RetentionPeriod),
	}

	return out
}

func QuantityBasedRetention_KRMToApi(ctx *MapContext, in *krm.ClusterQuantityBasedRetention) *api.QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &api.QuantityBasedRetention{
		Count: int64(ValueOf(in.Count)),
	}

	return out
}

func ContinuousBackupConfig_KRMToApi(ctx *MapContext, in *krm.ClusterContinuousBackupConfig) *api.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &api.ContinuousBackupConfig{
		Enabled:            ValueOf(in.Enabled),
		EncryptionConfig:   EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		RecoveryWindowDays: int64(ValueOf(in.RecoveryWindowDays)),
	}

	return out
}

func InitialUser_KRMToApi(ctx *MapContext, in *krm.ClusterInitialUser) *api.UserPassword {
	if in == nil {
		return nil
	}
	out := &api.UserPassword{
		User: ValueOf(in.User),
	}
	if in.Password.Value != nil {
		out.Password = ValueOf(in.Password.Value)
	} else {
		out.Password = in.Password.ValueFrom.SecretKeyRef.Key
	}
	return out
}

func ClusterStatusFromApi(ctx *MapContext, in *api.Cluster) *krm.AlloyDBClusterStatus {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterStatus{
		BackupSource:         BackupSource_KRMFromApi(ctx, in.BackupSource),
		ContinuousBackupInfo: ContinuousBackupInfo_KRMFromApi(ctx, in.ContinuousBackupInfo),
		DatabaseVersion:      LazyPtr(in.DatabaseVersion),
		EncryptionInfo:       EncryptionInfo_KRMFromApi(ctx, in.EncryptionInfo),
		MigrationSource:      MigrationSource_KRMFromApi(ctx, in.MigrationSource),
		Name:                 LazyPtr(in.Name),
		Uid:                  LazyPtr(in.Uid),
	}
	return out
}

func BackupSource_KRMFromApi(ctx *MapContext, in *api.BackupSource) []krm.ClusterBackupSourceStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterBackupSourceStatus{
		{
			BackupName: LazyPtr(in.BackupName),
		},
	}
	return out
}

func ContinuousBackupInfo_KRMFromApi(ctx *MapContext, in *api.ContinuousBackupInfo) []krm.ClusterContinuousBackupInfoStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterContinuousBackupInfoStatus{
		{
			EarliestRestorableTime: LazyPtr(in.EarliestRestorableTime),
			EnabledTime:            LazyPtr(in.EnabledTime),
			Schedule:               in.Schedule,
		},
	}
	return out
}

func EncryptionInfo_KRMFromApi(ctx *MapContext, in *api.EncryptionInfo) []krm.ClusterEncryptionInfoStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterEncryptionInfoStatus{
		{
			EncryptionType: LazyPtr(in.EncryptionType),
			KmsKeyVersions: in.KmsKeyVersions,
		},
	}
	return out
}

func MigrationSource_KRMFromApi(ctx *MapContext, in *api.MigrationSource) []krm.ClusterMigrationSourceStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterMigrationSourceStatus{
		{
			HostPort:    LazyPtr(in.HostPort),
			ReferenceId: LazyPtr(in.ReferenceId),
			SourceType:  LazyPtr(in.SourceType),
		},
	}
	return out
}
