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
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicyNegativeCachingPolicy) *krm.BackendBucketCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicyNegativeCachingPolicy{}
	if in.Code != nil {
		val := int(*in.Code)
		out.Code = &val
	}
	if in.Ttl != nil {
		val := int(*in.Ttl)
		out.Ttl = &val
	}
	return out
}

func BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendBucketCdnPolicyNegativeCachingPolicy) *pb.BackendBucketCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucketCdnPolicyNegativeCachingPolicy{}
	if in.Code != nil {
		val := int32(*in.Code)
		out.Code = &val
	}
	if in.Ttl != nil {
		val := int32(*in.Ttl)
		out.Ttl = &val
	}
	return out
}

func BackendBucketCdnPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucketCdnPolicy) *krm.BackendBucketCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendBucketCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_FromProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendBucketCdnPolicyBypassCacheOnRequestHeaders_v1beta1_FromProto)
	out.CacheKeyPolicy = BackendBucketCdnPolicyCacheKeyPolicy_v1beta1_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.NegativeCachingPolicy = direct.Slice_FromProto(mapCtx, in.NegativeCachingPolicy, BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_FromProto)
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		val := int(*in.ClientTtl)
		out.ClientTtl = &val
	}
	if in.DefaultTtl != nil {
		val := int(*in.DefaultTtl)
		out.DefaultTtl = &val
	}
	if in.MaxTtl != nil {
		val := int(*in.MaxTtl)
		out.MaxTtl = &val
	}
	out.NegativeCaching = in.NegativeCaching
	out.RequestCoalescing = in.RequestCoalescing
	if in.ServeWhileStale != nil {
		val := int(*in.ServeWhileStale)
		out.ServeWhileStale = &val
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		val := int(*in.SignedUrlCacheMaxAgeSec) // it is *int64 in proto
		out.SignedUrlCacheMaxAgeSec = &val
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
	out.NegativeCachingPolicy = direct.Slice_ToProto(mapCtx, in.NegativeCachingPolicy, BackendBucketCdnPolicyNegativeCachingPolicy_v1beta1_ToProto)
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		val := int32(*in.ClientTtl)
		out.ClientTtl = &val
	}
	if in.DefaultTtl != nil {
		val := int32(*in.DefaultTtl)
		out.DefaultTtl = &val
	}
	if in.MaxTtl != nil {
		val := int32(*in.MaxTtl)
		out.MaxTtl = &val
	}
	out.NegativeCaching = in.NegativeCaching
	out.RequestCoalescing = in.RequestCoalescing
	if in.ServeWhileStale != nil {
		val := int32(*in.ServeWhileStale)
		out.ServeWhileStale = &val
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		val := int64(*in.SignedUrlCacheMaxAgeSec)
		out.SignedUrlCacheMaxAgeSec = &val
	}
	return out
}

func ComputeBackendBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketSpec) *pb.BackendBucket {
	if in == nil {
		return nil
	}
	out := ComputeBackendBucketSpec_v1beta1_ToProto(mapCtx, in)

	out.CdnPolicy = BackendBucketCdnPolicy_v1beta1_ToProto(mapCtx, in.CdnPolicy)
	out.EnableCdn = in.EnableCdn

	if in.BucketRef != nil {
		out.BucketName = direct.LazyPtr(in.BucketRef.External)
	}

	return out
}

func ComputeBackendBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketSpec {
	if in == nil {
		return nil
	}
	out := ComputeBackendBucketSpec_v1beta1_FromProto(mapCtx, in)

	out.CdnPolicy = BackendBucketCdnPolicy_v1beta1_FromProto(mapCtx, in.CdnPolicy)
	out.EnableCdn = in.EnableCdn

	if in.BucketName != nil {
		out.BucketRef = &storagev1beta1.StorageBucketRef{
			External: *in.BucketName,
		}
	}

	return out
}

func ComputeBackendBucketStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketStatus) *pb.BackendBucket {
	if in == nil {
		return nil
	}
	out := ComputeBackendBucketStatus_v1beta1_ToProto(mapCtx, in)
	return out
}

func ComputeBackendBucketStatus_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketStatus {
	if in == nil {
		return nil
	}
	out := ComputeBackendBucketStatus_v1beta1_FromProto(mapCtx, in)
	return out
}
