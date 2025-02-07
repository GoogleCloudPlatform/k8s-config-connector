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

package apigeeregistry

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigeeregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApiVersion_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApiVersion {
	if in == nil {
		return nil
	}
	out := &krm.ApiVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.State = direct.LazyPtr(in.GetState())
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiVersion_ToProto(mapCtx *direct.MapContext, in *krm.ApiVersion) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.State = direct.ValueOf(in.State)
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApiVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApiVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiVersionObservedState) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigeeregistryApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigeeregistryApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiSpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApigeeregistryApiVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiVersionObservedState) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApigeeregistryApiVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiVersionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiVersionSpec) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
