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

package cloudcontrolspartner

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1beta/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
)
func EkmConnection_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.EkmConnection {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection{}
	out.ConnectionName = direct.LazyPtr(in.GetConnectionName())
	// MISSING: ConnectionState
	out.ConnectionError = EkmConnection_ConnectionError_FromProto(mapCtx, in.GetConnectionError())
	return out
}
func EkmConnection_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	out.ConnectionName = direct.ValueOf(in.ConnectionName)
	// MISSING: ConnectionState
	out.ConnectionError = EkmConnection_ConnectionError_ToProto(mapCtx, in.ConnectionError)
	return out
}
func EkmConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.EkmConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnectionObservedState{}
	// MISSING: ConnectionName
	out.ConnectionState = direct.Enum_FromProto(mapCtx, in.GetConnectionState())
	// MISSING: ConnectionError
	return out
}
func EkmConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnectionObservedState) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	// MISSING: ConnectionName
	out.ConnectionState = direct.Enum_ToProto[pb.EkmConnection_ConnectionState](mapCtx, in.ConnectionState)
	// MISSING: ConnectionError
	return out
}
func EkmConnection_ConnectionError_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection_ConnectionError) *krm.EkmConnection_ConnectionError {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection_ConnectionError{}
	out.ErrorDomain = direct.LazyPtr(in.GetErrorDomain())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func EkmConnection_ConnectionError_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection_ConnectionError) *pb.EkmConnection_ConnectionError {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection_ConnectionError{}
	out.ErrorDomain = direct.ValueOf(in.ErrorDomain)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
func EkmConnections_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.EkmConnections {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnections{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EkmConnections = direct.Slice_FromProto(mapCtx, in.EkmConnections, EkmConnection_FromProto)
	return out
}
func EkmConnections_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnections) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	out.Name = direct.ValueOf(in.Name)
	out.EkmConnections = direct.Slice_ToProto(mapCtx, in.EkmConnections, EkmConnection_ToProto)
	return out
}
func EkmConnectionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.EkmConnectionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnectionsObservedState{}
	// MISSING: Name
	out.EkmConnections = direct.Slice_FromProto(mapCtx, in.EkmConnections, EkmConnectionObservedState_FromProto)
	return out
}
func EkmConnectionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnectionsObservedState) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	out.EkmConnections = direct.Slice_ToProto(mapCtx, in.EkmConnections, EkmConnectionObservedState_ToProto)
	return out
}
