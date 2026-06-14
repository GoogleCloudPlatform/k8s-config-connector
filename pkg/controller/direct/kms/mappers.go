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
	"strings"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// KMSAutokeyConfigSpec_FromProto converts the protobuf AutokeyConfig to the KRM type.
// This is handcoded because the GCP proto for AutokeyConfig from the googleapis submodule does not contain KeyProjectResolutionMode,
// but the Go SDK's kmspb does, so we override it to include mapping of KeyProjectResolutionMode.
func KMSAutokeyConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigSpec{}
	if in.GetKeyProject() != "" {
		out.KeyProjectRef = &refsv1beta1.ProjectRef{External: in.GetKeyProject()}
	}
	out.KeyProjectResolutionMode = direct.Enum_FromProto(mapCtx, in.GetKeyProjectResolutionMode())
	return out
}

// KMSAutokeyConfigSpec_ToProto converts the KRM type KMSAutokeyConfigSpec to the protobuf representation.
// This is handcoded because the GCP proto for AutokeyConfig from the googleapis submodule does not contain KeyProjectResolutionMode,
// but the Go SDK's kmspb does, so we override it to include mapping of KeyProjectResolutionMode.
func KMSAutokeyConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigSpec) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	if in.KeyProjectRef != nil {
		out.KeyProject = in.KeyProjectRef.External
	}
	if in.KeyProjectResolutionMode != nil {
		out.KeyProjectResolutionMode = direct.Enum_ToProto[pb.AutokeyConfig_KeyProjectResolutionMode](mapCtx, in.KeyProjectResolutionMode)
	}
	return out
}

// CryptoKeyVersionTemplate_FromProto converts the protobuf CryptoKeyVersionTemplate to the KRM type.
// This is handcoded because the KRM's Algorithm field is a required non-pointer string,
// but the generated Enum_FromProto helper returns a pointer string.
func CryptoKeyVersionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersionTemplate) *krm.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	if algoPtr := direct.Enum_FromProto(mapCtx, in.GetAlgorithm()); algoPtr != nil {
		out.Algorithm = *algoPtr
	}
	return out
}

// CryptoKeyVersionTemplate_ToProto converts the KRM type CryptoKeyVersionTemplate to the protobuf representation.
// This is handcoded because the KRM's Algorithm field is a non-pointer string,
// but the generated Enum_ToProto helper expects a pointer string.
func CryptoKeyVersionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersionTemplate) *pb.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, &in.Algorithm)
	return out
}

// KMSImportJobSpec_FromProto converts the protobuf ImportJob to the KRM type KMSImportJobSpec.
// This is handcoded because it handles specific casing transformations or fields that are not automatically mapped.
func KMSImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSImportJobSpec{}
	// MISSING: Name
	out.ImportMethod = direct.Enum_FromProto(mapCtx, in.GetImportMethod())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	return out
}

// KMSImportJobSpec_ToProto converts the KRM type KMSImportJobSpec to the protobuf ImportJob representation.
// This is handcoded to transform enum strings to uppercase to handle mixed-case user input.
func KMSImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name

	if in.ImportMethod != nil {
		importMethodUpper := strings.ToUpper(*in.ImportMethod)
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, &importMethodUpper)
	} else {
		out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, in.ImportMethod)
	}
	if in.ProtectionLevel != nil {
		protectionLevelUpper := strings.ToUpper(*in.ProtectionLevel)
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, &protectionLevelUpper)
	} else {
		out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	}
	return out
}
