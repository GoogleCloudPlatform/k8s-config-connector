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

package connectors

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/connectors/apiv1/connectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ConnectorsProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.ConnectorsProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsProviderObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func ConnectorsProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsProviderObservedState) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func ConnectorsProviderSpec_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.ConnectorsProviderSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsProviderSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func ConnectorsProviderSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsProviderSpec) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func Provider_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.Provider {
	if in == nil {
		return nil
	}
	out := &krm.Provider{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func Provider_ToProto(mapCtx *direct.MapContext, in *krm.Provider) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DocumentationURI
	// MISSING: ExternalURI
	// MISSING: Description
	// MISSING: WebAssetsLocation
	// MISSING: DisplayName
	// MISSING: LaunchStage
	return out
}
func ProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.ProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProviderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Labels = in.Labels
	out.DocumentationURI = direct.LazyPtr(in.GetDocumentationUri())
	out.ExternalURI = direct.LazyPtr(in.GetExternalUri())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.WebAssetsLocation = direct.LazyPtr(in.GetWebAssetsLocation())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	return out
}
func ProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProviderObservedState) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Labels = in.Labels
	out.DocumentationUri = direct.ValueOf(in.DocumentationURI)
	out.ExternalUri = direct.ValueOf(in.ExternalURI)
	out.Description = direct.ValueOf(in.Description)
	out.WebAssetsLocation = direct.ValueOf(in.WebAssetsLocation)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.LaunchStage = direct.Enum_ToProto[pb.LaunchStage](mapCtx, in.LaunchStage)
	return out
}
