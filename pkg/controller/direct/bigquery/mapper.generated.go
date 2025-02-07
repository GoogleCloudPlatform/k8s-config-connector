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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/dataexchange/apiv1beta1/dataexchangepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataExchange_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.DataExchange {
	if in == nil {
		return nil
	}
	out := &krm.DataExchange{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: ListingCount
	out.Icon = in.GetIcon()
	return out
}
func DataExchange_ToProto(mapCtx *direct.MapContext, in *krm.DataExchange) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: ListingCount
	out.Icon = in.Icon
	return out
}
func DataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.DataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataExchangeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.ListingCount = direct.LazyPtr(in.GetListingCount())
	// MISSING: Icon
	return out
}
func DataExchangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataExchangeObservedState) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.ListingCount = direct.ValueOf(in.ListingCount)
	// MISSING: Icon
	return out
}
