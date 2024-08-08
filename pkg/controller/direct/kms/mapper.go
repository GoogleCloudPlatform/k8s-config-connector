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

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func KMSCryptoKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKey) *krm.KMSCryptoKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSCryptoKeySpec{}
	// MISSING: Name
	// MISSING: Primary
	out.Purpose = direct.Enum_FromProto(mapCtx, in.GetPurpose())
	// MISSING: CreateTime
	// MISSING: NextRotationTime
	out.RotationPeriod = CryptoKey_RotationPeriod_FromProto(mapCtx, in.GetRotationPeriod())
	out.VersionTemplate = CryptoKeyVersionTemplate_FromProto(mapCtx, in.GetVersionTemplate())
	// MISSING: Labels
	out.ImportOnly = direct.LazyPtr(in.GetImportOnly())
	out.DestroyScheduledDuration = CryptoKey_DestroyScheduledDuration_FromProto(mapCtx, in.GetDestroyScheduledDuration())
	// MISSING: CryptoKeyBackend
	// MISSING: KeyAccessJustificationsPolicy
	return out
}
func KMSCryptoKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSCryptoKeySpec) *pb.CryptoKey {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKey{}
	// MISSING: Name
	// MISSING: Primary
	out.Purpose = direct.Enum_ToProto[pb.CryptoKey_CryptoKeyPurpose](mapCtx, in.Purpose)
	// MISSING: CreateTime
	// MISSING: NextRotationTime
	if oneof := CryptoKey_RotationPeriod_ToProto(mapCtx, in.RotationPeriod); oneof != nil {
		out.RotationSchedule = &pb.CryptoKey_RotationPeriod{RotationPeriod: oneof}
	}
	out.VersionTemplate = CryptoKeyVersionTemplate_ToProto(mapCtx, in.VersionTemplate)
	// MISSING: Labels
	out.ImportOnly = direct.ValueOf(in.ImportOnly)
	out.DestroyScheduledDuration = CryptoKey_DestroyScheduledDuration_ToProto(mapCtx, in.DestroyScheduledDuration)
	// MISSING: CryptoKeyBackend
	// MISSING: KeyAccessJustificationsPolicy
	return out
}

func Certificate_NotAfterTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Certificate_NotAfterTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Certificate_NotBeforeTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Certificate_NotBeforeTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKey_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKey_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKey_NextRotationTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKey_NextRotationTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKeyVersion_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKeyVersion_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKeyVersion_GenerateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKeyVersion_GenerateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKeyVersion_DestroyTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKeyVersion_DestroyTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKeyVersion_DestroyEventTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKeyVersion_DestroyEventTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKeyVersion_ImportTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func CryptoKeyVersion_ImportTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func CryptoKey_RotationPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

func CryptoKey_RotationPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func CryptoKey_DestroyScheduledDuration_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

func CryptoKey_DestroyScheduledDuration_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func EkmConnection_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func EkmConnection_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ImportJob_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ImportJob_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ImportJob_GenerateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ImportJob_GenerateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ImportJob_ExpireEventTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ImportJob_ExpireEventTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ImportJob_ExpireTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ImportJob_ExpireTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func KeyRing_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func KeyRing_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Digest_Sha256_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Digest_Sha256 {
	mapCtx.NotImplemented()
	return nil
}
func Digest_Sha384_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Digest_Sha256 {
	mapCtx.NotImplemented()
	return nil
}
func Digest_Sha512_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Digest_Sha256 {
	mapCtx.NotImplemented()
	return nil
}

func int64_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int64Value) *int64 {
	mapCtx.NotImplemented()
	return nil
}
func int64_ToProto(mapCtx *direct.MapContext, in *int64) *wrapperspb.Int64Value {
	mapCtx.NotImplemented()
	return nil
}
