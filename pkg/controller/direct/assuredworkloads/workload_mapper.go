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
// krm.group: assuredworkloads.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.assuredworkloads.v1

package assuredworkloads

import (
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Workload_SaaEnrollmentResponse_FromProto(mapCtx *direct.MapContext, in *pb.Workload_SaaEnrollmentResponse) *krm.Workload_SaaEnrollmentResponse {
	if in == nil {
		return nil
	}
	out := &krm.Workload_SaaEnrollmentResponse{}
	if in.SetupErrors != nil {
		out.SetupErrors = []string{}
		for _, e := range in.SetupErrors {
			out.SetupErrors = append(out.SetupErrors, e.String())
		}
	}
	out.SetupStatus = direct.LazyPtr(in.SetupStatus.String())
	return out
}

func Workload_SaaEnrollmentResponse_ToProto(mapCtx *direct.MapContext, in *krm.Workload_SaaEnrollmentResponse) *pb.Workload_SaaEnrollmentResponse {
	if in == nil {
		return nil
	}
	out := &pb.Workload_SaaEnrollmentResponse{}
	out.SetupErrors = []pb.Workload_SaaEnrollmentResponse_SetupError{}
	for _, e := range in.SetupErrors {
		setupErr := pb.Workload_SaaEnrollmentResponse_SetupError_value[e]
		out.SetupErrors = append(out.SetupErrors, pb.Workload_SaaEnrollmentResponse_SetupError(setupErr))
	}

	setupState := pb.Workload_SaaEnrollmentResponse_SetupState_value[direct.ValueOf(in.SetupStatus)]
	out.SetupStatus = direct.LazyPtr(pb.Workload_SaaEnrollmentResponse_SetupState(setupState))
	return out
}
