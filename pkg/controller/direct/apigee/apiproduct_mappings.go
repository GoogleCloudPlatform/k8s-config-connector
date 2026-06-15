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

func GoogleCloudApigeeV1Attribute_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1Attribute) *api.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Attribute{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}

func GoogleCloudApigeeV1Attribute_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Attribute) *krm.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1Attribute{}
	out.Name = direct.LazyPtr(in.Name)
	out.Value = direct.LazyPtr(in.Value)
	return out
}

func GoogleCloudApigeeV1Attributes_ToAPI(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1Attribute) []*api.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1Attribute, len(in))
	for i, item := range in {
		out[i] = GoogleCloudApigeeV1Attribute_ToAPI(mapCtx, &item)
	}
	return out
}

func GoogleCloudApigeeV1Attributes_FromAPI(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1Attribute) []krm.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1Attribute, len(in))
	for i, item := range in {
		out[i] = *GoogleCloudApigeeV1Attribute_FromAPI(mapCtx, item)
	}
	return out
}

func ApigeeAPIProductObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductObservedState{}
	out.CreatedAt = direct.LazyPtr(in.CreatedAt)
	out.LastModifiedAt = direct.LazyPtr(in.LastModifiedAt)
	return out
}

func ApigeeAPIProductObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductObservedState) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	out.CreatedAt = direct.ValueOf(in.CreatedAt)
	out.LastModifiedAt = direct.ValueOf(in.LastModifiedAt)
	return out
}

func ApigeeAPIProductSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductSpec{}
	out.APIResources = in.ApiResources
	out.ApprovalType = direct.LazyPtr(in.ApprovalType)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromAPI(mapCtx, in.Attributes)
	out.Description = direct.LazyPtr(in.Description)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.Environments = in.Environments
	out.Proxies = in.Proxies
	out.Quota = direct.LazyPtr(in.Quota)
	out.QuotaCounterScope = direct.LazyPtr(in.QuotaCounterScope)
	out.QuotaInterval = direct.LazyPtr(in.QuotaInterval)
	out.QuotaTimeUnit = direct.LazyPtr(in.QuotaTimeUnit)
	out.Scopes = in.Scopes
	return out
}

func ApigeeAPIProductSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductSpec) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	out.ApiResources = in.APIResources
	out.ApprovalType = direct.ValueOf(in.ApprovalType)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToAPI(mapCtx, in.Attributes)
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Environments = in.Environments
	out.Proxies = in.Proxies
	out.Quota = direct.ValueOf(in.Quota)
	out.QuotaCounterScope = direct.ValueOf(in.QuotaCounterScope)
	out.QuotaInterval = direct.ValueOf(in.QuotaInterval)
	out.QuotaTimeUnit = direct.ValueOf(in.QuotaTimeUnit)
	out.Scopes = in.Scopes
	return out
}
