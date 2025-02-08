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
	pb "cloud.google.com/go/connectors/apiv1/connectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Connector_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.Connector {
	if in == nil {
		return nil
	}
	out := &krm.Connector{}
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
func Connector_ToProto(mapCtx *direct.MapContext, in *krm.Connector) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
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
func ConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.ConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorObservedState{}
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
func ConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorObservedState) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
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
func ConnectorsConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.ConnectorsConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectorObservedState{}
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
func ConnectorsConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectorObservedState) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
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
func ConnectorsConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.ConnectorsConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectorSpec{}
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
func ConnectorsConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectorSpec) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
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
