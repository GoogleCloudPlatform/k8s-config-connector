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

package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ProbingDetails_SingleEdgeResponseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails_SingleEdgeResponse) *krm.ProbingDetails_SingleEdgeResponseObservedState {
	if in == nil {
		return nil
	}
	mapCtx.NotImplemented()
	return nil
}

func ProbingDetails_SingleEdgeResponseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProbingDetails_SingleEdgeResponseObservedState) *pb.ProbingDetails_SingleEdgeResponse {
	if in == nil {
		return nil
	}
	mapCtx.NotImplemented()
	return nil
}
