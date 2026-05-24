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

package bigqueryanalyticshubdataexchange

import (
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analyticshub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeObservedState{}
	out.ListingCount = direct.LazyPtr(in.GetListingCount())
	out.SharingEnvironmentConfig = SharingEnvironmentConfigObservedState_FromProto(mapCtx, in.GetSharingEnvironmentConfig())
	return out
}

func BigQueryAnalyticsHubDataExchangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	out.Icon = in.GetIcon()
	out.SharingEnvironmentConfig = SharingEnvironmentConfig_FromProto(mapCtx, in.GetSharingEnvironmentConfig())
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	out.LogLinkedDatasetQueryUserEmail = in.LogLinkedDatasetQueryUserEmail
	return out
}

func BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeSpec) *pb.DataExchange {
	if in == nil {
		return nil
	}

	out := &pb.DataExchange{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	out.Icon = in.Icon
	out.SharingEnvironmentConfig = SharingEnvironmentConfig_ToProto(mapCtx, in.SharingEnvironmentConfig)
	dtype := direct.Enum_ToProto[pb.DiscoveryType](mapCtx, in.DiscoveryType)
	out.DiscoveryType = &dtype
	out.LogLinkedDatasetQueryUserEmail = in.LogLinkedDatasetQueryUserEmail

	return out
}

func BigQueryAnalyticsHubDataExchangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeObservedState) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	if in.ListingCount != nil {
		out.ListingCount = *in.ListingCount
	}
	return out
}

func SharingEnvironmentConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig) *krm.SharingEnvironmentConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfigObservedState{}
	switch env := in.Environment.(type) {
	case *pb.SharingEnvironmentConfig_DefaultExchangeConfig_:
		out.DefaultExchangeConfig = &krm.SharingEnvironmentConfig_DefaultExchangeConfigObservedState{}
	case *pb.SharingEnvironmentConfig_DcrExchangeConfig_:
		if env.DcrExchangeConfig != nil {
			out.DcrExchangeConfig = &krm.SharingEnvironmentConfig_DcrExchangeConfigObservedState{
				SingleSelectedResourceSharingRestriction: env.DcrExchangeConfig.SingleSelectedResourceSharingRestriction,
				SingleLinkedDatasetPerCleanroom:          env.DcrExchangeConfig.SingleLinkedDatasetPerCleanroom,
			}
		}
	}
	return out
}

func SharingEnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig) *krm.SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfig{}
	switch env := in.Environment.(type) {
	case *pb.SharingEnvironmentConfig_DefaultExchangeConfig_:
		out.DefaultExchangeConfig = &krm.SharingEnvironmentConfig_DefaultExchangeConfig{}
	case *pb.SharingEnvironmentConfig_DcrExchangeConfig_:
		if env.DcrExchangeConfig != nil {
			out.DcrExchangeConfig = &krm.SharingEnvironmentConfig_DcrExchangeConfig{
				SingleSelectedResourceSharingRestriction: env.DcrExchangeConfig.SingleSelectedResourceSharingRestriction,
				SingleLinkedDatasetPerCleanroom:          env.DcrExchangeConfig.SingleLinkedDatasetPerCleanroom,
			}
		}
	}
	return out
}

func SharingEnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfig) *pb.SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig{}
	if in.DefaultExchangeConfig != nil {
		out.Environment = &pb.SharingEnvironmentConfig_DefaultExchangeConfig_{
			DefaultExchangeConfig: &pb.SharingEnvironmentConfig_DefaultExchangeConfig{},
		}
	} else if in.DcrExchangeConfig != nil {
		out.Environment = &pb.SharingEnvironmentConfig_DcrExchangeConfig_{
			DcrExchangeConfig: &pb.SharingEnvironmentConfig_DcrExchangeConfig{
				SingleSelectedResourceSharingRestriction: in.DcrExchangeConfig.SingleSelectedResourceSharingRestriction,
				SingleLinkedDatasetPerCleanroom:          in.DcrExchangeConfig.SingleLinkedDatasetPerCleanroom,
			},
		}
	}
	return out
}
