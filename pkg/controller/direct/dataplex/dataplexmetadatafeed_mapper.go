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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexMetadataFeedSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexMetadataFeedSpec) *pb.MetadataFeed {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFeed{}
	// MISSING: Name
	out.Scope = MetadataFeed_Scope_ToProto(mapCtx, in.Scope)
	out.Filters = MetadataFeed_Filters_ToProto(mapCtx, in.Filters)
	out.Labels = in.Labels
	if in.PubsubTopicRef != nil {
		out.Endpoint = &pb.MetadataFeed_PubsubTopic{
			PubsubTopic: in.PubsubTopicRef.External,
		}
	}
	return out
}
