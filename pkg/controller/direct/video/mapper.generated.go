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

package video

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AkamaiCdnKey_FromProto(mapCtx *direct.MapContext, in *pb.AkamaiCdnKey) *krm.AkamaiCdnKey {
	if in == nil {
		return nil
	}
	out := &krm.AkamaiCdnKey{}
	out.TokenKey = in.GetTokenKey()
	return out
}
func AkamaiCdnKey_ToProto(mapCtx *direct.MapContext, in *krm.AkamaiCdnKey) *pb.AkamaiCdnKey {
	if in == nil {
		return nil
	}
	out := &pb.AkamaiCdnKey{}
	out.TokenKey = in.TokenKey
	return out
}
func CdnKey_FromProto(mapCtx *direct.MapContext, in *pb.CdnKey) *krm.CdnKey {
	if in == nil {
		return nil
	}
	out := &krm.CdnKey{}
	out.GoogleCdnKey = GoogleCdnKey_FromProto(mapCtx, in.GetGoogleCdnKey())
	out.AkamaiCdnKey = AkamaiCdnKey_FromProto(mapCtx, in.GetAkamaiCdnKey())
	out.MediaCdnKey = MediaCdnKey_FromProto(mapCtx, in.GetMediaCdnKey())
	out.Name = direct.LazyPtr(in.GetName())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	return out
}
func CdnKey_ToProto(mapCtx *direct.MapContext, in *krm.CdnKey) *pb.CdnKey {
	if in == nil {
		return nil
	}
	out := &pb.CdnKey{}
	if oneof := GoogleCdnKey_ToProto(mapCtx, in.GoogleCdnKey); oneof != nil {
		out.CdnKeyConfig = &pb.CdnKey_GoogleCdnKey{GoogleCdnKey: oneof}
	}
	if oneof := AkamaiCdnKey_ToProto(mapCtx, in.AkamaiCdnKey); oneof != nil {
		out.CdnKeyConfig = &pb.CdnKey_AkamaiCdnKey{AkamaiCdnKey: oneof}
	}
	if oneof := MediaCdnKey_ToProto(mapCtx, in.MediaCdnKey); oneof != nil {
		out.CdnKeyConfig = &pb.CdnKey_MediaCdnKey{MediaCdnKey: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.Hostname = direct.ValueOf(in.Hostname)
	return out
}
func GoogleCdnKey_FromProto(mapCtx *direct.MapContext, in *pb.GoogleCdnKey) *krm.GoogleCdnKey {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCdnKey{}
	out.PrivateKey = in.GetPrivateKey()
	out.KeyName = direct.LazyPtr(in.GetKeyName())
	return out
}
func GoogleCdnKey_ToProto(mapCtx *direct.MapContext, in *krm.GoogleCdnKey) *pb.GoogleCdnKey {
	if in == nil {
		return nil
	}
	out := &pb.GoogleCdnKey{}
	out.PrivateKey = in.PrivateKey
	out.KeyName = direct.ValueOf(in.KeyName)
	return out
}
func MediaCdnKey_FromProto(mapCtx *direct.MapContext, in *pb.MediaCdnKey) *krm.MediaCdnKey {
	if in == nil {
		return nil
	}
	out := &krm.MediaCdnKey{}
	out.PrivateKey = in.GetPrivateKey()
	out.KeyName = direct.LazyPtr(in.GetKeyName())
	out.TokenConfig = MediaCdnKey_TokenConfig_FromProto(mapCtx, in.GetTokenConfig())
	return out
}
func MediaCdnKey_ToProto(mapCtx *direct.MapContext, in *krm.MediaCdnKey) *pb.MediaCdnKey {
	if in == nil {
		return nil
	}
	out := &pb.MediaCdnKey{}
	out.PrivateKey = in.PrivateKey
	out.KeyName = direct.ValueOf(in.KeyName)
	out.TokenConfig = MediaCdnKey_TokenConfig_ToProto(mapCtx, in.TokenConfig)
	return out
}
func MediaCdnKey_TokenConfig_FromProto(mapCtx *direct.MapContext, in *pb.MediaCdnKey_TokenConfig) *krm.MediaCdnKey_TokenConfig {
	if in == nil {
		return nil
	}
	out := &krm.MediaCdnKey_TokenConfig{}
	out.QueryParameter = direct.LazyPtr(in.GetQueryParameter())
	return out
}
func MediaCdnKey_TokenConfig_ToProto(mapCtx *direct.MapContext, in *krm.MediaCdnKey_TokenConfig) *pb.MediaCdnKey_TokenConfig {
	if in == nil {
		return nil
	}
	out := &pb.MediaCdnKey_TokenConfig{}
	out.QueryParameter = direct.ValueOf(in.QueryParameter)
	return out
}
func VideoCdnKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CdnKey) *krm.VideoCdnKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoCdnKeyObservedState{}
	// MISSING: GoogleCdnKey
	// MISSING: AkamaiCdnKey
	// MISSING: MediaCdnKey
	// MISSING: Name
	// MISSING: Hostname
	return out
}
func VideoCdnKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoCdnKeyObservedState) *pb.CdnKey {
	if in == nil {
		return nil
	}
	out := &pb.CdnKey{}
	// MISSING: GoogleCdnKey
	// MISSING: AkamaiCdnKey
	// MISSING: MediaCdnKey
	// MISSING: Name
	// MISSING: Hostname
	return out
}
func VideoCdnKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.CdnKey) *krm.VideoCdnKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoCdnKeySpec{}
	// MISSING: GoogleCdnKey
	// MISSING: AkamaiCdnKey
	// MISSING: MediaCdnKey
	// MISSING: Name
	// MISSING: Hostname
	return out
}
func VideoCdnKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoCdnKeySpec) *pb.CdnKey {
	if in == nil {
		return nil
	}
	out := &pb.CdnKey{}
	// MISSING: GoogleCdnKey
	// MISSING: AkamaiCdnKey
	// MISSING: MediaCdnKey
	// MISSING: Name
	// MISSING: Hostname
	return out
}
