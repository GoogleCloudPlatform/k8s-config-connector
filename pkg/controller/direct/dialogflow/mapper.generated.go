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
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DialogflowParticipantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Participant) *krm.DialogflowParticipantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowParticipantObservedState{}
	// MISSING: Name
	// MISSING: Role
	// MISSING: SipRecordingMediaLabel
	// MISSING: ObfuscatedExternalUserID
	// MISSING: DocumentsMetadataFilters
	return out
}
func DialogflowParticipantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowParticipantObservedState) *pb.Participant {
	if in == nil {
		return nil
	}
	out := &pb.Participant{}
	// MISSING: Name
	// MISSING: Role
	// MISSING: SipRecordingMediaLabel
	// MISSING: ObfuscatedExternalUserID
	// MISSING: DocumentsMetadataFilters
	return out
}
func DialogflowParticipantSpec_FromProto(mapCtx *direct.MapContext, in *pb.Participant) *krm.DialogflowParticipantSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowParticipantSpec{}
	// MISSING: Name
	// MISSING: Role
	// MISSING: SipRecordingMediaLabel
	// MISSING: ObfuscatedExternalUserID
	// MISSING: DocumentsMetadataFilters
	return out
}
func DialogflowParticipantSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowParticipantSpec) *pb.Participant {
	if in == nil {
		return nil
	}
	out := &pb.Participant{}
	// MISSING: Name
	// MISSING: Role
	// MISSING: SipRecordingMediaLabel
	// MISSING: ObfuscatedExternalUserID
	// MISSING: DocumentsMetadataFilters
	return out
}
func Participant_FromProto(mapCtx *direct.MapContext, in *pb.Participant) *krm.Participant {
	if in == nil {
		return nil
	}
	out := &krm.Participant{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Role = direct.Enum_FromProto(mapCtx, in.GetRole())
	out.SipRecordingMediaLabel = direct.LazyPtr(in.GetSipRecordingMediaLabel())
	out.ObfuscatedExternalUserID = direct.LazyPtr(in.GetObfuscatedExternalUserId())
	out.DocumentsMetadataFilters = in.DocumentsMetadataFilters
	return out
}
func Participant_ToProto(mapCtx *direct.MapContext, in *krm.Participant) *pb.Participant {
	if in == nil {
		return nil
	}
	out := &pb.Participant{}
	out.Name = direct.ValueOf(in.Name)
	out.Role = direct.Enum_ToProto[pb.Participant_Role](mapCtx, in.Role)
	out.SipRecordingMediaLabel = direct.ValueOf(in.SipRecordingMediaLabel)
	out.ObfuscatedExternalUserId = direct.ValueOf(in.ObfuscatedExternalUserID)
	out.DocumentsMetadataFilters = in.DocumentsMetadataFilters
	return out
}
