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

package monitoring

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func MonitoringAlertChartObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlertChart) *krm.MonitoringAlertChartObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertChartObservedState{}
	// MISSING: Name
	return out
}
func MonitoringAlertChartObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertChartObservedState) *pb.AlertChart {
	if in == nil {
		return nil
	}
	out := &pb.AlertChart{}
	// MISSING: Name
	return out
}
func MonitoringAlertChartSpec_FromProto(mapCtx *direct.MapContext, in *pb.AlertChart) *krm.MonitoringAlertChartSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertChartSpec{}
	// MISSING: Name
	return out
}
func MonitoringAlertChartSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertChartSpec) *pb.AlertChart {
	if in == nil {
		return nil
	}
	out := &pb.AlertChart{}
	// MISSING: Name
	return out
}
