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

// +generated:mapper
// krm.group: parametermanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.parametermanager.v1

package parametermanager

import (
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ParameterManagerParameterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ParameterManagerParameterSpec) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	// MISSING: Labels
	out.Format = direct.Enum_ToProto[pb.ParameterFormat](mapCtx, in.Format)
	// MISSING: PolicyMember
	if in.KMSKeyRef != nil {
		out.KmsKey = &in.KMSKeyRef.External
	}
	return out
}

func ResourcePolicyMemberObservedState_FromProto(mapCtx *direct.MapContext, in *iampb.ResourcePolicyMember) *krm.ResourcePolicyMemberObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcePolicyMemberObservedState{}
	out.IAMPolicyNamePrincipal = direct.LazyPtr(in.GetIamPolicyNamePrincipal())
	out.IAMPolicyUidPrincipal = direct.LazyPtr(in.GetIamPolicyUidPrincipal())
	return out
}

func ResourcePolicyMemberObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcePolicyMemberObservedState) *iampb.ResourcePolicyMember {
	if in == nil {
		return nil
	}
	out := &iampb.ResourcePolicyMember{}
	out.IamPolicyNamePrincipal = direct.ValueOf(in.IAMPolicyNamePrincipal)
	out.IamPolicyUidPrincipal = direct.ValueOf(in.IAMPolicyUidPrincipal)
	return out
}
