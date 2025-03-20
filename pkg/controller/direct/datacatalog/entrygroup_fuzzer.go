// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.datacatalog.v1.EntryGroup
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogEntryGroupFuzzer())
}

func dataCatalogEntryGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EntryGroup{},
		DataCatalogEntryGroupSpec_FromProto, DataCatalogEntryGroupSpec_ToProto,
		DataCatalogEntryGroupObservedState_FromProto, DataCatalogEntryGroupObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".data_catalog_timestamps")

	f.UnimplementedFields.Insert(".name")                    // special field
	f.UnimplementedFields.Insert(".transferred_to_dataplex") // read only

	return f
}

func DataCatalogEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func DataCatalogEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func DataCatalogEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupObservedState{}
	out.DataCatalogTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetDataCatalogTimestamps())
	return out
}

func DataCatalogEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.DataCatalogTimestamps = SystemTimestamps_ToProto(mapCtx, in.DataCatalogTimestamps)
	return out
}

func SystemTimestamps_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}

func SystemTimestamps_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
