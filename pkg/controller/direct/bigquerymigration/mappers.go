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

package bigquerymigration

import (
	pb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerymigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

func Tasks_FromProto(mapCtx *direct.MapContext, in map[string]*pb.MigrationTask) map[string]krm.MigrationTask {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.MigrationTask)
	for k, v := range in {
		if v != nil {
			out[k] = *MigrationTask_FromProto(mapCtx, v)
		}
	}
	return out
}

func Tasks_ToProto(mapCtx *direct.MapContext, in map[string]krm.MigrationTask) map[string]*pb.MigrationTask {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.MigrationTask)
	for k, v := range in {
		out[k] = MigrationTask_ToProto(mapCtx, &v)
	}
	return out
}

func Any_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krm.Any {
	if in == nil {
		return nil
	}
	out := &krm.Any{}
	out.TypeURL = direct.LazyPtr(in.GetTypeUrl())
	out.Value = in.GetValue()
	return out
}

func Any_ToProto(mapCtx *direct.MapContext, in *krm.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	out := &anypb.Any{}
	out.TypeUrl = direct.ValueOf(in.TypeURL)
	out.Value = in.Value
	return out
}
