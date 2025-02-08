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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DiscoveryengineSitemapObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Sitemap) *krm.DiscoveryengineSitemapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSitemapObservedState{}
	// MISSING: URI
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSitemapObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSitemapObservedState) *pb.Sitemap {
	if in == nil {
		return nil
	}
	out := &pb.Sitemap{}
	// MISSING: URI
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSitemapSpec_FromProto(mapCtx *direct.MapContext, in *pb.Sitemap) *krm.DiscoveryengineSitemapSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSitemapSpec{}
	// MISSING: URI
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSitemapSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSitemapSpec) *pb.Sitemap {
	if in == nil {
		return nil
	}
	out := &pb.Sitemap{}
	// MISSING: URI
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func Sitemap_FromProto(mapCtx *direct.MapContext, in *pb.Sitemap) *krm.Sitemap {
	if in == nil {
		return nil
	}
	out := &krm.Sitemap{}
	out.URI = direct.LazyPtr(in.GetUri())
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func Sitemap_ToProto(mapCtx *direct.MapContext, in *krm.Sitemap) *pb.Sitemap {
	if in == nil {
		return nil
	}
	out := &pb.Sitemap{}
	if oneof := Sitemap_Uri_ToProto(mapCtx, in.URI); oneof != nil {
		out.Feed = oneof
	}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func SitemapObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Sitemap) *krm.SitemapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SitemapObservedState{}
	// MISSING: URI
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func SitemapObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SitemapObservedState) *pb.Sitemap {
	if in == nil {
		return nil
	}
	out := &pb.Sitemap{}
	// MISSING: URI
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
