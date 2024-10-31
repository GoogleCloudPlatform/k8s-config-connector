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

package logging

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "/home/tylerreid/dev/waze/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigQueryDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krm.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataset{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDataset_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataset) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func Link_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krm.Link {
	if in == nil {
		return nil
	}
	out := &krm.Link{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	out.BigqueryDataset = BigQueryDataset_FromProto(mapCtx, in.GetBigqueryDataset())
	return out
}
func Link_ToProto(mapCtx *direct.MapContext, in *krm.Link) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LifecycleState = direct.Enum_ToProto[pb.LifecycleState](mapCtx, in.LifecycleState)
	out.BigqueryDataset = BigQueryDataset_ToProto(mapCtx, in.BigqueryDataset)
	return out
}
func LoggingLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krm.LoggingLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLinkSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: LifecycleState
	// MISSING: BigqueryDataset
	return out
}
func LoggingLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLinkSpec) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: LifecycleState
	// MISSING: BigqueryDataset
	return out
}
