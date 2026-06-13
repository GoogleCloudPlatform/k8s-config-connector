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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// Why we need manual mapper: KeyValue is a custom structure with value/valueFrom in KRM,
// but maps to a simple string KeyValue in the proto. KeyName maps to the ResourceID/name.
// Since the KRM spec and proto schemas are structurally different, we handcode these.

func ComputeBackendBucketSignedURLKeySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SignedUrlKey) *krm.ComputeBackendBucketSignedURLKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendBucketSignedURLKeySpec{}
	if in.KeyName != nil {
		out.ResourceID = in.KeyName
	}
	if in.KeyValue != nil {
		out.KeyValue = krm.BackendbucketsignedurlkeyKeyValue{
			Value: in.KeyValue,
		}
	}
	return out
}

func ComputeBackendBucketSignedURLKeySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendBucketSignedURLKeySpec) *pb.SignedUrlKey {
	if in == nil {
		return nil
	}
	out := &pb.SignedUrlKey{}
	if in.ResourceID != nil {
		out.KeyName = in.ResourceID
	}
	if in.KeyValue.Value != nil {
		out.KeyValue = in.KeyValue.Value
	}
	return out
}
