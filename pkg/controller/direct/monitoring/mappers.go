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
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ServiceTelemetry_FromProto maps from protobuf Service_Telemetry to KRM ServiceTelemetry.
func ServiceTelemetry_FromProto(mapCtx *direct.MapContext, in *pb.Service_Telemetry) *krm.ServiceTelemetry {
	if in == nil {
		return nil
	}
	out := &krm.ServiceTelemetry{}
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	return out
}

// ServiceTelemetry_ToProto maps from KRM ServiceTelemetry to protobuf Service_Telemetry.
func ServiceTelemetry_ToProto(mapCtx *direct.MapContext, in *krm.ServiceTelemetry) *pb.Service_Telemetry {
	if in == nil {
		return nil
	}
	out := &pb.Service_Telemetry{}
	out.ResourceName = direct.ValueOf(in.ResourceName)
	return out
}
