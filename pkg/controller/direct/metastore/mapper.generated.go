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

// +generated:mapper
// krm.group: metastore.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.metastore.v1

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MetastoreBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MetastoreBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreBackupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceRevision = MetastoreServiceSpec_FromProto(mapCtx, in.GetServiceRevision())
	out.RestoringServices = in.RestoringServices
	return out
}
func MetastoreBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceRevision = MetastoreServiceSpec_ToProto(mapCtx, in.ServiceRevision)
	out.RestoringServices = in.RestoringServices
	return out
}
func MetastoreBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MetastoreBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreBackupSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func MetastoreBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	return out
}
