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

package transcoder

import (
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/transcoder/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Encryption_FromProto(mapCtx *direct.MapContext, in *pb.Encryption) *krm.Encryption {
	if in == nil {
		return nil
	}
	out := &krm.Encryption{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Aes128 = Encryption_Aes128Encryption_FromProto(mapCtx, in.GetAes_128())
	out.SampleAes = Encryption_SampleAesEncryption_FromProto(mapCtx, in.GetSampleAes())
	out.MpegCenc = Encryption_MpegCommonEncryption_FromProto(mapCtx, in.GetMpegCenc())
	out.SecretManagerKeySource = Encryption_SecretManagerSource_FromProto(mapCtx, in.GetSecretManagerKeySource())
	out.DrmSystems = Encryption_DrmSystems_FromProto(mapCtx, in.GetDrmSystems())
	return out
}

func Encryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption) *pb.Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := Encryption_Aes128Encryption_ToProto(mapCtx, in.Aes128); oneof != nil {
		out.EncryptionMode = &pb.Encryption_Aes_128{Aes_128: oneof}
	}
	if oneof := Encryption_SampleAesEncryption_ToProto(mapCtx, in.SampleAes); oneof != nil {
		out.EncryptionMode = &pb.Encryption_SampleAes{SampleAes: oneof}
	}
	if oneof := Encryption_MpegCommonEncryption_ToProto(mapCtx, in.MpegCenc); oneof != nil {
		out.EncryptionMode = &pb.Encryption_MpegCenc{MpegCenc: oneof}
	}
	if oneof := Encryption_SecretManagerSource_ToProto(mapCtx, in.SecretManagerKeySource); oneof != nil {
		out.SecretSource = &pb.Encryption_SecretManagerKeySource{SecretManagerKeySource: oneof}
	}
	out.DrmSystems = Encryption_DrmSystems_ToProto(mapCtx, in.DrmSystems)
	return out
}
