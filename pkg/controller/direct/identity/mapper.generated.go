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

package identity

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/identity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func IdentityServicePerimeterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeter) *krm.IdentityServicePerimeterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IdentityServicePerimeterObservedState{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PerimeterType
	// MISSING: Status
	// MISSING: Spec
	// MISSING: UseExplicitDryRunSpec
	return out
}
func IdentityServicePerimeterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IdentityServicePerimeterObservedState) *pb.ServicePerimeter {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeter{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PerimeterType
	// MISSING: Status
	// MISSING: Spec
	// MISSING: UseExplicitDryRunSpec
	return out
}
func IdentityServicePerimeterSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeter) *krm.IdentityServicePerimeterSpec {
	if in == nil {
		return nil
	}
	out := &krm.IdentityServicePerimeterSpec{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PerimeterType
	// MISSING: Status
	// MISSING: Spec
	// MISSING: UseExplicitDryRunSpec
	return out
}
func IdentityServicePerimeterSpec_ToProto(mapCtx *direct.MapContext, in *krm.IdentityServicePerimeterSpec) *pb.ServicePerimeter {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeter{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PerimeterType
	// MISSING: Status
	// MISSING: Spec
	// MISSING: UseExplicitDryRunSpec
	return out
}
func ServicePerimeter_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeter) *krm.ServicePerimeter {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PerimeterType = direct.Enum_FromProto(mapCtx, in.GetPerimeterType())
	out.Status = ServicePerimeterConfig_FromProto(mapCtx, in.GetStatus())
	out.Spec = ServicePerimeterConfig_FromProto(mapCtx, in.GetSpec())
	out.UseExplicitDryRunSpec = direct.LazyPtr(in.GetUseExplicitDryRunSpec())
	return out
}
func ServicePerimeter_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeter) *pb.ServicePerimeter {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeter{}
	out.Name = direct.ValueOf(in.Name)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PerimeterType = direct.Enum_ToProto[pb.ServicePerimeter_PerimeterType](mapCtx, in.PerimeterType)
	out.Status = ServicePerimeterConfig_ToProto(mapCtx, in.Status)
	out.Spec = ServicePerimeterConfig_ToProto(mapCtx, in.Spec)
	out.UseExplicitDryRunSpec = direct.ValueOf(in.UseExplicitDryRunSpec)
	return out
}
func ServicePerimeterConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig) *krm.ServicePerimeterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig{}
	out.Resources = in.Resources
	out.AccessLevels = in.AccessLevels
	out.RestrictedServices = in.RestrictedServices
	out.VpcAccessibleServices = ServicePerimeterConfig_VpcAccessibleServices_FromProto(mapCtx, in.GetVpcAccessibleServices())
	out.IngressPolicies = direct.Slice_FromProto(mapCtx, in.IngressPolicies, ServicePerimeterConfig_IngressPolicy_FromProto)
	out.EgressPolicies = direct.Slice_FromProto(mapCtx, in.EgressPolicies, ServicePerimeterConfig_EgressPolicy_FromProto)
	return out
}
func ServicePerimeterConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig) *pb.ServicePerimeterConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig{}
	out.Resources = in.Resources
	out.AccessLevels = in.AccessLevels
	out.RestrictedServices = in.RestrictedServices
	out.VpcAccessibleServices = ServicePerimeterConfig_VpcAccessibleServices_ToProto(mapCtx, in.VpcAccessibleServices)
	out.IngressPolicies = direct.Slice_ToProto(mapCtx, in.IngressPolicies, ServicePerimeterConfig_IngressPolicy_ToProto)
	out.EgressPolicies = direct.Slice_ToProto(mapCtx, in.EgressPolicies, ServicePerimeterConfig_EgressPolicy_ToProto)
	return out
}
func ServicePerimeterConfig_ApiOperation_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_ApiOperation) *krm.ServicePerimeterConfig_ApiOperation {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_ApiOperation{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.MethodSelectors = direct.Slice_FromProto(mapCtx, in.MethodSelectors, ServicePerimeterConfig_MethodSelector_FromProto)
	return out
}
func ServicePerimeterConfig_ApiOperation_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_ApiOperation) *pb.ServicePerimeterConfig_ApiOperation {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_ApiOperation{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.MethodSelectors = direct.Slice_ToProto(mapCtx, in.MethodSelectors, ServicePerimeterConfig_MethodSelector_ToProto)
	return out
}
func ServicePerimeterConfig_EgressFrom_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_EgressFrom) *krm.ServicePerimeterConfig_EgressFrom {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_EgressFrom{}
	out.Identities = in.Identities
	out.IdentityType = direct.Enum_FromProto(mapCtx, in.GetIdentityType())
	return out
}
func ServicePerimeterConfig_EgressFrom_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_EgressFrom) *pb.ServicePerimeterConfig_EgressFrom {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_EgressFrom{}
	out.Identities = in.Identities
	out.IdentityType = direct.Enum_ToProto[pb.ServicePerimeterConfig_IdentityType](mapCtx, in.IdentityType)
	return out
}
func ServicePerimeterConfig_EgressPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_EgressPolicy) *krm.ServicePerimeterConfig_EgressPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_EgressPolicy{}
	out.EgressFrom = ServicePerimeterConfig_EgressFrom_FromProto(mapCtx, in.GetEgressFrom())
	out.EgressTo = ServicePerimeterConfig_EgressTo_FromProto(mapCtx, in.GetEgressTo())
	return out
}
func ServicePerimeterConfig_EgressPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_EgressPolicy) *pb.ServicePerimeterConfig_EgressPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_EgressPolicy{}
	out.EgressFrom = ServicePerimeterConfig_EgressFrom_ToProto(mapCtx, in.EgressFrom)
	out.EgressTo = ServicePerimeterConfig_EgressTo_ToProto(mapCtx, in.EgressTo)
	return out
}
func ServicePerimeterConfig_EgressTo_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_EgressTo) *krm.ServicePerimeterConfig_EgressTo {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_EgressTo{}
	out.Resources = in.Resources
	out.Operations = direct.Slice_FromProto(mapCtx, in.Operations, ServicePerimeterConfig_ApiOperation_FromProto)
	out.ExternalResources = in.ExternalResources
	return out
}
func ServicePerimeterConfig_EgressTo_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_EgressTo) *pb.ServicePerimeterConfig_EgressTo {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_EgressTo{}
	out.Resources = in.Resources
	out.Operations = direct.Slice_ToProto(mapCtx, in.Operations, ServicePerimeterConfig_ApiOperation_ToProto)
	out.ExternalResources = in.ExternalResources
	return out
}
func ServicePerimeterConfig_IngressFrom_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_IngressFrom) *krm.ServicePerimeterConfig_IngressFrom {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_IngressFrom{}
	out.Sources = direct.Slice_FromProto(mapCtx, in.Sources, ServicePerimeterConfig_IngressSource_FromProto)
	out.Identities = in.Identities
	out.IdentityType = direct.Enum_FromProto(mapCtx, in.GetIdentityType())
	return out
}
func ServicePerimeterConfig_IngressFrom_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_IngressFrom) *pb.ServicePerimeterConfig_IngressFrom {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_IngressFrom{}
	out.Sources = direct.Slice_ToProto(mapCtx, in.Sources, ServicePerimeterConfig_IngressSource_ToProto)
	out.Identities = in.Identities
	out.IdentityType = direct.Enum_ToProto[pb.ServicePerimeterConfig_IdentityType](mapCtx, in.IdentityType)
	return out
}
func ServicePerimeterConfig_IngressPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_IngressPolicy) *krm.ServicePerimeterConfig_IngressPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_IngressPolicy{}
	out.IngressFrom = ServicePerimeterConfig_IngressFrom_FromProto(mapCtx, in.GetIngressFrom())
	out.IngressTo = ServicePerimeterConfig_IngressTo_FromProto(mapCtx, in.GetIngressTo())
	return out
}
func ServicePerimeterConfig_IngressPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_IngressPolicy) *pb.ServicePerimeterConfig_IngressPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_IngressPolicy{}
	out.IngressFrom = ServicePerimeterConfig_IngressFrom_ToProto(mapCtx, in.IngressFrom)
	out.IngressTo = ServicePerimeterConfig_IngressTo_ToProto(mapCtx, in.IngressTo)
	return out
}
func ServicePerimeterConfig_IngressSource_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_IngressSource) *krm.ServicePerimeterConfig_IngressSource {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_IngressSource{}
	out.AccessLevel = direct.LazyPtr(in.GetAccessLevel())
	out.Resource = direct.LazyPtr(in.GetResource())
	return out
}
func ServicePerimeterConfig_IngressSource_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_IngressSource) *pb.ServicePerimeterConfig_IngressSource {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_IngressSource{}
	if oneof := ServicePerimeterConfig_IngressSource_AccessLevel_ToProto(mapCtx, in.AccessLevel); oneof != nil {
		out.Source = oneof
	}
	if oneof := ServicePerimeterConfig_IngressSource_Resource_ToProto(mapCtx, in.Resource); oneof != nil {
		out.Source = oneof
	}
	return out
}
func ServicePerimeterConfig_IngressTo_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_IngressTo) *krm.ServicePerimeterConfig_IngressTo {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_IngressTo{}
	out.Operations = direct.Slice_FromProto(mapCtx, in.Operations, ServicePerimeterConfig_ApiOperation_FromProto)
	out.Resources = in.Resources
	return out
}
func ServicePerimeterConfig_IngressTo_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_IngressTo) *pb.ServicePerimeterConfig_IngressTo {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_IngressTo{}
	out.Operations = direct.Slice_ToProto(mapCtx, in.Operations, ServicePerimeterConfig_ApiOperation_ToProto)
	out.Resources = in.Resources
	return out
}
func ServicePerimeterConfig_MethodSelector_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_MethodSelector) *krm.ServicePerimeterConfig_MethodSelector {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_MethodSelector{}
	out.Method = direct.LazyPtr(in.GetMethod())
	out.Permission = direct.LazyPtr(in.GetPermission())
	return out
}
func ServicePerimeterConfig_MethodSelector_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_MethodSelector) *pb.ServicePerimeterConfig_MethodSelector {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_MethodSelector{}
	if oneof := ServicePerimeterConfig_MethodSelector_Method_ToProto(mapCtx, in.Method); oneof != nil {
		out.Kind = oneof
	}
	if oneof := ServicePerimeterConfig_MethodSelector_Permission_ToProto(mapCtx, in.Permission); oneof != nil {
		out.Kind = oneof
	}
	return out
}
func ServicePerimeterConfig_VpcAccessibleServices_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_VpcAccessibleServices) *krm.ServicePerimeterConfig_VpcAccessibleServices {
	if in == nil {
		return nil
	}
	out := &krm.ServicePerimeterConfig_VpcAccessibleServices{}
	out.EnableRestriction = direct.LazyPtr(in.GetEnableRestriction())
	out.AllowedServices = in.AllowedServices
	return out
}
func ServicePerimeterConfig_VpcAccessibleServices_ToProto(mapCtx *direct.MapContext, in *krm.ServicePerimeterConfig_VpcAccessibleServices) *pb.ServicePerimeterConfig_VpcAccessibleServices {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_VpcAccessibleServices{}
	out.EnableRestriction = direct.ValueOf(in.EnableRestriction)
	out.AllowedServices = in.AllowedServices
	return out
}
