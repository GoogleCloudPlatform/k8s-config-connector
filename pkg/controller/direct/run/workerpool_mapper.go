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

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	krmrunv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apipb "google.golang.org/genproto/googleapis/api"
)

func BinaryAuthorization_UseDefault_ToProto(mapCtx *direct.MapContext, in *bool) *pb.BinaryAuthorization_UseDefault {
	if in == nil {
		return nil
	}
	return &pb.BinaryAuthorization_UseDefault{UseDefault: *in}
}

func EnvVar_Value_ToProto(mapCtx *direct.MapContext, in *string) *pb.EnvVar_Value {
	if in == nil {
		return nil
	}
	return &pb.EnvVar_Value{Value: *in}
}

func RunWorkerPoolSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmrunv1alpha1.RunWorkerPoolSpec) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	// MISSING: Name (handled by controller)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Labels (handled by controller via metadata.labels)
	out.Annotations = direct.MapStringString_ToProto(mapCtx, in.Annotations)
	out.Client = direct.ValueOf(in.Client)
	out.ClientVersion = direct.ValueOf(in.ClientVersion)
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	out.BinaryAuthorization = BinaryAuthorization_v1alpha1_ToProto(mapCtx, in.BinaryAuthorization)
	out.Template = WorkerPoolRevisionTemplate_v1alpha1_ToProto(mapCtx, in.Template)
	out.InstanceSplits = direct.Slice_ToProto(mapCtx, in.InstanceSplits, InstanceSplit_v1alpha1_ToProto)
	out.Scaling = WorkerPoolScaling_v1alpha1_ToProto(mapCtx, in.Scaling)
	return out
}
