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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1beta/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
)
func PartnerPermissions_FromProto(mapCtx *direct.MapContext, in *pb.PartnerPermissions) *krm.PartnerPermissions {
	if in == nil {
		return nil
	}
	out := &krm.PartnerPermissions{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PartnerPermissions = direct.EnumSlice_FromProto(mapCtx, in.PartnerPermissions)
	return out
}
func PartnerPermissions_ToProto(mapCtx *direct.MapContext, in *krm.PartnerPermissions) *pb.PartnerPermissions {
	if in == nil {
		return nil
	}
	out := &pb.PartnerPermissions{}
	out.Name = direct.ValueOf(in.Name)
	out.PartnerPermissions = direct.EnumSlice_ToProto[pb.PartnerPermissions_Permission](mapCtx, in.PartnerPermissions)
	return out
}
