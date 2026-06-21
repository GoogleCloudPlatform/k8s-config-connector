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

// ComputeBackendServiceSignedURLKeySpec_v1alpha1_FromProto converts the Proto SignedUrlKey to the KRM Spec.
// We hand-code this function because the KeyValue field in KRM is a struct containing SecretKeyRef (sensitive),
// whereas in the GCP proto it is a plain string. Since we cannot retrieve a secret's value back from GCP,
// we only map KeyValue if it exists, and map the key name.
func ComputeBackendServiceSignedURLKeySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SignedUrlKey) *krm.ComputeBackendServiceSignedURLKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendServiceSignedURLKeySpec{}
	if in.KeyName != nil {
		out.ResourceID = in.KeyName
	}
	if in.KeyValue != nil {
		out.KeyValue.Value = in.KeyValue
	}
	return out
}

// ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto converts the KRM Spec to the Proto SignedUrlKey.
// We hand-code this because KeyValue is a struct with sensitive value/valueFrom fields in KRM,
// but is a plain string in the GCP proto. The actual secret resolution is handled during reconciliation
// rather than direct mapping.
func ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceSignedURLKeySpec) *pb.SignedUrlKey {
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
