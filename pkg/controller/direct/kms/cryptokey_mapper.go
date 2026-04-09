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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CryptoKeyVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersion) *krm.CryptoKeyVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func CryptoKeyVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersionObservedState) *pb.CryptoKeyVersion {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionState](mapCtx, in.State)
	return out
}
func CryptoKeyVersionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyVersionTemplate) *krm.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	return out
}
func CryptoKeyVersionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyVersionTemplate) *pb.CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyVersionTemplate{}
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, in.ProtectionLevel)
	out.Algorithm = direct.Enum_ToProto[pb.CryptoKeyVersion_CryptoKeyVersionAlgorithm](mapCtx, in.Algorithm)
	return out
}
func KMSCryptoKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krm.KMSCryptoKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSCryptoKeyObservedState{}
	out.Primary = CryptoKeyVersionObservedState_FromProto(mapCtx, in.GetPrimary())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func KMSCryptoKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSCryptoKeyObservedState) *pb.CryptoKey {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKey{}
	out.Primary = CryptoKeyVersionObservedState_ToProto(mapCtx, in.Primary)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func KMSCryptoKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krm.KMSCryptoKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSCryptoKeySpec{}
	out.Purpose = direct.Enum_FromProto(mapCtx, in.GetPurpose())
	out.RotationPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRotationPeriod())
	out.VersionTemplate = CryptoKeyVersionTemplate_FromProto(mapCtx, in.GetVersionTemplate())
	out.ImportOnly = direct.LazyPtr(in.GetImportOnly())
	out.DestroyScheduledDuration = direct.StringDuration_FromProto(mapCtx, in.GetDestroyScheduledDuration())
	return out
}
func KMSCryptoKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSCryptoKeySpec) *pb.CryptoKey {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKey{}
	out.Purpose = direct.Enum_ToProto[pb.CryptoKey_CryptoKeyPurpose](mapCtx, in.Purpose)
	if oneof := direct.StringDuration_ToProto(mapCtx, in.RotationPeriod); oneof != nil {
		out.RotationSchedule = &pb.CryptoKey_RotationPeriod{RotationPeriod: oneof}
	}
	out.VersionTemplate = CryptoKeyVersionTemplate_ToProto(mapCtx, in.VersionTemplate)
	out.ImportOnly = direct.ValueOf(in.ImportOnly)
	out.DestroyScheduledDuration = direct.StringDuration_ToProto(mapCtx, in.DestroyScheduledDuration)
	return out
}
