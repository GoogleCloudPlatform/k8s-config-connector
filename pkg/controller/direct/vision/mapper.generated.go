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

package vision

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vision/apiv1p4beta1/visionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vision/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ProductSet_FromProto(mapCtx *direct.MapContext, in *pb.ProductSet) *krm.ProductSet {
	if in == nil {
		return nil
	}
	out := &krm.ProductSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: IndexTime
	// MISSING: IndexError
	return out
}
func ProductSet_ToProto(mapCtx *direct.MapContext, in *krm.ProductSet) *pb.ProductSet {
	if in == nil {
		return nil
	}
	out := &pb.ProductSet{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: IndexTime
	// MISSING: IndexError
	return out
}
func ProductSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProductSet) *krm.ProductSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProductSetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.IndexTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIndexTime())
	out.IndexError = Status_FromProto(mapCtx, in.GetIndexError())
	return out
}
func ProductSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProductSetObservedState) *pb.ProductSet {
	if in == nil {
		return nil
	}
	out := &pb.ProductSet{}
	// MISSING: Name
	// MISSING: DisplayName
	out.IndexTime = direct.StringTimestamp_ToProto(mapCtx, in.IndexTime)
	out.IndexError = Status_ToProto(mapCtx, in.IndexError)
	return out
}
