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

package iap

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iap/apiv1/iappb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func IAPSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.IapSettings) *krm.IAPSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.IAPSettingsSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AccessSettings = AccessSettings_FromProto(mapCtx, in.GetAccessSettings())
	out.ApplicationSettings = ApplicationSettings_FromProto(mapCtx, in.GetApplicationSettings())
	return out
}
func IAPSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.IAPSettingsSpec) *pb.IapSettings {
	if in == nil {
		return nil
	}
	out := &pb.IapSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.AccessSettings = AccessSettings_ToProto(mapCtx, in.AccessSettings)
	out.ApplicationSettings = ApplicationSettings_ToProto(mapCtx, in.ApplicationSettings)
	return out
}
func IapTunnelDestGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TunnelDestGroup) *krm.IapTunnelDestGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IapTunnelDestGroupObservedState{}
	// MISSING: Name
	// MISSING: Cidrs
	// MISSING: Fqdns
	return out
}
func IapTunnelDestGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IapTunnelDestGroupObservedState) *pb.TunnelDestGroup {
	if in == nil {
		return nil
	}
	out := &pb.TunnelDestGroup{}
	// MISSING: Name
	// MISSING: Cidrs
	// MISSING: Fqdns
	return out
}
func IapTunnelDestGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.TunnelDestGroup) *krm.IapTunnelDestGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.IapTunnelDestGroupSpec{}
	// MISSING: Name
	// MISSING: Cidrs
	// MISSING: Fqdns
	return out
}
func IapTunnelDestGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.IapTunnelDestGroupSpec) *pb.TunnelDestGroup {
	if in == nil {
		return nil
	}
	out := &pb.TunnelDestGroup{}
	// MISSING: Name
	// MISSING: Cidrs
	// MISSING: Fqdns
	return out
}
func TunnelDestGroup_FromProto(mapCtx *direct.MapContext, in *pb.TunnelDestGroup) *krm.TunnelDestGroup {
	if in == nil {
		return nil
	}
	out := &krm.TunnelDestGroup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Cidrs = in.Cidrs
	out.Fqdns = in.Fqdns
	return out
}
func TunnelDestGroup_ToProto(mapCtx *direct.MapContext, in *krm.TunnelDestGroup) *pb.TunnelDestGroup {
	if in == nil {
		return nil
	}
	out := &pb.TunnelDestGroup{}
	out.Name = direct.ValueOf(in.Name)
	out.Cidrs = in.Cidrs
	out.Fqdns = in.Fqdns
	return out
}
