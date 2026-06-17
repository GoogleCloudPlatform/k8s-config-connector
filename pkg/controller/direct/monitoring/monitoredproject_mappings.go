// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitoring

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	metricsscopepb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// MonitoringMonitoredProjectSpec_FromProto converts the Proto representation of MonitoredProject to KRM Spec.
// This is handcoded because the KRM Spec separates the parent Metrics Scope path from the Monitored Project ID (resourceID),
// whereas the Proto MonitoredProject represents both as a single resource name.
func MonitoringMonitoredProjectSpec_FromProto(mapCtx *direct.MapContext, in *metricsscopepb.MonitoredProject) *krm.MonitoringMonitoredProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMonitoredProjectSpec{}
	tokens := strings.Split(in.GetName(), "/")
	if len(tokens) == 6 && tokens[4] == "projects" {
		out.MetricsScope = strings.Join(tokens[0:4], "/")
		out.ResourceID = direct.LazyPtr(tokens[5])
	} else {
		// Fallback to name as-is if pattern doesn't match
		out.MetricsScope = in.GetName()
	}
	return out
}

// MonitoringMonitoredProjectSpec_ToProto converts the KRM Spec to the Proto representation of MonitoredProject.
// This is handcoded because the KRM Spec separates the parent Metrics Scope path from the Monitored Project ID (resourceID),
// whereas the Proto MonitoredProject represents both as a single resource name.
func MonitoringMonitoredProjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMonitoredProjectSpec) *metricsscopepb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &metricsscopepb.MonitoredProject{}
	name := in.MetricsScope
	if in.ResourceID != nil && *in.ResourceID != "" {
		name = fmt.Sprintf("%s/projects/%s", name, *in.ResourceID)
	}
	out.Name = name
	return out
}

// MonitoringMonitoredProjectStatus_FromProto converts the Proto representation of MonitoredProject to KRM Status.
// This is handcoded to map the proto create_time Timestamp to the KRM Status createTime string in RFC3339 format.
func MonitoringMonitoredProjectStatus_FromProto(mapCtx *direct.MapContext, in *metricsscopepb.MonitoredProject) *krm.MonitoringMonitoredProjectStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMonitoredProjectStatus{}
	if in.CreateTime != nil {
		out.CreateTime = direct.LazyPtr(in.CreateTime.AsTime().Format(time.RFC3339Nano))
	}
	return out
}

// MonitoringMonitoredProjectStatus_ToProto converts the KRM Status to the Proto representation of MonitoredProject.
// This is handcoded to map the KRM Status createTime string in RFC3339 format to the proto create_time Timestamp.
func MonitoringMonitoredProjectStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMonitoredProjectStatus) *metricsscopepb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &metricsscopepb.MonitoredProject{}
	if in.CreateTime != nil {
		t, err := time.Parse(time.RFC3339Nano, *in.CreateTime)
		if err == nil {
			out.CreateTime = timestamppb.New(t)
		} else {
			mapCtx.Errorf("parsing createTime: %w", err)
		}
	}
	return out
}
