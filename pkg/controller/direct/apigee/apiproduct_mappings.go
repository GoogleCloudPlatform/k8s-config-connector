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

func ApigeeAPIProductObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductObservedState{}
	return out
}

func ApigeeAPIProductObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductObservedState) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	return out
}

func ApigeeAPIProductSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductSpec{}
	// Not fully implemented, as mock mappings are often stubbed out if not needed.
	return out
}

func ApigeeAPIProductSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductSpec) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	return out
}
