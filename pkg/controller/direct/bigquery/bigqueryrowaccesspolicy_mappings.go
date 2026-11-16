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

package bigquery

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/bigquery/v2"
)

func BigQueryRowAccessPolicySpec_FromAPI(mapCtx *direct.MapContext, in *api.RowAccessPolicy) *krm.BigQueryRowAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryRowAccessPolicySpec{}
	out.FilterPredicate = direct.LazyPtr(in.FilterPredicate)
	out.Grantees = in.Grantees
	return out
}

func BigQueryRowAccessPolicySpec_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryRowAccessPolicySpec) *api.RowAccessPolicy {
	if in == nil {
		return nil
	}
	out := &api.RowAccessPolicy{}
	out.FilterPredicate = direct.ValueOf(in.FilterPredicate)
	out.Grantees = in.Grantees
	return out
}

func BigQueryRowAccessPolicyStatus_FromAPI(mapCtx *direct.MapContext, in *api.RowAccessPolicy) *krm.BigQueryRowAccessPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryRowAccessPolicyStatus{}
	out.ObservedState = BigQueryRowAccessPolicyObservedState_FromAPI(mapCtx, in)
	return out
}

func BigQueryRowAccessPolicyStatus_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryRowAccessPolicyStatus) *api.RowAccessPolicy {
	if in == nil {
		return nil
	}
	out := &api.RowAccessPolicy{}
	if in.ObservedState != nil {
		out.CreationTime = direct.ValueOf(in.ObservedState.CreationTime)
		out.Etag = direct.ValueOf(in.ObservedState.Etag)
		out.LastModifiedTime = direct.ValueOf(in.ObservedState.LastModifiedTime)
	}
	return out
}

func BigQueryRowAccessPolicyObservedState_FromAPI(mapCtx *direct.MapContext, in *api.RowAccessPolicy) *krm.BigQueryRowAccessPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryRowAccessPolicyObservedState{}
	out.CreationTime = direct.LazyPtr(in.CreationTime)
	out.Etag = direct.LazyPtr(in.Etag)
	out.LastModifiedTime = direct.LazyPtr(in.LastModifiedTime)
	return out
}
