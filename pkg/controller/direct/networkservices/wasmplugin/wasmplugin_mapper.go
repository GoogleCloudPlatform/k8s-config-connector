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

package wasmplugin

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesWasmPluginSpec_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin) *krm.NetworkServicesWasmPluginSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesWasmPluginSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MainVersionID = direct.LazyPtr(in.GetMainVersionId())
	out.LogConfig = WasmPlugin_LogConfig_FromProto(mapCtx, in.GetLogConfig())
	if in.GetVersions() != nil {
		out.Versions = make(map[string]krm.WasmPlugin_VersionDetails)
		for k, v := range in.GetVersions() {
			mv := WasmPlugin_VersionDetails_FromProto(mapCtx, v)
			if mv != nil {
				out.Versions[k] = *mv
			}
		}
	}
	return out
}

func NetworkServicesWasmPluginSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesWasmPluginSpec) *pb.WasmPlugin {
	if in == nil {
		return nil
	}
	out := &pb.WasmPlugin{}
	out.Description = direct.ValueOf(in.Description)
	out.MainVersionId = direct.ValueOf(in.MainVersionID)
	out.LogConfig = WasmPlugin_LogConfig_ToProto(mapCtx, in.LogConfig)
	if in.Versions != nil {
		out.Versions = make(map[string]*pb.WasmPlugin_VersionDetails)
		for k, v := range in.Versions {
			pv := WasmPlugin_VersionDetails_ToProto(mapCtx, &v)
			if pv != nil {
				out.Versions[k] = pv
			}
		}
	}
	return out
}

func NetworkServicesWasmPluginObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin) *krm.NetworkServicesWasmPluginObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesWasmPluginObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.UsedBy = direct.Slice_FromProto(mapCtx, in.UsedBy, WasmPlugin_UsedByObservedState_FromProto)
	return out
}

func NetworkServicesWasmPluginObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesWasmPluginObservedState) *pb.WasmPlugin {
	if in == nil {
		return nil
	}
	out := &pb.WasmPlugin{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.UsedBy = direct.Slice_ToProto(mapCtx, in.UsedBy, WasmPlugin_UsedByObservedState_ToProto)
	return out
}

func WasmPlugin_LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin_LogConfig) *krm.WasmPlugin_LogConfig {
	if in == nil {
		return nil
	}
	out := &krm.WasmPlugin_LogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	out.MinLogLevel = direct.Enum_FromProto(mapCtx, in.GetMinLogLevel())
	out.SampleRate = direct.LazyPtr(in.GetSampleRate())
	return out
}

func WasmPlugin_LogConfig_ToProto(mapCtx *direct.MapContext, in *krm.WasmPlugin_LogConfig) *pb.WasmPlugin_LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.WasmPlugin_LogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	out.MinLogLevel = direct.Enum_ToProto[pb.WasmPlugin_LogConfig_LogLevel](mapCtx, in.MinLogLevel)
	out.SampleRate = direct.ValueOf(in.SampleRate)
	return out
}

func WasmPlugin_UsedByObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin_UsedBy) *krm.WasmPlugin_UsedByObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WasmPlugin_UsedByObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}

func WasmPlugin_UsedByObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WasmPlugin_UsedByObservedState) *pb.WasmPlugin_UsedBy {
	if in == nil {
		return nil
	}
	out := &pb.WasmPlugin_UsedBy{}
	out.Name = direct.ValueOf(in.Name)
	return out
}

func WasmPlugin_VersionDetails_FromProto(mapCtx *direct.MapContext, in *pb.WasmPlugin_VersionDetails) *krm.WasmPlugin_VersionDetails {
	if in == nil {
		return nil
	}
	out := &krm.WasmPlugin_VersionDetails{}
	if in.GetPluginConfigData() != nil {
		s := string(in.GetPluginConfigData())
		out.PluginConfigData = &s
	}
	out.PluginConfigURI = direct.LazyPtr(in.GetPluginConfigUri())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	return out
}

func WasmPlugin_VersionDetails_ToProto(mapCtx *direct.MapContext, in *krm.WasmPlugin_VersionDetails) *pb.WasmPlugin_VersionDetails {
	if in == nil {
		return nil
	}
	out := &pb.WasmPlugin_VersionDetails{}
	if in.PluginConfigData != nil {
		out.PluginConfigSource = &pb.WasmPlugin_VersionDetails_PluginConfigData{
			PluginConfigData: []byte(*in.PluginConfigData),
		}
	}
	if in.PluginConfigURI != nil {
		out.PluginConfigSource = &pb.WasmPlugin_VersionDetails_PluginConfigUri{
			PluginConfigUri: *in.PluginConfigURI,
		}
	}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.ImageUri = direct.ValueOf(in.ImageURI)
	return out
}

func WasmPlugin_VersionDetails_PluginConfigData_ToProto(mapCtx *direct.MapContext, in *string) *pb.WasmPlugin_VersionDetails_PluginConfigData {
	if in == nil {
		return nil
	}
	return &pb.WasmPlugin_VersionDetails_PluginConfigData{
		PluginConfigData: []byte(*in),
	}
}

func WasmPlugin_VersionDetails_PluginConfigUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.WasmPlugin_VersionDetails_PluginConfigUri {
	if in == nil {
		return nil
	}
	return &pb.WasmPlugin_VersionDetails_PluginConfigUri{
		PluginConfigUri: *in,
	}
}
