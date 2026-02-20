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

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Secret = SecretVolumeSource_FromProto(mapCtx, in.GetSecret())
	out.CloudSQLInstance = CloudSQLInstance_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.EmptyDir = EmptyDirVolumeSource_FromProto(mapCtx, in.GetEmptyDir())
	out.Nfs = NfsVolumeSource_FromProto(mapCtx, in.GetNfs())
	out.GCS = GCSVolumeSource_FromProto(mapCtx, in.GetGcs())
	return out
}

func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := SecretVolumeSource_ToProto(mapCtx, in.Secret); oneof != nil {
		out.VolumeType = &pb.Volume_Secret{Secret: oneof}
	}
	if oneof := CloudSQLInstance_ToProto(mapCtx, in.CloudSQLInstance); oneof != nil {
		out.VolumeType = &pb.Volume_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := EmptyDirVolumeSource_ToProto(mapCtx, in.EmptyDir); oneof != nil {
		out.VolumeType = &pb.Volume_EmptyDir{EmptyDir: oneof}
	}
	if oneof := NfsVolumeSource_ToProto(mapCtx, in.Nfs); oneof != nil {
		out.VolumeType = &pb.Volume_Nfs{Nfs: oneof}
	}
	if oneof := GCSVolumeSource_ToProto(mapCtx, in.GCS); oneof != nil {
		out.VolumeType = &pb.Volume_Gcs{Gcs: oneof}
	}
	return out
}
