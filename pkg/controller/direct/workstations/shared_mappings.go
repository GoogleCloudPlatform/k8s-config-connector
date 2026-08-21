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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

func WorkstationAnnotations_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationAnnotation) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for _, a := range in {
		out[a.Key] = a.Value
	}
	return out
}

func WorkstationLabels_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationLabel) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for _, l := range in {
		out[l.Key] = l.Value
	}
	return out
}

func WorkstationAnnotations_FromProto(mapCtx *direct.MapContext, in map[string]string) []krm.WorkstationAnnotation {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationAnnotation
	for k, v := range in {
		out = append(out, krm.WorkstationAnnotation{
			Key:   k,
			Value: v,
		})
	}
	return out
}

func WorkstationLabels_FromProto(mapCtx *direct.MapContext, in map[string]string) []krm.WorkstationLabel {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationLabel
	for k, v := range in {
		out = append(out, krm.WorkstationLabel{
			Key:   k,
			Value: v,
		})
	}
	return out
}

func WorkstationGCPConditions_FromProto(mapCtx *direct.MapContext, in []*status.Status) []krm.WorkstationServiceGCPCondition {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationServiceGCPCondition
	for _, c := range in {
		out = append(out, krm.WorkstationServiceGCPCondition{
			Code:    direct.LazyPtr(c.Code),
			Message: direct.LazyPtr(c.Message),
		})
	}
	return out
}

func WorkstationGCPConditions_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationServiceGCPCondition) []*status.Status {
	if in == nil {
		return nil
	}
	var out []*status.Status
	for _, c := range in {
		out = append(out, &status.Status{
			Code:    direct.ValueOf(c.Code),
			Message: direct.ValueOf(c.Message),
		})
	}
	return out
}
