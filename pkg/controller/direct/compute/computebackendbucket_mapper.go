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

func ComputeBackendBucketSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendBucket) *krm.ComputeBackendBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendBucketSpec{}
	out.BucketRef = StorageBucketRef_FromProto(mapCtx, in.GetBucketName())
	out.CDNPolicy = BackendBucketCDNPolicy_v1beta1_FromProto(mapCtx, in.GetCdnPolicy())
	out.CompressionMode = in.CompressionMode
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	out.EdgeSecurityPolicy = in.EdgeSecurityPolicy
	out.EnableCDN = in.EnableCdn
	out.ResourceID = direct.LazyPtr(in.GetName())
	return out
}

func ComputeBackendBucketSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketSpec) *pb.BackendBucket {
	if in == nil {
		return nil
	}
	out := &pb.BackendBucket{}
	out.BucketName = StorageBucketRef_ToProto(mapCtx, in.BucketRef)
	out.CdnPolicy = BackendBucketCDNPolicy_v1beta1_ToProto(mapCtx, in.CDNPolicy)
	out.CompressionMode = in.CompressionMode
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	out.EdgeSecurityPolicy = in.EdgeSecurityPolicy
	out.EnableCdn = in.EnableCDN
	return out
}

func StorageBucketRef_FromProto(mapCtx *direct.MapContext, in string) *storagev1beta1.StorageBucketRef {
	if in == "" {
		return nil
	}
	return &storagev1beta1.StorageBucketRef{
		External: in,
	}
}

func StorageBucketRef_ToProto(mapCtx *direct.MapContext, in *storagev1beta1.StorageBucketRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}
