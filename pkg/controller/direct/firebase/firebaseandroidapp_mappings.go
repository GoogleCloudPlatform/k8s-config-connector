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

func FirebaseAndroidAppSpec_FromAPI(mapCtx *direct.MapContext, in *pb.AndroidApp) *krm.FirebaseAndroidAppSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseAndroidAppSpec{}
	out.ApiKeyId = direct.LazyPtr(in.ApiKeyId)
	out.DisplayName = in.DisplayName
	out.PackageName = direct.LazyPtr(in.PackageName)
	out.Sha1Hashes = in.Sha1Hashes
	out.Sha256Hashes = in.Sha256Hashes
	return out
}

func FirebaseAndroidAppSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseAndroidAppSpec) *pb.AndroidApp {
	if in == nil {
		return nil
	}
	out := &pb.AndroidApp{}
	out.ApiKeyId = direct.ValueOf(in.ApiKeyId)
	out.DisplayName = in.DisplayName
	out.PackageName = direct.ValueOf(in.PackageName)
	out.Sha1Hashes = in.Sha1Hashes
	out.Sha256Hashes = in.Sha256Hashes
	return out
}

func FirebaseAndroidAppStatus_FromAPI(mapCtx *direct.MapContext, in *pb.AndroidApp) *krm.FirebaseAndroidAppStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseAndroidAppStatus{}
	out.AppId = direct.LazyPtr(in.AppId)
	out.Etag = direct.LazyPtr(in.Etag)
	out.Name = direct.LazyPtr(in.Name)
	return out
}

func FirebaseAndroidAppStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseAndroidAppStatus) *pb.AndroidApp {
	if in == nil {
		return nil
	}
	out := &pb.AndroidApp{}
	out.AppId = direct.ValueOf(in.AppId)
	out.Etag = direct.ValueOf(in.Etag)
	out.Name = direct.ValueOf(in.Name)
	return out
}
