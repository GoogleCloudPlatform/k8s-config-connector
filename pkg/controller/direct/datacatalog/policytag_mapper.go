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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.datacatalog.v1beta1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogPolicyTagObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyTag) *krmv1beta1.DataCatalogPolicyTagObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DataCatalogPolicyTagObservedState{}
	return out
}
func DataCatalogPolicyTagObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DataCatalogPolicyTagObservedState) *pb.PolicyTag {
	if in == nil {
		return nil
	}
	out := &pb.PolicyTag{}
	return out
}
func DataCatalogPolicyTagSpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicyTag) *krmv1beta1.DataCatalogPolicyTagSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DataCatalogPolicyTagSpec{}
	out.DisplayName = in.DisplayName
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.ParentPolicyTag != "" {
		out.ParentPolicyTagRef = &krmv1beta1.PolicyTagRef{External: in.GetParentPolicyTag()}
	}
	return out
}
func DataCatalogPolicyTagSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DataCatalogPolicyTagSpec) *pb.PolicyTag {
	if in == nil {
		return nil
	}
	out := &pb.PolicyTag{}
	out.DisplayName = in.DisplayName
	out.Description = direct.ValueOf(in.Description)
	if in.ParentPolicyTagRef != nil {
		out.ParentPolicyTag = in.ParentPolicyTagRef.External
	}
	return out
}
