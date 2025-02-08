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

package securitycenter

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/settings/apiv1beta1/settingspb"
)
func SecuritycenterServiceAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.SecuritycenterServiceAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterServiceAccountObservedState{}
	// MISSING: Name
	// MISSING: ServiceAccount
	return out
}
func SecuritycenterServiceAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterServiceAccountObservedState) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	// MISSING: Name
	// MISSING: ServiceAccount
	return out
}
func SecuritycenterServiceAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.SecuritycenterServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterServiceAccountSpec{}
	// MISSING: Name
	// MISSING: ServiceAccount
	return out
}
func SecuritycenterServiceAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterServiceAccountSpec) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	// MISSING: Name
	// MISSING: ServiceAccount
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Name = direct.ValueOf(in.Name)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	return out
}
