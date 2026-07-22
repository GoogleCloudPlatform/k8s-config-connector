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

package siptrunk

import (
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionObservedState{}
	out.ConnectionID = direct.LazyPtr(in.GetConnectionId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ErrorDetails = Connection_ErrorDetailsObservedState_FromProto(mapCtx, in.GetErrorDetails())
	return out
}

func ConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.ConnectionId = direct.ValueOf(in.ConnectionID)
	out.State = direct.Enum_ToProto[pb.Connection_State](mapCtx, in.State)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ErrorDetails = Connection_ErrorDetailsObservedState_ToProto(mapCtx, in.ErrorDetails)
	return out
}

func Connection_ErrorDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection_ErrorDetails) *krm.Connection_ErrorDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Connection_ErrorDetailsObservedState{}
	out.CertificateState = direct.Enum_FromProto(mapCtx, in.GetCertificateState())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}

func Connection_ErrorDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Connection_ErrorDetailsObservedState) *pb.Connection_ErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.Connection_ErrorDetails{}
	if in.CertificateState != nil {
		val := direct.Enum_ToProto[pb.Connection_CertificateState](mapCtx, in.CertificateState)
		out.CertificateState = &val
	}
	out.ErrorMessage = in.ErrorMessage
	return out
}
