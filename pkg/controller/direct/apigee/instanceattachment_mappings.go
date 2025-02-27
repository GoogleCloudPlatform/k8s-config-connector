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

func ApigeeInstanceAttachmentObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1InstanceAttachment) *krm.ApigeeInstanceAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeInstanceAttachmentObservedState{}
	out.CreatedAt = direct.LazyPtr(ConvertEpochMillisToTimestamp(in.CreatedAt))
	return out
}

func ApigeeInstanceAttachmentObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeInstanceAttachmentObservedState) *api.GoogleCloudApigeeV1InstanceAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1InstanceAttachment{}
	i, err := ConvertTimestampToEpochMillis(direct.ValueOf(in.CreatedAt))
	if err != nil {
		mapCtx.Errorf("could not convert CreatedAt value for ObservedState: %w", err)
	} else {
		out.CreatedAt = i
	}
	return out
}

func ApigeeInstanceAttachmentSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1InstanceAttachment) *krm.ApigeeInstanceAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeInstanceAttachmentSpec{}
	if in.Environment != "" {
		out.EnvironmentRef = &apigeev1beta1.ApigeeEnvironmentRef{External: in.Environment}
	}
	return out
}

func ApigeeInstanceAttachmentSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeInstanceAttachmentSpec) *api.GoogleCloudApigeeV1InstanceAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1InstanceAttachment{}
	if in.EnvironmentRef != nil {
		out.Environment = in.EnvironmentRef.External
	}
	return out
}
