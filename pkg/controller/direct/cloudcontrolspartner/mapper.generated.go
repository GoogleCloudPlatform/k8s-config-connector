// Copyright 2025 Google LLC
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

package cloudcontrolspartner

import (
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AccessReason_FromProto(mapCtx *direct.MapContext, in *pb.AccessReason) *krm.AccessReason {
	if in == nil {
		return nil
	}
	out := &krm.AccessReason{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Detail = direct.LazyPtr(in.GetDetail())
	return out
}
func AccessReason_ToProto(mapCtx *direct.MapContext, in *krm.AccessReason) *pb.AccessReason {
	if in == nil {
		return nil
	}
	out := &pb.AccessReason{}
	out.Type = direct.Enum_ToProto[pb.AccessReason_Type](mapCtx, in.Type)
	out.Detail = direct.ValueOf(in.Detail)
	return out
}
