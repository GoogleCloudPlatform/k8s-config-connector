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

package firebase

import (
	pb "google.golang.org/api/firebase/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseWebAppSpec_FromAPI(mapCtx *direct.MapContext, in *pb.WebApp) *krm.FirebaseWebAppSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseWebAppSpec{}
	out.ApiKeyId = direct.LazyPtr(in.ApiKeyId)
	out.DisplayName = in.DisplayName
	return out
}

func FirebaseWebAppSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseWebAppSpec) *pb.WebApp {
	if in == nil {
		return nil
	}
	out := &pb.WebApp{}
	out.ApiKeyId = direct.ValueOf(in.ApiKeyId)
	out.DisplayName = in.DisplayName
	return out
}

func FirebaseWebAppStatus_FromAPI(mapCtx *direct.MapContext, in *pb.WebApp) *krm.FirebaseWebAppStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseWebAppStatus{}
	out.AppId = direct.LazyPtr(in.AppId)
	out.AppUrls = in.AppUrls
	out.Name = direct.LazyPtr(in.Name)
	return out
}

func FirebaseWebAppStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseWebAppStatus) *pb.WebApp {
	if in == nil {
		return nil
	}
	out := &pb.WebApp{}
	out.AppId = direct.ValueOf(in.AppId)
	out.AppUrls = in.AppUrls
	out.Name = direct.ValueOf(in.Name)
	return out
}
