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

package accesscontextmanager

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

type AccessContextManagerServicePerimeterResourceAPI struct {
	PerimeterName string
	Resource      string
}

func AccessContextManagerServicePerimeterResourceSpec_FromAPI(mapCtx *direct.MapContext, in *AccessContextManagerServicePerimeterResourceAPI) *krm.AccessContextManagerServicePerimeterResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerServicePerimeterResourceSpec{}
	out.PerimeterNameRef.External = in.PerimeterName
	out.ResourceRef.External = in.Resource
	return out
}

func AccessContextManagerServicePerimeterResourceSpec_ToAPI(mapCtx *direct.MapContext, in *krm.AccessContextManagerServicePerimeterResourceSpec) *AccessContextManagerServicePerimeterResourceAPI {
	if in == nil {
		return nil
	}
	out := &AccessContextManagerServicePerimeterResourceAPI{}
	out.PerimeterName = in.PerimeterNameRef.External
	out.Resource = in.ResourceRef.External
	return out
}

func AccessContextManagerServicePerimeterResourceStatus_FromAPI(mapCtx *direct.MapContext, in *AccessContextManagerServicePerimeterResourceAPI) *krm.AccessContextManagerServicePerimeterResourceStatus {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerServicePerimeterResourceStatus{}
	return out
}

func AccessContextManagerServicePerimeterResourceStatus_ToAPI(mapCtx *direct.MapContext, in *krm.AccessContextManagerServicePerimeterResourceStatus) *AccessContextManagerServicePerimeterResourceAPI {
	if in == nil {
		return nil
	}
	out := &AccessContextManagerServicePerimeterResourceAPI{}
	return out
}
