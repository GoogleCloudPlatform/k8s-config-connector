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

// +generated:mapper
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDMSConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.CloudDMSConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConnectionProfileObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Mysql = MySQLConnectionProfileObservedState_FromProto(mapCtx, in.GetMysql())
	out.Postgresql = PostgreSQLConnectionProfileObservedState_FromProto(mapCtx, in.GetPostgresql())
	out.Cloudsql = CloudSQLConnectionProfileObservedState_FromProto(mapCtx, in.GetCloudsql())
	out.Error = CloudDMSPrivateConnectionStatus_FromProto(mapCtx, in.GetError())
	return out
}

func CloudDMSConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.ConnectionProfile_State](mapCtx, in.State)
	if oneof := MySQLConnectionProfileObservedState_ToProto(mapCtx, in.Mysql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Mysql{Mysql: oneof}
	}
	if oneof := PostgreSQLConnectionProfileObservedState_ToProto(mapCtx, in.Postgresql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Postgresql{Postgresql: oneof}
	}
	if oneof := CloudSQLConnectionProfileObservedState_ToProto(mapCtx, in.Cloudsql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Cloudsql{Cloudsql: oneof}
	}
	out.Error = CloudDMSPrivateConnectionStatus_ToProto(mapCtx, in.Error)
	return out
}
