// Copyright 2024 Google LLC
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

package bigqueryanalyticshub

import (
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeObservedState{}
	out.ListingCount = direct.LazyPtr(int64(in.GetListingCount()))
	// MISSING: SharingEnvironmentConfig // not yet
	return out
}

func BigQueryAnalyticsHubDataExchangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// s := string(in.GetIcon())
	// out.Icon = &s // not yet
	// MISSING: SharingEnvironmentConfig // not yet
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeSpec) *pb.DataExchange {
	if in == nil {
		return nil
	}

	out := &pb.DataExchange{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// out.Icon = []byte(direct.ValueOf(in.Icon)) // not yet
	// MISSING: SharingEnvironmentConfig // not yet
	dtype := direct.Enum_ToProto[pb.DiscoveryType](mapCtx, in.DiscoveryType)
	out.DiscoveryType = &dtype

	return out
}
