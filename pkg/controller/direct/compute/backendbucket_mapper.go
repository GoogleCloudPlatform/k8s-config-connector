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
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackendBucketCdnPolicyBypassCacheOnRequestHeaders_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader) *krm.BackendBucketCdnPolicyBypassCacheOnRequestHeaders {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicyBypassCacheOnRequestHeaders{}
	out.HeaderName = in.HeaderName
	return out
}

func BackendBucketCdnPolicyBypassCacheOnRequestHeaders_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendBucketCdnPolicyBypassCacheOnRequestHeaders) *pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyBypassCacheOnRequestHeader{}
	out.HeaderName = in.HeaderName
	return out
}

func BackendBucketCdnPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicy) *krm.BackendBucketCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_FromProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendBucketCdnPolicyBypassCacheOnRequestHeaders_v1beta1_FromProto)
	out.CacheKeyPolicy = BackendBucketCdnPolicyCacheKeyPolicy_v1beta1_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		out.ClientTtl = direct.LazyPtr(int(*in.ClientTtl))
	}
	if in.DefaultTtl != nil {
		out.DefaultTtl = direct.LazyPtr(int(*in.DefaultTtl))
	}
	if in.MaxTtl != nil {
		out.MaxTtl = direct.LazyPtr(int(*in.MaxTtl))
	}
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_FromProto(mapCtx, in.NegativeCachingPolicy, BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_FromProto)
	out.RequestCoalescing = in.RequestCoalescing
	if in.ServeWhileStale != nil {
		out.ServeWhileStale = direct.LazyPtr(int(*in.ServeWhileStale))
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		out.SignedUrlCacheMaxAgeSec = direct.LazyPtr(int(*in.SignedUrlCacheMaxAgeSec))
	}
	return out
}

func BackendBucketCdnPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendBucketCdnPolicy) *pb.BackendBucketCdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_ToProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendBucketCdnPolicyBypassCacheOnRequestHeaders_v1beta1_ToProto)
	out.CacheKeyPolicy = BackendBucketCdnPolicyCacheKeyPolicy_v1beta1_ToProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		out.ClientTtl = direct.LazyPtr(int32(*in.ClientTtl))
	}
	if in.DefaultTtl != nil {
		out.DefaultTtl = direct.LazyPtr(int32(*in.DefaultTtl))
	}
	if in.MaxTtl != nil {
		out.MaxTtl = direct.LazyPtr(int32(*in.MaxTtl))
	}
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_ToProto(mapCtx, in.NegativeCachingPolicy, BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_ToProto)
	out.RequestCoalescing = in.RequestCoalescing
	if in.ServeWhileStale != nil {
		out.ServeWhileStale = direct.LazyPtr(int32(*in.ServeWhileStale))
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		out.SignedUrlCacheMaxAgeSec = direct.LazyPtr(int64(*in.SignedUrlCacheMaxAgeSec))
	}
	return out
}

func BackendBucketCdnPolicyCacheKeyPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyCacheKeyPolicy) *krm.BackendBucketCdnPolicyCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicyCacheKeyPolicy{}
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

func BackendBucketCdnPolicyCacheKeyPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendBucketCdnPolicyCacheKeyPolicy) *pb.BackendBucketCdnPolicyCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyCacheKeyPolicy{}
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

func BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyNegativeCachingPolicy) *krm.BackendBucketCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicyNegativeCachingPolicy{}
	if in.Code != nil {
		out.Code = direct.LazyPtr(int(*in.Code))
	}
	if in.Ttl != nil {
		out.Ttl = direct.LazyPtr(int(*in.Ttl))
	}
	return out
}

func BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendBucketCdnPolicyNegativeCachingPolicy) *pb.BackendBucketCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyNegativeCachingPolicy{}
	if in.Code != nil {
		out.Code = direct.LazyPtr(int32(*in.Code))
	}
	if in.Ttl != nil {
		out.Ttl = direct.LazyPtr(int32(*in.Ttl))
	}
	return out
}

func ComputeBackendBucketSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendBucketSpec{}
	if in.BucketName != nil {
		out.BucketRef = &krm.BackendBucketBucketRef{External: *in.BucketName}
	}
	out.CdnPolicy = BackendBucketCdnPolicy_v1beta1_FromProto(mapCtx, in.GetCdnPolicy())
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
		out.BucketName = direct.LazyPtr(in.BucketRef.External)
	}
	out.CdnPolicy = BackendBucketCdnPolicy_v1beta1_ToProto(mapCtx, in.CdnPolicy)
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
