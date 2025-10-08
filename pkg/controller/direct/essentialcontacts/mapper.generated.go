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

// +generated:mapper
// krm.group: essentialcontacts.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.essentialcontacts.v1

package essentialcontacts

import (
	pb "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/essentialcontacts/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EssentialContactsContactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.EssentialContactsContactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EssentialContactsContactObservedState{}
	// MISSING: Name
	out.ValidationState = direct.Enum_FromProto(mapCtx, in.GetValidationState())
	out.ValidateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetValidateTime())
	return out
}
func EssentialContactsContactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EssentialContactsContactObservedState) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	// MISSING: Name
	out.ValidationState = direct.Enum_ToProto[pb.ValidationState](mapCtx, in.ValidationState)
	out.ValidateTime = direct.StringTimestamp_ToProto(mapCtx, in.ValidateTime)
	return out
}
func EssentialContactsContactSpec_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.EssentialContactsContactSpec {
	if in == nil {
		return nil
	}
	out := &krm.EssentialContactsContactSpec{}
	// MISSING: Name
	out.Email = direct.LazyPtr(in.GetEmail())
	out.NotificationCategorySubscriptions = direct.EnumSlice_FromProto(mapCtx, in.NotificationCategorySubscriptions)
	out.LanguageTag = direct.LazyPtr(in.GetLanguageTag())
	return out
}
func EssentialContactsContactSpec_ToProto(mapCtx *direct.MapContext, in *krm.EssentialContactsContactSpec) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	// MISSING: Name
	out.Email = direct.ValueOf(in.Email)
	out.NotificationCategorySubscriptions = direct.EnumSlice_ToProto[pb.NotificationCategory](mapCtx, in.NotificationCategorySubscriptions)
	out.LanguageTag = direct.ValueOf(in.LanguageTag)
	return out
}
