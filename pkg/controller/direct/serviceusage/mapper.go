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

package serviceusage

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1beta1"
)

func ServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ServiceSpec{}
	// ProjectRef and ResourceID are handled by the reconciler/resolver
	return out
}

func ServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// ProjectRef and ResourceID are handled by the reconciler/resolver
	return out
}

func ServiceStatus_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ServiceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ServiceStatus{}
	// Conditions and ObservedGeneration are handled by the reconciler
	return out
}

func ServiceStatus_ToProto(mapCtx *direct.MapContext, in *krm.ServiceStatus) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	return out
}
