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

// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataplex.v1

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krmv1alpha1.DataplexEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexEntryGroupObservedState{}
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.TransferStatus = direct.Enum_FromProto(mapCtx, in.GetTransferStatus())
	return out
}
func DataplexEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.TransferStatus = direct.Enum_ToProto[pb.TransferStatus](mapCtx, in.TransferStatus)
	return out
}
func DataplexEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krmv1alpha1.DataplexEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexEntryGroupSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func DataplexEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
