// Copyright 2024 Google LLC
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

package workstations

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ApplyWorkstationConfigGCPDefaults(mapCtx *direct.MapContext, in *krm.WorkstationConfigSpec, out *pb.WorkstationConfig, actual *pb.WorkstationConfig) {
	if in.IdleTimeout == nil {
		// GCP default IdleTimeout is 20 minutes.
		defaultIdleTimeout := "1200s"
		out.IdleTimeout = direct.StringDuration_ToProto(mapCtx, &defaultIdleTimeout)
	}
	if in.RunningTimeout == nil {
		// GCP default RunningTimeout is 12 hours.
		defaultRunningTimeout := "43200s"
		out.RunningTimeout = direct.StringDuration_ToProto(mapCtx, &defaultRunningTimeout)
	}
	if in.Host == nil && actual != nil {
		// If desired host config is not specified, assume user wants the actual.
		out.Host = actual.Host
	}
	if in.Container == nil && actual != nil {
		// If desired container config is not specified, assume user wants the actual.
		out.Container = actual.Container
	}
	if in.ReplicaZones == nil && actual != nil {
		// If desired replica zones are not specified, assume user wants the actual.
		out.ReplicaZones = actual.ReplicaZones
	}
}
