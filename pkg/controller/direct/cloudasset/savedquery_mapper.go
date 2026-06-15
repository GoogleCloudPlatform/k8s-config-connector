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

package cloudasset

import (
	"strings"

	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudasset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudAssetIAMPolicyAnalysisQuery_Scope_ToProto(mapCtx *direct.MapContext, in *krm.CloudAssetIAMPolicyAnalysisQuery_Scope) string {
	if in == nil {
		return ""
	}
	if in.ProjectRef != nil {
		return in.ProjectRef.External
	}
	if in.FolderRef != nil {
		return in.FolderRef.External
	}
	if in.OrganizationRef != nil {
		return in.OrganizationRef.External
	}
	return ""
}

func CloudAssetIAMPolicyAnalysisQuery_Scope_FromProto(mapCtx *direct.MapContext, in string) *krm.CloudAssetIAMPolicyAnalysisQuery_Scope {
	if in == "" {
		return nil
	}
	out := &krm.CloudAssetIAMPolicyAnalysisQuery_Scope{}
	if strings.HasPrefix(in, "projects/") {
		out.ProjectRef = &v1beta1.ProjectRef{External: in}
	} else if strings.HasPrefix(in, "folders/") {
		out.FolderRef = &v1beta1.FolderRef{External: in}
	} else if strings.HasPrefix(in, "organizations/") {
		out.OrganizationRef = &v1beta1.OrganizationRef{External: in}
	} else {
		// Fallback
		out.ProjectRef = &v1beta1.ProjectRef{External: in}
	}
	return out
}

func CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FullResourceName_ToProto(mapCtx *direct.MapContext, in *v1beta1.ExtendedProjectRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FullResourceName_FromProto(mapCtx *direct.MapContext, in string) *v1beta1.ExtendedProjectRef {
	if in == "" {
		return nil
	}
	return &v1beta1.ExtendedProjectRef{External: in}
}

func CloudAssetIAMPolicyAnalysisQuery_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery) *krm.CloudAssetIAMPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &krm.CloudAssetIAMPolicyAnalysisQuery{}
	out.Scope = CloudAssetIAMPolicyAnalysisQuery_Scope_FromProto(mapCtx, in.GetScope())
	out.ResourceSelector = CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx, in.GetResourceSelector())
	out.IdentitySelector = CloudAssetIAMPolicyAnalysisQuery_IdentitySelector_FromProto(mapCtx, in.GetIdentitySelector())
	out.AccessSelector = CloudAssetIAMPolicyAnalysisQuery_AccessSelector_FromProto(mapCtx, in.GetAccessSelector())
	out.Options = CloudAssetIAMPolicyAnalysisQuery_Options_FromProto(mapCtx, in.GetOptions())
	out.ConditionContext = CloudAssetIAMPolicyAnalysisQuery_ConditionContext_FromProto(mapCtx, in.GetConditionContext())
	return out
}

func CloudAssetIAMPolicyAnalysisQuery_ToProto(mapCtx *direct.MapContext, in *krm.CloudAssetIAMPolicyAnalysisQuery) *pb.IamPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery{}
	out.Scope = CloudAssetIAMPolicyAnalysisQuery_Scope_ToProto(mapCtx, in.Scope)
	out.ResourceSelector = CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx, in.ResourceSelector)
	out.IdentitySelector = CloudAssetIAMPolicyAnalysisQuery_IdentitySelector_ToProto(mapCtx, in.IdentitySelector)
	out.AccessSelector = CloudAssetIAMPolicyAnalysisQuery_AccessSelector_ToProto(mapCtx, in.AccessSelector)
	out.Options = CloudAssetIAMPolicyAnalysisQuery_Options_ToProto(mapCtx, in.Options)
	out.ConditionContext = CloudAssetIAMPolicyAnalysisQuery_ConditionContext_ToProto(mapCtx, in.ConditionContext)
	return out
}

func CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_ResourceSelector) *krm.CloudAssetIAMPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.CloudAssetIAMPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FullResourceName_FromProto(mapCtx, in.GetFullResourceName())
	return out
}

func CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.CloudAssetIAMPolicyAnalysisQuery_ResourceSelector) *pb.IamPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = CloudAssetIAMPolicyAnalysisQuery_ResourceSelector_FullResourceName_ToProto(mapCtx, in.FullResourceName)
	return out
}

// CloudAssetSavedQueryStatus_FromProto converts the CloudAssetSavedQueryStatus field from its Protobuf representation.
func CloudAssetSavedQueryStatus_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.CloudAssetSavedQueryStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudAssetSavedQueryStatus{}
	out.ObservedState = &krm.CloudAssetSavedQueryObservedState{}
	out.ObservedState.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ObservedState.Creator = direct.LazyPtr(in.GetCreator())
	out.ObservedState.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.ObservedState.LastUpdater = direct.LazyPtr(in.GetLastUpdater())
	return out
}
