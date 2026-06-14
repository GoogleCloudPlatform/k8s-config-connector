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

package monitoring

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apipb "google.golang.org/genproto/googleapis/api"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
)

// MonitoringMetricDescriptorSpec is hand-coded because of string-to-enum field mappings.
func MonitoringMetricDescriptorSpec_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.MonitoringMetricDescriptorSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMetricDescriptorSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = direct.Slice_FromProto(mapCtx, in.GetLabels(), MetricdescriptorLabels_FromProto)
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.Metadata = MetricdescriptorMetadata_FromProto(mapCtx, in.GetMetadata())
	out.MetricKind = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetMetricKind()))
	out.Type = in.GetType()
	out.Unit = direct.LazyPtr(in.GetUnit())
	out.ValueType = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetValueType()))
	return out
}

func MonitoringMetricDescriptorSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMetricDescriptorSpec) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = direct.Slice_ToProto(mapCtx, in.Labels, MetricdescriptorLabels_ToProto)
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	out.Metadata = MetricdescriptorMetadata_ToProto(mapCtx, in.Metadata)
	out.MetricKind = direct.Enum_ToProto[metricpb.MetricDescriptor_MetricKind](mapCtx, &in.MetricKind)
	out.Type = in.Type
	out.Unit = direct.ValueOf(in.Unit)
	out.ValueType = direct.Enum_ToProto[metricpb.MetricDescriptor_ValueType](mapCtx, &in.ValueType)
	return out
}

// MetricdescriptorMetadata is hand-coded because IngestDelay and SamplePeriod map between durationpb.Duration and KRM strings.
func MetricdescriptorMetadata_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor_MetricDescriptorMetadata) *krm.MetricdescriptorMetadata {
	if in == nil {
		return nil
	}
	out := &krm.MetricdescriptorMetadata{}
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.SamplePeriod = direct.Duration_FromProto(mapCtx, in.GetSamplePeriod())
	out.IngestDelay = direct.Duration_FromProto(mapCtx, in.GetIngestDelay())
	return out
}

func MetricdescriptorMetadata_ToProto(mapCtx *direct.MapContext, in *krm.MetricdescriptorMetadata) *metricpb.MetricDescriptor_MetricDescriptorMetadata {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor_MetricDescriptorMetadata{}
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	out.SamplePeriod = direct.Duration_ToProto(mapCtx, in.SamplePeriod)
	out.IngestDelay = direct.Duration_ToProto(mapCtx, in.IngestDelay)
	return out
}

// MonitoringMetricDescriptorStatus is hand-coded because the resource Name is translated to SelfLink.
func MonitoringMetricDescriptorStatus_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.MonitoringMetricDescriptorStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMetricDescriptorStatus{}
	out.MonitoredResourceTypes = in.GetMonitoredResourceTypes()
	out.SelfLink = direct.LazyPtr(in.GetName())
	return out
}

func MonitoringMetricDescriptorStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMetricDescriptorStatus) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	out.MonitoredResourceTypes = in.MonitoredResourceTypes
	out.Name = direct.ValueOf(in.SelfLink)
	return out
}
