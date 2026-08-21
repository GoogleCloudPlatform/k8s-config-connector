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

// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	alloydb "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, input *pb.Instance_ObservabilityInstanceConfig) *alloydb.Instance_ObservabilityInstanceConfigObservedState {
	if input == nil {
		return nil
	}
	out := &alloydb.Instance_ObservabilityInstanceConfigObservedState{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = input.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	return out
}

func Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, input *alloydb.Instance_ObservabilityInstanceConfigObservedState) *pb.Instance_ObservabilityInstanceConfig {
	if input == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = input.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	return out
}
