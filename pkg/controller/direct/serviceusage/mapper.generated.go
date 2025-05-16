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

// +generated:mapper
// krm.group: serviceusage.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.api.serviceusage.v1beta1

package serviceusage

import (
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1beta1"
)

func ServiceIdentityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceIdentity) *krmv1beta1.ServiceIdentityObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ServiceIdentityObservedState{}
	out.Email = direct.LazyPtr(in.GetEmail())
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	return out
}
func ServiceIdentityObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ServiceIdentityObservedState) *pb.ServiceIdentity {
	if in == nil {
		return nil
	}
	out := &pb.ServiceIdentity{}
	out.Email = direct.ValueOf(in.Email)
	out.UniqueId = direct.ValueOf(in.UniqueID)
	return out
}
func ServiceIdentitySpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceIdentity) *krmv1beta1.ServiceIdentitySpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ServiceIdentitySpec{}
	return out
}
func ServiceIdentitySpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ServiceIdentitySpec) *pb.ServiceIdentity {
	if in == nil {
		return nil
	}
	out := &pb.ServiceIdentity{}
	return out
}
