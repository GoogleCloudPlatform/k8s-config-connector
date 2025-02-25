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

package spannerbackupschedules

import (
	pb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spannerbackupschedules/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CreateBackupEncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.CreateBackupEncryptionConfig) *krm.CreateBackupEncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.CreateBackupEncryptionConfig{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	kmsKeyNameRefs := []*refs.KMSCryptoKeyRef{}
	if in.GetKmsKeyNames() != nil {
		for _, kmsKeyName := range in.GetKmsKeyNames() {
			kmsKeyNameRefs = append(kmsKeyNameRefs, &refs.KMSCryptoKeyRef{External: kmsKeyName})
		}
	}
	out.KMSKeyNameRefs = kmsKeyNameRefs
	return out
}
func CreateBackupEncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.CreateBackupEncryptionConfig) *pb.CreateBackupEncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.CreateBackupEncryptionConfig{}
	out.EncryptionType = direct.Enum_ToProto[pb.CreateBackupEncryptionConfig_EncryptionType](mapCtx, in.EncryptionType)
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
	kmsKeyNames := []string{}
	if in.KMSKeyNameRefs != nil {
		for _, kmsKeyNameRef := range in.KMSKeyNameRefs {
			if kmsKeyNameRef != nil {
				kmsKeyNames = append(kmsKeyNames, kmsKeyNameRef.External)
			}
		}
	}
	out.KmsKeyNames = kmsKeyNames
	return out
}
func CrontabSpec_FromProto(mapCtx *direct.MapContext, in *pb.CrontabSpec) *krm.CrontabSpec {
	if in == nil {
		return nil
	}
	out := &krm.CrontabSpec{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func CrontabSpec_ToProto(mapCtx *direct.MapContext, in *krm.CrontabSpec) *pb.CrontabSpec {
	if in == nil {
		return nil
	}
	out := &pb.CrontabSpec{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func CrontabSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrontabSpec) *krm.CrontabSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CrontabSpecObservedState{}
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.CreationWindow = direct.StringDuration_FromProto(mapCtx, in.GetCreationWindow())
	return out
}
func CrontabSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CrontabSpecObservedState) *pb.CrontabSpec {
	if in == nil {
		return nil
	}
	out := &pb.CrontabSpec{}
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.CreationWindow = direct.StringDuration_ToProto(mapCtx, in.CreationWindow)
	return out
}
func SpannerBackupScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.SpannerBackupScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerBackupScheduleObservedState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CronSpec = CrontabSpecObservedState_FromProto(mapCtx, in.Spec.GetCronSpec())
	return out
}
func SpannerBackupScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerBackupScheduleObservedState) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if in.CronSpec != nil {
		out.Spec = &pb.BackupScheduleSpec{
			ScheduleSpec: &pb.BackupScheduleSpec_CronSpec{
				CronSpec: CrontabSpecObservedState_ToProto(mapCtx, in.CronSpec),
			},
		}
	}
	return out
}
func SpannerBackupScheduleSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krm.SpannerBackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpannerBackupScheduleSpec{}
	out.Spec = BackupScheduleSpec_FromProto(mapCtx, in.GetSpec())
	out.RetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetRetentionDuration())
	out.EncryptionConfig = CreateBackupEncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.FullBackupSpec = FullBackupSpec_FromProto(mapCtx, in.GetFullBackupSpec())
	out.IncrementalBackupSpec = IncrementalBackupSpec_FromProto(mapCtx, in.GetIncrementalBackupSpec())
	return out
}
func SpannerBackupScheduleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpannerBackupScheduleSpec) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	out.Spec = BackupScheduleSpec_ToProto(mapCtx, in.Spec)
	out.RetentionDuration = direct.StringDuration_ToProto(mapCtx, in.RetentionDuration)
	out.EncryptionConfig = CreateBackupEncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	if oneof := FullBackupSpec_ToProto(mapCtx, in.FullBackupSpec); oneof != nil {
		out.BackupTypeSpec = &pb.BackupSchedule_FullBackupSpec{FullBackupSpec: oneof}
	}
	if oneof := IncrementalBackupSpec_ToProto(mapCtx, in.IncrementalBackupSpec); oneof != nil {
		out.BackupTypeSpec = &pb.BackupSchedule_IncrementalBackupSpec{IncrementalBackupSpec: oneof}
	}
	return out
}
