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
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	api "google.golang.org/api/apigee/v1"
)

func ApigeeEnvgroupAttachmentObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroupAttachment) *krm.ApigeeEnvgroupAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEnvgroupAttachmentObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	out.CreatedAt = direct.LazyPtr(ConvertEpochMillisToTimestamp(in.CreatedAt))
	return out
}

func ApigeeEnvgroupAttachmentObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroupAttachmentObservedState) *api.GoogleCloudApigeeV1EnvironmentGroupAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	out.Name = direct.ValueOf(in.Name)
	i, err := ConvertTimestampToEpochMillis(direct.ValueOf(in.CreatedAt))
	if err != nil {
		mapCtx.Errorf("could not convert CreatedAt value for ObservedState: %w", err)
	} else {
		out.CreatedAt = i
	}
	return out
}

func ApigeeEnvgroupAttachmentSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroupAttachment) *krm.ApigeeEnvgroupAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEnvgroupAttachmentSpec{}
	if in.Environment != "" {
		out.EnvironmentRef = &apigeev1beta1.ApigeeEnvironmentRef{External: in.Environment}
	}
	return out
}

func ApigeeEnvgroupAttachmentSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroupAttachmentSpec) *api.GoogleCloudApigeeV1EnvironmentGroupAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroupAttachment{}
	if in.EnvironmentRef != nil {
		out.Environment = in.EnvironmentRef.External
	}
	return out
}
