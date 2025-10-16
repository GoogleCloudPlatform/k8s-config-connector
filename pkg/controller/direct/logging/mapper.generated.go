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
// krm.group: logging.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.logging.v2

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryDataset_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krm.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataset{}
	// MISSING: DatasetID
	return out
}
func BigQueryDataset_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataset) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	// MISSING: DatasetID
	return out
}
func BigQueryDatasetObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krm.BigQueryDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetObservedState{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDatasetObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDatasetObservedState) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func LoggingLinkSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krm.LoggingLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLinkSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: LifecycleState
	// MISSING: BigqueryDataset
	return out
}
func LoggingLinkSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLinkSpec) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: LifecycleState
	// MISSING: BigqueryDataset
	return out
}
