// Copyright 2024 Google LLC
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

func ApigeeEnvgroupObservedState_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroup) *krm.ApigeeEnvgroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEnvgroupObservedState{}
	out.CreatedAt = direct.LazyPtr(in.CreatedAt)
	out.LastModifiedAt = direct.LazyPtr(in.LastModifiedAt)
	out.Name = direct.LazyPtr(in.Name)
	out.State = direct.LazyPtr(in.State)
	return out
}

func ApigeeEnvgroupObservedState_ToApi(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroupObservedState) *api.GoogleCloudApigeeV1EnvironmentGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroup{}
	out.CreatedAt = direct.ValueOf(in.CreatedAt)
	out.LastModifiedAt = direct.ValueOf(in.LastModifiedAt)
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.ValueOf(in.State)
	return out
}

func ApigeeEnvgroupSpec_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroup) *krm.ApigeeEnvgroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeEnvgroupSpec{}
	out.ResourceID = &in.Name
	out.Hostnames = in.Hostnames
	return out
}

func ApigeeEnvgroup_ToApi(mapCtx *direct.MapContext, in *krm.ApigeeEnvgroup) *api.GoogleCloudApigeeV1EnvironmentGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1EnvironmentGroup{}
	out.Hostnames = in.Spec.Hostnames

	if in.Spec.ResourceID != nil {
		out.Name = *in.Spec.ResourceID
	} else {
		out.Name = in.Name
	}
	return out
}
