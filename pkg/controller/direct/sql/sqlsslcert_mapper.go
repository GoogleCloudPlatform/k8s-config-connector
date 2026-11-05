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

package sql

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
)

func SQLSSLCertSpec_FromProto(mapCtx *direct.MapContext, in *pb.SslCert) *krm.SQLSSLCertSpec {
	if in == nil {
		return nil
	}
	out := &krm.SQLSSLCertSpec{}
	out.CommonName = in.GetCommonName()
	return out
}

func SQLSSLCertSpec_ToProto(mapCtx *direct.MapContext, in *krm.SQLSSLCertSpec) *pb.SslCert {
	if in == nil {
		return nil
	}
	out := &pb.SslCert{}
	out.CommonName = in.CommonName
	return out
}

func SQLSSLCertStatus_FromProto(mapCtx *direct.MapContext, in *pb.SslCert) *krm.SQLSSLCertStatus {
	if in == nil {
		return nil
	}
	out := &krm.SQLSSLCertStatus{}
	out.CertSerialNumber = direct.LazyPtr(in.GetCertSerialNumber())
	out.Cert = direct.LazyPtr(in.GetCert())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpirationTime())
	out.Sha1Fingerprint = direct.LazyPtr(in.GetSha1Fingerprint())
	return out
}

func SQLSSLCertStatus_ToProto(mapCtx *direct.MapContext, in *krm.SQLSSLCertStatus) *pb.SslCert {
	if in == nil {
		return nil
	}
	out := &pb.SslCert{}
	out.CertSerialNumber = direct.ValueOf(in.CertSerialNumber)
	out.Cert = direct.ValueOf(in.Cert)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpirationTime)
	out.Sha1Fingerprint = direct.ValueOf(in.Sha1Fingerprint)
	return out
}
