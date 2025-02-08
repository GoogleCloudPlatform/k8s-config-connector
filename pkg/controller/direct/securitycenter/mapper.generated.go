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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
)
func SecurityMarks_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &krm.SecurityMarks{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Marks = in.Marks
	out.CanonicalName = direct.LazyPtr(in.GetCanonicalName())
	return out
}
func SecurityMarks_ToProto(mapCtx *direct.MapContext, in *krm.SecurityMarks) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	out.Name = direct.ValueOf(in.Name)
	out.Marks = in.Marks
	out.CanonicalName = direct.ValueOf(in.CanonicalName)
	return out
}
func SecuritycenterSecurityMarksObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecuritycenterSecurityMarksObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSecurityMarksObservedState{}
	// MISSING: Name
	// MISSING: Marks
	// MISSING: CanonicalName
	return out
}
func SecuritycenterSecurityMarksObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSecurityMarksObservedState) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	// MISSING: Name
	// MISSING: Marks
	// MISSING: CanonicalName
	return out
}
func SecuritycenterSecurityMarksSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecuritycenterSecurityMarksSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSecurityMarksSpec{}
	// MISSING: Name
	// MISSING: Marks
	// MISSING: CanonicalName
	return out
}
func SecuritycenterSecurityMarksSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSecurityMarksSpec) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	// MISSING: Name
	// MISSING: Marks
	// MISSING: CanonicalName
	return out
}
