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

package api

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apikeys/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/api/v1alpha1"
)
func APIKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.APIKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIKeyObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func APIKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIKeyObservedState) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func APIKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.APIKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.APIKeySpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	out.Restrictions = Restrictions_FromProto(mapCtx, in.GetRestrictions())
	// MISSING: Etag
	return out
}
func APIKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.APIKeySpec) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	out.Restrictions = Restrictions_ToProto(mapCtx, in.Restrictions)
	// MISSING: Etag
	return out
}
func AndroidApplication_FromProto(mapCtx *direct.MapContext, in *pb.AndroidApplication) *krm.AndroidApplication {
	if in == nil {
		return nil
	}
	out := &krm.AndroidApplication{}
	out.Sha1Fingerprint = direct.LazyPtr(in.GetSha1Fingerprint())
	out.PackageName = direct.LazyPtr(in.GetPackageName())
	return out
}
func AndroidApplication_ToProto(mapCtx *direct.MapContext, in *krm.AndroidApplication) *pb.AndroidApplication {
	if in == nil {
		return nil
	}
	out := &pb.AndroidApplication{}
	out.Sha1Fingerprint = direct.ValueOf(in.Sha1Fingerprint)
	out.PackageName = direct.ValueOf(in.PackageName)
	return out
}
func AndroidKeyRestrictions_FromProto(mapCtx *direct.MapContext, in *pb.AndroidKeyRestrictions) *krm.AndroidKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &krm.AndroidKeyRestrictions{}
	out.AllowedApplications = direct.Slice_FromProto(mapCtx, in.AllowedApplications, AndroidApplication_FromProto)
	return out
}
func AndroidKeyRestrictions_ToProto(mapCtx *direct.MapContext, in *krm.AndroidKeyRestrictions) *pb.AndroidKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &pb.AndroidKeyRestrictions{}
	out.AllowedApplications = direct.Slice_ToProto(mapCtx, in.AllowedApplications, AndroidApplication_ToProto)
	return out
}
func ApiKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.ApiKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiKeyObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Restrictions
	// MISSING: Etag
	return out
}
func ApiKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiKeyObservedState) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Restrictions
	// MISSING: Etag
	return out
}
func ApiKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.ApiKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiKeySpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Restrictions
	// MISSING: Etag
	return out
}
func ApiKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiKeySpec) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: DisplayName
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Annotations
	// MISSING: Restrictions
	// MISSING: Etag
	return out
}
func ApiTarget_FromProto(mapCtx *direct.MapContext, in *pb.ApiTarget) *krm.ApiTarget {
	if in == nil {
		return nil
	}
	out := &krm.ApiTarget{}
	out.Service = direct.LazyPtr(in.GetService())
	out.Methods = in.Methods
	return out
}
func ApiTarget_ToProto(mapCtx *direct.MapContext, in *krm.ApiTarget) *pb.ApiTarget {
	if in == nil {
		return nil
	}
	out := &pb.ApiTarget{}
	out.Service = direct.ValueOf(in.Service)
	out.Methods = in.Methods
	return out
}
func BrowserKeyRestrictions_FromProto(mapCtx *direct.MapContext, in *pb.BrowserKeyRestrictions) *krm.BrowserKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &krm.BrowserKeyRestrictions{}
	out.AllowedReferrers = in.AllowedReferrers
	return out
}
func BrowserKeyRestrictions_ToProto(mapCtx *direct.MapContext, in *krm.BrowserKeyRestrictions) *pb.BrowserKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &pb.BrowserKeyRestrictions{}
	out.AllowedReferrers = in.AllowedReferrers
	return out
}
func IosKeyRestrictions_FromProto(mapCtx *direct.MapContext, in *pb.IosKeyRestrictions) *krm.IosKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &krm.IosKeyRestrictions{}
	out.AllowedBundleIds = in.AllowedBundleIds
	return out
}
func IosKeyRestrictions_ToProto(mapCtx *direct.MapContext, in *krm.IosKeyRestrictions) *pb.IosKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &pb.IosKeyRestrictions{}
	out.AllowedBundleIds = in.AllowedBundleIds
	return out
}
func Key_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.Key {
	if in == nil {
		return nil
	}
	out := &krm.Key{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Annotations = in.Annotations
	out.Restrictions = Restrictions_FromProto(mapCtx, in.GetRestrictions())
	// MISSING: Etag
	return out
}
func Key_ToProto(mapCtx *direct.MapContext, in *krm.Key) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: Uid
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: KeyString
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Annotations = in.Annotations
	out.Restrictions = Restrictions_ToProto(mapCtx, in.Restrictions)
	// MISSING: Etag
	return out
}
func KeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.KeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KeyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: DisplayName
	out.KeyString = direct.LazyPtr(in.GetKeyString())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Annotations
	// MISSING: Restrictions
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func KeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KeyObservedState) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: DisplayName
	out.KeyString = direct.ValueOf(in.KeyString)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Annotations
	// MISSING: Restrictions
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func Key_AnnotationsEntry_FromProto(mapCtx *direct.MapContext, in *pb.Key_AnnotationsEntry) *krm.Key_AnnotationsEntry {
	if in == nil {
		return nil
	}
	out := &krm.Key_AnnotationsEntry{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func Key_AnnotationsEntry_ToProto(mapCtx *direct.MapContext, in *krm.Key_AnnotationsEntry) *pb.Key_AnnotationsEntry {
	if in == nil {
		return nil
	}
	out := &pb.Key_AnnotationsEntry{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func Restrictions_FromProto(mapCtx *direct.MapContext, in *pb.Restrictions) *krm.Restrictions {
	if in == nil {
		return nil
	}
	out := &krm.Restrictions{}
	out.BrowserKeyRestrictions = BrowserKeyRestrictions_FromProto(mapCtx, in.GetBrowserKeyRestrictions())
	out.ServerKeyRestrictions = ServerKeyRestrictions_FromProto(mapCtx, in.GetServerKeyRestrictions())
	out.AndroidKeyRestrictions = AndroidKeyRestrictions_FromProto(mapCtx, in.GetAndroidKeyRestrictions())
	out.IosKeyRestrictions = IosKeyRestrictions_FromProto(mapCtx, in.GetIosKeyRestrictions())
	out.ApiTargets = direct.Slice_FromProto(mapCtx, in.ApiTargets, ApiTarget_FromProto)
	return out
}
func Restrictions_ToProto(mapCtx *direct.MapContext, in *krm.Restrictions) *pb.Restrictions {
	if in == nil {
		return nil
	}
	out := &pb.Restrictions{}
	if oneof := BrowserKeyRestrictions_ToProto(mapCtx, in.BrowserKeyRestrictions); oneof != nil {
		out.ClientRestrictions = &pb.Restrictions_BrowserKeyRestrictions{BrowserKeyRestrictions: oneof}
	}
	if oneof := ServerKeyRestrictions_ToProto(mapCtx, in.ServerKeyRestrictions); oneof != nil {
		out.ClientRestrictions = &pb.Restrictions_ServerKeyRestrictions{ServerKeyRestrictions: oneof}
	}
	if oneof := AndroidKeyRestrictions_ToProto(mapCtx, in.AndroidKeyRestrictions); oneof != nil {
		out.ClientRestrictions = &pb.Restrictions_AndroidKeyRestrictions{AndroidKeyRestrictions: oneof}
	}
	if oneof := IosKeyRestrictions_ToProto(mapCtx, in.IosKeyRestrictions); oneof != nil {
		out.ClientRestrictions = &pb.Restrictions_IosKeyRestrictions{IosKeyRestrictions: oneof}
	}
	out.ApiTargets = direct.Slice_ToProto(mapCtx, in.ApiTargets, ApiTarget_ToProto)
	return out
}
func ServerKeyRestrictions_FromProto(mapCtx *direct.MapContext, in *pb.ServerKeyRestrictions) *krm.ServerKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &krm.ServerKeyRestrictions{}
	out.AllowedIps = in.AllowedIps
	return out
}
func ServerKeyRestrictions_ToProto(mapCtx *direct.MapContext, in *krm.ServerKeyRestrictions) *pb.ServerKeyRestrictions {
	if in == nil {
		return nil
	}
	out := &pb.ServerKeyRestrictions{}
	out.AllowedIps = in.AllowedIps
	return out
}
