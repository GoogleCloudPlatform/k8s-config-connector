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

package dataplex

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
)
func DataplexPartitionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Partition) *krm.DataplexPartitionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexPartitionObservedState{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
func DataplexPartitionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexPartitionObservedState) *pb.Partition {
	if in == nil {
		return nil
	}
	out := &pb.Partition{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
func DataplexPartitionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Partition) *krm.DataplexPartitionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexPartitionSpec{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
func DataplexPartitionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexPartitionSpec) *pb.Partition {
	if in == nil {
		return nil
	}
	out := &pb.Partition{}
	// MISSING: Name
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
func Partition_FromProto(mapCtx *direct.MapContext, in *pb.Partition) *krm.Partition {
	if in == nil {
		return nil
	}
	out := &krm.Partition{}
	// MISSING: Name
	out.Values = in.Values
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Partition_ToProto(mapCtx *direct.MapContext, in *krm.Partition) *pb.Partition {
	if in == nil {
		return nil
	}
	out := &pb.Partition{}
	// MISSING: Name
	out.Values = in.Values
	out.Location = direct.ValueOf(in.Location)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func PartitionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Partition) *krm.PartitionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PartitionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
func PartitionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PartitionObservedState) *pb.Partition {
	if in == nil {
		return nil
	}
	out := &pb.Partition{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Values
	// MISSING: Location
	// MISSING: Etag
	return out
}
