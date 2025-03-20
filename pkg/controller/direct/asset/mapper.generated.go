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

package asset

import (
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FeedOutputConfig_FromProto(mapCtx *direct.MapContext, in *pb.FeedOutputConfig) *krm.FeedOutputConfig {
	if in == nil {
		return nil
	}
	out := &krm.FeedOutputConfig{}
	out.PubsubDestination = PubsubDestination_FromProto(mapCtx, in.GetPubsubDestination())
	return out
}
func FeedOutputConfig_ToProto(mapCtx *direct.MapContext, in *krm.FeedOutputConfig) *pb.FeedOutputConfig {
	if in == nil {
		return nil
	}
	out := &pb.FeedOutputConfig{}
	if oneof := PubsubDestination_ToProto(mapCtx, in.PubsubDestination); oneof != nil {
		out.Destination = &pb.FeedOutputConfig_PubsubDestination{PubsubDestination: oneof}
	}
	return out
}
func PubsubDestination_FromProto(mapCtx *direct.MapContext, in *pb.PubsubDestination) *krm.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &krm.PubsubDestination{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	return out
}
func PubsubDestination_ToProto(mapCtx *direct.MapContext, in *krm.PubsubDestination) *pb.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &pb.PubsubDestination{}
	out.Topic = direct.ValueOf(in.Topic)
	return out
}
