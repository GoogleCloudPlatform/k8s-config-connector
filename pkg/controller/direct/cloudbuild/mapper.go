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

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krmcloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func PrivatePoolV1Config_NetworkConfigSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krmcloudbuildv1beta1.PrivatePoolV1Config_NetworkConfigSpec {
	return PrivatePoolV1Config_NetworkConfigSpec_FromProto(mapCtx, in)
}

func PrivatePoolV1Config_NetworkConfigSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1beta1.PrivatePoolV1Config_NetworkConfigSpec) *pb.PrivatePoolV1Config_NetworkConfig {
	return PrivatePoolV1Config_NetworkConfigSpec_ToProto(mapCtx, in)
}

func CloudBuildWorkerPoolSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1beta1.CloudBuildWorkerPoolSpec) *pb.WorkerPool {
	return CloudBuildWorkerPoolSpec_ToProto(mapCtx, in)
}
