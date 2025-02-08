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

package securitycenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1beta1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Finding_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.Finding {
	if in == nil {
		return nil
	}
	out := &krm.Finding{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.ExternalURI = direct.LazyPtr(in.GetExternalUri())
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func Finding_ToProto(mapCtx *direct.MapContext, in *krm.Finding) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.ResourceName = direct.ValueOf(in.ResourceName)
	out.State = direct.Enum_ToProto[pb.Finding_State](mapCtx, in.State)
	out.Category = direct.ValueOf(in.Category)
	out.ExternalUri = direct.ValueOf(in.ExternalURI)
	// MISSING: SourceProperties
	// MISSING: SecurityMarks
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func FindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.FindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	out.SecurityMarks = SecurityMarks_FromProto(mapCtx, in.GetSecurityMarks())
	// MISSING: EventTime
	// MISSING: CreateTime
	return out
}
func FindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FindingObservedState) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: State
	// MISSING: Category
	// MISSING: ExternalURI
	// MISSING: SourceProperties
	out.SecurityMarks = SecurityMarks_ToProto(mapCtx, in.SecurityMarks)
	// MISSING: EventTime
	// MISSING: CreateTime
	return out
}
func SecurityMarks_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &krm.SecurityMarks{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Marks = in.Marks
	return out
}
func SecurityMarks_ToProto(mapCtx *direct.MapContext, in *krm.SecurityMarks) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	out.Name = direct.ValueOf(in.Name)
	out.Marks = in.Marks
	return out
}
