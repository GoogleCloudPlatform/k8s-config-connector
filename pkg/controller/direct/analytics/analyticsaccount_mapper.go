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

package analytics

import (
	pb "cloud.google.com/go/analytics/admin/apiv1alpha/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AnalyticsAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.Account) *v1alpha1.AnalyticsAccountSpec {
	if in == nil {
		return nil
	}
	out := &v1alpha1.AnalyticsAccountSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	return out
}

func AnalyticsAccountSpec_ToProto(mapCtx *direct.MapContext, in *v1alpha1.AnalyticsAccountSpec) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.RegionCode = direct.ValueOf(in.RegionCode)
	return out
}

func AnalyticsAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Account) *v1alpha1.AnalyticsAccountObservedState {
	if in == nil {
		return nil
	}
	out := &v1alpha1.AnalyticsAccountObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Deleted = direct.LazyPtr(in.GetDeleted())
	out.GMPOrganization = direct.LazyPtr(in.GetGmpOrganization())
	return out
}

func AnalyticsAccountObservedState_ToProto(mapCtx *direct.MapContext, in *v1alpha1.AnalyticsAccountObservedState) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}

	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Deleted = direct.ValueOf(in.Deleted)
	out.GmpOrganization = direct.ValueOf(in.GMPOrganization)
	return out
}

func AnalyticsAccountObservedState_FromAccountTicketID(mapCtx *direct.MapContext, in string) *v1alpha1.AnalyticsAccountObservedState {
	if in == "" {
		return nil
	}
	out := &v1alpha1.AnalyticsAccountObservedState{}
	out.AccountTicketID = direct.LazyPtr(in)
	return out
}
