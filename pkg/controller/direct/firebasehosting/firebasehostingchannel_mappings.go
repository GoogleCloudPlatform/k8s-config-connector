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

package firebasehosting

import (
	pb "google.golang.org/api/firebasehosting/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebasehosting/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseHostingChannelSpec_FromAPI(mapCtx *direct.MapContext, in *pb.Channel) *krm.FirebaseHostingChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseHostingChannelSpec{}
	out.ExpireTime = direct.LazyPtr(in.ExpireTime)
	if in.RetainedReleaseCount != 0 {
		out.RetainedReleaseCount = direct.LazyPtr(int(in.RetainedReleaseCount))
	}
	out.TTL = direct.LazyPtr(in.Ttl)
	return out
}

func FirebaseHostingChannelSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseHostingChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.ExpireTime = direct.ValueOf(in.ExpireTime)
	if in.RetainedReleaseCount != nil {
		out.RetainedReleaseCount = int64(*in.RetainedReleaseCount)
	}
	out.Ttl = direct.ValueOf(in.TTL)
	return out
}

func FirebaseHostingChannelStatus_FromAPI(mapCtx *direct.MapContext, in *pb.Channel) *krm.FirebaseHostingChannelStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseHostingChannelStatus{}
	out.Name = direct.LazyPtr(in.Name)
	return out
}

func FirebaseHostingChannelStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseHostingChannelStatus) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
