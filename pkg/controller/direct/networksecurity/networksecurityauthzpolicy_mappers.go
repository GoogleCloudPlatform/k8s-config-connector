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

package networksecurity

import (
	"fmt"
	"strconv"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	tagsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AuthzPolicy_Target_Resources_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.ForwardingRuleRef {
	out := make([]computev1beta1.ForwardingRuleRef, len(in))
	for i, r := range in {
		out[i] = computev1beta1.ForwardingRuleRef{External: r}
	}
	return out
}

func AuthzPolicy_Target_Resources_ToProto(mapCtx *direct.MapContext, in []computev1beta1.ForwardingRuleRef) []string {
	out := make([]string, len(in))
	for i, r := range in {
		out[i] = r.External
	}
	return out
}

func AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet) *pb.AuthzPolicy_AuthzRule_RequestResource_TagValueIdSet {
	if in == nil {
		return nil
	}
	out := &pb.AuthzPolicy_AuthzRule_RequestResource_TagValueIdSet{}
	for _, ref := range in.TagValues {
		if ref.External == "" {
			continue
		}
		identity := &tagsv1beta1.TagsTagValueIdentity{}
		if err := identity.FromExternal(ref.External); err != nil {
			mapCtx.Errorf("invalid TagsTagValue external reference %q: %v", ref.External, err)
			continue
		}
		id, err := strconv.ParseInt(identity.TagValue, 10, 64)
		if err != nil {
			mapCtx.Errorf("failed to parse tag value ID %q as int64: %v", identity.TagValue, err)
			continue
		}
		out.Ids = append(out.Ids, id)
	}
	return out
}

func AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthzPolicy_AuthzRule_RequestResource_TagValueIdSet) *krm.AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet {
	if in == nil {
		return nil
	}
	out := &krm.AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet{}
	for _, id := range in.Ids {
		externalRef := fmt.Sprintf("tagValues/%d", id)
		out.TagValues = append(out.TagValues, tagsv1beta1.TagsTagValueRef{
			External: externalRef,
		})
	}
	return out
}

func AuthzPolicy_CustomProvider_CloudIAP_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AuthzPolicy_CustomProvider_CloudIAP) *pb.AuthzPolicy_CustomProvider_CloudIap {
	if in == nil {
		return nil
	}
	if in.Enabled != nil && !*in.Enabled {
		return nil
	}
	return &pb.AuthzPolicy_CustomProvider_CloudIap{}
}

func AuthzPolicy_CustomProvider_CloudIAP_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuthzPolicy_CustomProvider_CloudIap) *krm.AuthzPolicy_CustomProvider_CloudIAP {
	if in == nil {
		return nil
	}
	return &krm.AuthzPolicy_CustomProvider_CloudIAP{
		Enabled: direct.LazyPtr(true),
	}
}
