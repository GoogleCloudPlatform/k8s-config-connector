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
	"encoding/base64"
	"strings"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// KMSKeyRingImportJobSpec_FromProto maps a pb.ImportJob spec fields to KRM.
// It is handcoded because KRM fields are non-pointer strings, which the generator
// cannot automatically map.
func KMSKeyRingImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSKeyRingImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingImportJobSpec{}
	out.ImportMethod = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetImportMethod()))
	out.ProtectionLevel = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetProtectionLevel()))
	return out
}

// KMSKeyRingImportJobSpec_ToProto maps a KRM KMSKeyRingImportJobSpec to pb.ImportJob.
// It is handcoded because KRM fields are non-pointer strings, requiring custom
// validation/conversion and uppercase transformation.
func KMSKeyRingImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	if in.ImportMethod != "" {
		importMethodUpper := strings.ToUpper(in.ImportMethod)
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, &importMethodUpper)
	} else {
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, nil)
	}
	if in.ProtectionLevel != "" {
		protectionLevelUpper := strings.ToUpper(in.ProtectionLevel)
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, &protectionLevelUpper)
	} else {
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, nil)
	}
	return out
}

// KMSKeyRingImportJobStatus_FromProto maps pb.ImportJob status fields to KRM.
// It is handcoded because KRM fields use slice-of-struct-values rather than
// slice-of-pointers, which the generator cannot automatically construct.
func KMSKeyRingImportJobStatus_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSKeyRingImportJobStatus {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingImportJobStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	if v := in.GetPublicKey(); v != nil {
		if pubKey := KeyringimportjobPublicKeyStatus_FromProto(mapCtx, v); pubKey != nil {
			out.PublicKey = []krm.KeyringimportjobPublicKeyStatus{*pubKey}
		}
	}
	if v := in.GetAttestation(); v != nil {
		if att := KeyringimportjobAttestationStatus_FromProto(mapCtx, v); att != nil {
			out.Attestation = []krm.KeyringimportjobAttestationStatus{*att}
		}
	}
	return out
}

// KMSKeyRingImportJobStatus_ToProto maps KRM KMSKeyRingImportJobStatus to pb.ImportJob.
// It is handcoded because of slice-of-struct-values handling for status fields.
func KMSKeyRingImportJobStatus_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingImportJobStatus) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	out.Name = direct.ValueOf(in.Name)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.State = direct.Enum_ToProto[pb.ImportJob_ImportJobState](mapCtx, in.State)
	if len(in.PublicKey) > 0 {
		out.PublicKey = KeyringimportjobPublicKeyStatus_ToProto(mapCtx, &in.PublicKey[0])
	}
	if len(in.Attestation) > 0 {
		out.Attestation = KeyringimportjobAttestationStatus_ToProto(mapCtx, &in.Attestation[0])
	}
	return out
}

// KeyringimportjobAttestationStatus_FromProto maps pb.KeyOperationAttestation to KRM.
// It is handcoded because KRM content is a base64-encoded string, whereas the proto field
// content is raw []byte.
func KeyringimportjobAttestationStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyOperationAttestation) *krm.KeyringimportjobAttestationStatus {
	if in == nil {
		return nil
	}
	out := &krm.KeyringimportjobAttestationStatus{}
	if in.Content != nil {
		encoded := base64.StdEncoding.EncodeToString(in.Content)
		out.Content = &encoded
	}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	return out
}

// KeyringimportjobAttestationStatus_ToProto maps KRM KeyringimportjobAttestationStatus to pb.KeyOperationAttestation.
// It is handcoded to perform base64 decoding on Content field.
func KeyringimportjobAttestationStatus_ToProto(mapCtx *direct.MapContext, in *krm.KeyringimportjobAttestationStatus) *pb.KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := &pb.KeyOperationAttestation{}
	if in.Content != nil {
		decoded, err := base64.StdEncoding.DecodeString(*in.Content)
		if err != nil {
			mapCtx.Errorf("decoding attestation content from base64: %w", err)
		} else {
			out.Content = decoded
		}
	}
	out.Format = direct.Enum_ToProto[pb.KeyOperationAttestation_AttestationFormat](mapCtx, in.Format)
	return out
}

// KeyringimportjobPublicKeyStatus_FromProto maps pb.ImportJob_WrappingPublicKey to KRM.
// Handcoded to keep signature consistent with attestation mappers.
func KeyringimportjobPublicKeyStatus_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob_WrappingPublicKey) *krm.KeyringimportjobPublicKeyStatus {
	if in == nil {
		return nil
	}
	out := &krm.KeyringimportjobPublicKeyStatus{}
	out.Pem = direct.LazyPtr(in.GetPem())
	return out
}

// KeyringimportjobPublicKeyStatus_ToProto maps KRM KeyringimportjobPublicKeyStatus to pb.ImportJob_WrappingPublicKey.
// Handcoded to keep signature consistent with attestation mappers.
func KeyringimportjobPublicKeyStatus_ToProto(mapCtx *direct.MapContext, in *krm.KeyringimportjobPublicKeyStatus) *pb.ImportJob_WrappingPublicKey {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob_WrappingPublicKey{}
	out.Pem = direct.ValueOf(in.Pem)
	return out
}
