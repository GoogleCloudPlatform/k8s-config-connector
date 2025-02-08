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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Query_FromProto(mapCtx *direct.MapContext, in *pb.Query) *krm.Query {
	if in == nil {
		return nil
	}
	out := &krm.Query{}
	out.Text = direct.LazyPtr(in.GetText())
	out.QueryID = direct.LazyPtr(in.GetQueryId())
	return out
}
func Query_ToProto(mapCtx *direct.MapContext, in *krm.Query) *pb.Query {
	if in == nil {
		return nil
	}
	out := &pb.Query{}
	if oneof := Query_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Content = oneof
	}
	out.QueryId = direct.ValueOf(in.QueryID)
	return out
}
func Session_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.Session {
	if in == nil {
		return nil
	}
	out := &krm.Session{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UserPseudoID = direct.LazyPtr(in.GetUserPseudoId())
	out.Turns = direct.Slice_FromProto(mapCtx, in.Turns, Session_Turn_FromProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func Session_ToProto(mapCtx *direct.MapContext, in *krm.Session) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Session_State](mapCtx, in.State)
	out.UserPseudoId = direct.ValueOf(in.UserPseudoID)
	out.Turns = direct.Slice_ToProto(mapCtx, in.Turns, Session_Turn_ToProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func SessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func SessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func Session_Turn_FromProto(mapCtx *direct.MapContext, in *pb.Session_Turn) *krm.Session_Turn {
	if in == nil {
		return nil
	}
	out := &krm.Session_Turn{}
	out.Query = Query_FromProto(mapCtx, in.GetQuery())
	out.Answer = direct.LazyPtr(in.GetAnswer())
	return out
}
func Session_Turn_ToProto(mapCtx *direct.MapContext, in *krm.Session_Turn) *pb.Session_Turn {
	if in == nil {
		return nil
	}
	out := &pb.Session_Turn{}
	out.Query = Query_ToProto(mapCtx, in.Query)
	out.Answer = direct.ValueOf(in.Answer)
	return out
}
