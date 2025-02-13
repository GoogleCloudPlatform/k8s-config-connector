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

package apigee

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/apigee/v1"
)

func EnvgroupAttachmentSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroupAttachment) *krm.ApigeeEnvgroupAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEnvgroupAttachmentSpec{}
	out.EnvironmentRef = EnvgroupAttachmentSpec_EnvironmentRef_FromAPI(mapCtx, in.Environment)
	out.EnvgroupRef = EnvgroupAttachmentSpec_EnvgroupRef_FromAPI(mapCtx, in.EnvironmentGroupId)

	return out
}

func EnvgroupAttachmentSpec_EnvgroupRef_FromAPI(mapCtx *direct.MapContext, in string) *krm.ApigeeEnvgroupRef {
	if in == "" {
		return nil
	}

	return &krm.ApigeeEnvgroupRef{External: in}
}

func EnvgroupAttachmentSpec_EnvironmentRef_FromAPI(mapCtx *direct.MapContext, in string) *krmv1beta1.ApigeeEnvironmentRef {
	if in == "" {
		return nil
	}

	return &krmv1beta1.ApigeeEnvironmentRef{External: in}
}

func EnvgroupAttachmentSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroupAttachmentSpec) *api.GoogleCloudApigeeV1EnvironmentGroupAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	out.EnvironmentGroupId = EnvgroupAttachmentSpec_EnvironmentGroupRef_ToAPI(mapCtx, in.EnvgroupRef)
	out.Environment = EnvgroupAttachmentSpec_EnvironmentRef_ToAPI(mapCtx, in.EnvironmentRef)

	return out
}

func EnvgroupAttachmentSpec_EnvironmentGroupRef_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroupRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func EnvgroupAttachmentSpec_EnvironmentRef_ToAPI(mapCtx *direct.MapContext, in *krmv1beta1.ApigeeEnvironmentRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func EnvgroupAttachmentObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroupAttachment) *krm.EnvgroupAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvgroupAttachmentObservedState{}
	out.CreatedAt = direct.LazyPtr(ConvertEpochMillisToTimestamp(in.CreatedAt))

	return out
}

func EnvgroupAttachmentObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.EnvgroupAttachmentObservedState) *api.GoogleCloudApigeeV1EnvironmentGroupAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	i, err := ConvertTimestampToEpochMillis(direct.ValueOf(in.CreatedAt))
	if err != nil {
		mapCtx.Errorf("could not convert CreatedAt value for ObservedState: %w", err)
	} else {
		out.CreatedAt = i
	}

	return out
}
