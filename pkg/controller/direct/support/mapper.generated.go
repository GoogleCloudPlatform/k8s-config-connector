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

package support

import (
	pb "cloud.google.com/go/support/apiv2/supportpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/support/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Actor_FromProto(mapCtx *direct.MapContext, in *pb.Actor) *krm.Actor {
	if in == nil {
		return nil
	}
	out := &krm.Actor{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	// MISSING: GoogleSupport
	return out
}
func Actor_ToProto(mapCtx *direct.MapContext, in *krm.Actor) *pb.Actor {
	if in == nil {
		return nil
	}
	out := &pb.Actor{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	// MISSING: GoogleSupport
	return out
}
func ActorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Actor) *krm.ActorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ActorObservedState{}
	// MISSING: DisplayName
	// MISSING: Email
	out.GoogleSupport = direct.LazyPtr(in.GetGoogleSupport())
	return out
}
func ActorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ActorObservedState) *pb.Actor {
	if in == nil {
		return nil
	}
	out := &pb.Actor{}
	// MISSING: DisplayName
	// MISSING: Email
	out.GoogleSupport = direct.ValueOf(in.GoogleSupport)
	return out
}
func Case_FromProto(mapCtx *direct.MapContext, in *pb.Case) *krm.Case {
	if in == nil {
		return nil
	}
	out := &krm.Case{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Classification = CaseClassification_FromProto(mapCtx, in.GetClassification())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.SubscriberEmailAddresses = in.SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Creator = Actor_FromProto(mapCtx, in.GetCreator())
	out.ContactEmail = direct.LazyPtr(in.GetContactEmail())
	out.Escalated = direct.LazyPtr(in.GetEscalated())
	out.TestCase = direct.LazyPtr(in.GetTestCase())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func Case_ToProto(mapCtx *direct.MapContext, in *krm.Case) *pb.Case {
	if in == nil {
		return nil
	}
	out := &pb.Case{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Classification = CaseClassification_ToProto(mapCtx, in.Classification)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.SubscriberEmailAddresses = in.SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Creator = Actor_ToProto(mapCtx, in.Creator)
	out.ContactEmail = direct.ValueOf(in.ContactEmail)
	out.Escalated = direct.ValueOf(in.Escalated)
	out.TestCase = direct.ValueOf(in.TestCase)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.Priority = direct.Enum_ToProto[pb.Case_Priority](mapCtx, in.Priority)
	return out
}
func CaseClassification_FromProto(mapCtx *direct.MapContext, in *pb.CaseClassification) *krm.CaseClassification {
	if in == nil {
		return nil
	}
	out := &krm.CaseClassification{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func CaseClassification_ToProto(mapCtx *direct.MapContext, in *krm.CaseClassification) *pb.CaseClassification {
	if in == nil {
		return nil
	}
	out := &pb.CaseClassification{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func CaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Case) *krm.CaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CaseObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Creator = ActorObservedState_FromProto(mapCtx, in.GetCreator())
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
func CaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CaseObservedState) *pb.Case {
	if in == nil {
		return nil
	}
	out := &pb.Case{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	out.State = direct.Enum_ToProto[pb.Case_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Creator = ActorObservedState_ToProto(mapCtx, in.Creator)
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
func SupportCaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Case) *krm.SupportCaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SupportCaseObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Creator
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
func SupportCaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SupportCaseObservedState) *pb.Case {
	if in == nil {
		return nil
	}
	out := &pb.Case{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Creator
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
func SupportCaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.Case) *krm.SupportCaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.SupportCaseSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Creator
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
func SupportCaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.SupportCaseSpec) *pb.Case {
	if in == nil {
		return nil
	}
	out := &pb.Case{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Classification
	// MISSING: TimeZone
	// MISSING: SubscriberEmailAddresses
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Creator
	// MISSING: ContactEmail
	// MISSING: Escalated
	// MISSING: TestCase
	// MISSING: LanguageCode
	// MISSING: Priority
	return out
}
