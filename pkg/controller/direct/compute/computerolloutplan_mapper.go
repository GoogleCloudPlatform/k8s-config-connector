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
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func RolloutPlanWaveValidationTimeBasedValidationMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.RolloutPlanWaveValidationTimeBasedValidationMetadata) *krmcomputev1alpha1.RolloutPlanWaveValidationTimeBasedValidationMetadata {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.RolloutPlanWaveValidationTimeBasedValidationMetadata{}
	out.WaitDuration = in.WaitDuration
	return out
}

func RolloutPlanWaveValidationTimeBasedValidationMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.RolloutPlanWaveValidationTimeBasedValidationMetadata) *pb.RolloutPlanWaveValidationTimeBasedValidationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.RolloutPlanWaveValidationTimeBasedValidationMetadata{}
	out.WaitDuration = in.WaitDuration
	return out
}
