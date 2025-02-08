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

package orchestration

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orchestration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func OrchestrationUserWorkloadsSecretObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsSecret) *krm.OrchestrationUserWorkloadsSecretObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationUserWorkloadsSecretObservedState{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsSecretObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationUserWorkloadsSecretObservedState) *pb.UserWorkloadsSecret {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsSecret{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsSecretSpec_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsSecret) *krm.OrchestrationUserWorkloadsSecretSpec {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationUserWorkloadsSecretSpec{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsSecretSpec_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationUserWorkloadsSecretSpec) *pb.UserWorkloadsSecret {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsSecret{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func UserWorkloadsSecret_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsSecret) *krm.UserWorkloadsSecret {
	if in == nil {
		return nil
	}
	out := &krm.UserWorkloadsSecret{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Data = in.Data
	return out
}
func UserWorkloadsSecret_ToProto(mapCtx *direct.MapContext, in *krm.UserWorkloadsSecret) *pb.UserWorkloadsSecret {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsSecret{}
	out.Name = direct.ValueOf(in.Name)
	out.Data = in.Data
	return out
}
