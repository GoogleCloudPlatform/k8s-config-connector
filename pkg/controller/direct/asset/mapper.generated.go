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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AssetFeedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.AssetFeedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetFeedObservedState{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetFeedObservedState) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedSpec_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.AssetFeedSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssetFeedSpec{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssetFeedSpec) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func Feed_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.Feed {
	if in == nil {
		return nil
	}
	out := &krm.Feed{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AssetNames = in.AssetNames
	out.AssetTypes = in.AssetTypes
	out.ContentType = direct.Enum_FromProto(mapCtx, in.GetContentType())
	out.FeedOutputConfig = FeedOutputConfig_FromProto(mapCtx, in.GetFeedOutputConfig())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	out.RelationshipTypes = in.RelationshipTypes
	return out
}
func Feed_ToProto(mapCtx *direct.MapContext, in *krm.Feed) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	out.Name = direct.ValueOf(in.Name)
	out.AssetNames = in.AssetNames
	out.AssetTypes = in.AssetTypes
	out.ContentType = direct.Enum_ToProto[pb.ContentType](mapCtx, in.ContentType)
	out.FeedOutputConfig = FeedOutputConfig_ToProto(mapCtx, in.FeedOutputConfig)
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	out.RelationshipTypes = in.RelationshipTypes
	return out
}
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
