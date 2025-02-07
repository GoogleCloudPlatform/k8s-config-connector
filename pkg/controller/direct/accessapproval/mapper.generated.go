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

package accessapproval

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accessapproval/v1alpha1"
)
func AccessLocations_FromProto(mapCtx *direct.MapContext, in *pb.AccessLocations) *krm.AccessLocations {
	if in == nil {
		return nil
	}
	out := &krm.AccessLocations{}
	out.PrincipalOfficeCountry = direct.LazyPtr(in.GetPrincipalOfficeCountry())
	out.PrincipalPhysicalLocationCountry = direct.LazyPtr(in.GetPrincipalPhysicalLocationCountry())
	return out
}
func AccessLocations_ToProto(mapCtx *direct.MapContext, in *krm.AccessLocations) *pb.AccessLocations {
	if in == nil {
		return nil
	}
	out := &pb.AccessLocations{}
	out.PrincipalOfficeCountry = direct.ValueOf(in.PrincipalOfficeCountry)
	out.PrincipalPhysicalLocationCountry = direct.ValueOf(in.PrincipalPhysicalLocationCountry)
	return out
}
func AccessReason_FromProto(mapCtx *direct.MapContext, in *pb.AccessReason) *krm.AccessReason {
	if in == nil {
		return nil
	}
	out := &krm.AccessReason{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Detail = direct.LazyPtr(in.GetDetail())
	return out
}
func AccessReason_ToProto(mapCtx *direct.MapContext, in *krm.AccessReason) *pb.AccessReason {
	if in == nil {
		return nil
	}
	out := &pb.AccessReason{}
	out.Type = direct.Enum_ToProto[pb.AccessReason_Type](mapCtx, in.Type)
	out.Detail = direct.ValueOf(in.Detail)
	return out
}
func ApproveDecision_FromProto(mapCtx *direct.MapContext, in *pb.ApproveDecision) *krm.ApproveDecision {
	if in == nil {
		return nil
	}
	out := &krm.ApproveDecision{}
	out.ApproveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetApproveTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.InvalidateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetInvalidateTime())
	out.SignatureInfo = SignatureInfo_FromProto(mapCtx, in.GetSignatureInfo())
	out.AutoApproved = direct.LazyPtr(in.GetAutoApproved())
	return out
}
func ApproveDecision_ToProto(mapCtx *direct.MapContext, in *krm.ApproveDecision) *pb.ApproveDecision {
	if in == nil {
		return nil
	}
	out := &pb.ApproveDecision{}
	out.ApproveTime = direct.StringTimestamp_ToProto(mapCtx, in.ApproveTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.InvalidateTime = direct.StringTimestamp_ToProto(mapCtx, in.InvalidateTime)
	out.SignatureInfo = SignatureInfo_ToProto(mapCtx, in.SignatureInfo)
	out.AutoApproved = direct.ValueOf(in.AutoApproved)
	return out
}
func DismissDecision_FromProto(mapCtx *direct.MapContext, in *pb.DismissDecision) *krm.DismissDecision {
	if in == nil {
		return nil
	}
	out := &krm.DismissDecision{}
	out.DismissTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDismissTime())
	out.Implicit = direct.LazyPtr(in.GetImplicit())
	return out
}
func DismissDecision_ToProto(mapCtx *direct.MapContext, in *krm.DismissDecision) *pb.DismissDecision {
	if in == nil {
		return nil
	}
	out := &pb.DismissDecision{}
	out.DismissTime = direct.StringTimestamp_ToProto(mapCtx, in.DismissTime)
	out.Implicit = direct.ValueOf(in.Implicit)
	return out
}
func ResourceProperties_FromProto(mapCtx *direct.MapContext, in *pb.ResourceProperties) *krm.ResourceProperties {
	if in == nil {
		return nil
	}
	out := &krm.ResourceProperties{}
	out.ExcludesDescendants = direct.LazyPtr(in.GetExcludesDescendants())
	return out
}
func ResourceProperties_ToProto(mapCtx *direct.MapContext, in *krm.ResourceProperties) *pb.ResourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.ResourceProperties{}
	out.ExcludesDescendants = direct.ValueOf(in.ExcludesDescendants)
	return out
}
func SignatureInfo_FromProto(mapCtx *direct.MapContext, in *pb.SignatureInfo) *krm.SignatureInfo {
	if in == nil {
		return nil
	}
	out := &krm.SignatureInfo{}
	out.Signature = in.GetSignature()
	out.GooglePublicKeyPem = direct.LazyPtr(in.GetGooglePublicKeyPem())
	out.CustomerKMSKeyVersion = direct.LazyPtr(in.GetCustomerKmsKeyVersion())
	return out
}
func SignatureInfo_ToProto(mapCtx *direct.MapContext, in *krm.SignatureInfo) *pb.SignatureInfo {
	if in == nil {
		return nil
	}
	out := &pb.SignatureInfo{}
	out.Signature = in.Signature
	if oneof := SignatureInfo_GooglePublicKeyPem_ToProto(mapCtx, in.GooglePublicKeyPem); oneof != nil {
		out.VerificationInfo = oneof
	}
	if oneof := SignatureInfo_CustomerKmsKeyVersion_ToProto(mapCtx, in.CustomerKMSKeyVersion); oneof != nil {
		out.VerificationInfo = oneof
	}
	return out
}
