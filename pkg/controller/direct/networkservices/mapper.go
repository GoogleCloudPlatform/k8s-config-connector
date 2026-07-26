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

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krmnetworkservicesv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	krmnetworkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func WasmPlugin_LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin_LogConfig) *krmnetworkservicesv1alpha1.WasmPlugin_LogConfig {
	return WasmPlugin_LogConfig_v1alpha1_FromProto(mapCtx, in)
}

func HttprouteRules_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RouteRule) *krmnetworkservicesv1beta1.HttprouteRules {
	return HttprouteRules_v1beta1_FromProto(mapCtx, in)
}

func NetworkServicesServiceBindingSpec_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1alpha1.NetworkServicesServiceBindingSpec) *pb.ServiceBinding {
	return NetworkServicesServiceBindingSpec_v1alpha1_ToProto(mapCtx, in)
}

func HttprouteRules_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRules) *pb.HttpRoute_RouteRule {
	return HttprouteRules_v1beta1_ToProto(mapCtx, in)
}

func NetworkServicesServiceBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1alpha1.NetworkServicesServiceBindingObservedState) *pb.ServiceBinding {
	return NetworkServicesServiceBindingObservedState_v1alpha1_ToProto(mapCtx, in)
}

func NetworkServicesServiceBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krmnetworkservicesv1alpha1.NetworkServicesServiceBindingSpec {
	return NetworkServicesServiceBindingSpec_v1alpha1_FromProto(mapCtx, in)
}

func NetworkServicesServiceBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krmnetworkservicesv1alpha1.NetworkServicesServiceBindingObservedState {
	return NetworkServicesServiceBindingObservedState_v1alpha1_FromProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func WasmPlugin_UsedByObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin_UsedBy) *krmnetworkservicesv1alpha1.WasmPlugin_UsedByObservedState {
	return WasmPlugin_UsedByObservedState_v1alpha1_FromProto(mapCtx, in)
}

func WasmPlugin_UsedByObservedState_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1alpha1.WasmPlugin_UsedByObservedState) *pb.WasmPlugin_UsedBy {
	return WasmPlugin_UsedByObservedState_v1alpha1_ToProto(mapCtx, in)
}

func WasmPlugin_LogConfig_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1alpha1.WasmPlugin_LogConfig) *pb.WasmPlugin_LogConfig {
	return WasmPlugin_LogConfig_v1alpha1_ToProto(mapCtx, in)
}
