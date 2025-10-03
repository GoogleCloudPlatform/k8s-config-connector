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

package essentialcontacts

import (
	pb "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/essentialcontacts/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Contact_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.Contact {
	if in == nil {
		return nil
	}
	out := &krm.Contact{}
	// MISSING: Name
	out.Email = direct.LazyPtr(in.GetEmail())
	out.NotificationCategorySubscriptions = direct.EnumSlice_FromProto(mapCtx, in.NotificationCategorySubscriptions)
	out.LanguageTag = direct.LazyPtr(in.GetLanguageTag())
	// MISSING: ValidationState
	out.ValidateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetValidateTime())
	return out
}
func Contact_ToProto(mapCtx *direct.MapContext, in *krm.Contact) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	// MISSING: Name
	out.Email = direct.ValueOf(in.Email)
	out.NotificationCategorySubscriptions = direct.EnumSlice_ToProto[pb.NotificationCategory](mapCtx, in.NotificationCategorySubscriptions)
	out.LanguageTag = direct.ValueOf(in.LanguageTag)
	// MISSING: ValidationState
	out.ValidateTime = direct.StringTimestamp_ToProto(mapCtx, in.ValidateTime)
	return out
}
func ContactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.ContactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Email
	// MISSING: NotificationCategorySubscriptions
	// MISSING: LanguageTag
	out.ValidationState = direct.Enum_FromProto(mapCtx, in.GetValidationState())
	// MISSING: ValidateTime
	return out
}
func ContactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactObservedState) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Email
	// MISSING: NotificationCategorySubscriptions
	// MISSING: LanguageTag
	out.ValidationState = direct.Enum_ToProto[pb.ValidationState](mapCtx, in.ValidationState)
	// MISSING: ValidateTime
	return out
}
func EssentialContactsContactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Contact) *krm.EssentialContactsContactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EssentialContactsContactObservedState{}
	// MISSING: Name
	out.ValidationState = direct.Enum_FromProto(mapCtx, in.GetValidationState())
	return out
}
func EssentialContactsContactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EssentialContactsContactObservedState) *pb.Contact {
	if in == nil {
		return nil
	}
	out := &pb.Contact{}
	// MISSING: Name
	out.ValidationState = direct.Enum_ToProto[pb.ValidationState](mapCtx, in.ValidationState)
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
	out.ValidateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetValidateTime())
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
	out.ValidateTime = direct.StringTimestamp_ToProto(mapCtx, in.ValidateTime)
	return out
}
