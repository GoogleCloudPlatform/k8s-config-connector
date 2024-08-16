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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/alloydb/v1beta"
)

func ClusterSpecToApi(ctx *direct.MapContext, in *krm.AlloyDBClusterSpec) *api.Cluster {
	if in == nil {
		return nil
	}
	out := &api.Cluster{
		AutomatedBackupPolicy:  AutomatedBackupPolicy_KRMToApi(ctx, in.AutomatedBackupPolicy),
		ClusterType:            direct.ValueOf(in.ClusterType),
		ContinuousBackupConfig: ContinuousBackupConfig_KRMToApi(ctx, in.ContinuousBackupConfig),
		DisplayName:            direct.ValueOf(in.DisplayName),
		EncryptionConfig:       EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		InitialUser:            InitialUser_KRMToApi(ctx, in.InitialUser),
		NetworkConfig:          NetworkConfig_KRMToApi(ctx, in.NetworkConfig),
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}

	return out
}

func NetworkConfig_KRMToApi(ctx *direct.MapContext, in *krm.ClusterNetworkConfig) *api.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &api.NetworkConfig{
		AllocatedIpRange: direct.ValueOf(in.AllocatedIpRange),
		Network:          in.NetworkRef.External,
	}
	return out

}

func AutomatedBackupPolicy_KRMToApi(ctx *direct.MapContext, in *krm.ClusterAutomatedBackupPolicy) *api.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &api.AutomatedBackupPolicy{
		BackupWindow:           direct.ValueOf(in.BackupWindow),
		Enabled:                direct.ValueOf(in.Enabled),
		EncryptionConfig:       EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		Labels:                 in.Labels,
		Location:               direct.ValueOf(in.Location),
		QuantityBasedRetention: QuantityBasedRetention_KRMToApi(ctx, in.QuantityBasedRetention),
		TimeBasedRetention:     TimeBasedRetention_KRMToApi(ctx, in.TimeBasedRetention),
		WeeklySchedule:         WeeklySchedule_KRMToApi(ctx, in.WeeklySchedule),
	}
	return out
}

func WeeklySchedule_KRMToApi(ctx *direct.MapContext, in *krm.ClusterWeeklySchedule) *api.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &api.WeeklySchedule{
		DaysOfWeek: in.DaysOfWeek,
		StartTimes: StartTimes_KRMToApi(ctx, in.StartTimes),
	}

	return out
}

func StartTimes_KRMToApi(ctx *direct.MapContext, in []krm.ClusterStartTimes) []*api.GoogleTypeTimeOfDay {
	out := make([]*api.GoogleTypeTimeOfDay, len(in))
	for i, v := range in {
		out[i] = Time_KRMToApi(ctx, direct.LazyPtr(v))
	}
	return out
}

func Time_KRMToApi(ctx *direct.MapContext, in *krm.ClusterStartTimes) *api.GoogleTypeTimeOfDay {
	if in == nil {
		return nil
	}
	out := &api.GoogleTypeTimeOfDay{
		Hours:   int64(direct.ValueOf(in.Hours)),
		Minutes: int64(direct.ValueOf(in.Minutes)),
		Nanos:   int64(direct.ValueOf(in.Nanos)),
		Seconds: int64(direct.ValueOf(in.Seconds)),
	}
	return out
}

func EncryptionConfig_KRMToApi(ctx *direct.MapContext, in *krm.ClusterEncryptionConfig) *api.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &api.EncryptionConfig{
		KmsKeyName: in.KmsKeyNameRef.External,
	}
	return out
}

func TimeBasedRetention_KRMToApi(ctx *direct.MapContext, in *krm.ClusterTimeBasedRetention) *api.TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &api.TimeBasedRetention{
		RetentionPeriod: direct.ValueOf(in.RetentionPeriod),
	}

	return out
}

func QuantityBasedRetention_KRMToApi(ctx *direct.MapContext, in *krm.ClusterQuantityBasedRetention) *api.QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &api.QuantityBasedRetention{
		Count: int64(direct.ValueOf(in.Count)),
	}

	return out
}

func ContinuousBackupConfig_KRMToApi(ctx *direct.MapContext, in *krm.ClusterContinuousBackupConfig) *api.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &api.ContinuousBackupConfig{
		Enabled:            direct.ValueOf(in.Enabled),
		EncryptionConfig:   EncryptionConfig_KRMToApi(ctx, in.EncryptionConfig),
		RecoveryWindowDays: int64(direct.ValueOf(in.RecoveryWindowDays)),
	}

	return out
}

func InitialUser_KRMToApi(ctx *direct.MapContext, in *krm.ClusterInitialUser) *api.UserPassword {
	if in == nil {
		return nil
	}
	out := &api.UserPassword{
		User: direct.ValueOf(in.User),
	}
	if in.Password.Value != nil {
		out.Password = direct.ValueOf(in.Password.Value)
	} else {
		out.Password = in.Password.ValueFrom.SecretKeyRef.Key
	}
	return out
}

func ClusterSpecFromAPI(ctx *direct.MapContext, in *api.Cluster) *krm.AlloyDBClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterSpec{
		AutomatedBackupPolicy:  AutomatedBackupPolicy_APIToKRM(ctx, in.AutomatedBackupPolicy),
		ClusterType:            direct.LazyPtr(in.ClusterType),
		ContinuousBackupConfig: ContinuousBackupConfig_APIToKRM(ctx, in.ContinuousBackupConfig),
		DisplayName:            direct.LazyPtr(in.DisplayName),
		EncryptionConfig:       EncryptionConfig_APIToKRM(ctx, in.EncryptionConfig),
		InitialUser:            InitialUser_APIToKRM(ctx, in.InitialUser),
		NetworkConfig:          NetworkConfig_APIToKRM(ctx, in.NetworkConfig),
	}
	if in.NetworkConfig != nil {
		out.NetworkRef = &v1alpha1.ResourceRef{
			External: in.NetworkConfig.Network,
		}
	}

	return out
}

func AutomatedBackupPolicy_APIToKRM(ctx *direct.MapContext, in *api.AutomatedBackupPolicy) *krm.ClusterAutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ClusterAutomatedBackupPolicy{
		BackupWindow:           direct.LazyPtr(in.BackupWindow),
		Enabled:                &in.Enabled,
		EncryptionConfig:       EncryptionConfig_APIToKRM(ctx, in.EncryptionConfig),
		Labels:                 in.Labels,
		Location:               direct.LazyPtr(in.Location),
		QuantityBasedRetention: QuantityBasedRetention_APIToKRM(ctx, in.QuantityBasedRetention),
		TimeBasedRetention:     TimeBasedRetention_APIToKRM(ctx, in.TimeBasedRetention),
		WeeklySchedule:         WeeklySchedule_APIToKRM(ctx, in.WeeklySchedule),
	}
	return out
}

func EncryptionConfig_APIToKRM(ctx *direct.MapContext, in *api.EncryptionConfig) *krm.ClusterEncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterEncryptionConfig{
		KmsKeyNameRef: &v1alpha1.ResourceRef{
			External: in.KmsKeyName,
		},
	}
	return out
}

func QuantityBasedRetention_APIToKRM(ctx *direct.MapContext, in *api.QuantityBasedRetention) *krm.ClusterQuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.ClusterQuantityBasedRetention{
		Count: direct.LazyPtr(in.Count),
	}

	return out
}

func TimeBasedRetention_APIToKRM(ctx *direct.MapContext, in *api.TimeBasedRetention) *krm.ClusterTimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.ClusterTimeBasedRetention{
		RetentionPeriod: direct.LazyPtr(in.RetentionPeriod),
	}

	return out
}

func WeeklySchedule_APIToKRM(ctx *direct.MapContext, in *api.WeeklySchedule) *krm.ClusterWeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.ClusterWeeklySchedule{
		DaysOfWeek: in.DaysOfWeek,
		StartTimes: StartTimes_APIToKRM(ctx, in.StartTimes),
	}

	return out
}

func StartTimes_APIToKRM(ctx *direct.MapContext, in []*api.GoogleTypeTimeOfDay) []krm.ClusterStartTimes {
	out := make([]krm.ClusterStartTimes, len(in))
	for i, v := range in {
		out[i] = direct.ValueOf(Time_APIToKRM(ctx, v))
	}
	return out
}

func Time_APIToKRM(ctx *direct.MapContext, in *api.GoogleTypeTimeOfDay) *krm.ClusterStartTimes {
	if in == nil {
		return nil
	}
	out := &krm.ClusterStartTimes{
		Hours:   &in.Hours,
		Minutes: &in.Minutes,
		Nanos:   &in.Nanos,
		Seconds: &in.Seconds,
	}
	return out
}

func ContinuousBackupConfig_APIToKRM(ctx *direct.MapContext, in *api.ContinuousBackupConfig) *krm.ClusterContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterContinuousBackupConfig{
		Enabled:            &in.Enabled,
		EncryptionConfig:   EncryptionConfig_APIToKRM(ctx, in.EncryptionConfig),
		RecoveryWindowDays: direct.LazyPtr(in.RecoveryWindowDays),
	}

	return out
}

func InitialUser_APIToKRM(ctx *direct.MapContext, in *api.UserPassword) *krm.ClusterInitialUser {
	if in == nil {
		return nil
	}
	out := &krm.ClusterInitialUser{
		User: direct.LazyPtr(in.User),
		Password: krm.ClusterPassword{
			Value: direct.LazyPtr(in.Password),
		},
	}
	return out
}

func NetworkConfig_APIToKRM(ctx *direct.MapContext, in *api.NetworkConfig) *krm.ClusterNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterNetworkConfig{
		AllocatedIpRange: direct.LazyPtr(in.AllocatedIpRange),
		NetworkRef: &v1alpha1.ResourceRef{
			External: in.Network,
		},
	}
	return out

}

func ClusterStatusFromApi(ctx *direct.MapContext, in *api.Cluster) *krm.AlloyDBClusterStatus {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterStatus{
		BackupSource:         BackupSource_KRMFromApi(ctx, in.BackupSource),
		ContinuousBackupInfo: ContinuousBackupInfo_KRMFromApi(ctx, in.ContinuousBackupInfo),
		DatabaseVersion:      direct.LazyPtr(in.DatabaseVersion),
		EncryptionInfo:       EncryptionInfo_KRMFromApi(ctx, in.EncryptionInfo),
		MigrationSource:      MigrationSource_KRMFromApi(ctx, in.MigrationSource),
		Name:                 direct.LazyPtr(in.Name),
		Uid:                  direct.LazyPtr(in.Uid),
	}
	return out
}

func BackupSource_KRMFromApi(ctx *direct.MapContext, in *api.BackupSource) []krm.ClusterBackupSourceStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterBackupSourceStatus{
		{
			BackupName: direct.LazyPtr(in.BackupName),
		},
	}
	return out
}

func ContinuousBackupInfo_KRMFromApi(ctx *direct.MapContext, in *api.ContinuousBackupInfo) []krm.ClusterContinuousBackupInfoStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterContinuousBackupInfoStatus{
		{
			EarliestRestorableTime: direct.LazyPtr(in.EarliestRestorableTime),
			EnabledTime:            direct.LazyPtr(in.EnabledTime),
			Schedule:               in.Schedule,
		},
	}
	return out
}

func EncryptionInfo_KRMFromApi(ctx *direct.MapContext, in *api.EncryptionInfo) []krm.ClusterEncryptionInfoStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterEncryptionInfoStatus{
		{
			EncryptionType: direct.LazyPtr(in.EncryptionType),
			KmsKeyVersions: in.KmsKeyVersions,
		},
	}
	return out
}

func MigrationSource_KRMFromApi(ctx *direct.MapContext, in *api.MigrationSource) []krm.ClusterMigrationSourceStatus {
	if in == nil {
		return nil
	}
	out := []krm.ClusterMigrationSourceStatus{
		{
			HostPort:    direct.LazyPtr(in.HostPort),
			ReferenceId: direct.LazyPtr(in.ReferenceId),
			SourceType:  direct.LazyPtr(in.SourceType),
		},
	}
	return out
}
