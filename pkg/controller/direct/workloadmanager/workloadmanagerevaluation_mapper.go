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

package workloadmanager

import (
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workloadmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	pb "cloud.google.com/go/workloadmanager/apiv1/workloadmanagerpb"
)

func ResourceStatus_FromProto(mapCtx *direct.MapContext, in *pb.ResourceStatus) *krm.ResourceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ResourceStatus{}
	if in.State != pb.ResourceStatus_STATE_UNSPECIFIED {
		out.State = direct.LazyPtr(in.State.String())
	}
	return out
}

func ResourceStatus_ToProto(mapCtx *direct.MapContext, in *krm.ResourceStatus) *pb.ResourceStatus {
	if in == nil {
		return nil
	}
	out := &pb.ResourceStatus{}
	if in.State != nil {
		if val, ok := pb.ResourceStatus_State_value[*in.State]; ok {
			out.State = pb.ResourceStatus_State(val)
		} else {
			mapCtx.Errorf("unknown ResourceStatus_State value: %q", *in.State)
		}
	}
	return out
}

func BigQueryDestination_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestination) *krm.BigQueryDestination {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestination{}
	out.DestinationDataset = direct.LazyPtr(in.GetDestinationDataset())
	out.CreateNewResultsTable = direct.LazyPtr(in.GetCreateNewResultsTable())
	return out
}

func BigQueryDestination_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestination) *pb.BigQueryDestination {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestination{}
	out.DestinationDataset = direct.ValueOf(in.DestinationDataset)
	out.CreateNewResultsTable = direct.ValueOf(in.CreateNewResultsTable)
	return out
}

func GceInstanceFilter_FromProto(mapCtx *direct.MapContext, in *pb.GceInstanceFilter) *krm.GCEInstanceFilter {
	if in == nil {
		return nil
	}
	out := &krm.GCEInstanceFilter{}
	out.ServiceAccounts = in.GetServiceAccounts()
	return out
}

func GceInstanceFilter_ToProto(mapCtx *direct.MapContext, in *krm.GCEInstanceFilter) *pb.GceInstanceFilter {
	if in == nil {
		return nil
	}
	out := &pb.GceInstanceFilter{}
	out.ServiceAccounts = in.ServiceAccounts
	return out
}

func ResourceFilter_FromProto(mapCtx *direct.MapContext, in *pb.ResourceFilter) *krm.ResourceFilter {
	if in == nil {
		return nil
	}
	out := &krm.ResourceFilter{}
	out.ResourceIDPatterns = in.GetResourceIdPatterns()
	out.InclusionLabels = in.GetInclusionLabels()
	out.GceInstanceFilter = GceInstanceFilter_FromProto(mapCtx, in.GetGceInstanceFilter())

	for _, scopeStr := range in.GetScopes() {
		scope := krm.Scope{}
		if strings.HasPrefix(scopeStr, "projects/") {
			scope.ProjectRef = &refsv1beta1.ProjectRef{External: strings.TrimPrefix(scopeStr, "projects/")}
		} else if strings.HasPrefix(scopeStr, "folders/") {
			scope.FolderRef = &refsv1beta1.FolderRef{External: strings.TrimPrefix(scopeStr, "folders/")}
		} else if strings.HasPrefix(scopeStr, "organizations/") {
			scope.OrganizationRef = &refsv1beta1.OrganizationRef{External: strings.TrimPrefix(scopeStr, "organizations/")}
		}
		out.ScopeRefs = append(out.ScopeRefs, scope)
	}

	return out
}

func ResourceFilter_ToProto(mapCtx *direct.MapContext, in *krm.ResourceFilter) *pb.ResourceFilter {
	if in == nil {
		return nil
	}
	out := &pb.ResourceFilter{}
	out.ResourceIdPatterns = in.ResourceIDPatterns
	out.InclusionLabels = in.InclusionLabels
	out.GceInstanceFilter = GceInstanceFilter_ToProto(mapCtx, in.GceInstanceFilter)

	for _, scope := range in.ScopeRefs {
		if scope.ProjectRef != nil {
			out.Scopes = append(out.Scopes, "projects/"+scope.ProjectRef.External)
		} else if scope.FolderRef != nil {
			out.Scopes = append(out.Scopes, "folders/"+scope.FolderRef.External)
		} else if scope.OrganizationRef != nil {
			out.Scopes = append(out.Scopes, "organizations/"+scope.OrganizationRef.External)
		}
	}

	return out
}
