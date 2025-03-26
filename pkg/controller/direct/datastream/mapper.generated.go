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
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryProfile_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryProfile) *krm.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryProfile{}
	return out
}
func BigQueryProfile_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryProfile) *pb.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryProfile{}
	return out
}
func DatastreamPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.DatastreamPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamPrivateConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Error_FromProto(mapCtx, in.GetError())
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = Error_ToProto(mapCtx, in.Error)
	// MISSING: VpcPeeringConfig
	return out
}
func GCSProfile_FromProto(mapCtx *direct.MapContext, in *pb.GcsProfile) *krm.GCSProfile {
	if in == nil {
		return nil
	}
	out := &krm.GCSProfile{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.RootPath = direct.LazyPtr(in.GetRootPath())
	return out
}
func GCSProfile_ToProto(mapCtx *direct.MapContext, in *krm.GCSProfile) *pb.GcsProfile {
	if in == nil {
		return nil
	}
	out := &pb.GcsProfile{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.RootPath = direct.ValueOf(in.RootPath)
	return out
}
func MysqlProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlProfile) *krm.MysqlProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MysqlProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SSLConfig = MysqlSSLConfigObservedState_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func MysqlProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MysqlProfileObservedState) *pb.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.MysqlProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SslConfig = MysqlSSLConfigObservedState_ToProto(mapCtx, in.SSLConfig)
	return out
}
func MysqlSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSSLConfig{}
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	// MISSING: ClientCertificateSet
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSSLConfig) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	out.ClientKey = direct.ValueOf(in.ClientKey)
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krm.MysqlSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSSLConfigObservedState{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.LazyPtr(in.GetClientKeySet())
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.LazyPtr(in.GetClientCertificateSet())
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func MysqlSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSSLConfigObservedState) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.ValueOf(in.ClientKeySet)
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.ValueOf(in.ClientCertificateSet)
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	return out
}
func OracleProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleProfile) *krm.OracleProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSSLConfig = OracleSSLConfigObservedState_FromProto(mapCtx, in.GetOracleSslConfig())
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleProfileObservedState) *pb.OracleProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSslConfig = OracleSSLConfigObservedState_ToProto(mapCtx, in.OracleSSLConfig)
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleSSLConfig{}
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	return out
}
func OracleSSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleSSLConfig) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	return out
}
func OracleSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krm.OracleSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleSSLConfigObservedState{}
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func OracleSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleSSLConfigObservedState) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	return out
}
func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	if in.PrivateConnectionRef != nil {
		out.PrivateConnection = in.PrivateConnectionRef.External
	}
	return out
}
func StaticServiceIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krm.StaticServiceIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticServiceIPConnectivity{}
	return out
}
func StaticServiceIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticServiceIPConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
