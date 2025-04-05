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

package speech

import (
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SpeechCustomClassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.SpeechCustomClassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechCustomClassObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func SpeechCustomClassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechCustomClassObservedState) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.State = direct.Enum_ToProto[pb.CustomClass_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	return out
}
