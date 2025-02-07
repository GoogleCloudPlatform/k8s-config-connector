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
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApiServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ApiServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiServiceObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Config
	// MISSING: State
	return out
}
func ApiServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Config
	// MISSING: State
	return out
}
func ApiServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ApiServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiServiceSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Config
	// MISSING: State
	return out
}
func ApiServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Config
	// MISSING: State
	return out
}
func Service_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.Service {
	if in == nil {
		return nil
	}
	out := &krm.Service{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.Config = ServiceConfig_FromProto(mapCtx, in.GetConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Service_ToProto(mapCtx *direct.MapContext, in *krm.Service) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.Config = ServiceConfig_ToProto(mapCtx, in.Config)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	return out
}
func ServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConfig) *krm.ServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Apis = direct.Slice_FromProto(mapCtx, in.Apis, Api_FromProto)
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.Quota = Quota_FromProto(mapCtx, in.GetQuota())
	out.Authentication = Authentication_FromProto(mapCtx, in.GetAuthentication())
	out.Usage = Usage_FromProto(mapCtx, in.GetUsage())
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Endpoint_FromProto)
	out.MonitoredResources = direct.Slice_FromProto(mapCtx, in.MonitoredResources, MonitoredResourceDescriptor_FromProto)
	out.Monitoring = Monitoring_FromProto(mapCtx, in.GetMonitoring())
	return out
}
func ServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConfig) *pb.ServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.Title = direct.ValueOf(in.Title)
	out.Apis = direct.Slice_ToProto(mapCtx, in.Apis, Api_ToProto)
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.Quota = Quota_ToProto(mapCtx, in.Quota)
	out.Authentication = Authentication_ToProto(mapCtx, in.Authentication)
	out.Usage = Usage_ToProto(mapCtx, in.Usage)
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Endpoint_ToProto)
	out.MonitoredResources = direct.Slice_ToProto(mapCtx, in.MonitoredResources, MonitoredResourceDescriptor_ToProto)
	out.Monitoring = Monitoring_ToProto(mapCtx, in.Monitoring)
	return out
}
