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

package apigee

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	api "google.golang.org/api/apigee/v1"
)

func ApigeeNATAddressSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1NatAddress) *krm.ApigeeNATAddressSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeNATAddressSpec{}
	return out
}

func ApigeeNATAddressSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeNATAddressSpec) *api.GoogleCloudApigeeV1NatAddress {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1NatAddress{}
	return out
}

func ApigeeNATAddressStatus_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1NatAddress) *krm.ApigeeNATAddressStatus {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeNATAddressStatus{}
	out.IpAddress = direct.LazyPtr(in.IpAddress)
	out.State = direct.LazyPtr(in.State)
	return out
}

func ApigeeNATAddressStatus_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeNATAddressStatus) *api.GoogleCloudApigeeV1NatAddress {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1NatAddress{}
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.State = direct.ValueOf(in.State)
	return out
}
