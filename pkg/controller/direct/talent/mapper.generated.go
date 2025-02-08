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

package talent

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/talent/apiv4/talentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/talent/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func TalentTenantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Tenant) *krm.TalentTenantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TalentTenantObservedState{}
	// MISSING: Name
	// MISSING: ExternalID
	return out
}
func TalentTenantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TalentTenantObservedState) *pb.Tenant {
	if in == nil {
		return nil
	}
	out := &pb.Tenant{}
	// MISSING: Name
	// MISSING: ExternalID
	return out
}
func TalentTenantSpec_FromProto(mapCtx *direct.MapContext, in *pb.Tenant) *krm.TalentTenantSpec {
	if in == nil {
		return nil
	}
	out := &krm.TalentTenantSpec{}
	// MISSING: Name
	// MISSING: ExternalID
	return out
}
func TalentTenantSpec_ToProto(mapCtx *direct.MapContext, in *krm.TalentTenantSpec) *pb.Tenant {
	if in == nil {
		return nil
	}
	out := &pb.Tenant{}
	// MISSING: Name
	// MISSING: ExternalID
	return out
}
func Tenant_FromProto(mapCtx *direct.MapContext, in *pb.Tenant) *krm.Tenant {
	if in == nil {
		return nil
	}
	out := &krm.Tenant{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExternalID = direct.LazyPtr(in.GetExternalId())
	return out
}
func Tenant_ToProto(mapCtx *direct.MapContext, in *krm.Tenant) *pb.Tenant {
	if in == nil {
		return nil
	}
	out := &pb.Tenant{}
	out.Name = direct.ValueOf(in.Name)
	out.ExternalId = direct.ValueOf(in.ExternalID)
	return out
}
