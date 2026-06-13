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

package compute

import (
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func getBucketName(external string) string {
	if idx := strings.Index(external, "/buckets/"); idx != -1 {
		return external[idx+9:]
	}
	if idx := strings.Index(external, "/b/"); idx != -1 {
		return external[idx+3:]
	}
	parts := strings.Split(external, "/")
	return parts[len(parts)-1]
}

func BackendbucketBypassCacheOnRequestHeaders_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader) *krm.BackendbucketBypassCacheOnRequestHeaders {
	if in == nil {
		return nil
	}
	out := &krm.BackendbucketBypassCacheOnRequestHeaders{}
	out.HeaderName = in.HeaderName
	return out
}

func BackendbucketBypassCacheOnRequestHeaders_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendbucketBypassCacheOnRequestHeaders) *pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader{}
	out.HeaderName = in.HeaderName
	return out
}

func BackendbucketCacheKeyPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyCacheKeyPolicy) *krm.BackendbucketCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendbucketCacheKeyPolicy{}
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

func BackendbucketCacheKeyPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendbucketCacheKeyPolicy) *pb.BackendBucketCdnPolicyCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyCacheKeyPolicy{}
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

func BackendbucketNegativeCachingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyNegativeCachingPolicy) *krm.BackendbucketNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendbucketNegativeCachingPolicy{}
	out.Code = direct.PtrInt32ToPtrInt64(in.Code)
	out.Ttl = direct.PtrInt32ToPtrInt64(in.Ttl)
	return out
}

func BackendbucketNegativeCachingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendbucketNegativeCachingPolicy) *pb.BackendBucketCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyNegativeCachingPolicy{}
	out.Code = direct.PtrInt64ToPtrInt32(in.Code)
	out.Ttl = direct.PtrInt64ToPtrInt32(in.Ttl)
	return out
}

func BackendbucketCdnPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicy) *krm.BackendbucketCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendbucketCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_FromProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendbucketBypassCacheOnRequestHeaders_v1beta1_FromProto)
	out.CacheKeyPolicy = BackendbucketCacheKeyPolicy_v1beta1_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.CacheMode = in.CacheMode
	out.ClientTtl = direct.PtrInt32ToPtrInt64(in.ClientTtl)
	out.DefaultTtl = direct.PtrInt32ToPtrInt64(in.DefaultTtl)
	out.MaxTtl = direct.PtrInt32ToPtrInt64(in.MaxTtl)
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_FromProto(mapCtx, in.NegativeCachingPolicy, BackendbucketNegativeCachingPolicy_v1beta1_FromProto)
	out.RequestCoalescing = in.RequestCoalescing
	out.ServeWhileStale = direct.PtrInt32ToPtrInt64(in.ServeWhileStale)
	out.SignedUrlCacheMaxAgeSec = in.SignedUrlCacheMaxAgeSec
	return out
}

func BackendbucketCdnPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendbucketCdnPolicy) *pb.BackendBucketCdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_ToProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendbucketBypassCacheOnRequestHeaders_v1beta1_ToProto)
	out.CacheKeyPolicy = BackendbucketCacheKeyPolicy_v1beta1_ToProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = in.CacheMode
	out.ClientTtl = direct.PtrInt64ToPtrInt32(in.ClientTtl)
	out.DefaultTtl = direct.PtrInt64ToPtrInt32(in.DefaultTtl)
	out.MaxTtl = direct.PtrInt64ToPtrInt32(in.MaxTtl)
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_ToProto(mapCtx, in.NegativeCachingPolicy, BackendbucketNegativeCachingPolicy_v1beta1_ToProto)
	out.RequestCoalescing = in.RequestCoalescing
	out.ServeWhileStale = direct.PtrInt64ToPtrInt32(in.ServeWhileStale)
	out.SignedUrlCacheMaxAgeSec = in.SignedUrlCacheMaxAgeSec
	return out
}

func ComputeBackendBucketSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendBucketSpec{}
	if in.BucketName != nil {
		out.BucketRef = &storagev1beta1.StorageBucketRef{
			External: direct.ValueOf(in.BucketName),
		}
	}
	out.CdnPolicy = BackendbucketCdnPolicy_v1beta1_FromProto(mapCtx, in.CdnPolicy)
	out.CompressionMode = in.CompressionMode
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	out.EdgeSecurityPolicy = in.EdgeSecurityPolicy
	out.EnableCdn = in.EnableCdn
	return out
}

func ComputeBackendBucketSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketSpec) *pb.BackendBucket {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucket{}
	if in.BucketRef != nil {
		out.BucketName = direct.LazyPtr(getBucketName(in.BucketRef.External))
	}
	out.CdnPolicy = BackendbucketCdnPolicy_v1beta1_ToProto(mapCtx, in.CdnPolicy)
	out.CompressionMode = in.CompressionMode
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	out.EdgeSecurityPolicy = in.EdgeSecurityPolicy
	out.EnableCdn = in.EnableCdn
	return out
}

func ComputeBackendBucketStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendBucketStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

func ComputeBackendBucketStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketStatus) *pb.BackendBucket {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucket{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}
