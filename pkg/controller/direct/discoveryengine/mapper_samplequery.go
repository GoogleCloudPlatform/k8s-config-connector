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

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiscoveryEngineSampleQuerySpec_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.DiscoveryEngineSampleQuerySpec {
	return DiscoveryEngineSampleQuerySpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSampleQuerySpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineSampleQuerySpec) *pb.SampleQuery {
	return DiscoveryEngineSampleQuerySpec_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineSampleQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.DiscoveryEngineSampleQueryObservedState {
	return DiscoveryEngineSampleQueryObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSampleQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineSampleQueryObservedState) *pb.SampleQuery {
	return DiscoveryEngineSampleQueryObservedState_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineSampleQuerySetSpec_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.DiscoveryEngineSampleQuerySetSpec {
	return DiscoveryEngineSampleQuerySetSpec_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSampleQuerySetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineSampleQuerySetSpec) *pb.SampleQuerySet {
	return DiscoveryEngineSampleQuerySetSpec_v1alpha1_ToProto(mapCtx, in)
}

func DiscoveryEngineSampleQuerySetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.DiscoveryEngineSampleQuerySetObservedState {
	return DiscoveryEngineSampleQuerySetObservedState_v1alpha1_FromProto(mapCtx, in)
}

func DiscoveryEngineSampleQuerySetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineSampleQuerySetObservedState) *pb.SampleQuerySet {
	return DiscoveryEngineSampleQuerySetObservedState_v1alpha1_ToProto(mapCtx, in)
}
