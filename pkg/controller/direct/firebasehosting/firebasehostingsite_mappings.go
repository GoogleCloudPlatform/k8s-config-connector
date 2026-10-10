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

package firebasehosting

import (
	pb "google.golang.org/api/firebasehosting/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebasehosting/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseHostingSiteSpec_FromAPI(mapCtx *direct.MapContext, in *pb.Site) *krm.FirebaseHostingSiteSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseHostingSiteSpec{}
	out.AppID = direct.LazyPtr(in.AppId)
	return out
}

func FirebaseHostingSiteSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseHostingSiteSpec) *pb.Site {
	if in == nil {
		return nil
	}
	out := &pb.Site{}
	out.AppId = direct.ValueOf(in.AppID)
	return out
}

func FirebaseHostingSiteStatus_FromAPI(mapCtx *direct.MapContext, in *pb.Site) *krm.FirebaseHostingSiteStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseHostingSiteStatus{}
	out.DefaultURL = direct.LazyPtr(in.DefaultUrl)
	out.Name = direct.LazyPtr(in.Name)
	return out
}

func FirebaseHostingSiteStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseHostingSiteStatus) *pb.Site {
	if in == nil {
		return nil
	}
	out := &pb.Site{}
	out.DefaultUrl = direct.ValueOf(in.DefaultURL)
	out.Name = direct.ValueOf(in.Name)
	return out
}
