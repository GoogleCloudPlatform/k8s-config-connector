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
	alloydbv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDbConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbConnectionProfile) *krm.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbConnectionProfile{}
	if in.GetClusterId() != "" {
		out.ClusterRef = &alloydbv1beta1.ClusterRef{
			External: in.GetClusterId(),
		}
	}
	out.Settings = AlloyDbSettings_FromProto(mapCtx, in.GetSettings())
	return out
}

func AlloyDbConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbConnectionProfile) *pb.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbConnectionProfile{}
	if in.ClusterRef != nil {
		if in.ClusterRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.ClusterRef.Name)
		}
		out.ClusterId = in.ClusterRef.External
	}
	out.Settings = AlloyDbSettings_ToProto(mapCtx, in.Settings)
	return out
}

func CloudSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlConnectionProfile) *krm.CloudSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLConnectionProfile{}
	if in.GetCloudSqlId() != "" {
		out.InstanceRef = &refsv1beta1.SQLInstanceRef{
			External: in.GetCloudSqlId(),
		}
	}
	out.Settings = CloudSQLSettings_FromProto(mapCtx, in.GetSettings())
	return out
}

func CloudSQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLConnectionProfile) *pb.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlConnectionProfile{}
	if in.InstanceRef != nil {
		if in.InstanceRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.InstanceRef.Name)
		}
		out.CloudSqlId = in.InstanceRef.External
	}
	out.Settings = CloudSQLSettings_ToProto(mapCtx, in.Settings)
	return out
}

func MySQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.MySQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	if in.GetCloudSqlId() != "" {
		out.InstanceRef = &refsv1beta1.SQLInstanceRef{
			External: in.GetCloudSqlId(),
		}
	}
	return out
}

func MySQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.MySQLConnectionProfile) *pb.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.MySqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	if in.InstanceRef != nil {
		if in.InstanceRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.InstanceRef.Name)
		}
		out.CloudSqlId = in.InstanceRef.External
	}
	return out
}

func PostgreSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	if in.GetCloudSqlId() != "" {
		out.InstanceRef = &refsv1beta1.SQLInstanceRef{
			External: in.GetCloudSqlId(),
		}
	}
	out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	out.PrivateServiceConnectConnectivity = PrivateServiceConnectConnectivity_FromProto(mapCtx, in.GetPrivateServiceConnectConnectivity())
	return out
}

func PostgreSQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.PostgreSQLConnectionProfile) *pb.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgreSqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	if in.InstanceRef != nil {
		if in.InstanceRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.InstanceRef.Name)
		}
		out.CloudSqlId = in.InstanceRef.External
	}
	if oneof := StaticIPConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	if oneof := PrivateServiceConnectConnectivity_ToProto(mapCtx, in.PrivateServiceConnectConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_PrivateServiceConnectConnectivity{PrivateServiceConnectConnectivity: oneof}
	}
	return out
}

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
	out.Oracle = OracleConnectionProfileObservedState_FromProto(mapCtx, in.GetOracle())
	out.Cloudsql = CloudSQLConnectionProfileObservedState_FromProto(mapCtx, in.GetCloudsql())
	out.Alloydb = AlloyDbConnectionProfileObservedState_FromProto(mapCtx, in.GetAlloydb())
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
	if oneof := OracleConnectionProfileObservedState_ToProto(mapCtx, in.Oracle); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Oracle{Oracle: oneof}
	}
	if oneof := CloudSQLConnectionProfileObservedState_ToProto(mapCtx, in.Cloudsql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Cloudsql{Cloudsql: oneof}
	}
	if oneof := AlloyDbConnectionProfileObservedState_ToProto(mapCtx, in.Alloydb); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Alloydb{Alloydb: oneof}
	}
	out.Error = CloudDMSPrivateConnectionStatus_ToProto(mapCtx, in.Error)
	return out
}
