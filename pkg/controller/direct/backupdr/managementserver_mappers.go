// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ManagementURIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementURI) *krm.ManagementURIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementURIObservedState{}
	out.WebUI = direct.LazyPtr(in.GetWebUi())
	out.API = direct.LazyPtr(in.GetApi())
	return out
}
func ManagementURIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementURIObservedState) *pb.ManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.ManagementURI{}
	out.WebUi = direct.ValueOf(in.WebUI)
	out.Api = direct.ValueOf(in.API)
	return out
}
