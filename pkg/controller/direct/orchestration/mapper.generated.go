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
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orchestration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func OrchestrationUserWorkloadsConfigMapObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsConfigMap) *krm.OrchestrationUserWorkloadsConfigMapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationUserWorkloadsConfigMapObservedState{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsConfigMapObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationUserWorkloadsConfigMapObservedState) *pb.UserWorkloadsConfigMap {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsConfigMap{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsConfigMapSpec_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsConfigMap) *krm.OrchestrationUserWorkloadsConfigMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationUserWorkloadsConfigMapSpec{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func OrchestrationUserWorkloadsConfigMapSpec_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationUserWorkloadsConfigMapSpec) *pb.UserWorkloadsConfigMap {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsConfigMap{}
	// MISSING: Name
	// MISSING: Data
	return out
}
func UserWorkloadsConfigMap_FromProto(mapCtx *direct.MapContext, in *pb.UserWorkloadsConfigMap) *krm.UserWorkloadsConfigMap {
	if in == nil {
		return nil
	}
	out := &krm.UserWorkloadsConfigMap{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Data = in.Data
	return out
}
func UserWorkloadsConfigMap_ToProto(mapCtx *direct.MapContext, in *krm.UserWorkloadsConfigMap) *pb.UserWorkloadsConfigMap {
	if in == nil {
		return nil
	}
	out := &pb.UserWorkloadsConfigMap{}
	out.Name = direct.ValueOf(in.Name)
	out.Data = in.Data
	return out
}
