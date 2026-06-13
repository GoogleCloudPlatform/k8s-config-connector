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

// +generated:mapper
// krm.group: pubsub.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.pubsub.v1

package pubsub

import (
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExpirationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ExpirationPolicy) *krm.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	return out
}

func ExpirationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ExpirationPolicy) *pb.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}

func PushConfig_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig) *krm.PushConfig {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig{}
	out.PushEndpoint = direct.LazyPtr(in.GetPushEndpoint())
	out.Attributes = in.Attributes
	out.OidcToken = PushConfigOidcToken_FromProto(mapCtx, in.GetOidcToken())
	out.NoWrapper = PushConfigNoWrapper_FromProto(mapCtx, in.GetNoWrapper())
	return out
}

func PushConfig_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig) *pb.PushConfig {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig{}
	out.PushEndpoint = direct.ValueOf(in.PushEndpoint)
	out.Attributes = in.Attributes
	if in.OidcToken != nil {
		out.AuthenticationMethod = &pb.PushConfig_OidcToken_{OidcToken: PushConfigOidcToken_ToProto(mapCtx, in.OidcToken)}
	}
	if oneof := PushConfigNoWrapper_ToProto(mapCtx, in.NoWrapper); oneof != nil {
		out.Wrapper = &pb.PushConfig_NoWrapper_{NoWrapper: oneof}
	}
	return out
}
