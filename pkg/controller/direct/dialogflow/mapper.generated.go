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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	// MISSING: ConnectionID
	// MISSING: State
	// MISSING: UpdateTime
	// MISSING: ErrorDetails
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: ConnectionID
	// MISSING: State
	// MISSING: UpdateTime
	// MISSING: ErrorDetails
	return out
}
func ConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionObservedState{}
	out.ConnectionID = direct.LazyPtr(in.GetConnectionId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ErrorDetails = Connection_ErrorDetails_FromProto(mapCtx, in.GetErrorDetails())
	return out
}
func ConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.ConnectionId = direct.ValueOf(in.ConnectionID)
	out.State = direct.Enum_ToProto[pb.Connection_State](mapCtx, in.State)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime); oneof != nil {
		out.UpdateTime = &pb.Connection_UpdateTime{UpdateTime: oneof}
	}
	if oneof := Connection_ErrorDetails_ToProto(mapCtx, in.ErrorDetails); oneof != nil {
		out.ErrorDetails = &pb.Connection_ErrorDetails_{ErrorDetails: oneof}
	}
	return out
}
func Connection_ErrorDetails_FromProto(mapCtx *direct.MapContext, in *pb.Connection_ErrorDetails) *krm.Connection_ErrorDetails {
	if in == nil {
		return nil
	}
	out := &krm.Connection_ErrorDetails{}
	// MISSING: CertificateState
	out.ErrorMessage = in.ErrorMessage
	return out
}
func Connection_ErrorDetails_ToProto(mapCtx *direct.MapContext, in *krm.Connection_ErrorDetails) *pb.Connection_ErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.Connection_ErrorDetails{}
	// MISSING: CertificateState
	out.ErrorMessage = in.ErrorMessage
	return out
}
func Connection_ErrorDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection_ErrorDetails) *krm.Connection_ErrorDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Connection_ErrorDetailsObservedState{}
	out.CertificateState = direct.Enum_FromProto(mapCtx, in.GetCertificateState())
	// MISSING: ErrorMessage
	return out
}
func Connection_ErrorDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Connection_ErrorDetailsObservedState) *pb.Connection_ErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.Connection_ErrorDetails{}
	if oneof := Connection_ErrorDetailsObservedState_CertificateState_ToProto(mapCtx, in.CertificateState); oneof != nil {
		out.CertificateState = oneof
	}
	// MISSING: ErrorMessage
	return out
}
func DialogflowSipTrunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SipTrunk) *krm.DialogflowSipTrunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowSipTrunkObservedState{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	// MISSING: Connections
	// MISSING: DisplayName
	return out
}
func DialogflowSipTrunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowSipTrunkObservedState) *pb.SipTrunk {
	if in == nil {
		return nil
	}
	out := &pb.SipTrunk{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	// MISSING: Connections
	// MISSING: DisplayName
	return out
}
func DialogflowSipTrunkSpec_FromProto(mapCtx *direct.MapContext, in *pb.SipTrunk) *krm.DialogflowSipTrunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowSipTrunkSpec{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	// MISSING: Connections
	// MISSING: DisplayName
	return out
}
func DialogflowSipTrunkSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowSipTrunkSpec) *pb.SipTrunk {
	if in == nil {
		return nil
	}
	out := &pb.SipTrunk{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	// MISSING: Connections
	// MISSING: DisplayName
	return out
}
func SipTrunk_FromProto(mapCtx *direct.MapContext, in *pb.SipTrunk) *krm.SipTrunk {
	if in == nil {
		return nil
	}
	out := &krm.SipTrunk{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExpectedHostname = in.ExpectedHostname
	// MISSING: Connections
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func SipTrunk_ToProto(mapCtx *direct.MapContext, in *krm.SipTrunk) *pb.SipTrunk {
	if in == nil {
		return nil
	}
	out := &pb.SipTrunk{}
	out.Name = direct.ValueOf(in.Name)
	out.ExpectedHostname = in.ExpectedHostname
	// MISSING: Connections
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func SipTrunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SipTrunk) *krm.SipTrunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SipTrunkObservedState{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Connection_FromProto)
	// MISSING: DisplayName
	return out
}
func SipTrunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SipTrunkObservedState) *pb.SipTrunk {
	if in == nil {
		return nil
	}
	out := &pb.SipTrunk{}
	// MISSING: Name
	// MISSING: ExpectedHostname
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Connection_ToProto)
	// MISSING: DisplayName
	return out
}
