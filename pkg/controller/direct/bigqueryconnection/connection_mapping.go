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

package bigqueryconnection

import (
	pb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudResourcePropertiesSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudResourcePropertiesSpec) *pb.CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudResourceProperties{}
	return out
}

func CloudResourcePropertiesSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudResourceProperties) *krm.CloudResourcePropertiesSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudResourcePropertiesSpec{}
	return out
}

func CloudSqlPropertiesSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlPropertiesSpec) *pb.CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlProperties{}
	// out.InstanceId = direct.ValueOf(in.InstanceID)
	out.Database = direct.ValueOf(in.Database)
	out.Type = direct.Enum_ToProto[pb.CloudSqlProperties_DatabaseType](mapCtx, in.Type)
	out.Credential = CloudSqlCredential_ToProto(mapCtx, in.Credential)
	if in.InstanceRef != nil {
		if in.InstanceRef.External == "" {
			mapCtx.Errorf("SQLInstance external reference was not pre-resolved")
		}
		out.InstanceId = in.InstanceRef.External
	}
	return out
}

func CloudSqlPropertiesSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlProperties) *krm.CloudSqlPropertiesSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlPropertiesSpec{}
	// out.InstanceID = direct.LazyPtr(in.InstanceId)
	out.Database = direct.LazyPtr(in.Database)
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Credential = CloudSqlCredential_FromProto(mapCtx, in.GetCredential())
	return out
}

func BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.BigQueryConnectionConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConnectionConnectionObservedState{}
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	if oneof := CloudResourcePropertiesStatus_FromProto(mapCtx, in.GetCloudResource()); oneof != nil {
		out.CloudResource = oneof
	}
	if oneof := CloudSqlPropertiesStatus_FromProto(mapCtx, in.GetCloudSql()); oneof != nil {
		out.CloudSql = oneof
	}
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}

func BigQueryConnectionConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConnectionConnectionSpec) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	if oneof := CloudResourcePropertiesSpec_ToProto(mapCtx, in.CloudResourceSpec); oneof != nil {
		out.Properties = &pb.Connection_CloudResource{}
	}
	if oneof := CloudSqlPropertiesSpec_ToProto(mapCtx, in.CloudSqlSpec); oneof != nil {
		out.Properties = &pb.Connection_CloudSql{CloudSql: oneof}
	}

	// MISSING: Aws
	// MISSING: Azure
	// MISSING: CloudSpanner
	// MISSING: Spark
	// MISSING: SalesforceDataCloud
	return out
}
