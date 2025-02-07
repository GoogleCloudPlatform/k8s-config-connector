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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/storage/apiv1beta1/storagepb"
)
func BigqueryStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.BigqueryStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryStreamObservedState{}
	// MISSING: Name
	return out
}
func BigqueryStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryStreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	return out
}
func BigqueryStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.BigqueryStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryStreamSpec{}
	// MISSING: Name
	return out
}
func BigqueryStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryStreamSpec) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	return out
}
func Stream_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.Stream {
	if in == nil {
		return nil
	}
	out := &krm.Stream{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Stream_ToProto(mapCtx *direct.MapContext, in *krm.Stream) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
