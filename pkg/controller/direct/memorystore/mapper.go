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

package memorystore

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	timeofdaypb "google.golang.org/genproto/googleapis/type/timeofday"
)

// Mappers that are not automatically generated

func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofdaypb.TimeOfDay) *krm.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krm.TimeOfDay{}
	out.Hours = direct.LazyPtr(in.GetHours())
	out.Minutes = direct.LazyPtr(in.GetMinutes())
	out.Seconds = direct.LazyPtr(in.GetSeconds())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofdaypb.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofdaypb.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}
