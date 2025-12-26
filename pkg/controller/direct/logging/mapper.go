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

// +generated:mapper
// krm.group: logging.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.logging.v2

package logging

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
)

func LogmetricMetricDescriptorStatus_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.LogmetricMetricDescriptorStatus {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricMetricDescriptorStatus{}
	out.Description = direct.LazyPtr(in.Description)
	out.MonitoredResourceTypes = in.MonitoredResourceTypes
	out.Name = direct.LazyPtr(in.Name)
	out.Type = direct.LazyPtr(in.Type)

	return out
}
func LogmetricMetricDescriptorStatus_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricMetricDescriptorStatus) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	out.Description = direct.ValueOf(in.Description)
	out.MonitoredResourceTypes = in.MonitoredResourceTypes
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)

	return out
}
