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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Project_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.Project {
	if in == nil {
		return nil
	}
	out := &krm.Project{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ProvisionCompletionTime
	// MISSING: ServiceTermsMap
	return out
}
func Project_ToProto(mapCtx *direct.MapContext, in *krm.Project) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ProvisionCompletionTime
	// MISSING: ServiceTermsMap
	return out
}
func ProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProjectObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ProvisionCompletionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetProvisionCompletionTime())
	// MISSING: ServiceTermsMap
	return out
}
func ProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProjectObservedState) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ProvisionCompletionTime = direct.StringTimestamp_ToProto(mapCtx, in.ProvisionCompletionTime)
	// MISSING: ServiceTermsMap
	return out
}
func Project_ServiceTerms_FromProto(mapCtx *direct.MapContext, in *pb.Project_ServiceTerms) *krm.Project_ServiceTerms {
	if in == nil {
		return nil
	}
	out := &krm.Project_ServiceTerms{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AcceptTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAcceptTime())
	out.DeclineTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeclineTime())
	return out
}
func Project_ServiceTerms_ToProto(mapCtx *direct.MapContext, in *krm.Project_ServiceTerms) *pb.Project_ServiceTerms {
	if in == nil {
		return nil
	}
	out := &pb.Project_ServiceTerms{}
	out.Id = direct.ValueOf(in.ID)
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.Project_ServiceTerms_State](mapCtx, in.State)
	out.AcceptTime = direct.StringTimestamp_ToProto(mapCtx, in.AcceptTime)
	out.DeclineTime = direct.StringTimestamp_ToProto(mapCtx, in.DeclineTime)
	return out
}
