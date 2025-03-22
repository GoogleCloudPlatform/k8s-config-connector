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
// krm.group: kms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.kms.v1

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ImportJob_WrappingPublicKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob_WrappingPublicKey) *krm.ImportJob_WrappingPublicKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImportJob_WrappingPublicKeyObservedState{}
	out.Pem = direct.LazyPtr(in.GetPem())
	return out
}
func ImportJob_WrappingPublicKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImportJob_WrappingPublicKeyObservedState) *pb.ImportJob_WrappingPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob_WrappingPublicKey{}
	out.Pem = direct.ValueOf(in.Pem)
	return out
}
func KMSImportJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSImportJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSImportJobObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.GenerateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetGenerateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.ExpireEventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireEventTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PublicKey = ImportJob_WrappingPublicKeyObservedState_FromProto(mapCtx, in.GetPublicKey())
	out.Attestation = KeyOperationAttestationObservedState_FromProto(mapCtx, in.GetAttestation())
	return out
}
func KMSImportJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobObservedState) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.GenerateTime = direct.StringTimestamp_ToProto(mapCtx, in.GenerateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.ExpireEventTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireEventTime)
	out.State = direct.Enum_ToProto[pb.ImportJob_ImportJobState](mapCtx, in.State)
	out.PublicKey = ImportJob_WrappingPublicKeyObservedState_ToProto(mapCtx, in.PublicKey)
	out.Attestation = KeyOperationAttestationObservedState_ToProto(mapCtx, in.Attestation)
	return out
}
func KeyOperationAttestationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krm.KeyOperationAttestationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KeyOperationAttestationObservedState{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Content = in.GetContent()
	out.CertChains = KeyOperationAttestation_CertificateChains_FromProto(mapCtx, in.GetCertChains())
	return out
}
func KeyOperationAttestationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KeyOperationAttestationObservedState) *pb.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation{}
	out.Format = direct.Enum_ToProto[pb.KeyOperationAttestation_AttestationFormat](mapCtx, in.Format)
	out.Content = in.Content
	out.CertChains = KeyOperationAttestation_CertificateChains_ToProto(mapCtx, in.CertChains)
	return out
}
func KeyOperationAttestation_CertificateChains_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation_CertificateChains) *krm.KeyOperationAttestation_CertificateChains {
	if in == nil {
		return nil
	}
	out := &krm.KeyOperationAttestation_CertificateChains{}
	out.CaviumCerts = in.CaviumCerts
	out.GoogleCardCerts = in.GoogleCardCerts
	out.GooglePartitionCerts = in.GooglePartitionCerts
	return out
}
func KeyOperationAttestation_CertificateChains_ToProto(mapCtx *direct.MapContext, in *krm.KeyOperationAttestation_CertificateChains) *pb.KeyOperationAttestation_CertificateChains {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation_CertificateChains{}
	out.CaviumCerts = in.CaviumCerts
	out.GoogleCardCerts = in.GoogleCardCerts
	out.GooglePartitionCerts = in.GooglePartitionCerts
	return out
}
