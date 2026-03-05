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
	"strings"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExtensionChain_Extension_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain_Extension) *krm.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain_Extension{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Authority = direct.LazyPtr(in.GetAuthority())

	service := in.GetService()
	if service != "" {
		if strings.Contains(service, "/backendServices/") {
			out.BackendServiceRef = &krmcomputev1beta1.ComputeBackendServiceRef{External: service}
		} else if strings.Contains(service, "/wasmPlugins/") {
			out.WasmPluginRef = &krm.NetworkServicesWasmPluginRef{External: service}
		}
	}

	out.SupportedEvents = direct.EnumSlice_FromProto(mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.FailOpen = direct.LazyPtr(in.GetFailOpen())
	out.ForwardHeaders = in.ForwardHeaders
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	return out
}

func ExtensionChain_Extension_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain_Extension) *pb.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain_Extension{}
	out.Name = direct.ValueOf(in.Name)
	out.Authority = direct.ValueOf(in.Authority)

	if in.BackendServiceRef != nil {
		out.Service = in.BackendServiceRef.External
	} else if in.WasmPluginRef != nil {
		out.Service = in.WasmPluginRef.External
	}

	out.SupportedEvents = direct.EnumSlice_ToProto[pb.EventType](mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.FailOpen = direct.ValueOf(in.FailOpen)
	out.ForwardHeaders = in.ForwardHeaders
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	return out
}
