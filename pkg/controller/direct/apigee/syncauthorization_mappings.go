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

func ApigeeSyncAuthorizationSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1SyncAuthorization) *krm.ApigeeSyncAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeSyncAuthorizationSpec{}
	out.Identities = in.Identities
	return out
}

func ApigeeSyncAuthorizationSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeSyncAuthorizationSpec) *api.GoogleCloudApigeeV1SyncAuthorization {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1SyncAuthorization{}
	out.Identities = in.Identities
	return out
}

func ApigeeSyncAuthorizationStatus_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1SyncAuthorization) *krm.ApigeeSyncAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeSyncAuthorizationStatus{}
	out.Etag = direct.LazyPtr(in.Etag)
	return out
}

func ApigeeSyncAuthorizationStatus_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeSyncAuthorizationStatus) *api.GoogleCloudApigeeV1SyncAuthorization {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1SyncAuthorization{}
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
