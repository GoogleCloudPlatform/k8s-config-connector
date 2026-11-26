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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krmbackupdrv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupSelectionConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.BackupSelectionConfig) *pb.BackupSelectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackupSelectionConfig{}
	if in.DataSourceRef != nil {
		out.Source = &pb.BackupSelectionConfig_DataSource{
			DataSource: in.DataSourceRef.External,
		}
	}
	return out
}

func PreRestoreConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.PreRestoreConfig) *pb.PreRestoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.PreRestoreConfig{}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	if in.CloudRunJobRef != nil {
		out.PreRestoreAction = &pb.PreRestoreConfig_CloudRunJob{
			CloudRunJob: in.CloudRunJobRef.External,
		}
	}
	return out
}

func PostRestoreConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.PostRestoreConfig) *pb.PostRestoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.PostRestoreConfig{}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	if in.CloudRunJobRef != nil {
		out.PostRestoreAction = &pb.PostRestoreConfig_CloudRunJob{
			CloudRunJob: in.CloudRunJobRef.External,
		}
	}
	return out
}

func VerificationConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.VerificationConfig) *pb.VerificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.VerificationConfig{}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	if in.CloudRunJobRef != nil {
		out.VerificationAction = &pb.VerificationConfig_CloudRunJob{
			CloudRunJob: in.CloudRunJobRef.External,
		}
	}
	return out
}
