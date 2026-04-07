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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.SignedUrlKey
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsecret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeBackendBucketSignedURLKeyFuzzer())
}

func computeBackendBucketSignedURLKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.SignedUrlKey{},
		computeBackendBucketSignedURLKeySpec_FromProto,
		computeBackendBucketSignedURLKeySpec_ToProto,
	)

	// key_name maps to spec.resourceID; key_value maps to spec.keyValue.value.
	f.SpecFields.Insert(".key_name")
	f.SpecFields.Insert(".key_value")

	return f
}

func computeBackendBucketSignedURLKeySpec_FromProto(_ *direct.MapContext, key *pb.SignedUrlKey) *krm.ComputeBackendBucketSignedURLKeySpec {
	if key == nil {
		return nil
	}
	spec := &krm.ComputeBackendBucketSignedURLKeySpec{}
	if v := key.GetKeyName(); v != "" {
		spec.ResourceID = &v
	}
	if v := key.GetKeyValue(); v != "" {
		spec.KeyValue = refsecret.Legacy{Value: &v}
	}
	return spec
}

func computeBackendBucketSignedURLKeySpec_ToProto(_ *direct.MapContext, spec *krm.ComputeBackendBucketSignedURLKeySpec) *pb.SignedUrlKey {
	if spec == nil {
		return nil
	}
	key := &pb.SignedUrlKey{}
	if spec.ResourceID != nil {
		keyName := *spec.ResourceID
		key.KeyName = &keyName
	}
	if spec.KeyValue.Value != nil {
		keyValue := *spec.KeyValue.Value
		key.KeyValue = &keyValue
	}
	return key
}
