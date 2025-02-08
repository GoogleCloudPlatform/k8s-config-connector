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

package dataplex

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
)
func DataplexSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DataplexSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexSessionObservedState{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DataplexSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexSessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DataplexSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DataplexSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexSessionSpec{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DataplexSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexSessionSpec) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func Session_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.Session {
	if in == nil {
		return nil
	}
	out := &krm.Session{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func Session_ToProto(mapCtx *direct.MapContext, in *krm.Session) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: UserID
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func SessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UserID = direct.LazyPtr(in.GetUserId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func SessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	out.Name = direct.ValueOf(in.Name)
	out.UserId = direct.ValueOf(in.UserID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	return out
}
