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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/apigee/v1"
)

func ApigeeEndpointAttachmentObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EndpointAttachment) *krm.ApigeeEndpointAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEndpointAttachmentObservedState{}
	out.ConnectionState = direct.LazyPtr(in.ConnectionState)
	out.Host = direct.LazyPtr(in.Host)
	out.State = direct.LazyPtr(in.State)
	return out
}

func ApigeeEndpointAttachmentObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEndpointAttachmentObservedState) *api.GoogleCloudApigeeV1EndpointAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EndpointAttachment{}
	out.ConnectionState = direct.ValueOf(in.ConnectionState)
	out.Host = direct.ValueOf(in.Host)
	out.State = direct.ValueOf(in.State)
	return out
}

func ApigeeEndpointAttachmentSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EndpointAttachment) *krm.ApigeeEndpointAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEndpointAttachmentSpec{}
	out.Location = direct.LazyPtr(in.Location)
	if in.ServiceAttachment != "" {
		out.ServiceAttachmentRef = &refs.ComputeServiceAttachmentRef{External: in.ServiceAttachment}
	}
	return out
}

func ApigeeEndpointAttachmentSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeEndpointAttachmentSpec) *api.GoogleCloudApigeeV1EndpointAttachment {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EndpointAttachment{}
	out.Location = direct.ValueOf(in.Location)
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	return out
}
