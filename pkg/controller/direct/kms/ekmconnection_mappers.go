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

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSEKMConnectionCertificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) krm.KMSEKMConnectionCertificate {
	if in == nil {
		return krm.KMSEKMConnectionCertificate{}
	}
	out := krm.KMSEKMConnectionCertificate{}
	out.RawDER = in.GetRawDer()
	return out
}
func KMSEKMConnectionCertificate_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionCertificate) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.RawDer = in.RawDER
	return out
}
func KMSEKMConnectionObservedCertificate_FromProto(mapCtx *direct.MapContext, in *pb.Certificate) krm.KMSEKMConnectionObservedCertificate {
	if in == nil {
		return krm.KMSEKMConnectionObservedCertificate{}
	}
	out := krm.KMSEKMConnectionObservedCertificate{}
	out.Parsed = direct.LazyPtr(in.GetParsed())
	out.Issuer = direct.LazyPtr(in.GetIssuer())
	out.Subject = direct.LazyPtr(in.GetSubject())
	out.SubjectAlternativeDNSNames = in.GetSubjectAlternativeDnsNames()
	out.NotBeforeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotBeforeTime())
	out.NotAfterTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNotAfterTime())
	out.SerialNumber = direct.LazyPtr(in.GetSerialNumber())
	out.Sha256Fingerprint = direct.LazyPtr(in.GetSha256Fingerprint())
	return out
}
func KMSEKMConnectionObservedCertificate_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionObservedCertificate) *pb.Certificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate{}
	out.Parsed = direct.ValueOf(in.Parsed)
	out.Issuer = direct.ValueOf(in.Issuer)
	out.Subject = direct.ValueOf(in.Subject)
	out.SubjectAlternativeDnsNames = in.SubjectAlternativeDNSNames
	out.NotBeforeTime = direct.StringTimestamp_ToProto(mapCtx, in.NotBeforeTime)
	out.NotAfterTime = direct.StringTimestamp_ToProto(mapCtx, in.NotAfterTime)
	out.SerialNumber = direct.ValueOf(in.SerialNumber)
	out.Sha256Fingerprint = direct.ValueOf(in.Sha256Fingerprint)
	return out
}
func KMSEKMConnectionObservedServiceResolver_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection_ServiceResolver) krm.KMSEKMConnectionObservedServiceResolver {
	if in == nil {
		return krm.KMSEKMConnectionObservedServiceResolver{}
	}
	out := krm.KMSEKMConnectionObservedServiceResolver{}
	out.ServerCertificates = make([]krm.KMSEKMConnectionObservedCertificate, len(in.ServerCertificates))
	for i, cert := range in.ServerCertificates {
		out.ServerCertificates[i] = KMSEKMConnectionObservedCertificate_FromProto(mapCtx, cert)
	}
	return out
}
func KMSEKMConnectionObservedServiceResolver_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionObservedServiceResolver) *pb.EkmConnection_ServiceResolver {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection_ServiceResolver{}
	out.ServerCertificates = make([]*pb.Certificate, len(in.ServerCertificates))
	for i, cert := range in.ServerCertificates {
		out.ServerCertificates[i] = KMSEKMConnectionObservedCertificate_ToProto(mapCtx, &cert)
	}
	return out
}
func KMSEKMConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.KMSEKMConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSEKMConnectionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ServiceResolvers = make([]krm.KMSEKMConnectionObservedServiceResolver, len(in.ServiceResolvers))
	for i, sr := range in.ServiceResolvers {
		out.ServiceResolvers[i] = KMSEKMConnectionObservedServiceResolver_FromProto(mapCtx, sr)
	}
	return out
}
func KMSEKMConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionObservedState) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.ServiceResolvers = make([]*pb.EkmConnection_ServiceResolver, len(in.ServiceResolvers))
	for i, sr := range in.ServiceResolvers {
		out.ServiceResolvers[i] = KMSEKMConnectionObservedServiceResolver_ToProto(mapCtx, &sr)
	}
	return out
}
func KMSEKMConnectionServiceResolver_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection_ServiceResolver) krm.KMSEKMConnectionServiceResolver {
	if in == nil {
		return krm.KMSEKMConnectionServiceResolver{}
	}
	out := krm.KMSEKMConnectionServiceResolver{}
	if in.GetServiceDirectoryService() != "" {
		out.ServiceDirectoryServiceRef = &krm.ServiceDirectoryServiceRef{
			External: in.GetServiceDirectoryService(),
		}
	}
	out.EndpointFilter = direct.LazyPtr(in.GetEndpointFilter())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.ServerCertificates = make([]krm.KMSEKMConnectionCertificate, len(in.ServerCertificates))
	for i, cert := range in.ServerCertificates {
		out.ServerCertificates[i] = KMSEKMConnectionCertificate_FromProto(mapCtx, cert)
	}
	return out
}
func KMSEKMConnectionServiceResolver_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionServiceResolver) *pb.EkmConnection_ServiceResolver {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection_ServiceResolver{}
	if in.ServiceDirectoryServiceRef != nil {
		out.ServiceDirectoryService = in.ServiceDirectoryServiceRef.External
	}
	out.EndpointFilter = direct.ValueOf(in.EndpointFilter)
	out.Hostname = direct.ValueOf(in.Hostname)
	out.ServerCertificates = make([]*pb.Certificate, len(in.ServerCertificates))
	for i, cert := range in.ServerCertificates {
		out.ServerCertificates[i] = KMSEKMConnectionCertificate_ToProto(mapCtx, &cert)
	}
	return out
}
func KMSEKMConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.KMSEKMConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSEKMConnectionSpec{}
	out.ServiceResolvers = make([]krm.KMSEKMConnectionServiceResolver, len(in.ServiceResolvers))
	for i, sr := range in.ServiceResolvers {
		out.ServiceResolvers[i] = KMSEKMConnectionServiceResolver_FromProto(mapCtx, sr)
	}
	out.KeyManagementMode = direct.Enum_FromProto(mapCtx, in.GetKeyManagementMode())
	out.CryptoSpacePath = direct.LazyPtr(in.GetCryptoSpacePath())
	return out
}
func KMSEKMConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSEKMConnectionSpec) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	out.ServiceResolvers = make([]*pb.EkmConnection_ServiceResolver, len(in.ServiceResolvers))
	for i, sr := range in.ServiceResolvers {
		out.ServiceResolvers[i] = KMSEKMConnectionServiceResolver_ToProto(mapCtx, &sr)
	}
	out.KeyManagementMode = direct.Enum_ToProto[pb.EkmConnection_KeyManagementMode](mapCtx, in.KeyManagementMode)
	out.CryptoSpacePath = direct.ValueOf(in.CryptoSpacePath)
	return out
}
