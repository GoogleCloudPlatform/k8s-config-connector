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

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DocumentAIProcessorDefaultVersionSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SetDefaultProcessorVersionRequest) *krmv1alpha1.DocumentAIProcessorDefaultVersionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DocumentAIProcessorDefaultVersionSpec{}
	out.Version = in.GetDefaultProcessorVersion()
	return out
}

func DocumentAIProcessorDefaultVersionSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DocumentAIProcessorDefaultVersionSpec) *pb.SetDefaultProcessorVersionRequest {
	if in == nil {
		return nil
	}
	out := &pb.SetDefaultProcessorVersionRequest{}
	out.DefaultProcessorVersion = in.Version
	return out
}
