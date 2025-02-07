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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/connection/apiv1beta1/connectionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudSqlCredential_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlCredential) *krm.CloudSqlCredential {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlCredential{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func CloudSqlCredential_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlCredential) *pb.CloudSqlCredential {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlCredential{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	return out
}
func CloudSqlProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlProperties) *krm.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlProperties{}
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Credential = CloudSqlCredential_FromProto(mapCtx, in.GetCredential())
	// MISSING: ServiceAccountID
	return out
}
func CloudSqlProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlProperties) *pb.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlProperties{}
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.Database = direct.ValueOf(in.Database)
	out.Type = direct.Enum_ToProto[pb.CloudSqlProperties_DatabaseType](mapCtx, in.Type)
	out.Credential = CloudSqlCredential_ToProto(mapCtx, in.Credential)
	// MISSING: ServiceAccountID
	return out
}
func CloudSqlPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlProperties) *krm.CloudSqlPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlPropertiesObservedState{}
	// MISSING: InstanceID
	// MISSING: Database
	// MISSING: Type
	// MISSING: Credential
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CloudSqlPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlPropertiesObservedState) *pb.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlProperties{}
	// MISSING: InstanceID
	// MISSING: Database
	// MISSING: Type
	// MISSING: Credential
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudSql = CloudSqlProperties_FromProto(mapCtx, in.GetCloudSql())
	// MISSING: CreationTime
	// MISSING: LastModifiedTime
	// MISSING: HasCredential
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.Name = direct.ValueOf(in.Name)
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	if oneof := CloudSqlProperties_ToProto(mapCtx, in.CloudSql); oneof != nil {
		out.Properties = &pb.Connection_CloudSql{CloudSql: oneof}
	}
	// MISSING: CreationTime
	// MISSING: LastModifiedTime
	// MISSING: HasCredential
	return out
}
func ConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionObservedState{}
	// MISSING: Name
	// MISSING: FriendlyName
	// MISSING: Description
	out.CloudSql = CloudSqlPropertiesObservedState_FromProto(mapCtx, in.GetCloudSql())
	out.CreationTime = direct.LazyPtr(in.GetCreationTime())
	out.LastModifiedTime = direct.LazyPtr(in.GetLastModifiedTime())
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}
func ConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: FriendlyName
	// MISSING: Description
	if oneof := CloudSqlPropertiesObservedState_ToProto(mapCtx, in.CloudSql); oneof != nil {
		out.Properties = &pb.Connection_CloudSql{CloudSql: oneof}
	}
	out.CreationTime = direct.ValueOf(in.CreationTime)
	out.LastModifiedTime = direct.ValueOf(in.LastModifiedTime)
	out.HasCredential = direct.ValueOf(in.HasCredential)
	return out
}
